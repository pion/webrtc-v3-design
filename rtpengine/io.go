package rtpengine

import (
	"context"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
)

// Reader is an interface to handle incoming RTP stream.
type Reader interface {
	FeedbackWriter
	ReadRTP(context.Context) (*rtp.Packet, error)
}

// Writer is an interface to handle outgoing RTP stream.
type Writer interface {
	FeedbackReader
	WriteRTP(context.Context, *rtp.Packet) error
}

// Copy rtp.Packet from Reader to Writer.
func Copy(context.Context, Writer, Reader) error {
	panic("unimplemented")
}

// FeedbackReader is an interface to read feedback RTCP packets.
type FeedbackReader interface {
	ReadRTCP(context.Context) (rtcp.Packet, error)
}

// FeedbackWriter is an interface to write feedback RTCP packets.
type FeedbackWriter interface {
	WriteRTCP(context.Context, rtcp.Packet) error
}

// Copy rtcp.Packet from FeedbackReader to FeedbackWriter.
func CopyFeedback(context.Context, FeedbackWriter, FeedbackReader) error {
	panic("unimplemented")
}
