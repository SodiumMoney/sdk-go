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

func TestPaymentUnicode(t *testing.T) {
	addr, err := GetPaymentAddress(PlatformTwitter, "Ansem 🐂🀄️", Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != "k7JEBhVHHWLrNdhJ87cPbpgJr98kDVQuqUb5HsVKW7S" {
		t.Fatalf("got %s", addr)
	}
}

func TestIdentityBlknoiz06(t *testing.T) {
	addr, err := GetIdentityAddress(PlatformTwitter, "@blknoiz06", Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != "T4gKVCxkB3c6pERaS67fLBgMyg1FpGmjeNcrGvbuLY1" {
		t.Fatalf("got %s", addr)
	}
}

func TestIdentityUnicode(t *testing.T) {
	addr, err := GetIdentityAddress(PlatformTwitter, "Ansem 🐂🀄️", Mainnet)
	if err != nil {
		t.Fatal(err)
	}
	if addr.String() != "83McthMovBoPHgXPVGsif61zkCZNEE7tmypdNkxh1cv5" {
		t.Fatalf("got %s", addr)
	}
}
