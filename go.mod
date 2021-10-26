module github.com/davcrypto/chain-indexing-app-example

go 1.16

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/rs/zerolog => github.com/rs/zerolog v1.23.0

replace github.com/crypto-com/chain-indexing => github.com/crypto-com/chain-indexing v0.0.0-20211026064551-da00242df23b

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/crypto-com/chain-indexing v0.0.0-00010101000000-000000000000
	github.com/valyala/fasthttp v1.17.0
)
