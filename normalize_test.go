package sodium

import "testing"

func TestNormalizeUsername(t *testing.T) {
	if got := NormalizeUsername("@blknoiz06"); got != "blknoiz06" {
		t.Fatalf("got %q", got)
	}
	if got := NormalizeUsername("Ansem 🐂🀄️"); got != "ansem 🐂🀄️" {
		t.Fatalf("got %q", got)
	}
}
