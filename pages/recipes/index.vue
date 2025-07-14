<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Search Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-4">Discover Recipes</h1>
      <p class="text-gray-600">Find the perfect recipe for any occasion</p>
    </div>

    <!-- Search and Filters -->
    <div class="bg-white rounded-2xl shadow-lg p-6 mb-8">
      <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <!-- Search Input -->
        <div class="lg:col-span-2">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Search Recipes
          </label>
          <div class="relative">
            <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
            <input
              v-model="filters.search"
              @input="debouncedSearch"
              type="text"
              class="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="Search by title, ingredients, or description..."
            />
          </div>
        </div>

        <!-- Category Filter -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Category
          </label>
          <select
            v-model="filters.category"
            @change="applyFilters"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
          >
            <option value="">All Categories</option>
            <option v-for="category in categories" :key="category.id" :value="category.id">
              {{ category.name }}
            </option>
          </select>
        </div>

        <!-- Difficulty Filter -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Difficulty
          </label>
          <select
            v-model="filters.difficulty"
            @change="applyFilters"
            class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
          >
            <option value="">All Levels</option>
            <option value="Easy">Easy</option>
            <option value="Medium">Medium</option>
            <option value="Hard">Hard</option>
          </select>
        </div>
      </div>

      <!-- Advanced Filters -->
      <div class="mt-6 pt-6 border-t border-gray-200">
        <button
          @click="showAdvancedFilters = !showAdvancedFilters"
          class="flex items-center text-orange-500 hover:text-orange-600 font-medium mb-4"
        >
          <Filter class="h-4 w-4 mr-2" />
          Advanced Filters
          <ChevronDown :class="['h-4 w-4 ml-2 transition-transform', showAdvancedFilters ? 'rotate-180' : '']" />
        </button>

        <div v-if="showAdvancedFilters" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
          <!-- Prep Time Filter -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Max Prep Time (minutes)
            </label>
            <select
              v-model="filters.maxPrepTime"
              @change="applyFilters"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="">Any Time</option>
              <option value="15">15 minutes</option>
              <option value="30">30 minutes</option>
              <option value="60">1 hour</option>
              <option value="120">2 hours</option>
            </select>
          </div>

          <!-- Cuisine Type -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Cuisine Type
            </label>
            <select
              v-model="filters.cuisineType"
              @change="applyFilters"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="">All Cuisines</option>
              <option value="Italian">Italian</option>
              <option value="Mexican">Mexican</option>
              <option value="Asian">Asian</option>
              <option value="Mediterranean">Mediterranean</option>
              <option value="American">American</option>
              <option value="French">French</option>
              <option value="Indian">Indian</option>
            </select>
          </div>

          <!-- Rating Filter -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Minimum Rating
            </label>
            <select
              v-model="filters.minRating"
              @change="applyFilters"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="">Any Rating</option>
              <option value="4">4+ Stars</option>
              <option value="3">3+ Stars</option>
              <option value="2">2+ Stars</option>
            </select>
          </div>

          <!-- Sort By -->
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Sort By
            </label>
            <select
              v-model="filters.sortBy"
              @change="applyFilters"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
            >
              <option value="createdAt_DESC">Newest First</option>
              <option value="createdAt_ASC">Oldest First</option>
              <option value="rating_DESC">Highest Rated</option>
              <option value="likes_DESC">Most Liked</option>
              <option value="prepTime_ASC">Quickest First</option>
              <option value="title_ASC">Alphabetical</option>
            </select>
          </div>
        </div>

        <!-- Ingredient Filter -->
        <div class="mt-6">
          <label class="block text-sm font-medium text-gray-700 mb-2">
            Filter by Ingredients
          </label>
          <div class="flex flex-wrap gap-2 mb-3">
            <span v-for="ingredient in filters.ingredients" :key="ingredient" 
              class="bg-orange-100 text-orange-800 px-3 py-1 rounded-full text-sm flex items-center">
              {{ ingredient }}
              <button @click="removeIngredientFilter(ingredient)" class="ml-2 hover:text-orange-600">
                <X class="h-3 w-3" />
              </button>
            </span>
          </div>
          <div class="flex space-x-2">
            <input
              v-model="newIngredient"
              @keyup.enter="addIngredientFilter"
              type="text"
              class="flex-1 px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              placeholder="Add ingredient (e.g., chicken, tomatoes)"
            />
            <button
              @click="addIngredientFilter"
              class="px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors"
            >
              Add
            </button>
          </div>
        </div>

        <!-- Clear Filters -->
        <div class="mt-6 flex justify-end">
          <button
            @click="clearFilters"
            class="px-4 py-2 text-gray-600 hover:text-gray-800 transition-colors"
          >
            Clear All Filters
          </button>
        </div>
      </div>
    </div>

    <!-- Results Header -->
    <div class="flex items-center justify-between mb-6">
      <div>
        <h2 class="text-xl font-semibold text-gray-900">
          {{ totalRecipes }} {{ totalRecipes === 1 ? 'Recipe' : 'Recipes' }} Found
        </h2>
        <p v-if="hasActiveFilters" class="text-sm text-gray-600 mt-1">
          Showing filtered results
        </p>
      </div>

      <!-- View Toggle -->
      <div class="flex items-center space-x-2">
        <button
          @click="viewMode = 'grid'"
          :class="[
            'p-2 rounded-lg transition-colors',
            viewMode === 'grid' ? 'bg-orange-500 text-white' : 'text-gray-600 hover:bg-gray-100'
          ]"
        >
          <Grid3X3 class="h-5 w-5" />
        </button>
        <button
          @click="viewMode = 'list'"
          :class="[
            'p-2 rounded-lg transition-colors',
            viewMode === 'list' ? 'bg-orange-500 text-white' : 'text-gray-600 hover:bg-gray-100'
          ]"
        >
          <List class="h-5 w-5" />
        </button>
      </div>
    </div>

    <!-- Loading State -->
    <div v-if="isLoading" class="flex justify-center py-12">
      <Loader2 class="h-8 w-8 animate-spin text-orange-500" />
    </div>

    <!-- No Results -->
    <div v-else-if="recipes.length === 0" class="text-center py-12">
      <ChefHat class="h-16 w-16 text-gray-400 mx-auto mb-4" />
      <h3 class="text-lg font-medium text-gray-900 mb-2">No recipes found</h3>
      <p class="text-gray-600 mb-4">Try adjusting your search criteria or filters</p>
      <button
        @click="clearFilters"
        class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600 transition-colors"
      >
        Clear Filters
      </button>
    </div>

    <!-- Recipe Results -->
    <div v-else>
      <!-- Grid View -->
      <div v-if="viewMode === 'grid'" class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <RecipeCard v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" />
      </div>

      <!-- List View -->
      <div v-else class="space-y-4">
        <RecipeListItem v-for="recipe in recipes" :key="recipe.id" :recipe="recipe" />
      </div>

      <!-- Pagination -->
      <div v-if="totalPages > 1" class="mt-12 flex justify-center">
        <nav class="flex items-center space-x-2">
          <button
            @click="goToPage(currentPage - 1)"
            :disabled="currentPage === 1"
            class="px-3 py-2 text-sm font-medium text-gray-500 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <ChevronLeft class="h-4 w-4" />
          </button>

          <button
            v-for="page in visiblePages"
            :key="page"
            @click="goToPage(page)"
            :class="[
              'px-3 py-2 text-sm font-medium rounded-lg transition-colors',
              page === currentPage
                ? 'bg-orange-500 text-white'
                : 'text-gray-700 hover:bg-gray-100'
            ]"
          >
            {{ page }}
          </button>

          <button
            @click="goToPage(currentPage + 1)"
            :disabled="currentPage === totalPages"
            class="px-3 py-2 text-sm font-medium text-gray-500 hover:text-gray-700 disabled:opacity-50 disabled:cursor-not-allowed"
          >
            <ChevronRight class="h-4 w-4" />
          </button>
        </nav>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, watch, onMounted } from 'vue'
