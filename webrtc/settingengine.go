package webrtc

type SettingEngine struct{}

// SetEncodings configures the codecs the newly created PeerConnection is willing to
// send and receive. If nothing is configured it will support everything that Pion WebRTC
// implements packetization for.
func (s *SettingEngine) SetEncodings([]*RTPCodecCapability) error { return nil }

// SetTrackLocalAdapters registers TrackLocalAdapters.
// TrackLocalRTP-RTPSenderPassthrough is registered by default.
func (s *SettingEngine) SetTrackLocalAdapters([]TrackLocalAdapter) error { return nil }

// SetTrackRemoteAdapters registers TrackRemoteAdapters.
// TrackRemoteRTP-RTPReceiverPassthrough is registered by default.
func (s *SettingEngine) SetTrackRemoteAdapters([]TrackRemoteAdapter) error { return nil }

// NewPeerConnection creates a NewPeerConnection
func (s *SettingEngine) NewPeerConnection(Configuration) (PeerConnection, error) {
	p := &peerConnection{}
	return p, nil
}
