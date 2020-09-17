package webrtc

import (
	"github.com/pion/webrtc-v3-design/rtpengine"
)

// TrackBase represents common MediaStreamTrack functionality of LocalTrack and RemoteTrack.
type TrackBase interface {
	ID() string
	Stop() error
}

// LocalTrack is an abstract interface of LocalTrack implementations.
type LocalTrack interface {
	TrackBase
}

// LocalTrackAdapter provides a method to connect specific pair of LocalTrack and RTPSender.
type LocalTrackAdapter interface {
	// Bind should implement the way how the media data flows from LocalTrack to RTPSender.
	// This will be called internally by RTPSender.ReplaceTrack() and
	// PeerConnection.AddTransceiverFromTrack().
	//
	// For example of LocalRTPTrack and PassthroughRTPSender, it might be like:
	//   func Bind(track LocalTrack, sender RTPSender) error {
	//     rtpTrack, ok := track.(LocalRTPTrack)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     ptSender, ok := sender.(PassthroughRTPSender)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     // RTP packets written via LocalRTPTrack.WriteRTP() will be
	//     // read by LocalRTPTrack.pipeReader.ReadRTP().
	//     return rtpengine.Copy(ptSender, rtpTrack.pipeReader)
	//   }
	Bind(LocalTrack, RTPSender) error
}

// LocalRTPTrack represents MediaStreamTrack which is fed by the local stream source.
// Unlike WebAPI's MediaStreamTrack, the track directly provides RTP stream.
type LocalRTPTrack interface {
	LocalTrack
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

// RemoteTrack is an abstract interface of RemoteTrack implementations.
type RemoteTrack interface {
	TrackBase
}

// RemoteTrackAdapter provides a method to connect specific pair of RemoteTrack and RTPReceiver.
type RemoteTrackAdapter interface {
	// Bind should implement the way how the media data flows from RTPReceiver to RemoteTrack.
	// PeerConnection.AddTransceiverFromKind().
	//
	// For example of RemoteRTPTrack and PassthroughRTPReceiver, it might be like:
	//   func Bind(track RemoteTrack, sender RTPReceiver) error {
	//     rtpTrack, ok := track.(RemoteRTPTrack)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     ptReceiver, ok := sender.(PassthroughRTPReceiver)
	//     if !ok {
	//       return ErrIncompatible
	//     }
	//     // RTP packets written to RemoteRTPTrack.pipeWriter.WriteRTP() will be read from
	//     // RemoteRTPTrack.ReadRTP().
	//     return rtpengine.Copy(rtpTrack.pipeWriter, ptReceiver)
	//   }
	Bind(RemoteTrack, RTPReceiver) error
}

// RemoteRTPTrack represents MediaStreamTrack which is fed by the remote peer.
// Unlike WebAPI's MediaStreamTrack, the track directly consumes RTP stream.
type RemoteRTPTrack interface {
	RemoteTrack
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

// PassthroughRTPSender is a RTPSender which can be used with LocalRTPTrack.
// RTP packets written through rtpengine.Writer interface will be directly
// passed to LocalRTPTrack.
type PassthroughRTPSender interface {
	RTPSender
	rtpengine.Writer
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

// PassthroughRTPReceiver is a RTPReceiver which can be used with RemoteRTPTrack.
// Received RTP packets can be directly read through rtpengine.Reader interface.
type PassthroughRTPReceiver interface {
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
