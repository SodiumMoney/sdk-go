package sodium

import "github.com/gagliardetto/solana-go"

func GetConfigAddress(network Network) (solana.PublicKey, error) {
	return findPDA([][]byte{[]byte("config")}, network)
}

func GetPaymentAddress(platform Platform, username string, network Network) (solana.PublicKey, error) {
	norm := NormalizeUsername(username)
	return findPDA([][]byte{[]byte("escrow"), []byte(string(platform)), []byte(norm)}, network)
}

func GetIdentityAddress(platform Platform, username string, network Network) (solana.PublicKey, error) {
	norm := NormalizeUsername(username)
	return findPDA([][]byte{[]byte("identity"), []byte(string(platform)), []byte(norm)}, network)
}

func GetUsedProofAddress(proofHash []byte, network Network) (solana.PublicKey, error) {
	return findPDA([][]byte{[]byte("used_proof"), proofHash}, network)
}

func GetLinkRequestAddress(usernameHash []byte, wallet solana.PublicKey, slot uint64, network Network) (solana.PublicKey, error) {
	seeds := [][]byte{[]byte("link_req"), usernameHash, wallet.Bytes(), slotSeed(slot)}
	return findPDA(seeds, network)
}

func GetTelegramLinkRequestAddress(usernameHash []byte, wallet solana.PublicKey, slot uint64, network Network) (solana.PublicKey, error) {
	seeds := [][]byte{[]byte("tg_link_req"), usernameHash, wallet.Bytes(), slotSeed(slot)}
	return findPDA(seeds, network)
}
