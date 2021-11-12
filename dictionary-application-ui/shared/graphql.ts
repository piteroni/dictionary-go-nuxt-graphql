export type QueryType<Union, Type> = Union extends { __typename: Type } ? Union : never
