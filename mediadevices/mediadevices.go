package mediadevices

import (
	mediadevices2 "github.com/pion/mediadevices"
	"github.com/pion/webrtc-v3-design/webrtc"
)

// MediaDevices is an interface that's defined on https://developer.mozilla.org/en-US/docs/Web/API/MediaDevices
type MediaDevices interface {
	GetDisplayMedia(constraints mediadevices2.MediaStreamConstraints) (MediaStream, error)
	GetUserMedia(constraints mediadevices2.MediaStreamConstraints) (MediaStream, error)
	EnumerateDevices() []mediadevices2.MediaDeviceInfo
}

// NewMediaDevices creates mediadevices interface.
func NewMediaDevices(opts ...mediadevices2.MediaDevicesOption) MediaDevices {
	panic("not implementd")
}

// MediaStream is an interface that represents a collection of existing tracks.
type MediaStream interface {
	// GetAudioTracks implements https://w3c.github.io/mediacapture-main/#dom-mediastream-getaudiotracks
	GetAudioTracks() []webrtc.TrackLocal
	// GetVideoTracks implements https://w3c.github.io/mediacapture-main/#dom-mediastream-getvideotracks
	GetVideoTracks() []webrtc.TrackLocal
	// GetTracks implements https://w3c.github.io/mediacapture-main/#dom-mediastream-gettracks
	GetTracks() []webrtc.TrackLocal
	// AddTrack implements https://w3c.github.io/mediacapture-main/#dom-mediastream-addtrack
	AddTrack(t webrtc.TrackLocal)
	// RemoveTrack implements https://w3c.github.io/mediacapture-main/#dom-mediastream-removetrack
	RemoveTrack(t webrtc.TrackLocal)
}
