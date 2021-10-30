package utils

import "math/big"

// ToEck number of FTM to Wei
func ToEck(ftm uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(ftm), big.NewInt(1e18))
}
