<template>
  <div class="bg-white rounded-2xl shadow-lg overflow-hidden hover:shadow-xl transition-shadow duration-300 group">
    <!-- Recipe Image -->
    <div class="relative overflow-hidden aspect-[4/3]">
      <img 
        :src="recipe.featured_image || '/placeholder.svg?height=300&width=400'" 
        :alt="recipe.title" 
        class="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300" 
      />
      
      <!-- Overlay Actions -->
      <div class="absolute top-4 right-4 flex space-x-2 opacity-0 group-hover:opacity-100 transition-opacity">
        <button 
          @click.stop="toggleBookmark" 
          class="p-2 bg-white/90 rounded-full hover:bg-white transition-colors"
        >
          <BookmarkIcon :class="['h-4 w-4', isBookmarked ? 'fill-orange-500 text-orange-500' : 'text-gray-600']" />
        </button>
        <button 
          @click.stop="toggleLike" 
          class="p-2 bg-white/90 rounded-full hover:bg-white transition-colors"
        >
          <HeartIcon :class="['h-4 w-4', isLiked ? 'fill-red-500 text-red-500' : 'text-gray-600']" />
        </button>
      </div>

      <!-- Category Badge -->
      <div class="absolute top-4 left-4">
        <span class="bg-orange-500 text-white px-3 py-1 rounded-full text-sm font-medium">
          {{ recipe.category?.name }}
        </span>
      </div>

      <!-- Difficulty Badge -->
      <div class="absolute bottom-4 left-4">
        <span :class="[
          'px-3 py-1 rounded-full text-sm font-medium',
          recipe.difficulty === 'Easy' ? 'bg-green-100 text-green-800' :
          recipe.difficulty === 'Medium' ? 'bg-yellow-100 text-yellow-800' :
          'bg-red-100 text-red-800'
        ]">
          {{ recipe.difficulty }}
        </span>
      </div>

      <!-- Price Badge -->
      <div v-if="recipe.price > 0" class="absolute bottom-4 right-4">
        <span class="bg-green-500 text-white px-3 py-1 rounded-full text-sm font-medium">
          ${{ recipe.price }}
        </span>
      </div>
    </div>

    <!-- Recipe Content -->
    <div class="p-6">
      <!-- Author Info -->
      <div class="flex items-center mb-3">
        <img 
          :src="recipe.author?.avatar || '/placeholder-user.jpg'" 
          :alt="recipe.author?.first_name" 
          class="w-8 h-8 rounded-full object-cover" 
        />
        <span class="ml-2 text-sm text-gray-600">
          {{ recipe.author?.first_name }} {{ recipe.author?.last_name }}
        </span>
        <div class="ml-auto flex items-center text-yellow-400">
          <StarIcon class="h-4 w-4 fill-current" />
          <span class="ml-1 text-sm text-gray-600">{{ recipe.average_rating || 0 }}</span>
        </div>
      </div>

      <!-- Recipe Title & Description -->
      <h3 
        class="text-xl font-bold text-gray-900 mb-2 line-clamp-2 cursor-pointer hover:text-orange-500 transition-colors"
        @click="navigateToRecipe"
      >
        {{ recipe.title }}
      </h3>
      <p class="text-gray-600 text-sm mb-4 line-clamp-2">{{ recipe.description }}</p>

      <!-- Recipe Meta -->
      <div class="flex items-center justify-between text-sm text-gray-500 mb-4">
        <div class="flex items-center">
          <ClockIcon class="h-4 w-4 mr-1" />
          <span>{{ recipe.prep_time }} min</span>
        </div>
        <div class="flex items-center">
          <UsersIcon class="h-4 w-4 mr-1" />
          <span>{{ recipe.servings }}</span>
        </div>
        <div class="flex items-center">
          <HeartIcon class="h-4 w-4 mr-1" />
          <span>{{ recipe.likes_count || 0 }}</span>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex space-x-2">
        <button 
          @click="navigateToRecipe" 
          class="flex-1 bg-orange-500 text-white py-2 px-4 rounded-lg hover:bg-orange-600 transition-colors font-medium"
        >
          View Recipe
        </button>
        <button 
          @click.stop="quickSave" 
          class="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50 transition-colors"
        >
          <BookmarkPlusIcon class="h-4 w-4" />
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  BookmarkIcon, 
  HeartIcon, 
  StarIcon, 
  ClockIcon, 
  UsersIcon,
  BookmarkPlusIcon 
} from '@heroicons/vue/24/outline'

const props = defineProps({
  recipe: {
    type: Object,
    required: true
  }
})

const isBookmarked = ref(false)
const isLiked = ref(false)

const toggleBookmark = async () => {
  try {
    const { toggleBookmark } = useRecipes()
    const result = await toggleBookmark(props.recipe.id)
    isBookmarked.value = result.isBookmarked
  } catch (error) {
    console.error('Error toggling bookmark:', error)
  }
}

const toggleLike = async () => {
  try {
    const { toggleLike } = useRecipes()
    const result = await toggleLike(props.recipe.id)
    isLiked.value = result.isLiked
  } catch (error) {
    console.error('Error toggling like:', error)
  }
}

const navigateToRecipe = () => {
  navigateTo(`/recipes/${props.recipe.id}`)
}

const quickSave = () => {
  toggleBookmark()
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