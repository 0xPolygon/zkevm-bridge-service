go run main.go deployClaimCompressor --url "http://localhost:8123" --bridgeAddress "0xFe12ABaa190Ef0c8638Ee0ba9F828BF41368Ca0E" --walletFile ../../test.keystore.claimtx --password "testonly"

go run main.go sendETH --url "http://localhost:8123" --destAddress "0x70997970c51812dc3a010c7d01b50e0d17dc79c8" --walletFile ../../test.keystore.claimtx --password "testonly"

go run main.go deploySovereignChainSMC --url "http://localhost:8123" --walletFile ../../test.keystore.aggregator --password "testonly"
