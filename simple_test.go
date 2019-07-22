package scrypt_playground
import "testing"

func TestBasic(t *testing.T) {
	total := 5 + 5
	if 9 != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
