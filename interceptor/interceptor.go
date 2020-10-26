package interceptor

import (
	"context"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/webrtc"
)

type IncomingRTPInterceptorFactory interface {
	CreateIncomingRTPInterceptor(*webrtc.PeerConnection) IncomingRTPInterceptor
	DeleteIncomingRTPInterceptor(*webrtc.PeerConnection)
}

type IncomingRTPInterceptor interface{
	InterceptIncomingRTP(ctx context.Context, packet *rtp.Packet, next func(*rtp.Packet)) error
}

type OutgoingRTPInterceptorFactory interface {
	CreateOutgoingRTPInterceptor(*webrtc.PeerConnection) OutgoingRTPInterceptor
	DeleteOutgoingRTPInterceptor(*webrtc.PeerConnection)
}

type OutgoingRTPInterceptor interface{
	InterceptOutgoingRTP(ctx context.Context, packet *rtp.Packet, next func(*rtp.Packet)) error
}

type RTCPInterceptorFactory interface {
	CreateRTCPInterceptor(*webrtc.PeerConnection) RTCPInterceptor
	DeleteRTCPInterceptor(*webrtc.PeerConnection)
}

type RTCPInterceptor interface {
	InterceptRTCP(ctx context.Context, packets []rtcp.Packet, next func([]rtcp.Packet)) error
}