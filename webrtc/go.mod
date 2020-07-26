module github.com/pion/webrtc-v3-design/webrtc

go 1.14

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/pion/randutil v0.1.0 // indirect
	github.com/pion/rtcp v1.2.3
	github.com/pion/rtp v1.6.0
	github.com/pion/webrtc-v3-design/rtpengine v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace github.com/pion/webrtc-v3-design/rtpengine => ../rtpengine
