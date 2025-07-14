-- Trigger to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Apply updated_at trigger to relevant tables
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_recipes_updated_at BEFORE UPDATE ON recipes
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_categories_updated_at BEFORE UPDATE ON categories
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_recipe_nutrition_updated_at BEFORE UPDATE ON recipe_nutrition
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_recipe_ratings_updated_at BEFORE UPDATE ON recipe_ratings
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_recipe_reviews_updated_at BEFORE UPDATE ON recipe_reviews
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Trigger to generate recipe slug from title
CREATE OR REPLACE FUNCTION generate_recipe_slug()
RETURNS TRIGGER AS $$
DECLARE
    base_slug TEXT;
    final_slug TEXT;
    counter INTEGER := 0;
BEGIN
    -- Generate base slug from title
    base_slug := lower(trim(regexp_replace(NEW.title, '[^a-zA-Z0-9\s]', '', 'g')));
    base_slug := regexp_replace(base_slug, '\s+', '-', 'g');
    
    final_slug := base_slug;
    
    -- Check for uniqueness and append counter if needed
    WHILE EXISTS (SELECT 1 FROM recipes WHERE slug = final_slug AND id != COALESCE(NEW.id, '00000000-0000-0000-0000-000000000000'::UUID)) LOOP
        counter := counter + 1;
        final_slug := base_slug || '-' || counter;
    END LOOP;
    
    NEW.slug := final_slug;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER generate_recipe_slug_trigger BEFORE INSERT OR UPDATE ON recipes
    FOR EACH ROW EXECUTE FUNCTION generate_recipe_slug();

-- Trigger to generate category slug from name
CREATE OR REPLACE FUNCTION generate_category_slug()
RETURNS TRIGGER AS $$
DECLARE
    base_slug TEXT;
    final_slug TEXT;
    counter INTEGER := 0;
BEGIN
    base_slug := lower(trim(regexp_replace(NEW.name, '[^a-zA-Z0-9\s]', '', 'g')));
    base_slug := regexp_replace(base_slug, '\s+', '-', 'g');
    
    final_slug := base_slug;
    
    WHILE EXISTS (SELECT 1 FROM categories WHERE slug = final_slug AND id != COALESCE(NEW.id, '00000000-0000-0000-0000-000000000000'::UUID)) LOOP
        counter := counter + 1;
        final_slug := base_slug || '-' || counter;
    END LOOP;
    
    NEW.slug := final_slug;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER generate_category_slug_trigger BEFORE INSERT OR UPDATE ON categories
    FOR EACH ROW EXECUTE FUNCTION generate_category_slug();

-- Trigger to track recipe views
CREATE TRIGGER track_recipe_views AFTER INSERT ON recipe_views
    FOR EACH ROW EXECUTE FUNCTION update_recipe_search_rank();

-- Trigger to update recipe popularity on likes
CREATE TRIGGER update_recipe_popularity_on_like AFTER INSERT OR DELETE ON recipe_likes
    FOR EACH ROW EXECUTE FUNCTION update_recipe_search_rank();

-- Trigger to award achievements
CREATE OR REPLACE FUNCTION check_user_achievements()
RETURNS TRIGGER AS $$
BEGIN
    -- Award "First Recipe" achievement
    IF TG_TABLE_NAME = 'recipes' AND TG_OP = 'INSERT' THEN
        IF (SELECT COUNT(*) FROM recipes WHERE author_id = NEW.author_id) = 1 THEN
            INSERT INTO user_achievements (user_id, achievement_type, achievement_name, description, icon)
            VALUES (NEW.author_id, 'first_recipe', 'First Recipe', 'Published your first recipe', 'chef-hat')
            ON CONFLICT DO NOTHING;
        END IF;
        
        -- Award "Recipe Master" achievement (10 recipes)
        IF (SELECT COUNT(*) FROM recipes WHERE author_id = NEW.author_id) = 10 THEN
            INSERT INTO user_achievements (user_id, achievement_type, achievement_name, description, icon)
            VALUES (NEW.author_id, 'recipe_master', 'Recipe Master', 'Published 10 recipes', 'crown')
            ON CONFLICT DO NOTHING;
        END IF;
    END IF;
    
    -- Award "Popular Recipe" achievement (100 likes)
    IF TG_TABLE_NAME = 'recipe_likes' AND TG_OP = 'INSERT' THEN
        IF (SELECT COUNT(*) FROM recipe_likes WHERE recipe_id = NEW.recipe_id) = 100 THEN
            INSERT INTO user_achievements (user_id, achievement_type, achievement_name, description, icon)
            SELECT r.author_id, 'popular_recipe', 'Popular Recipe', 'Recipe reached 100 likes', 'heart'
            FROM recipes r WHERE r.id = NEW.recipe_id
            ON CONFLICT DO NOTHING;
        END IF;
    END IF;
    
    RETURN COALESCE(NEW, OLD);
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER award_recipe_achievements AFTER INSERT ON recipes
    FOR EACH ROW EXECUTE FUNCTION check_user_achievements();

CREATE TRIGGER award_like_achievements AFTER INSERT ON recipe_likes
    FOR EACH ROW EXECUTE FUNCTION check_user_achievements();

-- Trigger to prevent self-following
CREATE OR REPLACE FUNCTION prevent_self_follow()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.follower_id = NEW.following_id THEN
        RAISE EXCEPTION 'Users cannot follow themselves';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER prevent_self_follow_trigger BEFORE INSERT OR UPDATE ON user_follows
    FOR EACH ROW EXECUTE FUNCTION prevent_self_follow();
