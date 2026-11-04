//go:build olcrtc_vkcalls

package builtin

import (
	authVKCalls "github.com/openlibrecommunity/olcrtc/internal/auth/vkcalls"
	"github.com/openlibrecommunity/olcrtc/internal/engine"
	engineVKCalls "github.com/openlibrecommunity/olcrtc/internal/engine/vkcalls"
)

// RegisterDefaults, under the olcrtc_vkcalls build tag, wires vkcalls and
// nothing else, for an app that embeds olcrtc only to reach VK Calls. The
// other providers are not imported, so their engines and SDKs - livekit's
// protobuf descriptors and CEL validator above all - stay out of the binary.
// Any other provider name, "none" included, fails with ErrProviderNotFound.
func RegisterDefaults() {
	engine.Register("vkcalls", engineVKCalls.New)
	register("vkcalls", authVKCalls.New())
}
