import { ref, computed } from 'vue'
import { useNuxtApp, navigateTo } from '#app'
import { gql } from 'graphql-tag'

export const useAuth = () => {
  const { $apollo } = useNuxtApp()
  
  const user = ref(null)
  const token = ref(null)
  const isLoading = ref(false)

  const isAuthenticated = computed(() => !!token.value)

  // Initialize auth state
  const initAuth = () => {
    if (process.client) {
      const savedToken = localStorage.getItem('auth_token')
      const savedUser = localStorage.getItem('auth_user')

      if (savedToken && savedUser) {
        token.value = savedToken
        user.value = JSON.parse(savedUser)
      }
    }
  }

  // Login mutation
  const login = async (credentials) => {
    isLoading.value = true

    try {
      const LOGIN_MUTATION = gql`
        mutation Login($input: LoginInput!) {
          login(input: $input) {
            success
            message
            access_token
            refresh_token
            user {
              id
              email
              firstName: first_name
              lastName: last_name
              username
              avatar
              bio
              isVerified: is_verified
              createdAt: created_at
            }
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: LOGIN_MUTATION,
        variables: {
          input: {
            email: credentials.email,
            password: credentials.password
          }
        }
      })

      if (data.login.success) {
        token.value = data.login.access_token
        user.value = data.login.user

        // Save to localStorage
        if (process.client) {
          localStorage.setItem('auth_token', data.login.access_token)
          localStorage.setItem('auth_user', JSON.stringify(data.login.user))
        }

        return data.login
      } else {
        throw new Error(data.login.message)
      }
    } catch (error) {
      console.error('Login error:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Signup mutation
  const signup = async (userData) => {
    isLoading.value = true

    try {
      const SIGNUP_MUTATION = gql`
        mutation Signup($input: SignupInput!) {
          signup(input: $input) {
            success
            message
            user {
              id
              email
              firstName: first_name
              lastName: last_name
              username
            }
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: SIGNUP_MUTATION,
        variables: {
          input: {
            email: userData.email,
            username: userData.username,
            first_name: userData.firstName,
            last_name: userData.lastName,
            password: userData.password,
            bio: userData.bio || '',
            avatar: userData.avatar || ''
          }
        }
      })

      return data.signup
    } catch (error) {
      console.error('Signup error:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Logout
  const logout = async () => {
    token.value = null
    user.value = null

    if (process.client) {
      localStorage.removeItem('auth_token')
      localStorage.removeItem('auth_user')
    }

    await navigateTo('/')
  }

  // Upload file
  const uploadFile = async (file, category = 'recipe') => {
    try {
      const formData = new FormData()
      formData.append('file', file)
      formData.append('category', category)

      const response = await $fetch('http://localhost:8000/upload/image', {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${token.value}`
        },
        body: formData
      })

      return `http://localhost:8000${response.url}`
    } catch (error) {
      console.error('File upload error:', error)
      throw error
    }
  }

  return {
    user,
    token,
    isLoading,
    isAuthenticated,
    initAuth,
    login,
    signup,
    logout,
    uploadFile
  }
}