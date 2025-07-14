-- Function to calculate recipe average rating
CREATE OR REPLACE FUNCTION get_recipe_average_rating(recipe_row recipes)
RETURNS DECIMAL AS $$
BEGIN
    RETURN (
        SELECT COALESCE(AVG(rating), 0)
        FROM recipe_ratings
        WHERE recipe_id = recipe_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get recipe total likes
CREATE OR REPLACE FUNCTION get_recipe_likes_count(recipe_row recipes)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM recipe_likes
        WHERE recipe_id = recipe_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get recipe total reviews
CREATE OR REPLACE FUNCTION get_recipe_reviews_count(recipe_row recipes)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM recipe_reviews
        WHERE recipe_id = recipe_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get recipe total views
CREATE OR REPLACE FUNCTION get_recipe_views_count(recipe_row recipes)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM recipe_views
        WHERE recipe_id = recipe_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get user recipe count
CREATE OR REPLACE FUNCTION get_user_recipe_count(user_row users)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM recipes
        WHERE author_id = user_row.id AND status = 'published'
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get user followers count
CREATE OR REPLACE FUNCTION get_user_followers_count(user_row users)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM user_follows
        WHERE following_id = user_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get user following count
CREATE OR REPLACE FUNCTION get_user_following_count(user_row users)
RETURNS INTEGER AS $$
BEGIN
    RETURN (
        SELECT COUNT(*)
        FROM user_follows
        WHERE follower_id = user_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get user average rating
CREATE OR REPLACE FUNCTION get_user_average_rating(user_row users)
RETURNS DECIMAL AS $$
BEGIN
    RETURN (
        SELECT COALESCE(AVG(rr.rating), 0)
        FROM recipe_reviews rr
        JOIN recipes r ON r.id = rr.recipe_id
        WHERE r.author_id = user_row.id
    );
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to search recipes with filters
CREATE OR REPLACE FUNCTION search_recipes(
    search_query TEXT DEFAULT NULL,
    category_ids UUID[] DEFAULT NULL,
    difficulty_levels TEXT[] DEFAULT NULL,
    max_prep_time INTEGER DEFAULT NULL,
    min_rating DECIMAL DEFAULT NULL,
    cuisine_types TEXT[] DEFAULT NULL,
    ingredient_names TEXT[] DEFAULT NULL,
    sort_by TEXT DEFAULT 'created_at',
    sort_order TEXT DEFAULT 'DESC',
    page_limit INTEGER DEFAULT 20,
    page_offset INTEGER DEFAULT 0
)
RETURNS TABLE (
    id UUID,
    title VARCHAR,
    description TEXT,
    featured_image TEXT,
    prep_time INTEGER,
    cook_time INTEGER,
    total_time INTEGER,
    servings INTEGER,
    difficulty VARCHAR,
    cuisine_type VARCHAR,
    price DECIMAL,
    author_id UUID,
    category_id UUID,
    created_at TIMESTAMP,
    average_rating DECIMAL,
    likes_count INTEGER,
    reviews_count INTEGER
) AS $$
DECLARE
    base_query TEXT;
    where_conditions TEXT[] := ARRAY[]::TEXT[];
    order_clause TEXT;
BEGIN
    -- Build WHERE conditions
    IF search_query IS NOT NULL AND search_query != '' THEN
        where_conditions := array_append(where_conditions, 
            format('to_tsvector(''english'', r.title || '' '' || r.description) @@ plainto_tsquery(''english'', %L)', search_query));
    END IF;
    
    IF category_ids IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            format('r.category_id = ANY(%L)', category_ids));
    END IF;
    
    IF difficulty_levels IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            format('r.difficulty = ANY(%L)', difficulty_levels));
    END IF;
    
    IF max_prep_time IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            format('r.prep_time <= %s', max_prep_time));
    END IF;
    
    IF cuisine_types IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            format('r.cuisine_type = ANY(%L)', cuisine_types));
    END IF;
    
    IF ingredient_names IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            'EXISTS (SELECT 1 FROM recipe_ingredients ri JOIN ingredients i ON i.id = ri.ingredient_id WHERE ri.recipe_id = r.id AND i.name = ANY(' || quote_literal(ingredient_names) || '))');
    END IF;
    
    IF min_rating IS NOT NULL THEN
        where_conditions := array_append(where_conditions, 
            format('(SELECT COALESCE(AVG(rating), 0) FROM recipe_ratings WHERE recipe_id = r.id) >= %s', min_rating));
    END IF;
    
    -- Always filter for published recipes
    where_conditions := array_append(where_conditions, 'r.status = ''published''');
    
    -- Build ORDER BY clause
    CASE sort_by
        WHEN 'rating' THEN order_clause := '(SELECT COALESCE(AVG(rating), 0) FROM recipe_ratings WHERE recipe_id = r.id)';
        WHEN 'likes' THEN order_clause := '(SELECT COUNT(*) FROM recipe_likes WHERE recipe_id = r.id)';
        WHEN 'prep_time' THEN order_clause := 'r.prep_time';
        WHEN 'title' THEN order_clause := 'r.title';
        ELSE order_clause := 'r.created_at';
    END CASE;
    
    order_clause := order_clause || ' ' || sort_order;
    
    -- Build and execute query
    base_query := format('
        SELECT 
            r.id,
            r.title,
            r.description,
            r.featured_image,
            r.prep_time,
            r.cook_time,
            r.total_time,
            r.servings,
            r.difficulty,
            r.cuisine_type,
            r.price,
            r.author_id,
            r.category_id,
            r.created_at,
            (SELECT COALESCE(AVG(rating), 0) FROM recipe_ratings WHERE recipe_id = r.id) as average_rating,
            (SELECT COUNT(*) FROM recipe_likes WHERE recipe_id = r.id) as likes_count,
            (SELECT COUNT(*) FROM recipe_reviews WHERE recipe_id = r.id) as reviews_count
        FROM recipes r
        WHERE %s
        ORDER BY %s
        LIMIT %s OFFSET %s',
        array_to_string(where_conditions, ' AND '),
        order_clause,
        page_limit,
        page_offset
    );
    
    RETURN QUERY EXECUTE base_query;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to get recipe recommendations
