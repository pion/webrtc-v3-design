package main

import (
	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media"
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

	uploadPeerConnection, _ := s.NewPeerConnection(webrtc.Configuration{})

	outboundTrack, _ := media.NewStaticLocalRTPTrack(webrtc.RTPCodecCapability{
		MimeType:  "video/vp8", // Should we make this a enum?
		ClockRate: 90000,       // Sholud we drop from API and just assume?
	}, "video", "desktop-capture")

	uploadPeerConnection.OnTrack(func(track webrtc.TrackRemote, receiver webrtc.RTPReceiver) {
		for {
			pkt, _ := track.ReadRTP()

			// WriteRTP sets the proper SSRC/PayloadType when pushing out to each PeerConnection
			_ = outboundTrack.WriteRTP(pkt)
		}

	})

	for i := 0; i <= 10; i++ {
		viewerPeerConnection, _ := s.NewPeerConnection(webrtc.Configuration{})
		_, _ = viewerPeerConnection.AddTransceiverFromTrack(outboundTrack, nil)
	}
}
