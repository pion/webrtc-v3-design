package rtpengine

// WriteInterceptor processes outbound RTP stream.
// For example, WriteInterceptor may implements congestion control and packet retransmission.
type WriteInterceptor interface {
	Intercept(Writer) Writer
}

// ReadInterceptor processes received RTP stream.
// For example, ReadInterceptor may implements RTP jitter buffer, RTCP report sender,
// and packet retransmission request.
type ReadInterceptor interface {
	Intercept(Reader) Reader
}
