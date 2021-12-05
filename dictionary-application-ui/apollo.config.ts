import { Context } from "@nuxt/types"
import { onError } from "@apollo/client/link/error"
import { InMemoryCache, IntrospectionFragmentMatcher } from "apollo-cache-inmemory"
import { HttpStatusCode } from "@/shared/http"
import introspection from "@/graphql/generated/introspection.json"

const fragmentMatcher = new IntrospectionFragmentMatcher({
  introspectionQueryResultData: introspection
})

export default function ({ error, $config }: Context) {
  const errorLink = onError(e => {
    const message = JSON.stringify({
      operationName: e.operation.operationName,
      graphQLErrors: e.graphQLErrors
    })

    console.error(message)

    error({ statusCode: HttpStatusCode.INTERNAL_SERVER_ERROR })
  })

  return {
    link: errorLink,
    httpEndpoint: $config.httpEndpoint,
    browserHttpEndpoint: $config.browserHttpEndpoint,
    cache: new InMemoryCache({ fragmentMatcher })
  }
}
