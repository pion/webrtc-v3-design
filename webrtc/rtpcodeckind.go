package webrtc

type RTPCodecKind int

const (
	RTPCodecKindVideo RTPCodecKind = iota + 1
	RTPCodecKindAudio
)
