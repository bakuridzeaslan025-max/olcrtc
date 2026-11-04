//go:build olcrtc_vkcalls

package client_test

// registeredBy narrows a full build's provider or engine names to this
// build's: the olcrtc_vkcalls build registers vkcalls alone, no "none".
func registeredBy([]string) []string { return []string{"vkcalls"} }