import { 
  Search, Filter, ChevronDown, X, Grid3X3, List, 
  Loader2, ChefHat, ChevronLeft, ChevronRight 
} from 'lucide-vue-next'
import { useRoute, useRouter } from '#app'
import { useDebounceFn } from '@vueuse/core'
import RecipeCard from '~/components/RecipeCard.vue'
import RecipeListItem from '~/components/RecipeListItem.vue'
import { useNuxtApp } from '#app'
import { gql } from 'graphql-tag'
import { useHead } from '#head'

const route = useRoute()
const router = useRouter()

const isLoading = ref(false)
const showAdvancedFilters = ref(false)
const viewMode = ref('grid')
const newIngredient = ref('')

const recipes = ref([])
const categories = ref([])
const totalRecipes = ref(0)
const currentPage = ref(1)
const totalPages = ref(1)
const pageSize = 12

const filters = reactive({
  search: '',
  category: '',
  difficulty: '',
  maxPrepTime: '',
  cuisineType: '',
  minRating: '',
  sortBy: 'createdAt_DESC',
  ingredients: []
})

const hasActiveFilters = computed(() => {
  return Object.values(filters).some(value => 
    Array.isArray(value) ? value.length > 0 : value !== ''
  )
})

const visiblePages = computed(() => {
  const pages = []
  const start = Math.max(1, currentPage.value - 2)
  const end = Math.min(totalPages.value, start + 4)
  
  for (let i = start; i <= end; i++) {
    pages.push(i)
  }
  
  return pages
})

