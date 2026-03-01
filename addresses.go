package sodium

import "github.com/gagliardetto/solana-go"

func GetConfigAddress(network Network) (solana.PublicKey, error) {
	return findPDA([][]byte{[]byte("config")}, network)
}

func GetPaymentAddress(platform Platform, username string, network Network) (solana.PublicKey, error) {
	norm := NormalizeUsername(username)
	seeds := [][]byte{
		[]byte("escrow"),
		[]byte(string(platform)),
		[]byte(norm),
	}
	return findPDA(seeds, network)
}
