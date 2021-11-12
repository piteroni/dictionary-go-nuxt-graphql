import { Context } from "@nuxt/types"
import { InMemoryCache, IntrospectionFragmentMatcher } from "apollo-cache-inmemory"
import introspection from "@/graphql/generated/introspection.json"

const fragmentMatcher = new IntrospectionFragmentMatcher({
  introspectionQueryResultData: introspection
})

export default function ({ $config }: Context) {
  return {
    httpEndpoint: $config.httpEndpoint,
    browserHttpEndpoint: $config.browserHttpEndpoint,
    cache: new InMemoryCache({ fragmentMatcher })
  }
}
