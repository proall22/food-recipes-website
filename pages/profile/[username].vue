<template>
  <div v-if="userProfile" class="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Profile Header -->
    <div class="bg-white rounded-2xl shadow-lg overflow-hidden mb-8">
      <div class="relative h-48 bg-gradient-to-r from-orange-500 to-red-500">
        <div class="absolute inset-0 bg-black opacity-20"></div>
      </div>
      
      <div class="relative px-6 pb-6">
        <!-- Profile Picture -->
        <div class="flex flex-col sm:flex-row sm:items-end sm:space-x-6 -mt-16">
          <div class="relative">
            <img :src="userProfile.avatar" :alt="userProfile.name" 
              class="w-32 h-32 rounded-full border-4 border-white object-cover" />
            <div v-if="userProfile.isVerified" 
              class="absolute bottom-2 right-2 bg-blue-500 text-white rounded-full p-2">
              <CheckCircle class="h-4 w-4" />
            </div>
          </div>
          
          <div class="mt-4 sm:mt-0 flex-1">
            <div class="flex flex-col sm:flex-row sm:items-center sm:justify-between">
              <div>
                <h1 class="text-2xl font-bold text-gray-900">
                  {{ userProfile.firstName }} {{ userProfile.lastName }}
                </h1>
                <p class="text-gray-600">@{{ userProfile.username }}</p>
                <div class="flex items-center mt-2 text-yellow-400">
                  <Star class="h-5 w-5 fill-current" />
                  <span class="ml-1 font-medium text-gray-900">{{ userProfile.averageRating }}</span>
                  <span class="ml-1 text-gray-600">({{ userProfile.totalReviews }} reviews)</span>
                </div>
              </div>
              
              <div class="mt-4 sm:mt-0 flex space-x-3">
                <button v-if="!isOwnProfile" @click="toggleFollow" 
                  :class="[
                    'px-6 py-2 rounded-lg font-medium transition-colors',
                    isFollowing 
                      ? 'bg-gray-200 text-gray-800 hover:bg-gray-300' 
                      : 'bg-orange-500 text-white hover:bg-orange-600'
                  ]">
                  {{ isFollowing ? 'Following' : 'Follow' }}
                </button>
                
                <button v-if="isOwnProfile" @click="editProfile" 
                  class="px-6 py-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors">
                  Edit Profile
                </button>
                
                <button class="p-2 border border-gray-300 rounded-lg text-gray-700 hover:bg-gray-50 transition-colors">
                  <Share2 class="h-5 w-5" />
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- Bio -->
        <div class="mt-6">
          <p class="text-gray-700">{{ userProfile.bio }}</p>
        </div>

        <!-- Stats -->
        <div class="mt-6 grid grid-cols-3 gap-6 text-center">
          <div>
            <div class="text-2xl font-bold text-gray-900">{{ userProfile.recipeCount }}</div>
            <div class="text-sm text-gray-600">Recipes</div>
          </div>
          <div>
            <div class="text-2xl font-bold text-gray-900">{{ userProfile.followersCount }}</div>
            <div class="text-sm text-gray-600">Followers</div>
          </div>
          <div>
            <div class="text-2xl font-bold text-gray-900">{{ userProfile.followingCount }}</div>
            <div class="text-sm text-gray-600">Following</div>
          </div>
        </div>

        <!-- Achievements -->
        <div v-if="userProfile.achievements?.length" class="mt-6">
          <h3 class="text-lg font-semibold text-gray-900 mb-3">Achievements</h3>
          <div class="flex flex-wrap gap-2">
            <span v-for="achievement in userProfile.achievements" :key="achievement.id"
              class="bg-yellow-100 text-yellow-800 px-3 py-1 rounded-full text-sm font-medium flex items-center">
              <Trophy class="h-4 w-4 mr-1" />
              {{ achievement.name }}
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Content Tabs -->
    <div class="bg-white rounded-2xl shadow-lg overflow-hidden">
      <div class="border-b border-gray-200">
        <nav class="flex space-x-8 px-6">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="[
              'py-4 px-1 border-b-2 font-medium text-sm transition-colors',
              activeTab === tab.id
                ? 'border-orange-500 text-orange-600'
                : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300'
            ]"
          >
            {{ tab.name }}
            <span class="ml-2 bg-gray-100 text-gray-600 py-1 px-2 rounded-full text-xs">
              {{ tab.count }}
            </span>
          </button>
        </nav>
      </div>

      <div class="p-6">
        <!-- Recipes Tab -->
        <div v-if="activeTab === 'recipes'">
          <div v-if="userRecipes.length === 0" class="text-center py-12">
            <ChefHat class="h-16 w-16 text-gray-400 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-gray-900 mb-2">No recipes yet</h3>
            <p class="text-gray-600">
              {{ isOwnProfile ? "You haven't shared any recipes yet." : "This user hasn't shared any recipes yet." }}
            </p>
          </div>
          
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <RecipeCard v-for="recipe in userRecipes" :key="recipe.id" :recipe="recipe" />
          </div>
        </div>

        <!-- Liked Recipes Tab -->
        <div v-if="activeTab === 'liked'">
          <div v-if="likedRecipes.length === 0" class="text-center py-12">
            <Heart class="h-16 w-16 text-gray-400 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-gray-900 mb-2">No liked recipes</h3>
            <p class="text-gray-600">
              {{ isOwnProfile ? "You haven't liked any recipes yet." : "This user hasn't liked any recipes yet." }}
            </p>
          </div>
          
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <RecipeCard v-for="recipe in likedRecipes" :key="recipe.id" :recipe="recipe" />
          </div>
        </div>

        <!-- Reviews Tab -->
        <div v-if="activeTab === 'reviews'">
          <div v-if="userReviews.length === 0" class="text-center py-12">
            <MessageSquare class="h-16 w-16 text-gray-400 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-gray-900 mb-2">No reviews yet</h3>
            <p class="text-gray-600">
              {{ isOwnProfile ? "You haven't written any reviews yet." : "This user hasn't written any reviews yet." }}
            </p>
          </div>
          
          <div v-else class="space-y-6">
            <div v-for="review in userReviews" :key="review.id" 
              class="border border-gray-200 rounded-lg p-4">
              <div class="flex items-start space-x-4">
                <img :src="review.recipe.featuredImage" :alt="review.recipe.title" 
                  class="w-16 h-16 rounded-lg object-cover" />
                <div class="flex-1">
                  <div class="flex items-center justify-between mb-2">
                    <h4 class="font-medium text-gray-900 hover:text-orange-500 cursor-pointer"
                      @click="navigateTo(`/recipes/${review.recipe.id}`)">
                      {{ review.recipe.title }}
                    </h4>
                    <div class="flex items-center">
                      <Star v-for="i in 5" :key="i" 
                        :class="['h-4 w-4', i <= review.rating ? 'text-yellow-400 fill-current' : 'text-gray-300']" />
                    </div>
                  </div>
                  <p class="text-gray-700 text-sm">{{ review.comment }}</p>
                  <p class="text-gray-500 text-xs mt-2">{{ formatDate(review.createdAt) }}</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Following Tab -->
        <div v-if="activeTab === 'following'">
          <div v-if="following.length === 0" class="text-center py-12">
            <Users class="h-16 w-16 text-gray-400 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-gray-900 mb-2">Not following anyone</h3>
            <p class="text-gray-600">
              {{ isOwnProfile ? "You're not following anyone yet." : "This user isn't following anyone yet." }}
            </p>
          </div>
          
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <UserCard v-for="user in following" :key="user.id" :user="user" />
          </div>
        </div>

        <!-- Followers Tab -->
        <div v-if="activeTab === 'followers'">
          <div v-if="followers.length === 0" class="text-center py-12">
            <Users class="h-16 w-16 text-gray-400 mx-auto mb-4" />
            <h3 class="text-lg font-medium text-gray-900 mb-2">No followers</h3>
            <p class="text-gray-600">
              {{ isOwnProfile ? "You don't have any followers yet." : "This user doesn't have any followers yet." }}
            </p>
          </div>
          
          <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            <UserCard v-for="user in followers" :key="user.id" :user="user" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { 
  CheckCircle, Star, Share2, Trophy, ChefHat, Heart, 
  MessageSquare, Users 
} from 'lucide-vue-next'
import { useRoute, navigateTo } from '#app'
import { useAuthStore } from '~/stores/auth'
import RecipeCard from '~/components/RecipeCard.vue'
import UserCard from '~/components/UserCard.vue'
import { useNuxtApp, gql, createError } from '#app'

