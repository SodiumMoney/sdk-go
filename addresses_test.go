package sodium

import "testing"

func TestConfigAddress(t *testing.T) {
	addr, err := GetConfigAddress(Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != "3WYJaf1ocPSbQAMjFhka5t8iusdH11Atw6FcxeuVYsTw" {
		t.Fatalf("got %s", addr)
	}
}

func TestPaymentBlknoiz06(t *testing.T) {
	addr, err := GetPaymentAddress(PlatformTwitter, "@blknoiz06", Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != "HPT9k5YWkKNQcZGewGmABjRu8A5Bnem6VsBmNrNsSgyd" {
		t.Fatalf("got %s", addr)
	}
}

func TestPaymentAlias(t *testing.T) {
	a, _ := GetPaymentAddress(PlatformTwitter, "@blknoiz06", Mainnet)
	b, _ := GetPaymentAddress(PlatformTwitter, "@BLKNOIZ06", Mainnet)
	if a != b {
		t.Fatal("alias mismatch")
	}
}
