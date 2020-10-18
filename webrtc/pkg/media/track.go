package media

import (
	"errors"

	"github.com/pion/rtp"
	"github.com/pion/webrtc-v3-design/webrtc"
)

var (
	ErrUnsupportedCodec = errors.New("unsupported codec")
	ErrUnbindFailed     = errors.New("failed to unbind Track from PeerConnection")
)

// trackBinding is a single bind for a Track
// Bind can be called multiple times, this stores the
// result for a single bind call so that it can be used when writing
type trackBinding struct {
	ssrc        webrtc.SSRC
	payloadType webrtc.PayloadType
	writeStream webrtc.TrackLocalWriter
}

// staticLocalRTPTrack is a track that has a pre-set codec
type localTrackStaticRTP struct {
	bindings     []trackBinding
	codec        webrtc.RTPCodecCapability
	id, streamId string
}

// NewStaticLocalTrack returns a TrackLocalRTP with a pre-set codec.
func NewLocalTrackStaticRTP(c webrtc.RTPCodecCapability, id, streamId string) (*localTrackStaticRTP, error) {
	return &localTrackStaticRTP{
		codec:    c,
		bindings: []trackBinding{},
		id:       id,
		streamId: streamId,
	}, nil
}

// Bind is called by the PeerConnection after negotation is complete
// This asserts that the code requested is supported by the remote peer.
// If so it setups all the state (SSRC and PayloadType) to have a call
func (s *localTrackStaticRTP) Bind(t webrtc.TrackLocalContext) error {
	for _, codec := range t.Parameters().Codecs {
		if codec.MimeType == s.codec.MimeType {
			s.bindings = append(s.bindings, trackBinding{
				ssrc:        t.Parameters().SSRC,
				payloadType: codec.PreferredPayloadType,
				writeStream: t.WriteStream(),
			})
			return nil
		}
	}
	return ErrUnsupportedCodec
}

func (s *localTrackStaticRTP) Unbind(t webrtc.TrackLocalContext) error {
	for i := range s.bindings {
		if s.bindings[i].writeStream == t.WriteStream() {
			s.bindings[i] = s.bindings[len(s.bindings)-1]
			s.bindings = s.bindings[:len(s.bindings)-1]
			return nil
		}
	}

	return ErrUnbindFailed
}

func (s *localTrackStaticRTP) ID() string       { return s.id }
func (s *localTrackStaticRTP) StreamID() string { return s.streamId }

// Loop each binding and set the proper SSRC/PayloadType before writing
func (s *localTrackStaticRTP) WriteRTP(p *rtp.Packet) error      { return nil }
func (s *localTrackStaticRTP) Write(b []byte) (n int, err error) { return }

type localTrackStaticSample struct {
	packetizer interface{}
	rtpTrack   *localTrackStaticRTP
}

func NewLocalTrackStaticSample(c webrtc.RTPCodecCapability, id, streamId string) (*localTrackStaticSample, error) {
	rtpTrack, err := NewLocalTrackStaticRTP(c, id, streamId)
	if err != nil {
		return nil, err
	}

	return &localTrackStaticSample{
		packetizer: nil,
		rtpTrack:   rtpTrack,
	}, nil
}

func (s *localTrackStaticSample) ID() string       { return s.rtpTrack.ID() }
func (s *localTrackStaticSample) StreamID() string { return s.rtpTrack.StreamID() }

// Call rtpTrack.Bind + setup packetizer
func (s *localTrackStaticSample) Bind(t webrtc.TrackLocalContext) error   { return nil }
func (s *localTrackStaticSample) Unbind(t webrtc.TrackLocalContext) error { return nil }

func (s *localTrackStaticSample) WriteSample(samp Sample) error { return nil }
