package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "leverage"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// KVStore key prefixes
var (
	KeyPrefixTokenDenom        = []byte{0x01}
	KeyPrefixUTokenDenom       = []byte{0x02}
	KeyPrefixRegisteredToken   = []byte{0x03}
	KeyPrefixLoanToken         = []byte{0x04}
	KeyPrefixCollateralSetting = []byte{0x05}
	KeyPrefixReserveAmount     = []byte{0x06}
	KeyPrefixLastInterestTime  = []byte{0x07}
	KeyPrefixExchangeRate      = []byte{0x08}
)

// CreateTokenDenomKey returns a KVStore key for getting and storing a token's
// associated uToken denomination.
func CreateTokenDenomKey(tokenDenom string) []byte {
	// tokenprefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixTokenDenom...)
	key = append(key, []byte(tokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateUTokenDenomKey returns a KVStore key for getting and storing a uToken's
// associated token denomination.
func CreateUTokenDenomKey(uTokenDenom string) []byte {
	// utokenprefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixUTokenDenom...)
	key = append(key, []byte(uTokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateRegisteredTokenKey returns a KVStore key for getting and setting a Token.
func CreateRegisteredTokenKey(baseTokenDenom string) []byte {
	// assetprefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixRegisteredToken...)
	key = append(key, []byte(baseTokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateLoanKey returns a KVStore key for getting and setting a Loan for a denom
// and borrower address.
func CreateLoanKey(borrowerAddr sdk.AccAddress, tokenDenom string) []byte {
	// loanprefix | lengthprefixed(borrowerAddr) | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixLoanToken...)
	key = append(key, address.MustLengthPrefix(borrowerAddr)...)
	key = append(key, []byte(tokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateLoanKeyNoDenom returns the common prefix used by all loans associated with a given
// borrower address.
func CreateLoanKeyNoDenom(borrowerAddr sdk.AccAddress) []byte {
	// loanprefix | lengthprefixed(borrowerAddr)
	var key []byte
	key = append(key, KeyPrefixLoanToken...)
	key = append(key, address.MustLengthPrefix(borrowerAddr)...)
	return key
}

// CreateLoanKeyNoAddress returns a safe copy of loanprefix
func CreateLoanKeyNoAddress() []byte {
	// loanprefix
	var key []byte
	key = append(key, KeyPrefixLoanToken...)
	return key
}

// CreateCollateralSettingKey returns a KVStore key for getting and setting a borrower's
// collateral setting for a single uToken
func CreateCollateralSettingKey(borrowerAddr sdk.AccAddress, uTokenDenom string) []byte {
	// collatprefix | lengthprefixed(borrowerAddr) | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixCollateralSetting...)
	key = append(key, address.MustLengthPrefix(borrowerAddr)...)
	key = append(key, []byte(uTokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateReserveAmountKey returns a KVStore key for getting and setting the amount reserved of a a given token.
func CreateReserveAmountKey(tokenDenom string) []byte {
	// reserveamountprefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixReserveAmount...)
	key = append(key, []byte(tokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateLastInterestTimeKey returns a KVStore key for getting and setting the amount reserved of a a given token.
func CreateLastInterestTimeKey() []byte {
	// lastinterestprefix
	var key []byte
	key = append(key, KeyPrefixLastInterestTime...)
	return key
}

// CreateExchangeRateKey returns a KVStore key for getting and setting the token:uToken rate for a a given token.
func CreateExchangeRateKey(tokenDenom string) []byte {
	// exchangeRatePrefix | denom | 0x00
	var key []byte
	key = append(key, KeyPrefixExchangeRate...)
	key = append(key, []byte(tokenDenom)...)
	return append(key, 0) // append 0 for null-termination
}

// CreateExchangeRateKeyNoDenom returns a safe copy of exchangeRatePrefix
func CreateExchangeRateKeyNoDenom() []byte {
	// exchangeRatePrefix
	var key []byte
	key = append(key, KeyPrefixExchangeRate...)
	return key
}
