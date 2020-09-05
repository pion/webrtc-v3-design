package webrtc

// PeerConnection represents RTCPeerConnection.
type PeerConnection interface {
	// AddTransceiverFromTrack creates a new RTPTransceiver from LocalTrack
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be a bidirectional stream by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromTrack(LocalRTPTrack, *RTPTransceiverInit) (RTPTransceiver, error)

	// AddTransceiverFromKind creates a new RTPTransceiver from RTPCodecType
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be a bidirectional stream by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromKind(RTPCodecKind, *RTPTransceiverInit) (RTPTransceiver, error)
}

// RTPTransceiverInit represents RTCRtpTransceiverInit dictionary.
type RTPTransceiverInit struct {
	Direction     RTPTransceiverDirection
	SendEncodings []RTPParameters
}
