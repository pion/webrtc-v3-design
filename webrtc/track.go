package webrtc

import (
	"context"

	"github.com/pion/webrtc-v3-design/rtpengine"
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

// TrackLocalAdapter provides a method to connect specific pair of TrackLocal and RTPSender.
type TrackLocalAdapter interface {
	// Bind should implement the way how the media data flows from TrackLocal to RTPSender.
	// This will be called internally by RTPSender.ReplaceTrack() and
	// PeerConnection.AddTransceiverFromTrack().
	//
	// For example of TrackLocalRTP and RTPSenderPassthrough, it might be like:
	//   func Bind(ctx context.Context, track TrackLocal, sender RTPSender) error {
	//     rtpTrack, ok := track.(TrackLocalRTP)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     ptSender, ok := sender.(RTPSenderPassthrough)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     // RTP packets written via TrackLocalRTP.WriteRTP() will be
	//     // read by TrackLocalRTP.pipeReader.ReadRTP().
	//     return rtpengine.Copy(ctx, ptSender, rtpTrack.pipeReader)
	//   }
	Bind(context.Context, TrackLocal, RTPSender) error
}

// TrackLocalRTP represents MediaStreamTrack which is fed by the local stream source.
// Unlike WebAPI's MediaStreamTrack, the track directly provides RTP stream.
type TrackLocalRTP interface {
	TrackLocal
	rtpengine.Writer

	// SetParameters sets information about how the data is to be encoded.
	// This will be called by PeerConnection according to the result of
	// SDP based negotiation.
	// It will be called via RTPSender.Parameters() by PeerConnection to
	// tell the negotiated media codec information.
	//
	// This is pion's extension to process data without having encoder/decoder
	// in webrtc package.
	SetParameters(RTPParameters) error
}

// TrackRemote is an abstract interface of TrackRemote implementations.
type TrackRemote interface {
	TrackBase
}

// TrackRemoteAdapter provides a method to connect specific pair of TrackRemote and RTPReceiver.
type TrackRemoteAdapter interface {
	// Bind should implement the way how the media data flows from RTPReceiver to TrackRemote.
	// PeerConnection.AddTransceiverFromKind().
	//
	// For example of TrackRemoteRTP and RTPReceiverPassthrough, it might be like:
	//   func Bind(ctx context.Context, track TrackRemote, sender RTPReceiver) error {
	//     rtpTrack, ok := track.(TrackRemoteRTP)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     ptReceiver, ok := sender.(RTPReceiverPassthrough)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     // RTP packets written to TrackRemoteRTP.pipeWriter.WriteRTP() will be read from
	//     // TrackRemoteRTP.ReadRTP().
	//     return rtpengine.Copy(ctx, rtpTrack.pipeWriter, ptReceiver)
	//   }
	Bind(context.Context, TrackRemote, RTPReceiver) error
}

// TrackRemoteRTP represents MediaStreamTrack which is fed by the remote peer.
// Unlike WebAPI's MediaStreamTrack, the track directly consumes RTP stream.
type TrackRemoteRTP interface {
	TrackRemote
	rtpengine.Reader

	// Parameters returns information about how the data is to be encoded.
	// Call of this function will be redirected to associated RTPReceiver
	// to get the negotiated media codec information.
	//
	// This is pion's extension to process data without having encoder/decoder
	// in webrtc package.
	Parameters() RTPParameters
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

// RTPSenderPassthrough is a RTPSender which can be used with TrackLocalRTP.
// RTP packets written through rtpengine.Writer interface will be directly
// passed to TrackLocalRTP.
type RTPSenderPassthrough interface {
	RTPSender
	rtpengine.Writer
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

// RTPReceiverPassthrough is a RTPReceiver which can be used with TrackRemoteRTP.
// Received RTP packets can be directly read through rtpengine.Reader interface.
type RTPReceiverPassthrough interface {
	RTPReceiver
	rtpengine.Reader
}

// RTPTransceiver represents RTCRtpTransceiver.
// It represents a combination of an RTCRtpSender and an RTCRtpReceiver that share a common mid.
//
// ref: https://www.w3.org/TR/webrtc/#rtcrtptransceiver-interface
type RTPTransceiver interface {
	RTPSender() RTPSender
	RTPReceiver() RTPReceiver
}
