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
