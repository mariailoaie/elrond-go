package process

import (
	"github.com/ElrondNetwork/elrond-go/genesis"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

// GenesisParser contains the parsed genesis json file and has some functionality regarding processed data
type GenesisParser interface {
	InitialAccountsSplitOnAddressesShards(shardCoordinator sharding.Coordinator) (map[uint32][]*genesis.InitialAccount, error)
	InitialAccountsSplitOnDelegationAddressesShards(shardCoordinator sharding.Coordinator) (map[uint32][]*genesis.InitialAccount, error)
	IsInterfaceNil() bool
}

// InitialNodesHandler contains the initial nodes setup
type InitialNodesHandler interface {
	InitialNodesInfo() (map[uint32][]sharding.GenesisNodeInfoHandler, map[uint32][]sharding.GenesisNodeInfoHandler)
	IsInterfaceNil() bool
}
