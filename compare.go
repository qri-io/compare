// compare implements comparison between common go types
package compare

import (
	"fmt"
)

func StringSlice(a, b []string) error {
	if a == nil && b == nil {
		return nil
	}
	if a == nil && b != nil || a != nil && b == nil {
		return fmt.Errorf("nil mismatch: %s != %s", a, b)
	}
	if len(a) != len(b) {
		return fmt.Errorf("length mismatch: %d != %d", len(a), len(b))
	}
	for i, s := range a {
		if b[i] != s {
			return fmt.Errorf("index %d mismatch: %s != %s", i, s, b[i])
		}
	}
	return nil
}
