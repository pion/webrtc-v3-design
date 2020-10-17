package main

import (
	"os"
	"time"

	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media"
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

	track, _ := media.NewStaticLocalSampleTrack(webrtc.RTPCodecCapability{
		MimeType:  "video/vp8", // Should we make this a enum?
		ClockRate: 90000,       // Sholud we drop from API and just assume?
	}, "video", "desktop-capture")

	_, _ = peerConnection.AddTransceiverFromTrack(track, nil)

	file, _ := os.Open("video.ivf")
	ivf, _, _ := ivfreader.NewWith(file)

	for i := 0; i <= 10; i++ {
		frame, _, _ := ivf.ParseNextFrame()
		_ = track.WriteSample(media.Sample{Data: frame, Duration: time.Second})
	}
}
