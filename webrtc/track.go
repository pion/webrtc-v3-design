package webrtc

import (
	"context"
	"errors"

	"github.com/pion/webrtc-v3-design/rtpengine"
)

var (
	// ErrIncompatible is an internal error returned from TrackLocal/RemoteAdapter.
	ErrIncompatible = errors.New("incompatible pair of Track and RTPSender/Receiver")
)

// TrackBase represents common MediaStreamTrack functionality of TrackLocal and TrackRemote.
type TrackBase interface {
	ID() string
	Stop() error
}

// TrackLocal is an abstract interface of TrackLocal implementations.
type TrackLocal interface {
	TrackBase
}

type TrackLocalAdapter interface {
	// Bind should implement the way how the media data flows from TrackLocal to RTPSender.
	// This will be called internally by RTPSender.ReplaceTrack() and
	// PeerConnection.AddTransceiverFromTrack().
	Bind(context.Context, TrackLocal, RTPSender) error
}

// TrackRemote is an abstract interface of TrackRemote implementations.
type TrackRemote interface {
	TrackBase
}

// TrackRemoteAdapter provides a method to connect specific pair of TrackRemote and RTPReceiver.
type TrackRemoteAdapter interface {
	// Bind should implement the way how the media data flows from RTPReceiver to TrackRemote.
	// PeerConnection.AddTransceiverFromKind().
	Bind(context.Context, TrackRemote, RTPReceiver) error
}

// RTPParameters represents RTCRtpParameters which contains information about
// how the RTC data is to be encoded/decoded.
//
// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCRtpSendParameters
type RTPParameters struct {
	SSRC          uint32
	SelectedCodec *RTPCodecCapability
	Codecs        []RTPCodecCapability
}

// RTPSender represents RTCRtpSender.
type RTPSender interface {
	// ReplaceTrackLocal registers given TrackLocal as a source of RTP packets.
	ReplaceTrack(TrackLocal) error
	// Track returns currently registered TrackLocal.
	Track() TrackLocal

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
	// Track returns associated TrackRemote.
	Track() TrackRemote

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
