export default defineNuxtConfig({
  devtools: { enabled: true },
  
  modules: [
    '@nuxtjs/tailwindcss',
    '@pinia/nuxt',
    '@vueuse/nuxt',
    '@nuxtjs/apollo'
  ],
  
  apollo: {
    clients: {
      default: {
        httpEndpoint: 'http://localhost:8080/v1/graphql',
        wsEndpoint: 'ws://localhost:8080/v1/graphql',
        httpLinkOptions: {
          headers: {
            'x-hasura-admin-secret': 'myadminsecretkey'
          }
        }
      }
    }
  },
  
  css: ['~/assets/css/main.css'],
  
  runtimeConfig: {
    public: {
      apiBase: 'http://localhost:8000',
      hasuraEndpoint: 'http://localhost:8080/v1/graphql',
      hasuraAdminSecret: 'myadminsecretkey'
    }
  },
  
  ssr: true,
  
  nitro: {
    preset: 'node-server'
  },
  
  app: {
    head: {
      title: 'RecipeHub - Discover Amazing Recipes',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Discover and share amazing recipes from around the world' }
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ]
    }
  },
  
  tailwindcss: {
    cssPath: '~/assets/css/main.css',
    configPath: 'tailwind.config.js'
  }
})