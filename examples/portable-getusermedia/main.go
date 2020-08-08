package main

import (
	"github.com/pion/mediadevices"
	"github.com/pion/mediadevices/pkg/codec/openh264"
	"github.com/pion/mediadevices/pkg/codec/x264"
	"github.com/pion/mediadevices/pkg/prop"
	"github.com/pion/webrtc-v3-design/webrtc"

	// Note: If you don't have a camera or microphone or your adapters are not supported,
	//       you can always swap your adapters with our dummy adapters below.
	// _ "github.com/pion/mediadevices/pkg/driver/videotest"
	_ "github.com/pion/mediadevices/pkg/driver/camera" // This is required to register camera adapter
)

func main() {
	var setting webrtc.SettingEngine

	x264Params, _ := x264.NewParams()
	openh264Params, _ := openh264.NewParams()
	rtpTracker, _ := mediadevices.NewRTPTracker(x264Params, openh264Params)
	rtpTracker.PopulateSetting(&setting)

	// Assume that we have configured the api and have proper config
	pc, _ := setting.NewPeerConnection(webrtc.Configuration{})

	mediaStream, _ := mediadevices.GetUserMedia(mediadevices.MediaStreamConstraints{
		Video: func(constraint *mediadevices.MediaTrackConstraints) {
			constraint.Width = prop.Int(600)
			constraint.Height = prop.Int(400)
		},
	})

	for _, mediaTrack := range mediaStream.GetTracks() {
		// rtpTracker.Track will create LocalRTPTrack, which later can be used for pulling video/audio frames.
		pc.AddTransceiverFromTrack(rtpTracker.Track(mediaTrack),
			webrtc.RtpTransceiverInit{
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
