package media

import (
	"errors"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/webrtc"
	"github.com/pion/webrtc-v3-design/webrtc/pkg/media/encodedframe"
)

var (
	ErrUnsupportedCodec = errors.New("unsupported codec")
)

type staticLocalRTPTrack struct {
	ssrc   uint32
	codecs []webrtc.RTPCodecCapability
}

// NewStaticLocalTrack returns a TrackLocalRTP with a pre-set codec.
func NewStaticLocalRTPTrack(c webrtc.RTPCodecCapability) *staticLocalRTPTrack {
	return &staticLocalRTPTrack{
		codecs: []webrtc.RTPCodecCapability{c},
	}
}

func (s *staticLocalRTPTrack) ID() string  { return "" }
func (s *staticLocalRTPTrack) Stop() error { return nil }

func (s *staticLocalRTPTrack) WriteRTP(*rtp.Packet) error     { return nil }
func (s *staticLocalRTPTrack) ReadRTCP() (rtcp.Packet, error) { return nil, nil }

// SetParameters proxies RTPSender.SetParameters() called by the PeerConnection.
// The Track should select matching codec from available codecs.
func (s *staticLocalRTPTrack) SetParameters(params webrtc.RTPParameters) error {
	if len(s.codecs) == 0 {
		return ErrUnsupportedCodec
	}
	for _, codec := range params.Codecs {
		if codec.MimeType == s.codecs[0].MimeType {
			// Found available codec.
			s.ssrc = params.SSRC
			return nil
		}
	}
	return ErrUnsupportedCodec
}

// Parameters returns available codecs and selected codec.
func (s *staticLocalRTPTrack) Parameters() webrtc.RTPParameters {
	return webrtc.RTPParameters{
		SSRC:          s.ssrc,
		SelectedCodec: &s.codecs[0],
		Codecs:        s.codecs,
	}
}

type staticLocalEncodedFrameTrack struct {
	ssrc   uint32
	codecs []webrtc.RTPCodecCapability
}

// NewStaticLocalEncodedFrameTrack returns a LocalEncodedFrameTrack with a pre-set codec.
func NewStaticLocalEncodedFrameTrack(c webrtc.RTPCodecCapability) *staticLocalEncodedFrameTrack {
	return &staticLocalEncodedFrameTrack{
		codecs: []webrtc.RTPCodecCapability{c},
	}
}

func (s *staticLocalEncodedFrameTrack) ID() string  { return "" }
func (s *staticLocalEncodedFrameTrack) Stop() error { return nil }

func (s *staticLocalEncodedFrameTrack) WriteEncodedFrame(*encodedframe.EncodedFrame) error {
	return nil
}
func (s *staticLocalEncodedFrameTrack) ReadRTCP() (rtcp.Packet, error) { return nil, nil }

// SetParameters proxies RTPSender.SetParameters() called by the PeerConnection.
// The Track should select matching codec from available codecs.
func (s *staticLocalEncodedFrameTrack) SetParameters(params webrtc.RTPParameters) error {
	if len(s.codecs) == 0 {
		return ErrUnsupportedCodec
	}
	for _, codec := range params.Codecs {
		if codec.MimeType == s.codecs[0].MimeType {
			// Found available codec.
			s.ssrc = params.SSRC
			return nil
		}
	}
	return ErrUnsupportedCodec
}

// Parameters returns available codecs and selected codec.
func (s *staticLocalEncodedFrameTrack) Parameters() webrtc.RTPParameters {
	return webrtc.RTPParameters{
		SSRC:          s.ssrc,
		SelectedCodec: &s.codecs[0],
		Codecs:        s.codecs,
	}
}
