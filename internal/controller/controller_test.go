package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	remote = "https://mainnet-tezos.giganode.io/chains/main/blocks/head"
)

func Test_getheader(t *testing.T) {
	h1, err := getheader(remote)
	assert.NoError(t, err)
	h2, err := getheader(remote)
	assert.NoError(t, err)
	assert.Equal(t, h1.Hash, h2.Hash)
}
