/* eslint-disable @typescript-eslint/explicit-module-boundary-types */

export function range(start: number, end: number): number[] {
  return Array.from({ length: (end - start + 1) }, (_, k) => k + start)
}
