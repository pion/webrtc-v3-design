package webrtc

// PayloadType is used to indicate the Payload format
// PayloadType type is used to identify an individual codec
// they are dynamic and determined by the offerer
type PayloadType uint8

// RTPCodecCapability is the configuration for one RTPCodec
type RTPCodecCapability struct {
	PreferredPayloadType PayloadType
	MimeType             string
	ClockRate            int
	Channels             int
	SdpFmtpLine          string
}
