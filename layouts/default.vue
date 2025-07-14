<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation Header -->
    <nav class="bg-white shadow-lg sticky top-0 z-50">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between items-center h-16">
          <!-- Logo -->
          <div class="flex items-center">
            <NuxtLink to="/" class="flex items-center space-x-2">
              <ChefHat class="h-8 w-8 text-orange-500" />
              <span class="text-2xl font-bold text-gray-900">RecipeHub</span>
            </NuxtLink>
          </div>

          <!-- Search Bar -->
          <div class="hidden md:flex flex-1 max-w-lg mx-8">
            <div class="relative w-full">
              <Search class="absolute left-3 top-1/2 transform -translate-y-1/2 h-4 w-4 text-gray-400" />
              <input
                v-model="searchQuery"
                @keyup.enter="handleSearch"
                type="text"
                placeholder="Search recipes..."
                class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-full focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              />
            </div>
          </div>

          <!-- Navigation Links -->
          <div class="hidden md:flex items-center space-x-6">
            <NuxtLink to="/recipes" class="text-gray-700 hover:text-orange-500 font-medium">
              Recipes
            </NuxtLink>
            <NuxtLink to="/categories" class="text-gray-700 hover:text-orange-500 font-medium">
              Categories
            </NuxtLink>
            <NuxtLink to="/creators" class="text-gray-700 hover:text-orange-500 font-medium">
              Creators
            </NuxtLink>
          </div>

          <!-- User Menu -->
          <div class="flex items-center space-x-4">
            <template v-if="user">
              <button @click="toggleBookmarks" class="p-2 text-gray-700 hover:text-orange-500">
                <Bookmark class="h-5 w-5" />
              </button>
              <div class="relative">
                <button @click="toggleUserMenu" class="flex items-center space-x-2">
                  <img :src="user.avatar || '/placeholder.svg?height=32&width=32'" 
                       class="h-8 w-8 rounded-full" />
                  <ChevronDown class="h-4 w-4 text-gray-700" />
                </button>
                <div v-if="showUserMenu" class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-50">
                  <NuxtLink to="/profile" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Profile
                  </NuxtLink>
                  <NuxtLink to="/my-recipes" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    My Recipes
                  </NuxtLink>
                  <NuxtLink to="/create-recipe" class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Create Recipe
                  </NuxtLink>
                  <button @click="logout" class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                    Logout
                  </button>
                </div>
              </div>
            </template>
            <template v-else>
              <NuxtLink to="/login" class="text-gray-700 hover:text-orange-500 font-medium">
                Login
              </NuxtLink>
              <NuxtLink to="/signup" 
                class="bg-orange-500 text-white px-4 py-2 rounded-full hover:bg-orange-600 transition-colors">
                Sign Up
              </NuxtLink>
            </template>
          </div>

          <!-- Mobile Menu Button -->
          <button @click="toggleMobileMenu" class="md:hidden p-2">
            <Menu class="h-6 w-6" />
          </button>
        </div>
      </div>

      <!-- Mobile Menu -->
      <div v-if="showMobileMenu" class="md:hidden bg-white border-t">
        <div class="px-4 py-2 space-y-2">
          <input
            v-model="searchQuery"
            @keyup.enter="handleSearch"
            type="text"
            placeholder="Search recipes..."
            class="w-full px-4 py-2 border border-gray-300 rounded-full focus:ring-2 focus:ring-orange-500"
          />
          <NuxtLink to="/recipes" class="block py-2 text-gray-700">Recipes</NuxtLink>
          <NuxtLink to="/categories" class="block py-2 text-gray-700">Categories</NuxtLink>
          <NuxtLink to="/creators" class="block py-2 text-gray-700">Creators</NuxtLink>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main>
      <slot />
    </main>

    <!-- Footer -->
    <footer class="bg-gray-900 text-white mt-16">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <div class="grid grid-cols-1 md:grid-cols-4 gap-8">
          <div>
            <div class="flex items-center space-x-2 mb-4">
              <ChefHat class="h-8 w-8 text-orange-500" />
              <span class="text-2xl font-bold">RecipeHub</span>
            </div>
            <p class="text-gray-400">Discover and share amazing recipes from around the world.</p>
          </div>
          <div>
            <h3 class="font-semibold mb-4">Recipes</h3>
            <ul class="space-y-2 text-gray-400">
              <li><NuxtLink to="/categories/breakfast" class="hover:text-white">Breakfast</NuxtLink></li>
              <li><NuxtLink to="/categories/lunch" class="hover:text-white">Lunch</NuxtLink></li>
              <li><NuxtLink to="/categories/dinner" class="hover:text-white">Dinner</NuxtLink></li>
              <li><NuxtLink to="/categories/dessert" class="hover:text-white">Desserts</NuxtLink></li>
            </ul>
          </div>
          <div>
            <h3 class="font-semibold mb-4">Community</h3>
            <ul class="space-y-2 text-gray-400">
              <li><NuxtLink to="/creators" class="hover:text-white">Top Creators</NuxtLink></li>
              <li><NuxtLink to="/about" class="hover:text-white">About Us</NuxtLink></li>
              <li><NuxtLink to="/contact" class="hover:text-white">Contact</NuxtLink></li>
            </ul>
          </div>
          <div>
            <h3 class="font-semibold mb-4">Follow Us</h3>
            <div class="flex space-x-4">
              <a href="#" class="text-gray-400 hover:text-white">
                <Facebook class="h-6 w-6" />
              </a>
              <a href="#" class="text-gray-400 hover:text-white">
                <Twitter class="h-6 w-6" />
              </a>
              <a href="#" class="text-gray-400 hover:text-white">
                <Instagram class="h-6 w-6" />
              </a>
            </div>
          </div>
        </div>
        <div class="border-t border-gray-800 mt-8 pt-8 text-center text-gray-400">
          <p>&copy; 2024 RecipeHub. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { 
  ChefHat, Search, Bookmark, ChevronDown, Menu, 
  Facebook, Twitter, Instagram 
} from 'lucide-vue-next'
import { navigateTo } from '#app' // Import navigateTo

const searchQuery = ref('')
const showUserMenu = ref(false)
const showMobileMenu = ref(false)
const user = ref(null) // Will be populated from auth store

const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value
}

const toggleMobileMenu = () => {
  showMobileMenu.value = !showMobileMenu.value
}

const toggleBookmarks = () => {
  // Navigate to bookmarks page
  navigateTo('/bookmarks')
}

const handleSearch = () => {
  if (searchQuery.value.trim()) {
    navigateTo(`/search?q=${encodeURIComponent(searchQuery.value)}`)
  }
}

const logout = () => {
  // Implement logout logic
  user.value = null
  showUserMenu.value = false
  navigateTo('/')
}

// Close menus when clicking outside
onMounted(() => {
  document.addEventListener('click', (e) => {
    if (!e.target.closest('.relative')) {
      showUserMenu.value = false
    }
  })
})
</script>
