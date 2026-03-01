package sodium

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

func findPDA(seeds [][]byte, network Network) (solana.PublicKey, error) {
	programID := solana.MustPublicKeyFromBase58(GetProgramID(network))
	pda, _, err := solana.FindProgramAddress(seeds, programID)
	return pda, err
}

func slotSeed(slot uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, slot)
	return b
}
