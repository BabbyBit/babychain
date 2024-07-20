package keeper

import (
	"math"

	cosmosmath "cosmossdk.io/math"
)

const basePrice = 10
const priceCoefficient = 1.09

func calculateIncome(level int64, exponent uint32) cosmosmath.Int {
	return cosmosmath.NewInt(level * int64(math.Pow10(int(exponent))))
}

func calculateLevelPrice(level int64, exponent uint32) cosmosmath.Int {
	// Цена = Базовая цена × 1.09^Сделано улучшений = 10*1.09^Level
	price := basePrice * math.Pow(priceCoefficient, float64(level))
	return cosmosmath.NewInt(int64(price) * int64(math.Pow10(int(exponent))))
}
