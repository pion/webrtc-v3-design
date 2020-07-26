package nackhandler

import (
	"io"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/rtpengine"
)

// NACKHandler is a stub implementation of Interceptor for NACK handler.
type NACKHandler struct {
}

// Intercept implements ReadInterceptor.
func (n *NACKHandler) Intercept(w rtpengine.Writer) rtpengine.Writer {
	chRTCP := make(chan rtcp.Packet)
	go func() {
		for {
			p, err := w.ReadRTCP()
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

func (n *nackHandlerWriter) WriteRTP(p *rtp.Packet) error {
	// Passthrough the packet and buffer some duration of the packets.
	return n.out.WriteRTP(p)
}

func (n *nackHandlerWriter) ReadRTCP() (rtcp.Packet, error) {
	p, ok := <-n.chRTCP
	if !ok {
		return nil, io.EOF
	}
	return p, nil
}

// Assert NACKHandler implements rtpengine.WriteInterceptor.
var _nackHandler rtpengine.WriteInterceptor = &NACKHandler{}
