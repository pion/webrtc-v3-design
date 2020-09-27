package main

import (
	mediadevices2 "github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec"
	"github.com/pion/mediadevices/pkg/codec/openh264"
	"github.com/pion/mediadevices/pkg/codec/vpx"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/webrtc-v3-design/mediadevices"
	"github.com/pion/webrtc-v3-design/webrtc"

	// Note: If you don't have a camera or microphone or your adapters are not supported,
	//       you can always swap your adapters with our dummy adapters below.
	// _ "github.com/pion/mediadevices/pkg/driver/videotest"
	_ "github.com/pion/mediadevices/pkg/driver/camera" // This is required to register camera adapter
)

func main() {
	var setting webrtc.SettingEngine

	openh264Params, _ := openh264.NewParams()
	vp8Params, _ := vpx.NewVP8Params()

	// Assume that we have configured the api and have proper config
	pc, _ := setting.NewPeerConnection(webrtc.Configuration{})

	md := mediadevices.NewMediaDevices()
	mediaStream, _ := md.GetUserMedia(mediadevices2.MediaStreamConstraints{
		Video: func(constraint *mediadevices2.MediaTrackConstraints) {
			constraint.Width = prop.Int(600)
			constraint.Height = prop.Int(400)
			constraint.VideoEncoderBuilders = []codec.VideoEncoderBuilder{&openh264Params, &vp8Params}
		},
	})

	for _, mediaTrack := range mediaStream.GetTracks() {
		// rtpTracker.Track will create TrackLocalRTP, which later can be used for pulling video/audio frames.
		pc.AddTransceiverFromTrack(mediaTrack,
			&webrtc.RTPTransceiverInit{
				Direction: webrtc.RTPTransceiverDirectionSendonly,
			},
		)
	}

	// peerconnection should negotiate with the other peer
	// peerconnection should set the negotiated codec with LocalRTPTrack.setParameters()

	// mediadevices then will set the codec and attach the encoder to the pipeline. PeerConnection now can start
	// pulling encoded frames from mediadevices.
	//
	// Question: How are we going to handle errors from the encoder, e.g. invalid resolutions, invalid sample rate, etc.?
	//           Should we return the error when the user calls SetRemoteDescription?
}
