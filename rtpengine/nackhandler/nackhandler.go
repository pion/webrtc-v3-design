package nackhandler

import (
	"context"
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/rtpengine"
)

// NACKHandler is a stub implementation of Interceptor for NACK handler.
type NACKHandler struct {
}

// Intercept implements ReadInterceptor.
func (n *NACKHandler) Intercept(ctx context.Context, w rtpengine.Writer) rtpengine.Writer {
	chRTCP := make(chan rtcp.Packet)
	go func() {
		for {
			p, err := w.ReadRTCP(ctx)
			if err != nil {
				return
			}
			// If received RTCP packet is NACK and requested buffered packet,
			// resend the packet by w.WriteRTP().

			// Passthrough the RTCP packet to the next writer.
			select {
			case chRTCP <- p:
			default:
			}
		}
	}()
	return &nackHandlerWriter{out: w, chRTCP: chRTCP}
}

type nackHandlerWriter struct {
	out    rtpengine.Writer
	chRTCP chan rtcp.Packet
}

func (n *nackHandlerWriter) WriteRTP(ctx context.Context, p *rtp.Packet) error {
	// Passthrough the packet and buffer some duration of the packets.
	return n.out.WriteRTP(ctx, p)
}

func (n *nackHandlerWriter) ReadRTCP(ctx context.Context) (rtcp.Packet, error) {
	select {
	case p, ok := <-n.chRTCP:
		if !ok {
			return nil, io.EOF
		}
		return p, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// Assert NACKHandler implements rtpengine.WriteInterceptor.
var _nackHandler rtpengine.WriteInterceptor = &NACKHandler{}
