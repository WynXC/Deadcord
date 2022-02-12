package core

import (
	"math/rand"
)

var (
	SpamFlag         = 0
	DeadcordVersion  = "1.0"
	RawTokensLoaded  []string
	BuiltTokenStruct map[int]map[string]string
)

func SetTokens(raw_token_list []string, built_token_list map[int]map[string]string) int {
	RawTokensLoaded = raw_token_list
	BuiltTokenStruct = built_token_list

	return len(RawTokensLoaded)
}

func RandomToken() string {
	random_token := RawTokensLoaded[rand.Intn(len(RawTokensLoaded))]
	return random_token
}
