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

export type Characteristic = {
  __typename?: 'Characteristic';
  description: Scalars['String'];
  name: Scalars['String'];
};

export type Description = {
  __typename?: 'Description';
  Series: Scalars['String'];
  text: Scalars['String'];
};

export type Gender = {
  __typename?: 'Gender';
  iconName: Scalars['String'];
  name: Scalars['String'];
};

export type Pokemon = {
  __typename?: 'Pokemon';
  characteristics: Array<Characteristic>;
  description: Description;
  genders: Array<Gender>;
  height: Scalars['String'];
  imageName: Scalars['String'];
  name: Scalars['String'];
  nationalNo: Scalars['Int'];
  species: Scalars['String'];
  types: Array<Type>;
  weight: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  pokemon: Pokemon;
};


export type QueryPokemonArgs = {
  pokemonId: Scalars['Int'];
};

export type Type = {
  __typename?: 'Type';
  iconName: Scalars['String'];
  name: Scalars['String'];
};

export type PokemonQueryVariables = Exact<{
  pokemonId: Scalars['Int'];
}>;


export type PokemonQuery = { __typename?: 'Query', pokemon: { __typename?: 'Pokemon', nationalNo: number, name: string, imageName: string, species: string, weight: string, height: string, genders: Array<{ __typename?: 'Gender', name: string, iconName: string }>, types: Array<{ __typename?: 'Type', name: string, iconName: string }>, characteristics: Array<{ __typename?: 'Characteristic', name: string, description: string }> } };


export const PokemonDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"pokemon"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pokemon"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pokemonId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nationalNo"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageName"}},{"kind":"Field","name":{"kind":"Name","value":"species"}},{"kind":"Field","name":{"kind":"Name","value":"weight"}},{"kind":"Field","name":{"kind":"Name","value":"height"}},{"kind":"Field","name":{"kind":"Name","value":"genders"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconName"}}]}},{"kind":"Field","name":{"kind":"Name","value":"types"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconName"}}]}},{"kind":"Field","name":{"kind":"Name","value":"characteristics"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}}]}}]}}]} as unknown as DocumentNode<PokemonQuery, PokemonQueryVariables>;