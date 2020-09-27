package encodedframe

import (
	"context"
	"time"
)

// EncodedFrame represents single complete data of a media frame.
// It replaces media.Sample in pion/webrtc v2 API.
type EncodedFrame struct {
	Data      []byte
	Timestamp time.Time
	Duration  time.Duration
}

// Reader is an interface to read EncodedFrame.
type Reader interface {
	Read(context.Context) (*EncodedFrame, error)
}

// Writer is an interface to write EncodedFrame.
type Writer interface {
	Write(context.Context, *EncodedFrame) error
}

// Copy EncodedFrame from Reader to Writer.
func Copy(context.Context, Writer, Reader) error {
	panic("unimplemented")
}