const route = useRoute()
const authStore = useAuthStore()

const username = route.params.username
const userProfile = ref(null)
const activeTab = ref('recipes')
const isFollowing = ref(false)

const userRecipes = ref([])
const likedRecipes = ref([])
const userReviews = ref([])
const following = ref([])
const followers = ref([])

const isOwnProfile = computed(() => {
  return authStore.user && authStore.user.username === username
})

const tabs = computed(() => [
  { id: 'recipes', name: 'Recipes', count: userProfile.value?.recipeCount || 0 },
  { id: 'liked', name: 'Liked', count: userProfile.value?.likedCount || 0 },
  { id: 'reviews', name: 'Reviews', count: userProfile.value?.reviewCount || 0 },
  { id: 'following', name: 'Following', count: userProfile.value?.followingCount || 0 },
  { id: 'followers', name: 'Followers', count: userProfile.value?.followersCount || 0 }
])

onMounted(async () => {
  await loadUserProfile()
  await loadTabContent()
})

const loadUserProfile = async () => {
  try {
    const { $apollo } = useNuxtApp()
    
    const GET_USER_PROFILE_QUERY = gql`
      query GetUserProfile($username: String!) {
        getUserByUsername(username: $username) {
          id
          firstName
          lastName
          username
          email
          avatar
          bio
          isVerified
          averageRating
          totalReviews
          recipeCount
          followersCount
          followingCount
          likedCount
          reviewCount
          achievements {
            id
            name
            description
            icon
          }
          createdAt
        }
      }
    `
    
    const { data } = await $apollo.query({
      query: GET_USER_PROFILE_QUERY,
      variables: { username }
    })
    
    userProfile.value = data.getUserByUsername
    
    // Check if current user is following this user
    if (authStore.isAuthenticated && !isOwnProfile.value) {
      await checkFollowStatus()
    }
    
  } catch (error) {
    console.error('Failed to load user profile:', error)
    // Handle user not found
    if (error.message.includes('not found')) {
      throw createError({ statusCode: 404, statusMessage: 'User not found' })
    }
  }
}

