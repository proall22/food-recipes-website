<template>
  <div v-if="recipe" class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Recipe Header -->
    <div class="mb-8">
      <div class="flex items-center text-sm text-gray-500 mb-4">
        <NuxtLink to="/" class="hover:text-orange-500">Home</NuxtLink>
        <ChevronRight class="h-4 w-4 mx-2" />
        <NuxtLink to="/recipes" class="hover:text-orange-500">Recipes</NuxtLink>
        <ChevronRight class="h-4 w-4 mx-2" />
        <span>{{ recipe.title }}</span>
      </div>

      <div class="flex flex-col lg:flex-row lg:items-center lg:justify-between mb-6">
        <div class="mb-4 lg:mb-0">
          <h1 class="text-3xl lg:text-4xl font-bold text-gray-900 mb-2">{{ recipe.title }}</h1>
          <p class="text-gray-600 text-lg">{{ recipe.description }}</p>
        </div>
        
        <div class="flex items-center space-x-4">
          <button @click="toggleBookmark" 
            class="flex items-center space-x-2 px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
            <Bookmark :class="['h-5 w-5', isBookmarked ? 'fill-orange-500 text-orange-500' : 'text-gray-600']" />
            <span>Save</span>
          </button>
          <button @click="toggleLike" 
            class="flex items-center space-x-2 px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50">
            <Heart :class="['h-5 w-5', isLiked ? 'fill-red-500 text-red-500' : 'text-gray-600']" />
            <span>{{ recipe.likes }}</span>
          </button>
          <button class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors">
            Buy Recipe (${{ recipe.price }})
          </button>
        </div>
      </div>

      <!-- Author & Meta Info -->
      <div class="flex flex-wrap items-center gap-6 p-4 bg-gray-50 rounded-lg">
        <div class="flex items-center">
          <img :src="recipe.author.avatar" :alt="recipe.author.name" 
            class="w-12 h-12 rounded-full object-cover" />
          <div class="ml-3">
            <p class="font-medium text-gray-900">{{ recipe.author.name }}</p>
            <p class="text-sm text-gray-500">{{ recipe.author.recipeCount }} recipes</p>
          </div>
        </div>
        
        <div class="flex items-center text-yellow-400">
          <Star class="h-5 w-5 fill-current" />
          <span class="ml-1 font-medium text-gray-900">{{ recipe.rating }}</span>
          <span class="ml-1 text-gray-500">({{ recipe.reviewCount }} reviews)</span>
        </div>
        
        <div class="flex items-center text-gray-600">
          <Clock class="h-5 w-5 mr-2" />
          <span>{{ recipe.prepTime }} minutes</span>
        </div>
        
        <div class="flex items-center text-gray-600">
          <Users class="h-5 w-5 mr-2" />
          <span>Serves {{ recipe.servings }}</span>
        </div>
      </div>
    </div>

    <!-- Recipe Image Gallery -->
    <div class="mb-8">
      <div class="relative aspect-[16/9] rounded-2xl overflow-hidden mb-4">
        <img :src="selectedImage" :alt="recipe.title" 
          class="w-full h-full object-cover" />
      </div>
      
      <div v-if="recipe.images.length > 1" class="flex space-x-2 overflow-x-auto">
        <button v-for="(image, index) in recipe.images" :key="index"
          @click="selectedImage = image"
          :class="[
            'flex-shrink-0 w-20 h-20 rounded-lg overflow-hidden border-2 transition-colors',
            selectedImage === image ? 'border-orange-500' : 'border-gray-200'
          ]">
          <img :src="image" :alt="`Recipe image ${index + 1}`" 
            class="w-full h-full object-cover" />
        </button>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Ingredients -->
      <div class="lg:col-span-1">
        <div class="bg-white rounded-2xl shadow-lg p-6 sticky top-24">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Ingredients</h2>
          
          <div class="mb-4">
            <label class="text-sm font-medium text-gray-700 mb-2 block">Servings:</label>
            <div class="flex items-center space-x-3">
              <button @click="adjustServings(-1)" 
                class="w-8 h-8 rounded-full border border-gray-300 flex items-center justify-center hover:bg-gray-50">
                <Minus class="h-4 w-4" />
              </button>
              <span class="font-medium">{{ currentServings }}</span>
              <button @click="adjustServings(1)" 
                class="w-8 h-8 rounded-full border border-gray-300 flex items-center justify-center hover:bg-gray-50">
                <Plus class="h-4 w-4" />
              </button>
            </div>
          </div>

          <ul class="space-y-3">
            <li v-for="ingredient in adjustedIngredients" :key="ingredient.id" 
              class="flex items-start space-x-3">
              <input type="checkbox" class="mt-1 rounded border-gray-300 text-orange-500 focus:ring-orange-500" />
              <span class="text-gray-700">
                <span class="font-medium">{{ ingredient.amount }}</span>
                {{ ingredient.unit }} {{ ingredient.name }}
              </span>
            </li>
          </ul>

          <button class="w-full mt-6 bg-orange-500 text-white py-3 rounded-lg hover:bg-orange-600 transition-colors font-medium">
            Add to Shopping List
          </button>
        </div>
      </div>

      <!-- Instructions -->
      <div class="lg:col-span-2">
        <div class="bg-white rounded-2xl shadow-lg p-6">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Instructions</h2>
          
          <div class="space-y-6">
            <div v-for="(step, index) in recipe.steps" :key="step.id" 
              class="flex space-x-4">
              <div class="flex-shrink-0">
                <div class="w-8 h-8 bg-orange-500 text-white rounded-full flex items-center justify-center font-bold text-sm">
                  {{ index + 1 }}
                </div>
              </div>
              <div class="flex-1">
                <p class="text-gray-700 leading-relaxed">{{ step.instruction }}</p>
                <div v-if="step.image" class="mt-3">
                  <img :src="step.image" :alt="`Step ${index + 1}`" 
                    class="w-full max-w-md rounded-lg" />
                </div>
                <div v-if="step.timer" class="mt-2 flex items-center text-sm text-gray-500">
                  <Timer class="h-4 w-4 mr-1" />
                  <span>{{ step.timer }} minutes</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Nutrition Info -->
        <div class="bg-white rounded-2xl shadow-lg p-6 mt-8">
          <h2 class="text-2xl font-bold text-gray-900 mb-6">Nutrition Information</h2>
          <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
            <div class="text-center p-4 bg-gray-50 rounded-lg">
              <p class="text-2xl font-bold text-orange-500">{{ recipe.nutrition.calories }}</p>
              <p class="text-sm text-gray-600">Calories</p>
            </div>
            <div class="text-center p-4 bg-gray-50 rounded-lg">
              <p class="text-2xl font-bold text-orange-500">{{ recipe.nutrition.protein }}g</p>
              <p class="text-sm text-gray-600">Protein</p>
            </div>
            <div class="text-center p-4 bg-gray-50 rounded-lg">
              <p class="text-2xl font-bold text-orange-500">{{ recipe.nutrition.carbs }}g</p>
              <p class="text-sm text-gray-600">Carbs</p>
            </div>
            <div class="text-center p-4 bg-gray-50 rounded-lg">
              <p class="text-2xl font-bold text-orange-500">{{ recipe.nutrition.fat }}g</p>
              <p class="text-sm text-gray-600">Fat</p>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Reviews Section -->
    <div class="mt-12">
      <div class="bg-white rounded-2xl shadow-lg p-6">
        <div class="flex items-center justify-between mb-6">
          <h2 class="text-2xl font-bold text-gray-900">Reviews & Ratings</h2>
          <button class="bg-orange-500 text-white px-4 py-2 rounded-lg hover:bg-orange-600 transition-colors">
            Write Review
          </button>
        </div>

        <!-- Rating Summary -->
        <div class="flex items-center space-x-8 mb-8 p-4 bg-gray-50 rounded-lg">
          <div class="text-center">
            <div class="text-4xl font-bold text-orange-500">{{ recipe.rating }}</div>
            <div class="flex items-center justify-center mt-1">
              <Star v-for="i in 5" :key="i" 
                :class="['h-4 w-4', i <= Math.floor(recipe.rating) ? 'text-yellow-400 fill-current' : 'text-gray-300']" />
            </div>
            <div class="text-sm text-gray-600 mt-1">{{ recipe.reviewCount }} reviews</div>
          </div>
          
          <div class="flex-1">
            <div v-for="rating in [5, 4, 3, 2, 1]" :key="rating" class="flex items-center mb-1">
              <span class="text-sm w-8">{{ rating }}</span>
              <Star class="h-3 w-3 text-yellow-400 fill-current mr-2" />
              <div class="flex-1 bg-gray-200 rounded-full h-2 mr-2">
                <div class="bg-yellow-400 h-2 rounded-full" 
                  :style="{ width: `${(recipe.ratingBreakdown[rating] / recipe.reviewCount) * 100}%` }"></div>
              </div>
              <span class="text-sm text-gray-600 w-8">{{ recipe.ratingBreakdown[rating] }}</span>
            </div>
          </div>
        </div>

        <!-- Individual Reviews -->
        <div class="space-y-6">
          <div v-for="review in recipe.reviews" :key="review.id" 
            class="border-b border-gray-200 pb-6 last:border-b-0">
            <div class="flex items-start space-x-4">
              <img :src="review.user.avatar" :alt="review.user.name" 
                class="w-10 h-10 rounded-full object-cover" />
              <div class="flex-1">
                <div class="flex items-center space-x-2 mb-2">
                  <h4 class="font-medium text-gray-900">{{ review.user.name }}</h4>
                  <div class="flex items-center">
                    <Star v-for="i in 5" :key="i" 
                      :class="['h-3 w-3', i <= review.rating ? 'text-yellow-400 fill-current' : 'text-gray-300']" />
                  </div>
                  <span class="text-sm text-gray-500">{{ formatDate(review.createdAt) }}</span>
                </div>
                <p class="text-gray-700">{{ review.comment }}</p>
                <div v-if="review.images?.length" class="flex space-x-2 mt-3">
                  <img v-for="image in review.images" :key="image" 
                    :src="image" alt="Review image" 
                    class="w-16 h-16 rounded-lg object-cover" />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Related Recipes -->
    <div class="mt-12">
      <h2 class="text-2xl font-bold text-gray-900 mb-6">You Might Also Like</h2>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <RecipeCard v-for="relatedRecipe in relatedRecipes" :key="relatedRecipe.id" :recipe="relatedRecipe" />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { 
  ChevronRight, Bookmark, Heart, Star, Clock, Users, 
  Minus, Plus, Timer 
} from 'lucide-vue-next'
import { useRoute } from '#app'
import { useHead } from '#head'
import RecipeCard from '~/components/RecipeCard.vue'

