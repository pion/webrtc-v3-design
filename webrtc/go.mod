module github.com/pion/webrtc-v3-design/webrtc

go 1.14

require (
	github.com/pion/rtcp v1.2.3
	github.com/pion/rtp v1.6.0
	github.com/pion/webrtc-v3-design/rtpengine v0.0.0-20200905201212-4337232b67dc
)

replace github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
