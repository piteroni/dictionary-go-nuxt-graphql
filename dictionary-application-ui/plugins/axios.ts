import { Context } from "@nuxt/types/app"
import { HttpStatusCode } from "@/shared/http"
import { range } from "@/shared/utils"

const assumedClientErrors = [
  HttpStatusCode.UNAUTHORIZED,
  HttpStatusCode.FORBIDDEN,
  HttpStatusCode.UNPROCESSABLE_ENTITY
]

export default function ({ $axios, error }: Context): void {
  $axios.setHeader("Content-Type", "application/json")

  $axios.onError(e => {
    // when network error
    if (!e.response) {
      console.error(e)
      error({ message: e.message })

      return
    }

    const status = e.response.status

    if (!assumedClientErrors.includes(status) && range(400, 499).includes(status)) {
      console.error(e)
      error({ message: "HTTP ERROR" })
    }
  })
}
