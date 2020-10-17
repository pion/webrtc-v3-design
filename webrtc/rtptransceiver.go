package webrtc

type SSRC uint32

// RTPSender represents RTCRtpSender.
type RTPSender interface {
	// ReplaceTrackLocal registers given TrackLocal as a source of RTP packets.
	ReplaceTrack(TrackLocal) error

	// Track returns currently bound TrackLocal.
	Track() TrackLocal

	// Parameters returns information about how the data is to be encoded.
	Parameters() RTPParameters

	// SetParameters sets information about how the data is to be encoded.
	// This will be called by PeerConnection according to the result of
	// SDP based negotiation.
	SetParameters(RTPParameters) error
}

// RTPReceiver represents RTCRtpReceiver.
type RTPReceiver interface {
	// Track returns associated TrackRemote.
	Track() TrackRemote

	// Parameters returns information about how the data is to be decoded.
	Parameters() RTPParameters
}

// RTPTransceiver represents RTCRtpTransceiver.
// It represents a combination of an RTCRtpSender and an RTCRtpReceiver that share a common mid.
//
// ref: https://www.w3.org/TR/webrtc/#rtcrtptransceiver-interface
type RTPTransceiver interface {
	RTPSender() RTPSender
	RTPReceiver() RTPReceiver
}
