package webrtc

// RTPTransceiverDirection indicates the direction of the RTPTransceiver.
type RTPTransceiverDirection int

const (
	// RTPTransceiverDirectionSendrecv indicates the RTPSender will offer
	// to send RTP and RTPReceiver the will offer to receive RTP.
	RTPTransceiverDirectionSendrecv RTPTransceiverDirection = iota + 1

	// RTPTransceiverDirectionSendonly indicates the RTPSender will offer
	// to send RTP.
	RTPTransceiverDirectionSendonly

	// RTPTransceiverDirectionRecvonly indicates the RTPReceiver the will
	// offer to receive RTP.
	RTPTransceiverDirectionRecvonly

	// RTPTransceiverDirectionInactive indicates the RTPSender won't offer
	// to send RTP and RTPReceiver the won't offer to receive RTP.
	RTPTransceiverDirectionInactive
)
