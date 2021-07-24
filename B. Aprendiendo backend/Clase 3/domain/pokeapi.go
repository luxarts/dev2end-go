package domain

type PokeAPIGetPokemonByNameResponse struct {
	Types []pokeType `json:"types"`
}

type pokeType struct {
	Type TypeName`json:"type"`
}

type TypeName struct {
	Name string	`json:"name"`
}

type DamageRelations struct {
	DoubleDamageFrom []TypeName `json:"double_damage_from"`
	DoubleDamageTo []TypeName `json:"double_damage_to"`
	HalfDamageFrom []TypeName `json:"half_damage_from"`
	HalfDamageTo []TypeName `json:"half_damage_to"`
	NoDamageFrom []TypeName `json:"no_damage_from"`
	NoDamageTo []TypeName `json:"no_damage_to"`
}

type PokeAPIGetTypeResponse struct {
	DamageRelations DamageRelations `json:"damage_relations"`
}