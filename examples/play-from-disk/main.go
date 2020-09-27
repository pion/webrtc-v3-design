package main

import (
	"context"

	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/encodedframe"
)

func main() {
	var s webrtc.SettingEngine

	// During Offer/Answer exchange the only codec we support will be VP8
	// If the remote doesn't support VP8 signaling will fail
	s.SetEncodings([]*webrtc.RTPCodecCapability{
		{
			MimeType:  "video/vp8", // Should we make this a enum?
			ClockRate: 90000,       // Sholud we drop from API and just assume?
		},
	})

	peerConnection, _ := s.NewPeerConnection(webrtc.Configuration{})

	track := media.NewStaticLocalEncodedFrameTrack(webrtc.RTPCodecCapability{
		MimeType:  "video/vp8", // Should we make this a enum?
		ClockRate: 90000,       // Sholud we drop from API and just assume?
	})

	peerConnection.AddTransceiverFromTrack(track, nil)

	ctx := context.TODO()

	go func() {
		fb, _ := track.ReadRTCP()
		_ = fb // do nothing in this example
	}()

	r := &encodedFrameReader{}
	for {
		frame, _ := r.Read(ctx)
		_ = track.WriteEncodedFrame(frame)
	}
}

type encodedFrameReader struct {
}

func (*encodedFrameReader) Read(ctx context.Context) (*encodedframe.EncodedFrame, error) {
	panic("not implemented")
}
