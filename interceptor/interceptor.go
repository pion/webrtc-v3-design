package interceptor

import (
	"context"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

type Interceptor interface {
	InterceptWriteRTP(context.Context, *webrtc.PeerConnection, Writer) Writer
	InterceptReadRTP(context.Context, *webrtc.PeerConnection, Reader) Reader
	InterceptWriteRTCP(context.Context, *webrtc.PeerConnection, FeedbackWriter) FeedbackWriter
	InterceptReadRTCP(context.Context, *webrtc.PeerConnection, FeedbackReader) FeedbackReader
	Delete(context.Context, *webrtc.PeerConnection)
}

// Reader is an interface to handle incoming RTP stream.
type Reader interface {
	ReadRTP(context.Context) (*rtp.Packet, map[interface{}]interface{}, error)
}

// Writer is an interface to handle outgoing RTP stream.
type Writer interface {
	WriteRTP(context.Context, *rtp.Packet, map[interface{}]interface{}) error
}

// FeedbackReader is an interface to read feedback RTCP packets.
type FeedbackReader interface {
	ReadRTCP(context.Context) ([]rtcp.Packet, error)
}

// FeedbackWriter is an interface to write feedback RTCP packets.
type FeedbackWriter interface {
	WriteRTCP(context.Context, []rtcp.Packet) error
}
