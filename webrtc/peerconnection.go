package webrtc

type Configuration struct{}

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
