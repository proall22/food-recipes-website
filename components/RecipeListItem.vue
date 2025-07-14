<template>
  <div class="bg-white rounded-2xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300">
    <div class="flex flex-col md:flex-row">
      <!-- Recipe Image -->
      <div class="md:w-1/3 relative">
        <div class="aspect-[4/3] md:aspect-square overflow-hidden">
          <img :src="recipe.featuredImage" :alt="recipe.title" 
            class="w-full h-full object-cover hover:scale-105 transition-transform duration-300" />
        </div>
        
        <!-- Category Badge -->
        <div class="absolute top-4 left-4">
          <span class="bg-orange-500 text-white px-3 py-1 rounded-full text-sm font-medium">
            {{ recipe.category.name }}
          </span>
        </div>

        <!-- Price Badge -->
        <div v-if="recipe.price" class="absolute top-4 right-4">
          <span class="bg-green-500 text-white px-3 py-1 rounded-full text-sm font-medium">
            ${{ recipe.price }}
          </span>
        </div>
      </div>

      <!-- Recipe Content -->
      <div class="md:w-2/3 p-6 flex flex-col justify-between">
        <div>
          <!-- Author Info -->
          <div class="flex items-center mb-3">
            <img :src="recipe.author.avatar" :alt="recipe.author.name" 
              class="w-8 h-8 rounded-full object-cover" />
            <span class="ml-2 text-sm text-gray-600">
              {{ recipe.author.firstName }} {{ recipe.author.lastName }}
            </span>
            <div class="ml-auto flex items-center text-yellow-400">
              <Star class="h-4 w-4 fill-current" />
              <span class="ml-1 text-sm text-gray-600">{{ recipe.rating }}</span>
              <span class="ml-1 text-sm text-gray-500">({{ recipe.reviewCount }})</span>
            </div>
          </div>

          <!-- Recipe Title & Description -->
          <h3 class="text-xl font-bold text-gray-900 mb-2 cursor-pointer hover:text-orange-500 transition-colors"
            @click="navigateToRecipe">
            {{ recipe.title }}
          </h3>
          <p class="text-gray-600 text-sm mb-4 line-clamp-2">{{ recipe.description }}</p>

          <!-- Recipe Meta -->
          <div class="flex flex-wrap items-center gap-4 text-sm text-gray-500 mb-4">
            <div class="flex items-center">
              <Clock class="h-4 w-4 mr-1" />
              <span>{{ recipe.prepTime }} min</span>
            </div>
            <div class="flex items-center">
              <Users class="h-4 w-4 mr-1" />
              <span>{{ recipe.servings }} servings</span>
            </div>
            <div class="flex items-center">
              <span :class="[
                'px-2 py-1 rounded-full text-xs font-medium',
                recipe.difficulty === 'Easy' ? 'bg-green-100 text-green-800' :
                recipe.difficulty === 'Medium' ? 'bg-yellow-100 text-yellow-800' :
                'bg-red-100 text-red-800'
              ]">
                {{ recipe.difficulty }}
              </span>
            </div>
            <div class="flex items-center">
              <Heart class="h-4 w-4 mr-1" />
              <span>{{ recipe.likes }}</span>
            </div>
          </div>
        </div>

        <!-- Action Buttons -->
        <div class="flex items-center justify-between">
          <div class="flex space-x-2">
            <button @click.stop="toggleBookmark" 
              class="p-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
              <Bookmark :class="['h-4 w-4', isBookmarked ? 'fill-orange-500 text-orange-500' : 'text-gray-600']" />
            </button>
            <button @click.stop="toggleLike" 
              class="p-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors">
              <Heart :class="['h-4 w-4', isLiked ? 'fill-red-500 text-red-500' : 'text-gray-600']" />
            </button>
          </div>
          
          <button @click="navigateToRecipe" 
            class="bg-orange-500 text-white px-6 py-2 rounded-lg hover:bg-orange-600 transition-colors font-medium">
            View Recipe
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { Bookmark, Heart, Star, Clock, Users } from 'lucide-vue-next'
import { navigateTo } from '#app'

const props = defineProps({
  recipe: {
    type: Object,
    required: true
  }
})

const isBookmarked = ref(false)
const isLiked = ref(false)

const toggleBookmark = () => {
  isBookmarked.value = !isBookmarked.value
  // Implement bookmark API call
}

const toggleLike = () => {
  isLiked.value = !isLiked.value
  // Implement like API call
}

const navigateToRecipe = () => {
  navigateTo(`/recipes/${props.recipe.id}`)
}
</script>

<style scoped>
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>
