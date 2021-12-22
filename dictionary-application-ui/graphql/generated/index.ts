import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
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

export type Ability = {
  __typename?: 'Ability';
  attack: Scalars['Int'];
  defense: Scalars['Int'];
  heart: Scalars['Int'];
  specialAttack: Scalars['Int'];
  specialDefense: Scalars['Int'];
  speed: Scalars['Int'];
};

export type Characteristic = {
  __typename?: 'Characteristic';
  description: Scalars['String'];
  name: Scalars['String'];
};

export type Description = {
  __typename?: 'Description';
  series: Scalars['String'];
  text: Scalars['String'];
};

export type Evolutions = {
  __typename?: 'Evolutions';
  pokemons: Array<Pokemon>;
};

export type EvolutionsResult = Evolutions | IllegalArguments | PokemonNotFound;

export type Gender = {
  __typename?: 'Gender';
  iconURL: Scalars['String'];
  name: Scalars['String'];
};

export type IllegalArguments = {
  __typename?: 'IllegalArguments';
  message: Scalars['String'];
};

export type PageInfo = {
  __typename?: 'PageInfo';
  hasNext: Scalars['Boolean'];
  hasPrev: Scalars['Boolean'];
  nextId: Scalars['String'];
  prevId: Scalars['String'];
};

export type PageInfoResult = IllegalArguments | PageInfo | PokemonNotFound;

export type Pokemon = {
  __typename?: 'Pokemon';
  ability: Ability;
  canEvolution: Scalars['Boolean'];
  characteristics: Array<Characteristic>;
  description: Description;
  genders: Array<Gender>;
  height: Scalars['String'];
  id: Scalars['String'];
  imageURL: Scalars['String'];
  name: Scalars['String'];
  nationalNo: Scalars['Int'];
  species: Scalars['String'];
  types: Array<Type>;
  weight: Scalars['String'];
};

export type PokemonConnection = {
  __typename?: 'PokemonConnection';
  endCursor: Scalars['String'];
  hasNext: Scalars['Boolean'];
  items: Array<Pokemon>;
};

export type PokemonConnectionResult = IllegalArguments | PokemonConnection | PokemonNotFound;

export type PokemonNotFound = {
  __typename?: 'PokemonNotFound';
  message: Scalars['String'];
};

export type PokemonResult = IllegalArguments | Pokemon | PokemonNotFound;

export type Query = {
  __typename?: 'Query';
  evolutions: EvolutionsResult;
  pageInfo: PageInfoResult;
  pokemon: PokemonResult;
  pokemons: PokemonConnectionResult;
};


export type QueryEvolutionsArgs = {
  pokemonId: Scalars['String'];
};


export type QueryPageInfoArgs = {
  pokemonId: Scalars['String'];
};


export type QueryPokemonArgs = {
  pokemonId: Scalars['String'];
};


export type QueryPokemonsArgs = {
  after?: InputMaybe<Scalars['String']>;
  first?: InputMaybe<Scalars['Int']>;
};

export type Type = {
  __typename?: 'Type';
  iconURL: Scalars['String'];
  name: Scalars['String'];
};

export type PokemonQueryVariables = Exact<{
  pokemonId: Scalars['String'];
}>;


export type PokemonQuery = { __typename?: 'Query', pokemon: { __typename: 'IllegalArguments', message: string } | { __typename: 'Pokemon', nationalNo: number, name: string, imageURL: string, species: string, weight: string, height: string, genders: Array<{ __typename?: 'Gender', name: string, iconURL: string }>, types: Array<{ __typename?: 'Type', name: string, iconURL: string }>, characteristics: Array<{ __typename?: 'Characteristic', name: string, description: string }>, description: { __typename?: 'Description', text: string, series: string }, ability: { __typename?: 'Ability', heart: number, attack: number, defense: number, specialAttack: number, specialDefense: number, speed: number } } | { __typename: 'PokemonNotFound', message: string }, pageInfo: { __typename: 'IllegalArguments' } | { __typename: 'PageInfo', hasPrev: boolean, hasNext: boolean, prevId: string, nextId: string } | { __typename: 'PokemonNotFound' }, evolutions: { __typename: 'Evolutions', pokemons: Array<{ __typename?: 'Pokemon', id: string, nationalNo: number, name: string, imageURL: string, canEvolution: boolean, types: Array<{ __typename?: 'Type', name: string, iconURL: string }> }> } | { __typename: 'IllegalArguments' } | { __typename: 'PokemonNotFound' } };

export type PokemonsQueryVariables = Exact<{
  first?: InputMaybe<Scalars['Int']>;
  after?: InputMaybe<Scalars['String']>;
}>;


export type PokemonsQuery = { __typename?: 'Query', pokemons: { __typename: 'IllegalArguments', message: string } | { __typename: 'PokemonConnection', endCursor: string, hasNext: boolean, items: Array<{ __typename?: 'Pokemon', id: string, name: string, imageURL: string }> } | { __typename: 'PokemonNotFound', message: string } };


export const PokemonDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"pokemon"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pokemon"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pokemonId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"__typename"}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Pokemon"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"nationalNo"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageURL"}},{"kind":"Field","name":{"kind":"Name","value":"species"}},{"kind":"Field","name":{"kind":"Name","value":"weight"}},{"kind":"Field","name":{"kind":"Name","value":"height"}},{"kind":"Field","name":{"kind":"Name","value":"genders"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconURL"}}]}},{"kind":"Field","name":{"kind":"Name","value":"types"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconURL"}}]}},{"kind":"Field","name":{"kind":"Name","value":"characteristics"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"description"}}]}},{"kind":"Field","name":{"kind":"Name","value":"description"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"text"}},{"kind":"Field","name":{"kind":"Name","value":"series"}}]}},{"kind":"Field","name":{"kind":"Name","value":"ability"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"heart"}},{"kind":"Field","name":{"kind":"Name","value":"attack"}},{"kind":"Field","name":{"kind":"Name","value":"defense"}},{"kind":"Field","name":{"kind":"Name","value":"specialAttack"}},{"kind":"Field","name":{"kind":"Name","value":"specialDefense"}},{"kind":"Field","name":{"kind":"Name","value":"speed"}}]}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"IllegalArguments"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"PokemonNotFound"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"pageInfo"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pokemonId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"__typename"}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"PageInfo"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hasPrev"}},{"kind":"Field","name":{"kind":"Name","value":"hasNext"}},{"kind":"Field","name":{"kind":"Name","value":"prevId"}},{"kind":"Field","name":{"kind":"Name","value":"nextId"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"evolutions"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pokemonId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pokemonId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"__typename"}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Evolutions"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pokemons"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"nationalNo"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageURL"}},{"kind":"Field","name":{"kind":"Name","value":"types"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"iconURL"}}]}},{"kind":"Field","name":{"kind":"Name","value":"canEvolution"}}]}}]}}]}}]}}]} as unknown as DocumentNode<PokemonQuery, PokemonQueryVariables>;
export const PokemonsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"pokemons"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"first"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"after"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pokemons"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"first"},"value":{"kind":"Variable","name":{"kind":"Name","value":"first"}}},{"kind":"Argument","name":{"kind":"Name","value":"after"},"value":{"kind":"Variable","name":{"kind":"Name","value":"after"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"__typename"}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"PokemonConnection"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"endCursor"}},{"kind":"Field","name":{"kind":"Name","value":"hasNext"}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"imageURL"}}]}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"IllegalArguments"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}},{"kind":"InlineFragment","typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"PokemonNotFound"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"message"}}]}}]}}]}}]} as unknown as DocumentNode<PokemonsQuery, PokemonsQueryVariables>;