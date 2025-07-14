-- Insert default categories
INSERT INTO categories (name, description, image) VALUES
('Breakfast', 'Start your day with delicious breakfast recipes', '/placeholder.svg?height=200&width=200'),
('Lunch', 'Quick and satisfying lunch ideas', '/placeholder.svg?height=200&width=200'),
('Dinner', 'Hearty dinner recipes for the whole family', '/placeholder.svg?height=200&width=200'),
('Desserts', 'Sweet treats and desserts', '/placeholder.svg?height=200&width=200'),
('Appetizers', 'Perfect starters and appetizers', '/placeholder.svg?height=200&width=200'),
('Beverages', 'Refreshing drinks and beverages', '/placeholder.svg?height=200&width=200'),
('Snacks', 'Quick and easy snack recipes', '/placeholder.svg?height=200&width=200'),
('Salads', 'Fresh and healthy salad recipes', '/placeholder.svg?height=200&width=200'),
('Soups', 'Comforting soups and stews', '/placeholder.svg?height=200&width=200'),
('Vegetarian', 'Delicious meat-free recipes', '/placeholder.svg?height=200&width=200'),
('Vegan', 'Plant-based recipes', '/placeholder.svg?height=200&width=200'),
('Gluten-Free', 'Gluten-free recipe options', '/placeholder.svg?height=200&width=200');

-- Insert common ingredients
INSERT INTO ingredients (name, category) VALUES
-- Proteins
('Chicken Breast', 'Protein'),
('Ground Beef', 'Protein'),
('Salmon', 'Protein'),
('Eggs', 'Protein'),
('Tofu', 'Protein'),
('Black Beans', 'Protein'),

-- Vegetables
('Onion', 'Vegetable'),
('Garlic', 'Vegetable'),
('Tomatoes', 'Vegetable'),
('Bell Peppers', 'Vegetable'),
('Carrots', 'Vegetable'),
('Broccoli', 'Vegetable'),
('Spinach', 'Vegetable'),
('Mushrooms', 'Vegetable'),
('Potatoes', 'Vegetable'),
('Sweet Potatoes', 'Vegetable'),

-- Grains & Starches
('Rice', 'Grain'),
('Pasta', 'Grain'),
('Bread', 'Grain'),
('Quinoa', 'Grain'),
('Oats', 'Grain'),
('Flour', 'Grain'),

-- Dairy
('Milk', 'Dairy'),
('Cheese', 'Dairy'),
('Butter', 'Dairy'),
('Yogurt', 'Dairy'),
('Cream', 'Dairy'),

-- Spices & Herbs
('Salt', 'Spice'),
('Black Pepper', 'Spice'),
('Olive Oil', 'Oil'),
('Basil', 'Herb'),
('Oregano', 'Herb'),
('Thyme', 'Herb'),
('Rosemary', 'Herb'),
('Paprika', 'Spice'),
('Cumin', 'Spice'),
('Garlic Powder', 'Spice'),

-- Pantry Items
('Sugar', 'Pantry'),
('Brown Sugar', 'Pantry'),
('Honey', 'Pantry'),
('Vanilla Extract', 'Pantry'),
('Baking Powder', 'Pantry'),
('Baking Soda', 'Pantry'),
('Soy Sauce', 'Pantry'),
('Vinegar', 'Pantry'),
('Lemon Juice', 'Pantry'),
('Coconut Oil', 'Oil');

-- Create a demo user (password: 'password123')
INSERT INTO users (email, username, first_name, last_name, password_hash, bio, avatar, is_verified) VALUES
('demo@recipehub.com', 'demo_chef', 'Demo', 'Chef', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi', 'Passionate home cook sharing delicious recipes with the world!', '/placeholder.svg?height=96&width=96', true);

-- Get the demo user ID for sample recipes
DO $$
DECLARE
    demo_user_id UUID;
    breakfast_cat_id UUID;
    dinner_cat_id UUID;
    dessert_cat_id UUID;
    recipe_id UUID;
    ingredient_id UUID;
BEGIN
    -- Get IDs
    SELECT id INTO demo_user_id FROM users WHERE username = 'demo_chef';
    SELECT id INTO breakfast_cat_id FROM categories WHERE slug = 'breakfast';
    SELECT id INTO dinner_cat_id FROM categories WHERE slug = 'dinner';
    SELECT id INTO dessert_cat_id FROM categories WHERE slug = 'desserts';

    -- Insert sample recipe 1: Pancakes
    INSERT INTO recipes (title, description, featured_image, prep_time, cook_time, servings, difficulty, cuisine_type, author_id, category_id)
    VALUES ('Fluffy Pancakes', 'Light and fluffy pancakes perfect for weekend breakfast', '/placeholder.svg?height=400&width=600', 15, 20, 4, 'Easy', 'American', demo_user_id, breakfast_cat_id)
    RETURNING id INTO recipe_id;

    -- Add pancake ingredients
    INSERT INTO recipe_ingredients (recipe_id, ingredient_id, amount, unit, sort_order)
    SELECT recipe_id, i.id, amounts.amount, amounts.unit, amounts.sort_order
    FROM (VALUES
        ('Flour', 2, 'cups', 1),
        ('Sugar', 2, 'tbsp', 2),
        ('Baking Powder', 2, 'tsp', 3),
        ('Salt', 0.5, 'tsp', 4),
        ('Milk', 1.75, 'cups', 5),
        ('Eggs', 2, 'large', 6),
        ('Butter', 4, 'tbsp', 7)
    ) AS amounts(ingredient_name, amount, unit, sort_order)
    JOIN ingredients i ON i.name = amounts.ingredient_name;

    -- Add pancake steps
    INSERT INTO recipe_steps (recipe_id, step_number, instruction, timer_minutes) VALUES
    (recipe_id, 1, 'In a large bowl, whisk together flour, sugar, baking powder, and salt.', NULL),
    (recipe_id, 2, 'In another bowl, whisk together milk, eggs, and melted butter.', NULL),
    (recipe_id, 3, 'Pour the wet ingredients into the dry ingredients and stir until just combined. Do not overmix.', NULL),
    (recipe_id, 4, 'Heat a griddle or large skillet over medium heat. Lightly grease with butter.', 2),
    (recipe_id, 5, 'Pour 1/4 cup of batter for each pancake onto the griddle.', NULL),
    (recipe_id, 6, 'Cook until bubbles form on the surface and edges look set, then flip and cook until golden brown.', 3),
    (recipe_id, 7, 'Serve hot with maple syrup and butter.', NULL);

    -- Add nutrition info
    INSERT INTO recipe_nutrition (recipe_id, calories, protein, carbohydrates, fat, fiber, sugar)
    VALUES (recipe_id, 285, 8.5, 42, 9.2, 1.8, 12);

    -- Insert sample recipe 2: Spaghetti Carbonara
    INSERT INTO recipes (title, description, featured_image, prep_time, cook_time, servings, difficulty, cuisine_type, author_id, category_id)
    VALUES ('Classic Spaghetti Carbonara', 'Authentic Italian carbonara with eggs, cheese, and pancetta', '/placeholder.svg?height=400&width=600', 10, 15, 4, 'Medium', 'Italian', demo_user_id, dinner_cat_id)
    RETURNING id INTO recipe_id;

    -- Add more sample recipes as needed...

END $$;
