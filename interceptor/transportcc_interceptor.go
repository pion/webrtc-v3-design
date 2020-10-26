package interceptor

import (
	"context"
	"sync"

	"github.com/pion/webrtc/v3"
)

type TransportCCInterceptor struct {
	m            *sync.Mutex
	transportCCs map[*webrtc.PeerConnection]*TransportCC
}

func (t *TransportCCInterceptor) InterceptWriteRTP(_ context.Context, pc *webrtc.PeerConnection, writer Writer) Writer {
	return &transportCCRTPWriter{transportCC: t.getOrCreateTransportCC(pc), wrapped: writer}
}

func (t *TransportCCInterceptor) InterceptReadRTP(ctx context.Context, pc *webrtc.PeerConnection, reader Reader) Reader {
	return &transportCCRTPReader{transportCC: t.getOrCreateTransportCC(pc), wrapped: reader}
}

func (t *TransportCCInterceptor) InterceptWriteRTCP(ctx context.Context, pc *webrtc.PeerConnection, writer FeedbackWriter) FeedbackWriter {
	transportCC := t.getOrCreateTransportCC(pc)
	transportCC.feedbackWriter = writer
	return writer
}

func (t *TransportCCInterceptor) InterceptReadRTCP(ctx context.Context, pc *webrtc.PeerConnection, reader FeedbackReader) FeedbackReader {
	return &transportCCRTCPReader{transportCC: t.getOrCreateTransportCC(pc), wrapped: reader}
}

func (t *TransportCCInterceptor) Delete(ctx context.Context, pc *webrtc.PeerConnection) {
	t.m.Lock()
	defer t.m.Unlock()

	if transportCC, ok := t.transportCCs[pc]; ok {
		transportCC.stop()
		delete(t.transportCCs, pc)
	}
}
