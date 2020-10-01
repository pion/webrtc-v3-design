module github.com/pion/webrtc-v3-design/examples
//Hey My first commit
go 1.14

require (
	github.com/pion/mediadevices v0.0.0-20200929170321-f3e3dc9589ca
	github.com/pion/webrtc-v3-design/mediadevices v0.0.0-20200905201212-4337232b67dc
	github.com/pion/webrtc-v3-design/webrtc v0.0.0-20200927053335-2e0f9748bf9e
)

replace (
	github.com/pion/webrtc-v3-design/mediadevices => ../mediadevices
	github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
	github.com/pion/webrtc-v3-design/webrtc => ../webrtc
)
