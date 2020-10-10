package validator

import (
	"testing"
)

type Dummy struct {
	Name string `vd:"not_blank"`
	ID   string `vd:"uuidv4"`
}

func BenchmarkFieldSuccess(b *testing.B) {
	// validate := New()

	p := &Dummy{
		ID:   "s",
		Name: "My",
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = Validate(p)
	}
}
