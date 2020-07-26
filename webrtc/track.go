package webrtc

import (
	"github.com/pion/webrtc-v3-design/rtpengine"
	"github.com/pion/webrtc/v2"
)

// TrackBase represents common MediaStreamTrack functionality of LocalTrack and RemoteTrack.
type TrackBase interface {
	ID() string
}

// LocalTrack represents MediaStreamTrack which is fed by the local stream source.
type LocalTrack interface {
	TrackBase
}

// RemoteTrack represents MediaStreamTrack which is fed by the remote peer.
type RemoteTrack interface {
	TrackBase
}

// Track represents bi-directional MediaStreamTrack.
type Track interface {
	RemoteTrack
	LocalTrack
}

// RTPParameters represents RTCRtpParameters which contains information about
// how the RTC data is to be encoded/decoded.
//
// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCRtpSendParameters
type RTPParameters struct {
	SSRC          uint32
	SelectedCodec *webrtc.RTPCodec
	Codecs        []*webrtc.RTPCodec
}

// RTPSender represents RTCRtpSender.
type RTPSender interface {
	// ReplaceLocalTrack registers given LocalTrack as a source of RTP packets.
	ReplaceTrack(LocalTrack) error
	// Track returns currently registered LocalTrack.
	Track() LocalTrack

	// Parameters returns information about how the data is to be encoded.
	Parameters() RTPParameters
	// SetParameters sets information about how the data is to be encoded.
	// This will be called by PeerConnection according to the result of
	// SDP based negotiation.
	SetParameters(RTPParameters) error

	// SetRTPInterceptor inserts given rtpengine.Interceptor to the sending RTP stream.
	// This is pion extension of WebRTC to customize packet processing algorithms like
	// packet retransmission and congestion control.
	SetRTPInterceptor(rtpengine.WriteInterceptor)
}

// RTPReceiver represents RTCRtpReceiver.
type RTPReceiver interface {
	// Track returns associated RemoteTrack.
	Track() RemoteTrack

	// Parameters returns information about how the data is to be decoded.
	Parameters() RTPParameters

	// SetRTPInterceptor inserts given rtpengine.Interceptor to the received RTP stream.
	// This is pion extension of WebRTC to customize packet processing algorithms like
	// jitter buffer and congestion control.
	SetRTPInterceptor(rtpengine.ReadInterceptor)
}

// RTPTransceiver represents RTCRtpTransceiver.
// It represents a combination of an RTCRtpSender and an RTCRtpReceiver that share a common mid.
//
// ref: https://www.w3.org/TR/webrtc/#rtcrtptransceiver-interface
type RTPTransceiver interface {
	RTPSender() RTPSender
	RTPReceiver() RTPReceiver
}
