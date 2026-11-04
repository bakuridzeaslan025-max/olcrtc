//go:build !olcrtc_vkcalls

package mobile

const (
	// salutejazz is the provider the lean bind once lost (olcbox#22).
	providerSureToRegister = "salutejazz"
	// testProvider is the provider the runtime tests configure.
	testProvider = "jitsi"
)
