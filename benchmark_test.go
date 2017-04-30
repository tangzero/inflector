package inflector_test

import (
	"testing"

	"github.com/tangzero/inflector"
)

func BenchmarkPluralize(b *testing.B) {
	inflector.ShouldCache = false
	inflector.ClearCache()
	for n := 0; n < b.N; n++ {
		for singular := range SingularToPlural {
			inflector.Pluralize(singular)
		}
	}
}

func BenchmarkCachedPluralize(b *testing.B) {
	inflector.ShouldCache = true
	inflector.ClearCache()
	for n := 0; n < b.N; n++ {
		for singular := range SingularToPlural {
			inflector.Pluralize(singular)
		}
	}
}

func BenchmarkSingularize(b *testing.B) {
	inflector.ShouldCache = false
	inflector.ClearCache()
	for n := 0; n < b.N; n++ {
		for _, plural := range SingularToPlural {
			inflector.Singularize(plural)
		}
	}
}

func BenchmarkCachedSingularize(b *testing.B) {
	inflector.ShouldCache = true
	inflector.ClearCache()
	for n := 0; n < b.N; n++ {
		for _, plural := range SingularToPlural {
			inflector.Singularize(plural)
		}
	}
}
