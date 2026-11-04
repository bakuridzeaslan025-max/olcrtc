//go:build !olcrtc_vkcalls

package client_test

// registeredBy narrows a full build's provider or engine names to this build's.
func registeredBy(names []string) []string { return names }
