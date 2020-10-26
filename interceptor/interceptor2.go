package interceptor

import (
	"context"
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

type Interceptor2 interface {
	Intercept(context.Context, *webrtc.PeerConnection, ReadWriter) ReadWriter
}

// Reader is an interface to handle incoming RTP stream.
type ReadWriter interface {
	ReadRTP(context.Context) (*rtp.Packet, map[interface{}]interface{}, error)
	WriteRTP(context.Context, *rtp.Packet, map[interface{}]interface{}) error
	ReadRTCP(context.Context) ([]rtcp.Packet, error)
	WriteRTCP(context.Context, []rtcp.Packet) error
	io.Closer
}

type ExampleInterceptor struct{}

func (d *ExampleInterceptor) Intercept(ctx context.Context, connection *webrtc.PeerConnection, readWriter ReadWriter) ReadWriter {
	return &ExampleReadWriter{readWriter}
}

type ExampleReadWriter struct {
	ReadWriter
}

func (e *ExampleReadWriter) ReadRTP(ctx context.Context) (*rtp.Packet, map[interface{}]interface{}, error) {
	packet, meta, err := e.ReadWriter.ReadRTP(ctx)

	// do stuff

	return packet, meta, err
}