package webrtc

import (
	"github.com/pion/rtp"
)

// Track represents MediaStreamTrack.
type Track interface {
	ID() string
}

// LocalTrack represents MediaStreamTrack which is fed by the local stream source.
type LocalTrack interface {
	Track
	WriteRTP(*rtp.Packet) error
}

// RemoteTrack represents MediaStreamTrack which is fed by the remote peer.
type RemoteTrack interface {
	Track
	ReadRTP() (*rtp.Packet, error)
}

// RTPSender represents RTCRtpSender.
type RTPSender interface {
	ReplaceTrack(LocalTrack) error
}

// RTPReceiver represents RTCRtpReceiver.
type RTPReceiver interface {
}

// RTPTransceiver represents RTCRtpTransceiver.
type RTPTransceiver interface {
	RTPSender
	RTPReceiver
}
