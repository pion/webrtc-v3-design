package webrtc

import (
	"errors"
)

var (
	// ErrIncompatible is an internal error returned from TrackLocal/RemoteAdapter.
	ErrIncompatible = errors.New("incompatible pair of Track and RTPSender/Receiver")
)

// TrackLocal is an interface that controls how the user can send media
// The user can provide their own TrackLocal implementatiosn, or use
// the implementations in pkg/media
type TrackLocal interface {
	// Bind should implement the way how the media data flows from the Track to the PeerConnection
	// This will be called internally after signaling is complete and the list of available
	// codecs has been determined
	Bind() error

	// Unbind should implement the teardown logic when the track is no longer needed. This happens
	// because a track has been stopped.
	Unbind() error
}

// Track represents a single media track from a remote peer
type TrackRemote struct{}

// Read reads data from the track.
func (t *TrackRemote) Read(b []byte) (n int, err error) { return }
