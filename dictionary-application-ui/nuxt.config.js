export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: "%s - dictionary-application-ui",
    title: "dictionary-application-ui",
    htmlAttrs: {
      lang: "ja"
    },
    meta: [
      { charset: "utf-8" },
      { name: "viewport", content: "width=device-width, initial-scale=1" },
      { hid: "description", name: "description", content: "" },
      { name: "format-detection", content: "telephone=no" }
    ],
    link: [
      { rel: "icon", type: "image/x-icon", href: "/favicon.ico" }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: false,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    "@nuxt/typescript-build",
    // https://go.nuxtjs.dev/tailwindcss
    "@nuxtjs/tailwindcss",
    // https://typed-vuex.roe.dev
    "nuxt-typed-vuex",
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    "@nuxtjs/apollo"
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },

  apollo: {
    clientConfigs: {
      default: "~/apollo.config.ts"
    }
  },

  privateRuntimeConfig: {
    httpEndpoint: process.env.HTTP_ENDPOINT
  },

  publicRuntimeConfig: {
    browserHttpEndpoint: process.env.BROWSER_HTTP_ENDPOINT,
  },

  router: {
    extendRoutes(routes) {
      routes.push(
        { path: "/", redirect: "/pokemons" },
      )
    }
  }
}
