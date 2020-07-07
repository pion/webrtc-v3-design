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

// RTPParameters represents RTCRtpParameters which contains information about
// how the RTC data is to be encoded/decoded.
//
// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCRtpSendParameters
type RTPParameters struct {
	SSRC        uint32
	PayloadType uint16
	Encodings   []*webrtc.RTPCodec
}

// RTPSender represents RTCRtpSender.
type RTPSender interface {
	ReplaceTrack(LocalTrack) error

	// Parameters returns information about how the data is to be encoded.
	Parameters() RTPParameters
	// SetParameters sets information about how the data is to be encoded.
	// This will be called by PeerConnection according to the result of
	// SDP based negotiation.
	SetParameters(RTPParameters) error
}

// RTPReceiver represents RTCRtpReceiver.
type RTPReceiver interface {
	// Parameters returns information about how the data is to be decoded.
	Parameters() RTPParameters
}

// RTPTransceiver represents RTCRtpTransceiver.
type RTPTransceiver interface {
	RTPSender
	RTPReceiver
}
