package webrtc

type PayloadType uint8

type RTPCodecCapability struct {
	PreferredPayloadType PayloadType
	MimeType             string
	ClockRate            int
	Channels             int
	SdpFmtpLine          string
}
