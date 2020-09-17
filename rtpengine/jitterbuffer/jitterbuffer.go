package jitterbuffer

import (
	"context"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/rtpengine"
)

// JitterBuffer is a stub implementation of Interceptor for jitter buffering.
type JitterBuffer struct {
	maxLate int
}

// Intercept implements ReadInterceptor.
func (j *JitterBuffer) Intercept(ctx context.Context, r rtpengine.Reader) rtpengine.Reader {
	return &jitterBufferReader{r}
}

type jitterBufferReader struct {
	in rtpengine.Reader
}

func (j *jitterBufferReader) WriteRTCP(ctx context.Context, p rtcp.Packet) error {
	// This interceptor doesn't use RTCP packet.
	return j.in.WriteRTCP(ctx, p)
}

func (j *jitterBufferReader) ReadRTP(ctx context.Context) (*rtp.Packet, error) {
	p, err := j.in.ReadRTP(ctx)
	if err != nil {
		return nil, err
	}
	// Do jitter buffering here.

	return p, nil
}

// Assert JitterBuffer implements rtpengine.ReadInterceptor.
var _jitterBuffer rtpengine.ReadInterceptor = &JitterBuffer{}
