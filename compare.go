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

func MapStringInterface(a, b map[string]interface{}) error {
	if err := Interface(a, b); err != nil {
		return err
	}

	if len(a) != len(b) {
		return fmt.Errorf("length mismatch: %d != %d", len(a), len(b))
	}

	for key, av := range a {
		if err := Interface(av, b[key]); err != nil {
			return fmt.Errorf("key %s error: %s", key, err.Error())
		}
	}

	return nil
}

func Interface(a, b interface{}) error {
	if a == nil && b == nil {
		return nil
	}

	if a == nil && b != nil || a != nil && b == nil {
		return fmt.Errorf("nil mismatch: %v != %v", a, b)
	}

	// switch on common data types
	switch at := a.(type) {
	case string:
		if bt, ok := b.(string); ok {
			if at != bt {
				return fmt.Errorf("string mismatch: %s != %s", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case bool:
		if bt, ok := b.(bool); ok {
			if at != bt {
				return fmt.Errorf("boolean mismatch: %s != %s", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case int:
		if bt, ok := b.(int); ok {
			if at != bt {
				return fmt.Errorf("int mismatch: %d != %d", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case int32:
		if bt, ok := b.(int32); ok {
			if at != bt {
				return fmt.Errorf("int mismatch: %d != %d", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case int64:
		if bt, ok := b.(int64); ok {
			if at != bt {
				return fmt.Errorf("int mismatch: %d != %d", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case float32:
		if bt, ok := b.(float32); ok {
			if at != bt {
				return fmt.Errorf("int mismatch: %f != %f", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	case float64:
		if bt, ok := b.(float64); ok {
			if at != bt {
				return fmt.Errorf("int mismatch: %f != %f", at, bt)
			}
		} else {
			return fmt.Errorf("type mismatch: %v != %v", a, b)
		}
	}

	return nil
}
