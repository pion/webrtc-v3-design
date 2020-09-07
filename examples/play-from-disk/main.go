package main

import (
	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media"
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

	track := media.NewStaticLocalRTPTrack(webrtc.RTPCodecCapability{
		MimeType:  "video/vp8", // Should we make this a enum?
		ClockRate: 90000,       // Sholud we drop from API and just assume?
	})

	peerConnection.AddTransceiverFromTrack(track, nil)

	// TODO
	// Do we expose a packetizer as a public API?
	// Do we add WriteSample?
	track.WriteRTP(nil)
}
