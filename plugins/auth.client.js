export default defineNuxtPlugin(() => {
  const { initAuth } = useAuth()
  
  // Initialize auth state on client side
  if (process.client) {
    initAuth()
  }
})