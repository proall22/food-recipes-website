<template>
  <div class="min-h-screen bg-gradient-to-br from-orange-50 to-red-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div class="text-center">
        <div class="flex justify-center">
          <ChefHat class="h-12 w-12 text-orange-500" />
        </div>
        <h2 class="mt-6 text-3xl font-extrabold text-gray-900">
          Join RecipeHub Community
        </h2>
        <p class="mt-2 text-sm text-gray-600">
          Create your account and start sharing delicious recipes
        </p>
      </div>

      <div class="bg-white rounded-2xl shadow-xl p-8">
        <form @submit.prevent="handleSignup" class="space-y-6">
          <!-- Profile Picture Upload -->
          <div class="flex justify-center">
            <div class="relative">
              <div class="w-24 h-24 rounded-full bg-gray-200 flex items-center justify-center overflow-hidden">
                <img v-if="form.profilePicture" :src="profilePicturePreview" 
                  class="w-full h-full object-cover" />
                <User v-else class="h-12 w-12 text-gray-400" />
              </div>
              <button
                type="button"
                @click="$refs.profilePictureInput.click()"
                class="absolute bottom-0 right-0 bg-orange-500 text-white rounded-full p-2 hover:bg-orange-600 transition-colors"
              >
                <Camera class="h-4 w-4" />
              </button>
              <input
                ref="profilePictureInput"
                type="file"
                accept="image/*"
                @change="handleProfilePictureChange"
                class="hidden"
              />
            </div>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                First Name
              </label>
              <Field
                name="firstName"
                type="text"
                v-model="form.firstName"
                :class="[
                  'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.firstName ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="John"
              />
              <ErrorMessage name="firstName" class="text-red-500 text-sm mt-1" />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Last Name
              </label>
              <Field
                name="lastName"
                type="text"
                v-model="form.lastName"
                :class="[
                  'w-full px-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.lastName ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="Doe"
              />
              <ErrorMessage name="lastName" class="text-red-500 text-sm mt-1" />
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Username
            </label>
            <div class="relative">
              <AtSign class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <Field
                name="username"
                type="text"
                v-model="form.username"
                :class="[
                  'w-full pl-10 pr-4 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.username ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="johndoe"
              />
            </div>
            <ErrorMessage name="username" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
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
                placeholder="john@example.com"
              />
            </div>
            <ErrorMessage name="email" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
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
                placeholder="Create a strong password"
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
            
            <!-- Password Strength Indicator -->
            <div class="mt-2">
              <div class="flex space-x-1">
                <div v-for="i in 4" :key="i" 
                  :class="[
                    'h-1 flex-1 rounded',
                    passwordStrength >= i ? getStrengthColor(passwordStrength) : 'bg-gray-200'
                  ]">
                </div>
              </div>
              <p class="text-xs text-gray-600 mt-1">
                {{ getStrengthText(passwordStrength) }}
              </p>
            </div>
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Confirm Password
            </label>
            <div class="relative">
              <Lock class="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
              <Field
                name="confirmPassword"
                :type="showConfirmPassword ? 'text' : 'password'"
                v-model="form.confirmPassword"
                :class="[
                  'w-full pl-10 pr-12 py-3 border rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors',
                  errors.confirmPassword ? 'border-red-500' : 'border-gray-300'
                ]"
                placeholder="Confirm your password"
              />
              <button
                type="button"
                @click="showConfirmPassword = !showConfirmPassword"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
              >
                <Eye v-if="!showConfirmPassword" class="h-5 w-5" />
                <EyeOff v-else class="h-5 w-5" />
              </button>
            </div>
            <ErrorMessage name="confirmPassword" class="text-red-500 text-sm mt-1" />
          </div>

          <div>
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Bio (Optional)
            </label>
            <textarea
              v-model="form.bio"
              rows="3"
              class="w-full px-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent transition-colors resize-none"
              placeholder="Tell us about yourself and your cooking style..."
            ></textarea>
          </div>

          <div class="flex items-center">
            <input
              id="agree-terms"
              name="agree-terms"
              type="checkbox"
              v-model="form.agreeTerms"
              class="h-4 w-4 text-orange-600 focus:ring-orange-500 border-gray-300 rounded"
            />
            <label for="agree-terms" class="ml-2 block text-sm text-gray-900">
              I agree to the 
              <NuxtLink to="/terms" class="text-orange-600 hover:text-orange-500">Terms of Service</NuxtLink>
              and 
              <NuxtLink to="/privacy" class="text-orange-600 hover:text-orange-500">Privacy Policy</NuxtLink>
            </label>
          </div>

          <div>
            <button
              type="submit"
              :disabled="isLoading || !form.agreeTerms"
              class="group relative w-full flex justify-center py-3 px-4 border border-transparent text-sm font-medium rounded-lg text-white bg-orange-600 hover:bg-orange-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-orange-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              <span v-if="isLoading" class="absolute left-0 inset-y-0 flex items-center pl-3">
                <Loader2 class="h-5 w-5 animate-spin" />
              </span>
              {{ isLoading ? 'Creating account...' : 'Create Account' }}
            </button>
          </div>

          <div class="text-center">
            <span class="text-sm text-gray-600">
              Already have an account?
              <NuxtLink to="/login" class="font-medium text-orange-600 hover:text-orange-500">
                Sign in here
              </NuxtLink>
            </span>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useForm, Field, ErrorMessage } from 'vee-validate'