const route = useRoute()
const recipeId = route.params.id

const recipe = ref(null)
const selectedImage = ref('')
const currentServings = ref(4)
const isBookmarked = ref(false)
const isLiked = ref(false)
const relatedRecipes = ref([])

// Sample recipe data - will be replaced with GraphQL query
const sampleRecipe = {
  id: 1,
  title: 'Creamy Mushroom Risotto',
  description: 'A rich and creamy Italian classic with wild mushrooms and parmesan cheese',
  images: [
    '/placeholder.svg?height=400&width=600',
    '/placeholder.svg?height=400&width=600',
    '/placeholder.svg?height=400&width=600'
  ],
  prepTime: 45,
  servings: 4,
  difficulty: 'Medium',
  rating: 4.8,
  reviewCount: 127,
  likes: 234,
  price: 4.99,
  category: 'Dinner',
  author: {
    name: 'Maria Rodriguez',
    avatar: '/placeholder.svg?height=48&width=48',
    recipeCount: 45
  },
  ingredients: [
    { id: 1, name: 'Arborio rice', amount: 1.5, unit: 'cups' },
    { id: 2, name: 'Mixed mushrooms', amount: 300, unit: 'g' },
    { id: 3, name: 'Vegetable broth', amount: 4, unit: 'cups' },
    { id: 4, name: 'White wine', amount: 0.5, unit: 'cup' },
    { id: 5, name: 'Parmesan cheese', amount: 100, unit: 'g' },
    { id: 6, name: 'Butter', amount: 3, unit: 'tbsp' },
    { id: 7, name: 'Onion', amount: 1, unit: 'medium' },
    { id: 8, name: 'Garlic cloves', amount: 3, unit: 'cloves' }
  ],
  steps: [
    {
      id: 1,
      instruction: 'Heat the vegetable broth in a saucepan and keep it warm over low heat.',
      timer: null,
      image: null
    },
    {
      id: 2,
      instruction: 'In a large pan, melt 2 tablespoons of butter over medium heat. Add the sliced mushrooms and cook until golden brown, about 5-7 minutes. Season with salt and pepper, then set aside.',
      timer: 7,
      image: '/placeholder.svg?height=200&width=300'
    },
    {
      id: 3,
      instruction: 'In the same pan, add the remaining butter. Add the diced onion and cook until translucent, about 3-4 minutes. Add the minced garlic and cook for another minute.',
      timer: 4,
      image: null
    },
    {
      id: 4,
      instruction: 'Add the Arborio rice to the pan and stir to coat with the butter. Cook for 1-2 minutes until the rice is lightly toasted.',
      timer: 2,
      image: null
    },
    {
      id: 5,
      instruction: 'Pour in the white wine and stir until it\'s mostly absorbed by the rice.',
      timer: null,
      image: null
    },
    {
      id: 6,
      instruction: 'Begin adding the warm broth one ladle at a time, stirring constantly. Wait until each addition is almost completely absorbed before adding the next. This process should take about 18-20 minutes.',
      timer: 20,
      image: '/placeholder.svg?height=200&width=300'
    },
    {
      id: 7,
      instruction: 'When the rice is creamy and tender, stir in the cooked mushrooms and grated Parmesan cheese. Season with salt and pepper to taste.',
      timer: null,
      image: null
    },
    {
      id: 8,
      instruction: 'Serve immediately, garnished with additional Parmesan cheese and fresh herbs if desired.',
      timer: null,
      image: '/placeholder.svg?height=200&width=300'
    }
  ],
  nutrition: {
    calories: 385,
    protein: 12,
    carbs: 58,
    fat: 14
  },
  ratingBreakdown: {
    5: 89,
    4: 28,
    3: 7,
    2: 2,
    1: 1
  },
  reviews: [
    {
      id: 1,
      user: { name: 'John Smith', avatar: '/placeholder.svg?height=40&width=40' },
      rating: 5,
      comment: 'Absolutely delicious! The mushrooms were perfectly cooked and the risotto was so creamy. Will definitely make this again.',
      createdAt: '2024-01-15',
      images: ['/placeholder.svg?height=64&width=64']
    },
    {
      id: 2,
      user: { name: 'Emma Wilson', avatar: '/placeholder.svg?height=40&width=40' },
      rating: 4,
      comment: 'Great recipe! Took a bit longer than expected but the result was worth it. My family loved it.',
      createdAt: '2024-01-10',
      images: []
    }
  ]
}

