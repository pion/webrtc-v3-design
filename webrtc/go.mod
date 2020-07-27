module github.com/pion/webrtc-v3-design/webrtc

go 1.14

require (
	github.com/pion/webrtc-v3-design/rtpengine v0.0.0-00010101000000-000000000000
	github.com/pion/webrtc/v2 v2.2.21
	github.com/pion/webrtc/v3 v3.0.0-20200724080143-6ee528d349a5
)

replace github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
