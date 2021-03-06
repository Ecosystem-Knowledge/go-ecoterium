package opera

import (
	"math/big"

	"github.com/Ecosystem-Knowledge/lachesis-base/hash"
	"github.com/Ecosystem-Knowledge/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/Ecosystem-Knowledge/go-ecoterium/inter"
	"github.com/Ecosystem-Knowledge/go-ecoterium/opera/genesis"
	"github.com/Ecosystem-Knowledge/go-ecoterium/opera/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
