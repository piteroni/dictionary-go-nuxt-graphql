export default {
  head: {
    title: "dictionary-go-nuxt-graphql",
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
    "~/plugins/axios",
    '~/plugins/apollo'
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: false,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    "@nuxt/typescript-build",
    // https://go.nuxtjs.dev/tailwindcss
    "@nuxtjs/tailwindcss",
    // https://composition-api.nuxtjs.org
    "@nuxtjs/composition-api/module"
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    "@nuxtjs/axios",
    "@nuxtjs/apollo"
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },

  // https://axios.nuxtjs.org/options/
  // @deprecated
  axios: {
    baseURL: "http://localhost:8080/api/i",
  },

  apollo: {
    clientConfigs: {
      default: {
        httpEndpoint: "http://localhost:8080/api/i/query",
      }
    }
  }
}
