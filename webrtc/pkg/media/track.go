package media

import (
	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/webrtc"
)

type staticLocalRTPTrack struct{}

// NewStaticTrack returns a LocalRTPTrack with a pre-set codec.
func NewStaticLocalRTPTrack(webrtc.RTPCodecCapability) *staticLocalRTPTrack { return nil }

func (s *staticLocalRTPTrack) ID() string  { return "" }
func (s *staticLocalRTPTrack) Stop() error { return nil }

func (s *staticLocalRTPTrack) WriteRTP(*rtp.Packet) error     { return nil }
func (s *staticLocalRTPTrack) ReadRTCP() (rtcp.Packet, error) { return nil, nil }

// SetParameters asserts that requested codec is available from the other side
func (s *staticLocalRTPTrack) SetParameters(webrtc.RTPParameters) error { return nil }
