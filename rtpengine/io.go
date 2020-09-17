package rtpengine

import (
	"context"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// Reader is an interface to handle incoming RTP stream.
type Reader interface {
	ReadRTP(context.Context) (*rtp.Packet, error)
	WriteRTCP(context.Context, rtcp.Packet) error
}

// Writer is an interface to handle outgoing RTP stream.
type Writer interface {
	WriteRTP(context.Context, *rtp.Packet) error
	ReadRTCP(context.Context) (rtcp.Packet, error)
}

// Copy rtp.Packet from Reader to Writer and rtcp.Packet from Writer to Reader.
func Copy(context.Context, Writer, Reader) error {
	panic("unimplemented")
}
