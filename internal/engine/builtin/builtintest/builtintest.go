// Package builtintest lets a test that names particular providers skip in a
// build that does not register them (the olcrtc_vkcalls tag).
package builtintest

import (
	"slices"
	"testing"

	"github.com/openlibrecommunity/olcrtc/internal/engine/builtin"
)

// RequireProviders registers the built-in providers and skips t unless every
// one of names is among them.
func RequireProviders(t testing.TB, names ...string) {
	t.Helper()
	builtin.RegisterDefaults()
	available := builtin.Available()
	for _, name := range names {
		if !slices.Contains(available, name) {
			t.Skipf("provider %q is not in this build (have %v)", name, available)
		}
	}
}
