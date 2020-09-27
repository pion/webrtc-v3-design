package webrtc

type Configuration struct{}

type SettingEngine interface {
	// SetEncodings configures the codecs the newly created PeerConnection is willing to
	// send and receive. If nothing is configured it will support everything that Pion WebRTC
	// implements packetization for.
	SetEncodings([]*RTPCodecCapability) error

	// SetTrackLocalAdapters registers TrackLocalAdapters.
	// TrackLocalRTP-RTPSenderPassthrough is registered by default.
	SetTrackLocalAdapters([]TrackLocalAdapter) error

	// SetTrackRemoteAdapters registers TrackRemoteAdapters.
	// TrackRemoteRTP-RTPReceiverPassthrough is registered by default.
	SetTrackRemoteAdapters([]TrackRemoteAdapter) error

	// NewPeerConnection creates a NewPeerConnection
	NewPeerConnection(Configuration) (PeerConnection, error)
}

// PeerConnection represents RTCPeerConnection.
type PeerConnection interface {
	// AddTransceiverFromTrack creates a new RTPTransceiver from TrackLocal
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be a bidirectional stream by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromTrack(TrackLocal, *RTPTransceiverInit) (RTPTransceiver, error)

	// AddTransceiverFromKind creates a new RTPTransceiver from RTPCodecType
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be a bidirectional stream by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromKind(RTPCodecKind, *RTPTransceiverInit) (RTPTransceiver, error)

	// OnTrack handles an incoming media feed.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/ontrack
	OnTrack(func(TrackRemote, RTPReceiver))
}

// RTPTransceiverInit represents RTCRtpTransceiverInit dictionary.
type RTPTransceiverInit struct {
	Direction     RTPTransceiverDirection
	SendEncodings []RTPParameters
}
