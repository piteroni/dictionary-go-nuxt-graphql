import { TypedDocumentNode as DocumentNode } from "@graphql-typed-document-node/core"
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

export type Ability = {
  __typename?: "Ability";
  attack: Scalars["Int"];
  defense: Scalars["Int"];
  heart: Scalars["Int"];
  specialAttack: Scalars["Int"];
  specialDefense: Scalars["Int"];
  speed: Scalars["Int"];
};

export type Characteristic = {
  __typename?: "Characteristic";
  description: Scalars["String"];
  name: Scalars["String"];
};

export type Description = {
  __typename?: "Description";
  series: Scalars["String"];
  text: Scalars["String"];
};

export type Gender = {
  __typename?: "Gender";
  iconURL: Scalars["String"];
  name: Scalars["String"];
};

export type LinkInfo = {
  __typename?: "LinkInfo";
  hasNext: Scalars["Boolean"];
  hasPrev: Scalars["Boolean"];
  nextNationalNo: Scalars["Int"];
  prevNationalNo: Scalars["Int"];
};

export type Pokemon = {
  __typename?: "Pokemon";
  ability: Ability;
  canEvolution: Scalars["Boolean"];
  characteristics: Array<Characteristic>;
  description: Description;
  evolutions: Array<Pokemon>;
  genders: Array<Gender>;
  height: Scalars["String"];
  imageURL: Scalars["String"];
  linkInfo: LinkInfo;
  name: Scalars["String"];
  nationalNo: Scalars["Int"];
  species: Scalars["String"];
  types: Array<Type>;
  weight: Scalars["String"];
};

export type Query = {
  __typename?: "Query";
  pokemon: Pokemon;
};

export type QueryPokemonArgs = {
  pokemonId: Scalars["Int"];
};

export type Type = {
  __typename?: "Type";
  iconURL: Scalars["String"];
  name: Scalars["String"];
};

export type PokemonQueryVariables = Exact<{
  pokemonId: Scalars["Int"];
}>;

export type PokemonQuery = { __typename?: "Query", pokemon: { __typename?: "Pokemon", nationalNo: number, name: string, imageURL: string, species: string, weight: string, height: string, genders: Array<{ __typename?: "Gender", name: string, iconURL: string }>, types: Array<{ __typename?: "Type", name: string, iconURL: string }>, characteristics: Array<{ __typename?: "Characteristic", name: string, description: string }>, description: { __typename?: "Description", text: string, series: string }, ability: { __typename?: "Ability", heart: number, attack: number, defense: number, specialAttack: number, specialDefense: number, speed: number }, linkInfo: { __typename?: "LinkInfo", prevNationalNo: number, nextNationalNo: number, hasPrev: boolean, hasNext: boolean }, evolutions: Array<{ __typename?: "Pokemon", nationalNo: number, name: string, imageURL: string, canEvolution: boolean, types: Array<{ __typename?: "Type", name: string, iconURL: string }> }> } };

export const PokemonDocument = { "kind": "Document", "definitions": [{ "kind": "OperationDefinition", "operation": "query", "name": { "kind": "Name", "value": "pokemon" }, "variableDefinitions": [{ "kind": "VariableDefinition", "variable": { "kind": "Variable", "name": { "kind": "Name", "value": "pokemonId" } }, "type": { "kind": "NonNullType", "type": { "kind": "NamedType", "name": { "kind": "Name", "value": "Int" } } } }], "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "pokemon" }, "arguments": [{ "kind": "Argument", "name": { "kind": "Name", "value": "pokemonId" }, "value": { "kind": "Variable", "name": { "kind": "Name", "value": "pokemonId" } } }], "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "nationalNo" } }, { "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "imageURL" } }, { "kind": "Field", "name": { "kind": "Name", "value": "species" } }, { "kind": "Field", "name": { "kind": "Name", "value": "weight" } }, { "kind": "Field", "name": { "kind": "Name", "value": "height" } }, { "kind": "Field", "name": { "kind": "Name", "value": "genders" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "iconURL" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "types" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "iconURL" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "characteristics" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "description" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "description" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "text" } }, { "kind": "Field", "name": { "kind": "Name", "value": "series" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "ability" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "heart" } }, { "kind": "Field", "name": { "kind": "Name", "value": "attack" } }, { "kind": "Field", "name": { "kind": "Name", "value": "defense" } }, { "kind": "Field", "name": { "kind": "Name", "value": "specialAttack" } }, { "kind": "Field", "name": { "kind": "Name", "value": "specialDefense" } }, { "kind": "Field", "name": { "kind": "Name", "value": "speed" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "linkInfo" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "prevNationalNo" } }, { "kind": "Field", "name": { "kind": "Name", "value": "nextNationalNo" } }, { "kind": "Field", "name": { "kind": "Name", "value": "hasPrev" } }, { "kind": "Field", "name": { "kind": "Name", "value": "hasNext" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "evolutions" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "nationalNo" } }, { "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "imageURL" } }, { "kind": "Field", "name": { "kind": "Name", "value": "types" }, "selectionSet": { "kind": "SelectionSet", "selections": [{ "kind": "Field", "name": { "kind": "Name", "value": "name" } }, { "kind": "Field", "name": { "kind": "Name", "value": "iconURL" } }] } }, { "kind": "Field", "name": { "kind": "Name", "value": "canEvolution" } }] } }] } }] } }] } as unknown as DocumentNode<PokemonQuery, PokemonQueryVariables>
