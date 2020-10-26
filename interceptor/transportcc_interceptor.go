package interceptor

import (
	"context"
	"log"
	"time"

	"github.com/pion/rtcp"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
)

type TransportCCInterceptor struct {
	peerConnection         *webrtc.PeerConnection
	outgoinSequenceNumbers map[uint32]uint16
	incomingStats          interface{}
	close                  chan struct{}
	incomingExtensionIDs   map[uint32]uint8
	outgoingExtensionIDs   map[uint32]uint8
}

func (t *TransportCCInterceptor) InterceptRTCP(ctx context.Context, packets []rtcp.Packet, next func([]rtcp.Packet)) error {
	for _, p := range packets {
		if tcc, ok := p.(*rtcp.TransportLayerCC); ok {
			log.Print(tcc.String()) // TODO: calculate and report suggested bandwidth
		}
	}

	next(packets)

	return nil
}

func (t *TransportCCInterceptor) InterceptOutgoingRTP(ctx context.Context, packet *rtp.Packet, next func(*rtp.Packet)) error {
	t.outgoinSequenceNumbers[packet.SSRC]++
	seqNum := t.outgoinSequenceNumbers[packet.SSRC]

	b, err := (&rtp.TransportCCExtension{TransportSequence: seqNum}).Marshal()
	if err != nil {
		return err
	}

	err = packet.SetExtension(t.outgoingExtensionIDs[packet.SSRC], b) // use extension id instead of 0
	if err != nil {
		return err
	}

	next(packet)

	return nil
}

func (t *TransportCCInterceptor) InterceptIncomingRTP(ctx context.Context, packet *rtp.Packet, next func(*rtp.Packet)) error {
	ext := &rtp.TransportCCExtension{}
	err := ext.Unmarshal(packet.GetExtension(t.incomingExtensionIDs[packet.SSRC])) // use extension id instead of 0
	if err != nil {
		return err
	}

	t.incomingStats.Add(ext.TransportSequence, time.Now())

	return nil
}

func (t *TransportCCInterceptor) sendReport() {
	t.incomingStats.GetReport() // get report

	t.peerConnection.WriteRTCP([]rtcp.Packet{
		&rtcp.TransportLayerCC{}, // fill this
	})
}

func (t *TransportCCInterceptor) start() {
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

func (t *TransportCCInterceptor) stop() {
	close(t.close)
}
