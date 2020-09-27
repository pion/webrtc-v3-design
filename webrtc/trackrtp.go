package webrtc

import (
	"context"
	"golang.org/x/sync/errgroup"

	"github.com/pion/webrtc-v3-design/rtpengine"
)

// TrackLocalRTP_RTPSenderPassthrough provides a method to connect TrackLocalRTP and RTPSenderPassthrough.
type TrackLocalRTP_RTPSenderPassthrough struct {
}

// Bind implements the way how the media data flows from TrackLocalRTP to RTPSenderPassthrough.
// This will be called internally by RTPSender.ReplaceTrack() and
// PeerConnection.AddTransceiverFromTrack().
func (*TrackLocalRTP_RTPSenderPassthrough) Bind(ctx context.Context, track TrackLocal, sender RTPSender) error {
	rtpTrack, ok := track.(TrackLocalRTP)
	if !ok {
		return ErrIncompatible
	}
	ptSender, ok := sender.(RTPSenderPassthrough)
	if !ok {
		return ErrIncompatible
	}
	eg, ctx2 := errgroup.WithContext(ctx)
	// RTP packets written via TrackLocalRTP.WriteRTP() will be
	// read by TrackLocalRTP.pipeReader.ReadRTP().
	eg.Go(func() error {
		return rtpengine.Copy(ctx2, ptSender, rtpTrack.pipeReader())
	})
	eg.Go(func() error {
		return rtpengine.CopyFeedback(ctx2, rtpTrack.pipeReader(), ptSender)
	})
	return eg.Wait()
}

// TrackLocalRTP represents MediaStreamTrack which is fed by the local stream source.
// Unlike WebAPI's MediaStreamTrack, the track directly provides RTP stream.
type TrackLocalRTP interface {
	TrackLocal
	rtpengine.Writer

	// SetParameters sets information about how the data is to be encoded.
	// This will be called by PeerConnection according to the result of
	// SDP based negotiation.
	// It will be called via RTPSender.Parameters() by PeerConnection to
	// tell the negotiated media codec information.
	//
	// This is pion's extension to process data without having encoder/decoder
	// in webrtc package.
	SetParameters(RTPParameters) error

	// pipeReader returns pipe reader connected to the Writer.
	// For internal implementation.
	pipeReader() rtpengine.Reader
}

// TrackRemoteRTP_RTPReceiverPassthrough provides a method to connect TrackRemoteRTP and RTPReceiverPassthrough.
type TrackRemoteRTP_RTPReceiverPassthrough struct {
}

// Bind implements the way how the media data flows from RTPReceiverPassthrough to TrackRemoteRTP.
func (*TrackRemoteRTP_RTPReceiverPassthrough) Bind(ctx context.Context, track TrackRemote, sender RTPReceiver) error {
	rtpTrack, ok := track.(TrackRemoteRTP)
	if !ok {
		return ErrIncompatible
	}
	ptReceiver, ok := sender.(RTPReceiverPassthrough)
	if !ok {
		return ErrIncompatible
	}
	eg, ctx2 := errgroup.WithContext(ctx)
	// RTP packets written to TrackRemoteRTP.pipeWriter.WriteRTP() will be read from
	// TrackRemoteRTP.ReadRTP().
	eg.Go(func() error {
		return rtpengine.Copy(ctx2, rtpTrack.pipeWriter(), ptReceiver)
	})
	eg.Go(func() error {
		return rtpengine.CopyFeedback(ctx2, ptReceiver, rtpTrack.pipeWriter())
	})
	return eg.Wait()
}

// TrackRemoteRTP represents MediaStreamTrack which is fed by the remote peer.
// Unlike WebAPI's MediaStreamTrack, the track directly consumes RTP stream.
type TrackRemoteRTP interface {
	TrackRemote
	rtpengine.Reader

	// Parameters returns information about how the data is to be encoded.
	// Call of this function will be redirected to associated RTPReceiver
	// to get the negotiated media codec information.
	//
	// This is pion's extension to process data without having encoder/decoder
	// in webrtc package.
	Parameters() RTPParameters

	// pipeWriter returns pipe writer connected to the Reader.
	pipeWriter() rtpengine.Writer
}

// RTPSenderPassthrough is a RTPSender which can be used with TrackLocalRTP.
// RTP packets written through rtpengine.Writer interface will be directly
// passed to TrackLocalRTP.
type RTPSenderPassthrough interface {
	RTPSender
	rtpengine.Writer
}

// RTPReceiverPassthrough is a RTPReceiver which can be used with TrackRemoteRTP.
// Received RTP packets can be directly read through rtpengine.Reader interface.
type RTPReceiverPassthrough interface {
	RTPReceiver
	rtpengine.Reader
}
