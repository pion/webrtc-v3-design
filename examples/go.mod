module github.com/pion/webrtc-v3-design/examples

go 1.14

require (
	github.com/pion/mediadevices v0.0.0-20200903041225-86e3a3f14ca8
	github.com/pion/webrtc-v3-design/webrtc v0.0.0-20200905201212-4337232b67dc
)

replace (
	github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
	github.com/pion/webrtc-v3-design/webrtc => ../webrtc
)
