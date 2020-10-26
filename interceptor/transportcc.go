package interceptor

import (
	"context"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/sdp/v3"
	"github.com/pion/webrtc/v3"
)

type TransportCC struct {
	peerConnection         *webrtc.PeerConnection
	outgoinSequenceNumbers map[uint32]uint16
	incomingStats          interface{}
	close                  chan struct{}
	incomingExtensionIDs   map[uint32]uint8
	outgoingExtensionIDs   map[uint32]uint8
	feedbackWriter         FeedbackWriter
}

func ConfigureTransportCC(se *webrtc.SettingEngine, me *webrtc.MediaEngine) {
	transportCCURL, _ := url.Parse(sdp.TransportCCURI)
	se.AddSDPExtensions(webrtc.SDPSectionVideo, []sdp.ExtMap{
		{
			URI: transportCCURL,
		},
	})
	for _, codec := range me.GetCodecsByKind(webrtc.RTPCodecTypeVideo) {
		// TODO: this work similarly to extension, in case the offer didn't contain a feedback, neither should the answer
		codec.RTCPFeedback = append(codec.RTCPFeedback, webrtc.RTCPFeedback{Type: webrtc.TypeRTCPFBTransportCC})
	}

	t := &TransportCCInterceptor{
		transportCCs: make(map[*webrtc.PeerConnection]*TransportCC),
		m:            &sync.Mutex{},
	}
	se.SetInterceptors(append(se.GetInterceptors(), t))
}

func (t *TransportCCInterceptor) getOrCreateTransportCC(pc *webrtc.PeerConnection) *TransportCC {
	t.m.Lock()
	defer t.m.Unlock()

	if transportCC, ok := t.transportCCs[pc]; ok {
		return transportCC
	}

	transportCC := &TransportCC{peerConnection: pc}
	transportCC.start()
	t.transportCCs[pc] = transportCC

	return transportCC
}

func (t *TransportCC) sendReport() {
	t.incomingStats.GetReport() // get report

	t.feedbackWriter.WriteRTCP(context.Background(), []rtcp.Packet{
		&rtcp.TransportLayerCC{}, // fill this
	})
}

func (t *TransportCC) start() {
	t.peerConnection.OnConnectionStateChange(func(state webrtc.PeerConnectionState) {
		if state == webrtc.PeerConnectionStateConnected {
			remoteSDP := t.peerConnection.CurrentRemoteDescription()
			t.incomingExtensionIDs // TODO: get extension number from sdp and fill this

			localSDP := t.peerConnection.CurrentLocalDescription()
			t.outgoingExtensionIDs // TODO: get extension number from sdp and fill this
		}
	})

	go func() {
		ticker := time.Tick(100 * time.Millisecond)
		select {
		case <-ticker:
			t.sendReport()
		case <-t.close:
			return
		}
	}()
}

func (t *TransportCC) stop() {
	close(t.close)
}

type transportCCRTPWriter struct {
	wrapped     Writer
	transportCC *TransportCC
}

func (t *transportCCRTPWriter) WriteRTP(ctx context.Context, packet *rtp.Packet, m map[interface{}]interface{}) error {
	t.transportCC.outgoinSequenceNumbers[packet.SSRC]++
	seqNum := t.transportCC.outgoinSequenceNumbers[packet.SSRC]

	b, err := (&rtp.TransportCCExtension{TransportSequence: seqNum}).Marshal()
	if err != nil {
		return err
	}

	err = packet.SetExtension(t.transportCC.outgoingExtensionIDs[packet.SSRC], b) // use extension id instead of 0
	if err != nil {
		return err
	}

	return t.wrapped.WriteRTP(ctx, packet, m)
}

type transportCCRTPReader struct {
	wrapped     Reader
	transportCC *TransportCC
}

func (t *transportCCRTPReader) ReadRTP(ctx context.Context) (*rtp.Packet, map[interface{}]interface{}, error) {
	packet, meta, err := t.wrapped.ReadRTP(ctx)
	if err != nil {
		return nil, nil, err
	}

	ext := &rtp.TransportCCExtension{}
	err = ext.Unmarshal(packet.GetExtension(t.transportCC.incomingExtensionIDs[packet.SSRC])) // use extension id instead of 0
	if err != nil {
		return nil, nil, err
	}

	t.transportCC.incomingStats.Add(ext.TransportSequence, time.Now())

	return packet, meta, nil
}

type transportCCRTCPReader struct {
	wrapped     FeedbackReader
	transportCC *TransportCC
}

func (t *transportCCRTCPReader) ReadRTCP(ctx context.Context) ([]rtcp.Packet, error) {
	packets, err := t.wrapped.ReadRTCP(ctx)
	if err != nil {
		return nil, err
	}

	for _, p := range packets {
		if tcc, ok := p.(*rtcp.TransportLayerCC); ok {
			log.Print(tcc.String()) // TODO: calculate and report suggested bandwidth
		}
	}

	return packets, nil
}
