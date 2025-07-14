import { defineNuxtConfig } from "nuxt"

export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@nuxtjs/tailwindcss", "@pinia/nuxt", "@vueuse/nuxt", "@nuxtjs/apollo"],
  apollo: {
    clients: {
      default: {
        httpEndpoint: "http://localhost:8080/v1/graphql",
        wsEndpoint: "ws://localhost:8080/v1/graphql",
      },
    },
  },
  css: ["~/assets/css/main.css"],
  runtimeConfig: {
    public: {
      apiBase: "http://localhost:8000",
      hasuraEndpoint: "http://localhost:8080/v1/graphql",
    },
  },
  ssr: true,
  nitro: {
    preset: "node-server",
  },
})
