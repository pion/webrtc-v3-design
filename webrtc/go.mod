module github.com/pion/webrtc-v3-design/webrtc

go 1.14

require (
	github.com/pion/rtcp v1.2.4
	github.com/pion/rtp v1.6.1
	github.com/pion/srtp v1.5.2
	github.com/pion/webrtc-v3-design/rtpengine v0.0.0-20200905201212-4337232b67dc
	golang.org/x/sync v0.0.0-20200625203802-6e8e738ad208
)

replace github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
