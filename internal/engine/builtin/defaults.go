//go:build !olcrtc_vkcalls

package builtin

import (
	authJitsi "github.com/openlibrecommunity/olcrtc/internal/auth/jitsi"
	authMaxCalls "github.com/openlibrecommunity/olcrtc/internal/auth/maxcalls"
	authMtsLink "github.com/openlibrecommunity/olcrtc/internal/auth/mtslink"
	authSaluteJazz "github.com/openlibrecommunity/olcrtc/internal/auth/salutejazz"
	authTelemost "github.com/openlibrecommunity/olcrtc/internal/auth/telemost"
	authVKCalls "github.com/openlibrecommunity/olcrtc/internal/auth/vkcalls"
	authWBStream "github.com/openlibrecommunity/olcrtc/internal/auth/wbstream"
	"github.com/openlibrecommunity/olcrtc/internal/engine"
	"github.com/openlibrecommunity/olcrtc/internal/engine/goolom"
	engineJitsi "github.com/openlibrecommunity/olcrtc/internal/engine/jitsi"
	engineMaxCalls "github.com/openlibrecommunity/olcrtc/internal/engine/maxcalls"
	engineMtsLink "github.com/openlibrecommunity/olcrtc/internal/engine/mtslink"
	engineSaluteJazz "github.com/openlibrecommunity/olcrtc/internal/engine/salutejazz"
	engineVKCalls "github.com/openlibrecommunity/olcrtc/internal/engine/vkcalls"
)

// RegisterDefaults wires the built-in providers: jitsi, telemost, wbstream,
// salutejazz, vkcalls, mtslink, max and "none" (direct engine access).
func RegisterDefaults() {
	registerLivekitEngine()
	engine.Register("goolom", goolom.New)
	engine.Register("jitsi", engineJitsi.New)
	engine.Register("salutejazz", engineSaluteJazz.New)
	engine.Register("vkcalls", engineVKCalls.New)
	engine.Register("mtslink", engineMtsLink.New)
	engine.Register("max", engineMaxCalls.New)
	register("wbstream", authWBStream.Provider{})
	register("telemost", authTelemost.Provider{})
	register("jitsi", authJitsi.Provider{})
	register("salutejazz", authSaluteJazz.New())
	register("vkcalls", authVKCalls.New())
	register("mtslink", authMtsLink.New())
	register("max", authMaxCalls.Provider{})
	register("none", nil)
}
