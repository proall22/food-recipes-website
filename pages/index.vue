<template>
  <div>
    <!-- Hero Section -->
    <section class="relative bg-gradient-to-r from-orange-500 to-red-500 text-white">
      <div class="absolute inset-0 bg-black opacity-20"></div>
      <div class="relative max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-24">
        <div class="text-center">
          <h1 class="text-4xl md:text-6xl font-bold mb-6">
            Discover Amazing Recipes
          </h1>
          <p class="text-xl md:text-2xl mb-8 max-w-3xl mx-auto">
            Share your culinary creations and explore thousands of delicious recipes from home cooks around the world
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <NuxtLink to="/recipes" 
              class="bg-white text-orange-500 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition-colors">
              Browse Recipes
            </NuxtLink>
            <NuxtLink to="/create-recipe" 
              class="border-2 border-white text-white px-8 py-3 rounded-full font-semibold hover:bg-white hover:text-orange-500 transition-colors">
              Share Your Recipe
            </NuxtLink>
          </div>
        </div>
      </div>
    </section>

    <!-- Categories Section -->
    <section class="py-16 bg-white">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold text-gray-900 mb-4">Browse by Category</h2>
          <p class="text-gray-600 max-w-2xl mx-auto">
            Find the perfect recipe for any occasion, from quick breakfast ideas to elaborate dinner parties
          </p>
        </div>
        
        <div class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-6">
          <div v-for="category in categories" :key="category.id" 
            class="group cursor-pointer" @click="navigateToCategory(category.slug)">
            <div class="relative overflow-hidden rounded-2xl aspect-square mb-3 group-hover:scale-105 transition-transform duration-300">
              <img :src="category.image || '/placeholder.svg?height=200&width=200'" :alt="category.name" 
                class="w-full h-full object-cover" />
              <div class="absolute inset-0 bg-gradient-to-t from-black/50 to-transparent"></div>
              <div class="absolute bottom-3 left-3 text-white">
                <h3 class="font-semibold">{{ category.name }}</h3>
                <p class="text-sm opacity-90">{{ category.count }} recipes</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Recipes -->
    <section class="py-16 bg-gray-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center mb-12">
          <div>
            <h2 class="text-3xl font-bold text-gray-900 mb-4">Featured Recipes</h2>
            <p class="text-gray-600">Hand-picked recipes from our community</p>
          </div>
          <NuxtLink to="/recipes" 
            class="text-orange-500 hover:text-orange-600 font-medium flex items-center">
            View All 
            <ChevronRightIcon class="ml-1 h-4 w-4" />
          </NuxtLink>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          <RecipeCard v-for="recipe in featuredRecipes" :key="recipe.id" :recipe="recipe" />
        </div>
      </div>
    </section>

    <!-- Newsletter Section -->
    <section class="py-16 bg-gradient-to-r from-orange-500 to-red-500">
      <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
        <h2 class="text-3xl font-bold text-white mb-4">Stay Updated</h2>
        <p class="text-white/90 mb-8 text-lg">
          Get the latest recipes and cooking tips delivered to your inbox
        </p>
        <form @submit.prevent="subscribeNewsletter" class="flex flex-col sm:flex-row gap-4 max-w-md mx-auto">
          <input
            v-model="email"
            type="email"
            placeholder="Enter your email"
            class="flex-1 px-4 py-3 rounded-full focus:ring-2 focus:ring-white focus:outline-none"
            required
          />
          <button type="submit" 
            class="bg-white text-orange-500 px-8 py-3 rounded-full font-semibold hover:bg-gray-100 transition-colors">
            Subscribe
          </button>
        </form>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { ChevronRightIcon } from '@heroicons/vue/24/outline'

const { getCategories, getRecipes } = useRecipes()

const email = ref('')
const categories = ref([])
const featuredRecipes = ref([])

const navigateToCategory = (slug) => {
  navigateTo(`/categories/${slug}`)
}

const subscribeNewsletter = () => {
  // Implement newsletter subscription
  console.log('Subscribing email:', email.value)
  email.value = ''
}

onMounted(async () => {
  try {
    // Load categories
    categories.value = await getCategories()
    
    // Load featured recipes
    featuredRecipes.value = await getRecipes({ limit: 6 })
  } catch (error) {
    console.error('Error loading data:', error)
  }
})

// SEO
useHead({
  title: 'RecipeHub - Discover Amazing Recipes',
  meta: [
    { name: 'description', content: 'Discover and share amazing recipes from around the world. Browse thousands of recipes by category, creator, and cooking time.' }
  ]
})
</script>