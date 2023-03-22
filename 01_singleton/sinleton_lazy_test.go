package singleton_test

import (
	"github.com/stretchr/testify/assert"
	"singleton"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, singleton.GetLazyInstance(), singleton.GetLazyInstance())
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		if singleton.GetInstance() != singleton.GetLazyInstance() {
			b.Errorf("test fail")
		}
	})
}
