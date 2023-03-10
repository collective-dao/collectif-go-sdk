package fvm

import (
	"collective-go-sdk/config"
	"context"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLotusClient(t *testing.T) {
	ctx := context.Background()
	config, err := config.LoadConfig("../config")

	if err != nil {
		assert.Error(t, err)
	}

	client, err := NewLotusClient(config)

	chainId, err := client.client.ChainID(ctx)
	if err != nil {
		assert.Error(t, err)
	}

	assert.Equal(t, big.NewInt(3141), chainId)
	assert.Equal(t, config.RPCAddress, client.host)
}
