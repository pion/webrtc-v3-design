module github.com/pion/webrtc-v3-design/examples

go 1.14

require (
	github.com/pion/mediadevices v0.0.0-20200921200001-a44240be5fb0
	github.com/pion/webrtc-v3-design/mediadevices v0.0.0-20200905201212-4337232b67dc
	github.com/pion/webrtc-v3-design/webrtc v0.0.0-20200927053335-2e0f9748bf9e
)

replace (
	github.com/pion/webrtc-v3-design/mediadevices => ../mediadevices
	github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
	github.com/pion/webrtc-v3-design/webrtc => ../webrtc
)
