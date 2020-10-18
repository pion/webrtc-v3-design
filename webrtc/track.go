package webrtc

import (
	"errors"

	"github.com/pion/rtp"
)

var (
	// ErrIncompatible is an internal error returned from TrackLocal/RemoteAdapter.
	ErrIncompatible = errors.New("incompatible pair of Track and RTPSender/Receiver")
)

// TrackLocalContext is the Context passed when a TrackLocal has been Binded/Unbinded from a PeerConnection
type TrackLocalContext struct{}

// Parameters returns the negotiated Parameters. These are the codecs supported by both
// PeerConnections and the SSRC/PayloadTypes
func (t *TrackLocalContext) Parameters() RTPParameters {
	return RTPParameters{}
}

// TrackLocalWriter is the Writer for outbound RTP Packets
type TrackLocalWriter interface {
	// WriteRTP encrypts a RTP packet and writes to the connection
	WriteRTP(header *rtp.Header, payload []byte) (int, error)

	// Write encrypts and writes a full RTP packet
	Write(b []byte) (int, error)
}

// WriteStream returns the WriteStream for this TrackLocal. The implementer writes the outbound
// media packets to it
func (t *TrackLocalContext) WriteStream() TrackLocalWriter {
	return nil
}

// TrackLocal is an interface that controls how the user can send media
// The user can provide their own TrackLocal implementatiosn, or use
// the implementations in pkg/media
type TrackLocal interface {
	// Bind should implement the way how the media data flows from the Track to the PeerConnection
	// This will be called internally after signaling is complete and the list of available
	// codecs has been determined
	Bind(TrackLocalContext) error

	// Unbind should implement the teardown logic when the track is no longer needed. This happens
	// because a track has been stopped.
	Unbind(TrackLocalContext) error

	// ID is the unique identifier for this Track. This should be unique for the
	// stream, but doesn't have to globally unique. A common example would be 'audio' or 'video'
	// and StreamID would be 'desktop' or 'webcam'
	ID() string

	// StreamID is the group this track belongs too. This must be unique
	StreamID() string
}

// Track represents a single media track from a remote peer
type TrackRemote struct{}

// Read reads data from the track.
func (t *TrackRemote) Read(b []byte) (n int, err error) { return }

// ReadRTP is a convenience method that wraps Read and unmarshals for you
func (t *TrackRemote) ReadRTP() (*rtp.Packet, error) { return nil, nil }
