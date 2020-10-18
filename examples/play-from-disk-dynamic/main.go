package main

import (
	"errors"

	"github.com/pion/webrtc-v3-design/webrtc"
)

// dynamicTrackLocal is a track in which the codec is determined by what
// codecs the remote supports
type dynamicTrackLocal struct {
	ssrc        webrtc.SSRC
	payloadType webrtc.PayloadType
	writeStream webrtc.TrackLocalWriter
}

// Users can define their own Track types and decide codecs as they wish
func (d *dynamicTrackLocal) Bind(t webrtc.TrackLocalContext) error {
	writeFileToStream := func(fileName string) {
		// loop file and write to the writeStream
	}

	for _, codec := range t.Parameters().Codecs {
		if codec.MimeType == "h264" || codec.MimeType == "vp8" {
			d.ssrc = t.Parameters().SSRC
			d.payloadType = codec.PreferredPayloadType
			d.writeStream = t.WriteStream()
			go writeFileToStream("out." + codec.MimeType)
			return nil
		}
	}

	return errors.New("remote Peer supports neither H264 or VP8")
}
func (d *dynamicTrackLocal) Unbind(c webrtc.TrackLocalContext) error { return nil }
func (d *dynamicTrackLocal) ID() string                              { return "video" }
func (d *dynamicTrackLocal) StreamID() string                        { return "desktop-capture" }

func main() {
	var s webrtc.SettingEngine

	// We support both VP8 and H264
	_ = s.SetEncodings([]*webrtc.RTPCodecCapability{
		{
			MimeType:  "video/vp8",
			ClockRate: 90000,
		},
		{
			MimeType:  "video/h264",
			ClockRate: 90000,
		},
	})

	pc, _ := s.NewPeerConnection(webrtc.Configuration{})
	track := &dynamicTrackLocal{}

	_, _ = pc.AddTransceiverFromTrack(track,
		&webrtc.RTPTransceiverInit{
			Direction: webrtc.RTPTransceiverDirectionSendonly,
		},
	)
}