const adjustedIngredients = computed(() => {
  const ratio = currentServings.value / recipe.value.servings
  return recipe.value.ingredients.map(ingredient => ({
    ...ingredient,
    amount: (ingredient.amount * ratio).toFixed(1)
  }))
})

const adjustServings = (change) => {
  const newServings = currentServings.value + change
  if (newServings > 0 && newServings <= 20) {
    currentServings.value = newServings
  }
}

const toggleBookmark = () => {
  isBookmarked.value = !isBookmarked.value
}

const toggleLike = () => {
  isLiked.value = !isLiked.value
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

onMounted(() => {
  // Load recipe data - replace with GraphQL query
  recipe.value = sampleRecipe
  selectedImage.value = recipe.value.images[0]
  currentServings.value = recipe.value.servings
  
  // Load related recipes
  relatedRecipes.value = [
    {
      id: 2,
      title: 'Truffle Pasta',
      description: 'Luxurious pasta with truffle oil',
      image: '/placeholder.svg?height=300&width=400',
      prepTime: 25,
      difficulty: 'Easy',
      rating: 4.7,
      author: { name: 'Chef Antonio', avatar: '/placeholder.svg?height=40&width=40' },
      likes: 156,
      category: 'Dinner'
    }
  ]
})

// SEO
useHead({
  title: computed(() => recipe.value ? `${recipe.value.title} - RecipeHub` : 'Recipe - RecipeHub'),
  meta: [
    { name: 'description', content: computed(() => recipe.value?.description || 'Delicious recipe on RecipeHub') }
  ]
})
</script>
