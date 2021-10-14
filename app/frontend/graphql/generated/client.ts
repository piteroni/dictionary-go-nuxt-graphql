import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Gender = {
  __typename?: 'Gender';
  iconName: Scalars['String'];
  name: Scalars['String'];
};

export type Pokemon = {
  __typename?: 'Pokemon';
  genders: Array<Gender>;
  imageName: Scalars['String'];
  name: Scalars['String'];
  nationalNo: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  pokemon: Pokemon;
};


export type QueryPokemonArgs = {
  pokemonId: Scalars['Int'];
};

export type PokemonQueryVariables = Exact<{
  pokemonId: Scalars['Int'];
}>;


export type PokemonQuery = { __typename?: 'Query', pokemon: { __typename?: 'Pokemon', nationalNo: number, name: string, imageName: string, genders: Array<{ __typename?: 'Gender', name: string, iconName: string }> } };


export const PokemonDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"pokemon"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pokemon"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pokemonId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nationalNo"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageName"}},{"kind":"Field","name":{"kind":"Name","value":"genders"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconName"}}]}}]}}]}}]} as unknown as DocumentNode<PokemonQuery, PokemonQueryVariables>;