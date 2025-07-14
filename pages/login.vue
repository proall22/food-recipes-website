<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-red-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <div class="flex justify-center">
          <ChefHat class="h-12 w-12 text-orange-500" />
        </div>
        <h2 class="mt-6 text-3xl font-extrabold text-gray-900">
          Welcome back to RecipeHub
        </h2>
        <p class="mt-2 text-sm text-gray-600">
          Sign in to your account to continue sharing amazing recipes
        </p>
      </div>

      <div class="bg-white rounded-2xl shadow-xl p-8">
        <form @submit.prevent="handleLogin" class="space-y-6">
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
              Email Address
            </label>
            <div class="relative">
              <Mail class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <Field
                name="email"
                type="email"
                v-model="form.email"
                :class="[
                  'w-full pl-10 pr-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.email ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="Enter your email"
              />
            </div>
            <ErrorMessage name="email" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label for="password" class="block text-sm font-medium text-gray-700 mb-2">
              Password
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <Field
                name="password"
                :type="showPassword ? 'text' : 'password'"
                v-model="form.password"
                :class="[
                  'w-full pl-10 pr-12 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.password ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="Enter your password"
              />
              <button
                type="button"
                @click="showPassword = !showPassword"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <Eye v-if="!showPassword" class="h-5 w-5" />
                <EyeOff v-else class="h-5 w-5" />
              </button>
            </div>
            <ErrorMessage name="password" class="text-red-500 text-sm mt-1" />
          </div>

          <div class="flex items-center justify-between">
            <div class="flex items-center">
              <input
                id="remember-me"
                name="remember-me"
                type="checkbox"
                v-model="form.rememberMe"
                class="h-4 w-4 text-orange-600 focus:ring-orange-500 border-gray-300 rounded"
              />
              <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                Remember me
              </label>
            </div>

            <div class="text-sm">
              <NuxtLink to="/forgot-password" class="font-medium text-orange-600 hover:text-orange-500">
                Forgot your password?
              </NuxtLink>
            </div>
          </div>

          <div>
            <button
              type="submit"
              :disabled="isLoading"
              class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-orange-600 hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              <span v-if="isLoading" class="absolute left-0 inset-y-0 flex items-center pl-3">
                <Loader2 class="h-5 w-5 animate-spin" />
              </span>
              {{ isLoading ? 'Signing in...' : 'Sign in' }}
            </button>
          </div>

          <div class="text-center">
            <span class="text-sm text-gray-600">
              Don't have an account?
              <NuxtLink to="/signup" class="font-medium text-orange-600 hover:text-orange-500">
                Sign up here
              </NuxtLink>
            </span>
          </div>
        </form>

        <!-- Social Login -->
        <div class="mt-6">
          <div class="relative">
            <div class="absolute inset-0 flex items-center">
              <div class="w-full border-t border-gray-300" />
            </div>
            <div class="relative flex justify-center text-sm">
              <span class="px-2 bg-white text-gray-500">Or continue with</span>
            </div>
          </div>

          <div class="mt-6 grid grid-cols-2 gap-3">
            <button
              @click="loginWithGoogle"
              class="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
            >
              <svg class="h-5 w-5" viewBox="0 0 24 24">
                <path fill="#4285F4" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
                <path fill="#34A853" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23s.43 3.45 1.18 4.93l3.66 2.84.81-.62z"/>
                <path fill="#FBBC05" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
                <path fill="#EA4335" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
              </svg>
              <span class="ml-2">Google</span>
            </button>

            <button
              @click="loginWithFacebook"
              class="w-full inline-flex justify-center py-2 px-4 border border-gray-300 rounded-md shadow-sm bg-white text-sm font-medium text-gray-500 hover:bg-gray-50"
            >
              <Facebook class="h-5 w-5 text-blue-600" />
              <span class="ml-2">Facebook</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useForm, Field, ErrorMessage } from 'vee-validate'
import * as yup from 'yup'
import { ChefHat, Mail, Lock, Eye, EyeOff, Loader2, Facebook } from 'lucide-vue-next'
import { navigateTo, useRoute } from '#app'
import { useAuthStore } from '~/stores/auth'
import { useHead } from '#head'

const authStore = useAuthStore()
const route = useRoute()

const showPassword = ref(false)
const isLoading = ref(false)

const form = reactive({
  email: '',
  password: '',
  rememberMe: false
})

// Validation schema
const schema = yup.object({
  email: yup.string().email('Invalid email format').required('Email is required'),
  password: yup.string().min(6, 'Password must be at least 6 characters').required('Password is required')
})

const { errors, validate } = useForm({
  validationSchema: schema
})

const handleLogin = async () => {
  const { valid } = await validate()
  if (!valid) return

  isLoading.value = true
  
  try {
    await authStore.login({
      email: form.email,
      password: form.password,
      rememberMe: form.rememberMe
    })
    
    // Redirect to intended page or home
    const redirect = route.query.redirect || '/'
    await navigateTo(redirect)
  } catch (error) {
    console.error('Login failed:', error)
    // Show error message
  } finally {
    isLoading.value = false
  }
}

const loginWithGoogle = async () => {
  try {
    await authStore.loginWithGoogle()
    await navigateTo('/')
  } catch (error) {
    console.error('Google login failed:', error)
  }
}

const loginWithFacebook = async () => {
  try {
    await authStore.loginWithFacebook()
    await navigateTo('/')
  } catch (error) {
    console.error('Facebook login failed:', error)
  }
}

// SEO
useHead({
  title: 'Login - RecipeHub',
  meta: [
    { name: 'description', content: 'Sign in to your RecipeHub account to share and discover amazing recipes' }
  ]
})
</script>
