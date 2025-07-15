<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-red-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <div class="flex justify-center">
          <ChefHatIcon class="h-12 w-12 text-orange-500" />
        </div>
        <h2 class="mt-6 text-3xl font-extrabold text-gray-900">
          Welcome back to RecipeHub
        </h2>
        <p class="mt-2 text-sm text-gray-600">
          Sign in to your account to continue sharing amazing recipes
        </p>
      </div>

      <div class="bg-white rounded-2xl shadow-xl p-8">
        <Form @submit="handleLogin" :validation-schema="schema" v-slot="{ errors }">
          <div class="space-y-6">
            <div>
              <label for="email" class="block text-sm font-medium text-gray-700 mb-2">
                Email Address
              </label>
              <div class="relative">
                <EnvelopeIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
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
                <LockClosedIcon class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
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
                  <EyeIcon v-if="!showPassword" class="h-5 w-5" />
                  <EyeSlashIcon v-else class="h-5 w-5" />
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
                  <ArrowPathIcon class="h-5 w-5 animate-spin" />
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
          </div>
        </Form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { Form, Field, ErrorMessage } from 'vee-validate'
import * as yup from 'yup'
import { 
  ChefHatIcon, 
  EnvelopeIcon, 
  LockClosedIcon, 
  EyeIcon, 
  EyeSlashIcon,
  ArrowPathIcon 
} from '@heroicons/vue/24/outline'

definePageMeta({
  middleware: 'guest',
  layout: false
})

const { login } = useAuth()
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

const handleLogin = async () => {
  isLoading.value = true
  
  try {
    await login({
      email: form.email,
      password: form.password,
      rememberMe: form.rememberMe
    })
    
    await navigateTo('/')
  } catch (error) {
    console.error('Login failed:', error)
    // Show error message to user
  } finally {
    isLoading.value = false
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