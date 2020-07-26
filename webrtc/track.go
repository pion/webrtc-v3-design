package webrtc

import (
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/rtpengine"
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

	// SetRTPInterceptor inserts given rtpengine.Interceptor to the sending RTP stream.
	// This is pion extension of WebRTC to customize packet processing algorithms like
	// packet retransmission and congestion control.
	SetRTPInterceptor(rtpengine.WriteInterceptor)
}

// RTPReceiver represents RTCRtpReceiver.
type RTPReceiver interface {
	// SetRTPInterceptor inserts given rtpengine.Interceptor to the received RTP stream.
	// This is pion extension of WebRTC to customize packet processing algorithms like
	// jitter buffer and congestion control.
	SetRTPInterceptor(rtpengine.ReadInterceptor)
}

// RTPTransceiver represents RTCRtpTransceiver.
type RTPTransceiver interface {
	RTPSender() RTPSender
	RTPReceiver() RTPReceiver
}
