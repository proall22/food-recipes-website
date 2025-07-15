import { ref } from 'vue'
import { useNuxtApp } from '#app'
import { gql } from 'graphql-tag'

export const useRecipes = () => {
  const { $apollo } = useNuxtApp()
  
  const recipes = ref([])
  const recipe = ref(null)
  const categories = ref([])
  const isLoading = ref(false)
  const totalRecipes = ref(0)

  // Get all recipes with filters
  const getRecipes = async (filters = {}) => {
    isLoading.value = true

    try {
      const GET_RECIPES_QUERY = gql`
        query GetRecipes(
          $where: recipes_bool_exp
          $orderBy: [recipes_order_by!]
          $limit: Int
          $offset: Int
        ) {
          recipes(
            where: $where
            order_by: $orderBy
            limit: $limit
            offset: $offset
          ) {
            id
            title
            description
            featured_image
            prep_time
            cook_time
            total_time
            servings
            difficulty
            cuisine_type
            price
            status
            created_at
            updated_at
            average_rating
            likes_count
            reviews_count
            category {
              id
              name
              slug
            }
            author {
              id
              first_name
              last_name
              username
              avatar
            }
          }
          recipes_aggregate(where: $where) {
            aggregate {
              count
            }
          }
        }
      `

      // Build where clause based on filters
      const where = { status: { _eq: 'published' } }
      
      if (filters.search) {
        where._or = [
          { title: { _ilike: `%${filters.search}%` } },
          { description: { _ilike: `%${filters.search}%` } }
        ]
      }
      
      if (filters.category) {
        where.category_id = { _eq: filters.category }
      }
      
      if (filters.difficulty) {
        where.difficulty = { _eq: filters.difficulty }
      }
      
      if (filters.maxPrepTime) {
        where.prep_time = { _lte: filters.maxPrepTime }
      }
      
      if (filters.cuisineType) {
        where.cuisine_type = { _eq: filters.cuisineType }
      }

      // Build order by clause
      let orderBy = [{ created_at: 'desc' }]
      if (filters.sortBy) {
        const [field, direction] = filters.sortBy.split('_')
        orderBy = [{ [field]: direction.toLowerCase() }]
      }

      const { data } = await $apollo.query({
        query: GET_RECIPES_QUERY,
        variables: {
          where,
          orderBy,
          limit: filters.limit || 12,
          offset: filters.offset || 0
        }
      })

      recipes.value = data.recipes
      totalRecipes.value = data.recipes_aggregate.aggregate.count

      return data.recipes
    } catch (error) {
      console.error('Error fetching recipes:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Get single recipe by ID
  const getRecipe = async (id) => {
    isLoading.value = true

    try {
      const GET_RECIPE_QUERY = gql`
        query GetRecipe($id: uuid!) {
          recipes_by_pk(id: $id) {
            id
            title
            slug
            description
            featured_image
            prep_time
            cook_time
            total_time
            servings
            difficulty
            cuisine_type
            price
            status
            created_at
            updated_at
            average_rating
            likes_count
            reviews_count
            views_count
            category {
              id
              name
              slug
            }
            author {
              id
              first_name
              last_name
              username
              avatar
              recipe_count
            }
            images {
              id
              image_url
              alt_text
              sort_order
            }
            ingredients {
              id
              amount
              unit
              notes
              sort_order
              ingredient {
                id
                name
                category
              }
            }
            steps {
              id
              step_number
              instruction
              image_url
              timer_minutes
              temperature
            }
            nutrition {
              id
              calories
              protein
              carbohydrates
              fat
              fiber
              sugar
              sodium
            }
            reviews(order_by: { created_at: desc }, limit: 10) {
              id
              rating
              comment
              created_at
              user {
                id
                first_name
                last_name
                username
                avatar
              }
              images {
                id
                image_url
              }
            }
          }
        }
      `

      const { data } = await $apollo.query({
        query: GET_RECIPE_QUERY,
        variables: { id }
      })

      recipe.value = data.recipes_by_pk
      return data.recipes_by_pk
    } catch (error) {
      console.error('Error fetching recipe:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Get categories
  const getCategories = async () => {
    try {
      const GET_CATEGORIES_QUERY = gql`
        query GetCategories {
          categories(where: { is_active: { _eq: true } }) {
            id
            name
            slug
            description
            image
            recipes_aggregate {
              aggregate {
                count
              }
            }
          }
        }
      `

      const { data } = await $apollo.query({
        query: GET_CATEGORIES_QUERY
      })

      categories.value = data.categories.map(category => ({
        ...category,
        count: category.recipes_aggregate.aggregate.count
      }))

      return categories.value
    } catch (error) {
      console.error('Error fetching categories:', error)
      throw error
    }
  }

  // Create recipe
  const createRecipe = async (recipeData) => {
    try {
      const CREATE_RECIPE_MUTATION = gql`
        mutation CreateRecipe($recipe: recipes_insert_input!) {
          insert_recipes_one(object: $recipe) {
            id
            slug
            title
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: CREATE_RECIPE_MUTATION,
        variables: {
          recipe: recipeData
        }
      })

      return data.insert_recipes_one
    } catch (error) {
      console.error('Error creating recipe:', error)
      throw error
    }
  }

  // Like/Unlike recipe
  const toggleLike = async (recipeId) => {
    try {
      const TOGGLE_LIKE_MUTATION = gql`
        mutation ToggleLike($recipeId: uuid!) {
          toggleRecipeLike(recipeId: $recipeId) {
            success
            isLiked
            likesCount
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: TOGGLE_LIKE_MUTATION,
        variables: { recipeId }
      })

      return data.toggleRecipeLike
    } catch (error) {
      console.error('Error toggling like:', error)
      throw error
    }
  }

  // Bookmark/Unbookmark recipe
  const toggleBookmark = async (recipeId) => {
    try {
      const TOGGLE_BOOKMARK_MUTATION = gql`
        mutation ToggleBookmark($recipeId: uuid!) {
          toggleRecipeBookmark(recipeId: $recipeId) {
            success
            isBookmarked
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: TOGGLE_BOOKMARK_MUTATION,
        variables: { recipeId }
      })

      return data.toggleRecipeBookmark
    } catch (error) {
      console.error('Error toggling bookmark:', error)
      throw error
    }
  }

  return {
    recipes,
    recipe,
    categories,
    isLoading,
    totalRecipes,
    getRecipes,
    getRecipe,
    getCategories,
    createRecipe,
    toggleLike,
    toggleBookmark
  }
}