package main

import (
	"os"
	"time"

	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/encodedframe"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/ivfreader"
)

func main() {
	var s webrtc.SettingEngine

	// During Offer/Answer exchange the only codec we support will be VP8
	// If the remote doesn't support VP8 signaling will fail
	_ = s.SetEncodings([]*webrtc.RTPCodecCapability{
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

	_, _ = peerConnection.AddTransceiverFromTrack(track, nil)

	go func() {
		fb, _ := track.ReadRTCP()
		_ = fb // do nothing in this example
	}()

	file, _ := os.Open("video.ivf")
	ivf, _, _ := ivfreader.NewWith(file)

	for i := 0; i <= 10; i++ {
		frame, _, _ := ivf.ParseNextFrame()
		_ = track.WriteEncodedFrame(&encodedframe.EncodedFrame{Data: frame, Duration: time.Second})
	}
}
