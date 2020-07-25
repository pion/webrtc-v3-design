package webrtc

import (
	"github.com/pion/webrtc/v3"
)

// PeerConnection represents RTCPeerConnection.
type PeerConnection interface {
	// AddTransceiverFromTrack creates a new RTPTransceiver from LocalTrack
	// and register it to the PeerConnection.
	// RTPTransceiver represents a bidirectional stream.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromTrack(LocalTrack, ...RTPTransceiverInit) (RTPTransceiver, error)
	// AddTransceiverFromTrack creates a new RTPTransceiver from RTPCodecType
	// and register it to the PeerConnection.
	// RTPTransceiver represents a bidirectional stream.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromKind(webrtc.RTPCodecType, ...RTPTransceiverInit) (RTPTransceiver, error)
}

// RTPTransceiverInitDict represents RTCRtpTransceiverInit dictionary.
type RTPTransceiverInitDict struct {
	Direction     webrtc.RTPTransceiverDirection
	SendEncodings []RTPParameters
}

// RTPTransceiverInit is a functional option type to set RTPTransceiverInitDict to AddTransceiverFromKind/Track.
type RTPTransceiverInit func(*RTPTransceiverInitDict)

// RTPTransceiverInitSendonly sets RTPTransceriver direction to sendonly.
func RTPTransceiverInitSendonly() RTPTransceiverInit {
	return func(d *RTPTransceiverInitDict) {
		d.Direction = webrtc.RTPTransceiverDirectionSendonly
	}
}

// RTPTransceiverInitSendrecv sets RTPTransceriver direction to sendrecv.
func RTPTransceiverInitSendrecv() RTPTransceiverInit {
	return func(d *RTPTransceiverInitDict) {
		d.Direction = webrtc.RTPTransceiverDirectionSendrecv
	}
}

// RTPTransceiverInitRecvonly sets RTPTransceriver direction to recvonly.
func RTPTransceiverInitRecvonly() RTPTransceiverInit {
	return func(d *RTPTransceiverInitDict) {
		d.Direction = webrtc.RTPTransceiverDirectionRecvonly
	}
}

// RTPTransceiverInitSendEncodings sets a list of encodings to allow when sending RTP media.
func RTPTransceiverInitSendEncodings(encodings ...RTPParameters) RTPTransceiverInit {
	return func(d *RTPTransceiverInitDict) {
		d.SendEncodings = encodings
	}
}
