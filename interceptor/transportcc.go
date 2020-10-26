package interceptor

import (
	"net/url"
	"sync"

	"github.com/pion/sdp/v3"
	"github.com/pion/webrtc/v3"
)

type TransportCC struct {
	interceptors map[*webrtc.PeerConnection]*TransportCCInterceptor
	m            *sync.Mutex
}

func ConfigureTransportCC(se *webrtc.SettingEngine, me webrtc.MediaEngine) {
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

	t := &TransportCC{
		interceptors: make(map[*webrtc.PeerConnection]*TransportCCInterceptor),
		m:            &sync.Mutex{},
	}
	se.SetLocalRTPInterceptors(append(se.GetLocalRTPInterceptors(), t))
	se.SetRemoteRTPInterceptors(append(se.GetRemoteRTPInterceptors(), t))
	se.SetRTCPInterceptors(append(se.GetRTCPInterceptors(), t))
}

func (t *TransportCC) CreateIncomingRTPInterceptor(pc *webrtc.PeerConnection) IncomingRTPInterceptor {
	return t.getOrCreateInterceptor(pc)
}

func (t *TransportCC) DeleteIncomingRTPInterceptor(pc *webrtc.PeerConnection) {
	t.deleteInterceptor(pc)
}

func (t *TransportCC) CreateOutgoingRTPInterceptor(pc *webrtc.PeerConnection) OutgoingRTPInterceptor {
	return t.getOrCreateInterceptor(pc)
}

func (t *TransportCC) DeleteOutgoingRTPInterceptor(pc *webrtc.PeerConnection) {
	t.deleteInterceptor(pc)
}

func (t *TransportCC) CreateRTCPInterceptor(pc *webrtc.PeerConnection) RTCPInterceptor {
	return t.getOrCreateInterceptor(pc)
}

func (t *TransportCC) DeleteRTCPInterceptor(pc *webrtc.PeerConnection) {
	t.deleteInterceptor(pc)
}

func (t *TransportCC) getOrCreateInterceptor(pc *webrtc.PeerConnection) *TransportCCInterceptor {
	t.m.Lock()
	defer t.m.Unlock()

	if interceptor, ok := t.interceptors[pc]; ok {
		return interceptor
	}

	interceptor := &TransportCCInterceptor{
		peerConnection: pc,
	}
	interceptor.start()
	t.interceptors[pc] = interceptor

	return interceptor
}

func (t *TransportCC) deleteInterceptor(pc *webrtc.PeerConnection) {
	t.m.Lock()
	defer t.m.Unlock()

	if interceptor, ok := t.interceptors[pc]; ok {
		interceptor.stop()
		delete(t.interceptors, pc)
	}
}
