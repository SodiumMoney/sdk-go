package main

import (
	"fmt"
	sodium "github.com/SodiumMoney/sdk-go"
)

func main() {
	addr, err := sodium.GetPaymentAddress(sodium.PlatformTwitter, "@blknoiz06", sodium.Mainnet)
	if err != nil {
		panic(err)
	}
	fmt.Println(addr)
}
