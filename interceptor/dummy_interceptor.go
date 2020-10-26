package interceptor

import (
	"context"

	"github.com/pion/webrtc/v3"
)

type DummyInterceptor struct {
}

func (d *DummyInterceptor) InterceptWriteRTP(_ context.Context, _ *webrtc.PeerConnection, writer Writer) Writer {
	return writer
}

func (d *DummyInterceptor) InterceptReadRTP(_ context.Context, _ *webrtc.PeerConnection, reader Reader) Reader {
	return reader
}

func (d *DummyInterceptor) InterceptWriteRTCP(_ context.Context, _ *webrtc.PeerConnection, writer FeedbackWriter) FeedbackWriter {
	return writer
}

func (d *DummyInterceptor) InterceptReadRTCP(_ context.Context, _ *webrtc.PeerConnection, reader FeedbackReader) FeedbackReader {
	return reader
}

func (d *DummyInterceptor) Delete(_ context.Context, _ *webrtc.PeerConnection) {
	// do nothing
}