// Initialize filters from URL query params
onMounted(() => {
  const query = route.query
  
  if (query.q) filters.search = query.q
  if (query.category) filters.category = query.category
  if (query.difficulty) filters.difficulty = query.difficulty
  if (query.maxPrepTime) filters.maxPrepTime = query.maxPrepTime
  if (query.cuisineType) filters.cuisineType = query.cuisineType
  if (query.minRating) filters.minRating = query.minRating
  if (query.sortBy) filters.sortBy = query.sortBy
  if (query.ingredients) {
    filters.ingredients = Array.isArray(query.ingredients) 
      ? query.ingredients 
      : [query.ingredients]
  }
  if (query.page) currentPage.value = parseInt(query.page)
  
  loadCategories()
  loadRecipes()
})

// Watch for filter changes
watch(filters, () => {
  currentPage.value = 1
  updateURL()
}, { deep: true })

watch(currentPage, () => {
  updateURL()
  loadRecipes()
})

const loadCategories = async () => {
  try {
    const { $apollo } = useNuxtApp()
    
    const GET_CATEGORIES_QUERY = gql`
      query GetCategories {
        categories {
          id
          name
          slug
        }
      }
    `
    
    const { data } = await $apollo.query({
      query: GET_CATEGORIES_QUERY
    })
    
    categories.value = data.categories
  } catch (error) {
    console.error('Failed to load categories:', error)
  }
}

const loadRecipes = async () => {
  isLoading.value = true
  
  try {
    const { $apollo } = useNuxtApp()
    
    const SEARCH_RECIPES_QUERY = gql`
      query SearchRecipes($input: RecipeSearchInput!) {
        searchRecipes(input: $input) {
          recipes {
            id
            title
            description
            featuredImage
            prepTime
            servings
            difficulty
            rating
            reviewCount
            likes
            price
            category {
              id
              name
            }
            author {
              id
              firstName
              lastName
              username
              avatar
            }
            createdAt
          }
          totalCount
          totalPages
        }
      }
    `
    
    const searchInput = {
      search: filters.search || undefined,
      categoryId: filters.category || undefined,
      difficulty: filters.difficulty || undefined,
      maxPrepTime: filters.maxPrepTime ? parseInt(filters.maxPrepTime) : undefined,
      cuisineType: filters.cuisineType || undefined,
      minRating: filters.minRating ? parseFloat(filters.minRating) : undefined,
      ingredients: filters.ingredients.length > 0 ? filters.ingredients : undefined,
      sortBy: filters.sortBy,
      page: currentPage.value,
      limit: pageSize
    }
    
    const { data } = await $apollo.query({
      query: SEARCH_RECIPES_QUERY,
      variables: { input: searchInput },
      fetchPolicy: 'cache-and-network'
    })
    
    recipes.value = data.searchRecipes.recipes
    totalRecipes.value = data.searchRecipes.totalCount
    totalPages.value = data.searchRecipes.totalPages
    
  } catch (error) {
    console.error('Failed to load recipes:', error)
  } finally {
    isLoading.value = false
  }
}

const debouncedSearch = useDebounceFn(() => {
  applyFilters()
}, 500)

const applyFilters = () => {
  currentPage.value = 1
  loadRecipes()
}

const addIngredientFilter = () => {
  if (newIngredient.value.trim() && !filters.ingredients.includes(newIngredient.value.trim())) {
    filters.ingredients.push(newIngredient.value.trim())
    newIngredient.value = ''
    applyFilters()
  }
}

const removeIngredientFilter = (ingredient) => {
  const index = filters.ingredients.indexOf(ingredient)
  if (index > -1) {
    filters.ingredients.splice(index, 1)
    applyFilters()
  }
}

const clearFilters = () => {
  Object.keys(filters).forEach(key => {
    if (Array.isArray(filters[key])) {
      filters[key] = []
    } else {
      filters[key] = ''
    }
  })
  filters.sortBy = 'createdAt_DESC'
  currentPage.value = 1
  applyFilters()
}

const goToPage = (page) => {
  if (page >= 1 && page <= totalPages.value) {
    currentPage.value = page
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }
}

const updateURL = () => {
  const query = {}
  
  if (filters.search) query.q = filters.search
  if (filters.category) query.category = filters.category
  if (filters.difficulty) query.difficulty = filters.difficulty
  if (filters.maxPrepTime) query.maxPrepTime = filters.maxPrepTime
  if (filters.cuisineType) query.cuisineType = filters.cuisineType
  if (filters.minRating) query.minRating = filters.minRating
  if (filters.sortBy !== 'createdAt_DESC') query.sortBy = filters.sortBy
  if (filters.ingredients.length > 0) query.ingredients = filters.ingredients
  if (currentPage.value > 1) query.page = currentPage.value
  
  router.replace({ query })
}

// SEO
useHead({
  title: 'Browse Recipes - RecipeHub',
  meta: [
    { name: 'description', content: 'Discover thousands of delicious recipes. Search by ingredients, category, cooking time, and more.' }
  ]
})
</script>
