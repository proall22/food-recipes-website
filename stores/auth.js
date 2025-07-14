import { defineStore } from "pinia"
import { ref, computed } from "vue"
import { useNuxtApp } from "#app"
import { gql } from "graphql-tag"

const { $apollo } = useNuxtApp()

export const useAuthStore = defineStore("auth", () => {
  const user = ref(null)
  const token = ref(null)
  const isLoading = ref(false)

  const isAuthenticated = computed(() => !!token.value)
  const userProfile = computed(() => user.value)

  // Initialize auth state from localStorage
  const initAuth = () => {
    if (process.client) {
      const savedToken = localStorage.getItem("auth_token")
      const savedUser = localStorage.getItem("auth_user")

      if (savedToken && savedUser) {
        token.value = savedToken
        user.value = JSON.parse(savedUser)

        // Set Apollo client auth header
        $apollo.defaultClient.setHeader("Authorization", `Bearer ${savedToken}`)
      }
    }
  }

  // Login with email and password
  const login = async (credentials) => {
    isLoading.value = true

    try {
      const LOGIN_ACTION = gql`
        mutation Login($input: LoginInput!) {
          login(input: $input) {
            success
            message
            access_token
            refresh_token
            user {
              id
              email
              first_name
              last_name
              username
              avatar
              bio
              is_verified
              created_at
            }
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: LOGIN_ACTION,
        variables: {
          input: {
            email: credentials.email,
            password: credentials.password,
          },
        },
      })

      if (data.login.success) {
        token.value = data.login.access_token
        user.value = data.login.user

        // Save to localStorage
        if (process.client) {
          localStorage.setItem("auth_token", data.login.access_token)
          localStorage.setItem("auth_user", JSON.stringify(data.login.user))

          if (credentials.rememberMe) {
            localStorage.setItem("remember_me", "true")
          }
        }

        // Set Apollo client auth header
        $apollo.defaultClient.setHeader("Authorization", `Bearer ${data.login.access_token}`)
      } else {
        throw new Error(data.login.message)
      }

      return data.login
    } catch (error) {
      console.error("Login error:", error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Signup new user
  const signup = async (userData) => {
    isLoading.value = true

    try {
      const SIGNUP_ACTION = gql`
        mutation Signup($input: SignupInput!) {
          signup(input: $input) {
            success
            message
            user {
              id
              email
              first_name
              last_name
              username
            }
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: SIGNUP_ACTION,
        variables: {
          input: {
            email: userData.email,
            username: userData.username,
            first_name: userData.firstName,
            last_name: userData.lastName,
            password: userData.password,
            bio: userData.bio || "",
            avatar: userData.avatar || "",
          },
        },
      })

      return data.signup
    } catch (error) {
      console.error("Signup error:", error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  // Logout user
  const logout = async () => {
    try {
      // Clear local state
      token.value = null
      user.value = null

      if (process.client) {
        localStorage.removeItem("auth_token")
        localStorage.removeItem("auth_user")
        localStorage.removeItem("remember_me")
      }

      // Clear Apollo cache and headers
      $apollo.defaultClient.clearStore()
      $apollo.defaultClient.setHeader("Authorization", "")
    } catch (error) {
      console.error("Logout error:", error)
    }
  }

  // Upload file helper
  const uploadFile = async (file, category = "recipe") => {
    try {
      const formData = new FormData()
      formData.append("file", file)
      formData.append("category", category)

      const response = await fetch("http://localhost:8000/upload/image", {
        method: "POST",
        headers: {
          Authorization: `Bearer ${token.value}`,
        },
        body: formData,
      })

      if (!response.ok) {
        throw new Error("Upload failed")
      }

      const result = await response.json()
      return `http://localhost:8000${result.url}`
    } catch (error) {
      console.error("File upload error:", error)
      throw error
    }
  }

  // Initialize payment
  const initializePayment = async (recipeId, amount) => {
    try {
      const INITIALIZE_PAYMENT = gql`
        mutation InitializePayment($input: PaymentInput!) {
          initializePayment(input: $input) {
            success
            message
            checkout_url
          }
        }
      `

      const { data } = await $apollo.mutate({
        mutation: INITIALIZE_PAYMENT,
        variables: {
          input: {
            recipe_id: recipeId,
            amount: amount,
          },
        },
      })

      return data.initializePayment
    } catch (error) {
      console.error("Payment initialization error:", error)
      throw error
    }
  }

  return {
    // State
    user,
    token,
    isLoading,

    // Getters
    isAuthenticated,
    userProfile,

    // Actions
    initAuth,
    login,
    signup,
    logout,
    uploadFile,
    initializePayment,
  }
})
