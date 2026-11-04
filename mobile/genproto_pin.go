//go:build !olcrtc_vkcalls

package mobile

// keep gomobile on post-split genproto modules
//
// Not under olcrtc_vkcalls: that build links no grpc or livekit, so nothing
// needs the pin, and the generated package would bring the protobuf runtime,
// whose reflection makes the linker keep every exported method of every
// reachable type.
import _ "google.golang.org/genproto/protobuf/field_mask"