const checkFollowStatus = async () => {
  try {
    const { $apollo } = useNuxtApp()
    
    const CHECK_FOLLOW_STATUS_QUERY = gql`
      query CheckFollowStatus($userId: ID!) {
        isFollowing(userId: $userId)
      }
    `
    
    const { data } = await $apollo.query({
      query: CHECK_FOLLOW_STATUS_QUERY,
      variables: { userId: userProfile.value.id }
    })
    
    isFollowing.value = data.isFollowing
  } catch (error) {
    console.error('Failed to check follow status:', error)
  }
}

const loadTabContent = async () => {
  switch (activeTab.value) {
    case 'recipes':
      await loadUserRecipes()
      break
    case 'liked':
      await loadLikedRecipes()
      break
    case 'reviews':
      await loadUserReviews()
      break
    case 'following':
      await loadFollowing()
      break
    case 'followers':
      await loadFollowers()
      break
  }
}

const loadUserRecipes = async () => {
  try {
    const { $apollo } = useNuxtApp()
    
    const GET_USER_RECIPES_QUERY = gql`
      query GetUserRecipes($userId: ID!) {
        getUserRecipes(userId: $userId) {
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
      }
    `
    
    const { data } = await $apollo.query({
      query: GET_USER_RECIPES_QUERY,
      variables: { userId: userProfile.value.id }
    })
    
    userRecipes.value = data.getUserRecipes
  } catch (error) {
    console.error('Failed to load user recipes:', error)
  }
}

const loadLikedRecipes = async () => {
  // Similar implementation for liked recipes
}

const loadUserReviews = async () => {
  // Similar implementation for user reviews
}

const loadFollowing = async () => {
  // Similar implementation for following list
}

const loadFollowers = async () => {
  // Similar implementation for followers list
}

const toggleFollow = async () => {
  if (!authStore.isAuthenticated) {
    await navigateTo('/login')
    return
  }
  
  try {
    const { $apollo } = useNuxtApp()
    
    const TOGGLE_FOLLOW_MUTATION = gql`
      mutation ToggleFollow($userId: ID!) {
        toggleFollow(userId: $userId) {
          success
          isFollowing
        }
      }
    `
    
    const { data } = await $apollo.mutate({
      mutation: TOGGLE_FOLLOW_MUTATION,
      variables: { userId: userProfile.value.id }
    })
    
    if (data.toggleFollow.success) {
      isFollowing.value = data.toggleFollow.isFollowing
      // Update follower count
      userProfile.value.followersCount += isFollowing.value ? 1 : -1
    }
  } catch (error) {
    console.error('Failed to toggle follow:', error)
  }
}

const editProfile = () => {
  navigateTo('/profile/edit')
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

// Watch for tab changes
watch(activeTab, loadTabContent)

// SEO
useHead({
  title: computed(() => userProfile.value ? `${userProfile.value.firstName} ${userProfile.value.lastName} - RecipeHub` : 'User Profile - RecipeHub'),
  meta: [
    { name: 'description', content: computed(() => userProfile.value?.bio || 'User profile on RecipeHub') }
  ]
})
</script>
