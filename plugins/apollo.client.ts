import { ApolloClient, InMemoryCache, createHttpLink, from } from "@apollo/client/core"
import { setContext } from "@apollo/client/link/context"
import { onError } from "@apollo/client/link/error"
import { defineNuxtPlugin } from "#app"

export default defineNuxtPlugin(() => {
  const httpLink = createHttpLink({
    uri: "http://localhost:8080/v1/graphql",
  })

  const authLink = setContext((_, { headers }) => {
    let token = null

    if (process.client) {
      token = localStorage.getItem("auth_token")
    }

    return {
      headers: {
        ...headers,
        ...(token ? { authorization: `Bearer ${token}` } : {}),
      },
    }
  })

  const errorLink = onError(({ graphQLErrors, networkError, operation, forward }) => {
    if (graphQLErrors) {
      graphQLErrors.forEach(({ message, locations, path }) => {
        console.error(`GraphQL error: Message: ${message}, Location: ${locations}, Path: ${path}`)
      })
    }

    if (networkError) {
      console.error(`Network error: ${networkError}`)

      // Handle authentication errors
      if (networkError.statusCode === 401) {
        if (process.client) {
          localStorage.removeItem("auth_token")
          localStorage.removeItem("auth_user")
          window.location.href = "/login"
        }
      }
    }
  })

  const apolloClient = new ApolloClient({
    link: from([errorLink, authLink, httpLink]),
    cache: new InMemoryCache(),
    defaultOptions: {
      watchQuery: {
        errorPolicy: "all",
      },
      query: {
        errorPolicy: "all",
      },
    },
  })

  return {
    provide: {
      apollo: apolloClient,
    },
  }
})
