package webrtc

import (
	"context"
	"golang.org/x/sync/errgroup"

	"github.com/pion/webrtc-v3-design/rtpengine"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/encodedframe"
)

// TrackLocalEncodedFrame_RTPSenderPacketizer provides a method to connect TrackLocalEncodedFrame and RTPSenderPacketizer.
type TrackLocalEncodedFrame_RTPSenderPacketizer struct {
}

// Bind implements the way how the media data flows from TrackLocalEncodedFrame to RTPSenderPacketizer.
// This will be called internally by RTPSender.ReplaceTrack() and
// PeerConnection.AddTransceiverFromTrack().
func (*TrackLocalEncodedFrame_RTPSenderPacketizer) Bind(ctx context.Context, track TrackLocal, sender RTPSender) error {
	encodedFrameTrack, ok := track.(TrackLocalEncodedFrame)
	if !ok {
		return ErrIncompatible
	}
	packetizeSender, ok := sender.(RTPSenderPacketizer)
	if !ok {
		return ErrIncompatible
	}
	eg, ctx2 := errgroup.WithContext(ctx)
	// Encoded frames written via TrackLocalEncodedFrame.Write() will be
	// read by TrackLocalEncodedFrame.pipeReader.ReadRTP().
	eg.Go(func() error {
		return encodedframe.Copy(ctx, packetizeSender, encodedFrameTrack.pipeReader())
	})
	eg.Go(func() error {
		return rtpengine.CopyFeedback(ctx2, encodedFrameTrack.pipeFeedbackWriter(), packetizeSender)
	})
	return eg.Wait()
}

// TrackLocalEncodedFrame represents MediaStreamTrack which is fed by the local stream source.
// Unlike WebAPI's MediaStreamTrack, the track directly provides encoded frames.
type TrackLocalEncodedFrame interface {
	TrackLocal
	encodedframe.Writer
	rtpengine.FeedbackReader

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
	pipeReader() encodedframe.Reader
	// pipeFeedbackWriter returns pipe writer connected to the FeedbackWriter.
	pipeFeedbackWriter() rtpengine.FeedbackWriter
}

// TrackRemoteEncodedFrame_RTPReceiverDepacketizer provides a method to connect TrackRemoteEncodedFrame and RTPReceiverDepacketizer.
type TrackRemoteEncodedFrame_RTPReceiverDepacketizer struct {
}

// Bind implements the way how the media data flows from RTPReceiverDepacketizer to TrackRemoteEncodedFrame.
func (*TrackRemoteEncodedFrame_RTPReceiverDepacketizer) Bind(ctx context.Context, track TrackRemote, sender RTPReceiver) error {
	encodedFrameTrack, ok := track.(TrackRemoteEncodedFrame)
	if !ok {
		return ErrIncompatible
	}
	depacketizeReceiver, ok := sender.(RTPReceiverDepacketizer)
	if !ok {
		return ErrIncompatible
	}
	eg, ctx2 := errgroup.WithContext(ctx)
	// Encoded frames written to TrackRemoteEncodedFrame.pipeWriter.WriteRTP() will be read from
	// TrackRemoteEncodedFrame.ReadRTP().
	eg.Go(func() error {
		return encodedframe.Copy(ctx2, encodedFrameTrack.pipeWriter(), depacketizeReceiver)
	})
	eg.Go(func() error {
		return rtpengine.CopyFeedback(ctx2, depacketizeReceiver, encodedFrameTrack.pipeFeedbackReader())
	})
	return eg.Wait()
}

// TrackRemoteEncodedFrame represents MediaStreamTrack which is fed by the remote peer.
// Unlike WebAPI's MediaStreamTrack, the track directly consumes RTP stream.
type TrackRemoteEncodedFrame interface {
	TrackRemote
	encodedframe.Reader
	rtpengine.FeedbackWriter

	// Parameters returns information about how the data is to be encoded.
	// Call of this function will be redirected to associated RTPReceiver
	// to get the negotiated media codec information.
	//
	// This is pion's extension to process data without having encoder/decoder
	// in webrtc package.
	Parameters() RTPParameters

	// pipeWriter returns pipe writer connected to the Reader.
	pipeWriter() encodedframe.Writer
	// pipeFeedbackReader returns pipe reader connected to the FeedbackReader.
	pipeFeedbackReader() rtpengine.FeedbackReader
}

// RTPSenderPacketizer is a RTPSender which can be used with TrackLocalEncodedFrame.
// Encoded frame written throught encodedframe.Writer interface will be packetized and sent.
type RTPSenderPacketizer interface {
	RTPSender
	encodedframe.Writer
	rtpengine.FeedbackReader
}

// RTPReceiverDepacketizer is a RTPReceiver which can be used with TrackRemoteEncodedFrame.
// Received RTP packets will be depacketized and read through encodedframe.Reader interface.
type RTPReceiverDepacketizer interface {
	RTPReceiver
	encodedframe.Reader
	rtpengine.FeedbackWriter
}
