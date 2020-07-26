package rtpengine

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// Reader is an interface to handle incoming RTP stream.
type Reader interface {
	ReadRTP() (*rtp.Packet, error)
	WriteRTCP(rtcp.Packet) error
}

// Writer is an interface to handle outgoing RTP stream.
type Writer interface {
	WriteRTP(*rtp.Packet) error
	ReadRTCP() (rtcp.Packet, error)
}
