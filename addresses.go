package sodium

import "github.com/gagliardetto/solana-go"

func GetConfigAddress(network Network) (solana.PublicKey, error) {
	return findPDA([][]byte{[]byte("config")}, network)
}
