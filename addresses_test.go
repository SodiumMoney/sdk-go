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
