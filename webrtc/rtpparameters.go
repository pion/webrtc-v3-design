package webrtc

// RTPParameters represents RTCRtpParameters which contains information about
// how the RTC data is to be encoded/decoded.
//
// ref: https://developer.mozilla.org/en-US/docs/Web/API/RTCRtpSendParameters
type RTPParameters struct {
	SSRC          SSRC
	SelectedCodec *RTPCodecCapability
	Codecs        []RTPCodecCapability
}