CREATE OR REPLACE FUNCTION get_recipe_recommendations(
    user_id_param UUID,
    limit_param INTEGER DEFAULT 10
)
RETURNS TABLE (
    recipe_id UUID,
    title VARCHAR,
    featured_image TEXT,
    average_rating DECIMAL,
    recommendation_score DECIMAL
) AS $$
BEGIN
    RETURN QUERY
    WITH user_preferences AS (
        -- Get user's liked categories and cuisines
        SELECT 
            r.category_id,
            r.cuisine_type,
            COUNT(*) as preference_weight
        FROM recipe_likes rl
        JOIN recipes r ON r.id = rl.recipe_id
        WHERE rl.user_id = user_id_param
        GROUP BY r.category_id, r.cuisine_type
    ),
    recipe_scores AS (
        SELECT 
            r.id as recipe_id,
            r.title,
            r.featured_image,
            (SELECT COALESCE(AVG(rating), 0) FROM recipe_ratings WHERE recipe_id = r.id) as average_rating,
            -- Calculate recommendation score based on various factors
            (
                COALESCE((SELECT preference_weight FROM user_preferences WHERE category_id = r.category_id), 0) * 0.3 +
                COALESCE((SELECT preference_weight FROM user_preferences WHERE cuisine_type = r.cuisine_type), 0) * 0.2 +
                (SELECT COALESCE(AVG(rating), 0) FROM recipe_ratings WHERE recipe_id = r.id) * 0.3 +
                (SELECT COUNT(*) FROM recipe_likes WHERE recipe_id = r.id) * 0.1 +
                CASE WHEN r.created_at > NOW() - INTERVAL '30 days' THEN 0.1 ELSE 0 END
            ) as recommendation_score
        FROM recipes r
        WHERE r.status = 'published'
        AND r.author_id != user_id_param
        AND NOT EXISTS (
            SELECT 1 FROM recipe_likes rl WHERE rl.recipe_id = r.id AND rl.user_id = user_id_param
        )
    )
    SELECT 
        rs.recipe_id,
        rs.title,
        rs.featured_image,
        rs.average_rating,
        rs.recommendation_score
    FROM recipe_scores rs
    ORDER BY rs.recommendation_score DESC, rs.average_rating DESC
    LIMIT limit_param;
END;
$$ LANGUAGE plpgsql STABLE;

-- Function to update recipe search rank (for analytics)
CREATE OR REPLACE FUNCTION update_recipe_search_rank()
RETURNS TRIGGER AS $$
BEGIN
    -- Update recipe popularity score based on views, likes, ratings
    UPDATE recipes 
    SET updated_at = NOW()
    WHERE id = NEW.recipe_id;
    
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
