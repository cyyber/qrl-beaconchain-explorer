package config

import _ "embed"

//go:embed default.config.yml
var DefaultConfigYml string

//go:embed mainnet.chain.yml
var MainnetChainYml string

//go:embed testnet.chain.yml
var TestnetChainYml string
