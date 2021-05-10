package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AttoTokn defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	AttoTokn string = "atokn"

	// BaseDenomUnit defines the base denomination unit for Tokns.
	// 1 tokn = 1x10^{BaseDenomUnit} atokn
	BaseDenomUnit = 18
)

// NewToknCoin is a utility function that returns an "atokn" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewToknCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(AttoTokn, amount)
}

// NewToknDecCoin is a utility function that returns an "atokn" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewToknDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(AttoTokn, amount)
}

// NewToknCoinInt64 is a utility function that returns an "atokn" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewToknCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(AttoTokn, amount)
}