import * as yup from 'yup'
import { 
  ChefHat, Mail, Lock, Eye, EyeOff, Loader2, User, Camera, AtSign 
} from 'lucide-vue-next'
import { navigateTo } from '#app'
import { useAuthStore } from '~/stores/auth'
import { useHead } from '#head'

const authStore = useAuthStore()

const showPassword = ref(false)
const showConfirmPassword = ref(false)
const isLoading = ref(false)
const profilePicturePreview = ref('')

const form = reactive({
  firstName: '',
  lastName: '',
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  bio: '',
  profilePicture: null,
  agreeTerms: false
})

// Validation schema
const schema = yup.object({
  firstName: yup.string().required('First name is required'),
  lastName: yup.string().required('Last name is required'),
  username: yup.string().min(3, 'Username must be at least 3 characters').required('Username is required'),
  email: yup.string().email('Invalid email format').required('Email is required'),
  password: yup.string().min(8, 'Password must be at least 8 characters').required('Password is required'),
  confirmPassword: yup.string()
    .oneOf([yup.ref('password')], 'Passwords must match')
    .required('Please confirm your password')
})

const { errors, validate } = useForm({
  validationSchema: schema
})

const passwordStrength = computed(() => {
  const password = form.password
  let strength = 0
  
  if (password.length >= 8) strength++
  if (/[A-Z]/.test(password)) strength++
  if (/[0-9]/.test(password)) strength++
  if (/[^A-Za-z0-9]/.test(password)) strength++
  
  return strength
})

const getStrengthColor = (strength) => {
  switch (strength) {
    case 1: return 'bg-red-500'
    case 2: return 'bg-yellow-500'
    case 3: return 'bg-blue-500'
    case 4: return 'bg-green-500'
    default: return 'bg-gray-200'
  }
}

const getStrengthText = (strength) => {
  switch (strength) {
    case 1: return 'Weak password'
    case 2: return 'Fair password'
    case 3: return 'Good password'
    case 4: return 'Strong password'
    default: return 'Enter a password'
  }
}

const handleProfilePictureChange = (event) => {
  const file = event.target.files[0]
  if (file) {
    form.profilePicture = file
    const reader = new FileReader()
    reader.onload = (e) => {
      profilePicturePreview.value = e.target.result
    }
    reader.readAsDataURL(file)
  }
}

const handleSignup = async () => {
  const { valid } = await validate()
  if (!valid || !form.agreeTerms) return

  isLoading.value = true
  
  try {
    await authStore.signup({
      firstName: form.firstName,
      lastName: form.lastName,
      username: form.username,
      email: form.email,
      password: form.password,
      bio: form.bio,
      profilePicture: form.profilePicture
    })
    
    await navigateTo('/login?message=Account created successfully')
  } catch (error) {
    console.error('Signup failed:', error)
    // Show error message
  } finally {
    isLoading.value = false
  }
}

// SEO
useHead({
  title: 'Sign Up - RecipeHub',
  meta: [
    { name: 'description', content: 'Join RecipeHub community and start sharing your favorite recipes with food lovers around the world' }
  ]
})
</script>
