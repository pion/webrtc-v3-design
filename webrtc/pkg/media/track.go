package media

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/encodedframe"
)

type staticLocalRTPTrack struct{}

// NewStaticLocalTrack returns a TrackLocalRTP with a pre-set codec.
func NewStaticLocalRTPTrack(webrtc.RTPCodecCapability) *staticLocalRTPTrack { return nil }

func (s *staticLocalRTPTrack) ID() string  { return "" }
func (s *staticLocalRTPTrack) Stop() error { return nil }

func (s *staticLocalRTPTrack) WriteRTP(*rtp.Packet) error     { return nil }
func (s *staticLocalRTPTrack) ReadRTCP() (rtcp.Packet, error) { return nil, nil }

// SetParameters asserts that requested codec is available from the other side
func (s *staticLocalRTPTrack) SetParameters(webrtc.RTPParameters) error { return nil }

type staticLocalEncodedFrameTrack struct{}

// NewStaticLocalEncodedFrameTrack returns a LocalEncodedFrameTrack with a pre-set codec.
func NewStaticLocalEncodedFrameTrack(webrtc.RTPCodecCapability) *staticLocalEncodedFrameTrack {
	return nil
}

func (s *staticLocalEncodedFrameTrack) ID() string  { return "" }
func (s *staticLocalEncodedFrameTrack) Stop() error { return nil }

func (s *staticLocalEncodedFrameTrack) WriteEncodedFrame(*encodedframe.EncodedFrame) error {
	return nil
}
func (s *staticLocalEncodedFrameTrack) ReadRTCP() (rtcp.Packet, error) { return nil, nil }

// SetParameters asserts that requested codec is available from the other side
func (s *staticLocalEncodedFrameTrack) SetParameters(webrtc.RTPParameters) error { return nil }
