package webrtc

// A Configuration defines how peer-to-peer communication via PeerConnection
// is established or re-established.
// Configurations may be set up once and reused across multiple connections.
// Configurations are treated as readonly. As long as they are unmodified,
// they are safe for concurrent use.
type Configuration struct{}

// NewPeerConnection creates a new WebRTC PeerConnection. This uses a uninitialized
// SettingEngine. If you wish to access any Pion specific behaviors you should create a
// PeerConnection using `NewPeerConnection` method of the SettingEngine
func NewPeerConnection(c Configuration) (PeerConnection, error) {
	var s SettingEngine
	return s.NewPeerConnection(c)
}

// PeerConnection represents RTCPeerConnection.
type PeerConnection interface {
	// AddTransceiverFromTrack creates a new RTPTransceiver from TrackLocal
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be sendrecv by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromTrack(TrackLocal, *RTPTransceiverInit) (RTPTransceiver, error)

	// AddTransceiverFromKind creates a new RTPTransceiver from RTPCodecType
	// and register it to the PeerConnection.
	// Pass nil as a second argument to use default setting.
	// Returned RTPTransceiver will be sendrecv by default.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/addTransceiver
	AddTransceiverFromKind(RTPCodecKind, *RTPTransceiverInit) (RTPTransceiver, error)

	// OnTrack handles an incoming media feed.
	//
	// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCPeerConnection/ontrack
	OnTrack(func(TrackRemote, RTPReceiver))
}

type peerConnection struct{}

func (p *peerConnection) AddTransceiverFromTrack(TrackLocal, *RTPTransceiverInit) (RTPTransceiver, error) {
	return nil, nil
}
func (p *peerConnection) AddTransceiverFromKind(RTPCodecKind, *RTPTransceiverInit) (RTPTransceiver, error) {
	return nil, nil
}
func (p *peerConnection) OnTrack(func(TrackRemote, RTPReceiver)) {}

// RTPTransceiverInit represents RTCRtpTransceiverInit dictionary.
type RTPTransceiverInit struct {
	Direction     RTPTransceiverDirection
	SendEncodings []RTPParameters
}
