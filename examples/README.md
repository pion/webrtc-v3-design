#### [Play Files from Disk](examples/play-from-disk)
Example of playing file from disk.

* User declares codecs they are reading from disk
* On each negotiated PeerConnection they are notified if the receiver can accept what they are offering
* Read packets from disk and send until completed

#### [Save files to Disk](examples/save-to-disk)
Example of saving user input to disk,

* User declares codecs they are saving to disk
* On each negotiated PeerConnection they are notified if the receiver can send what they are offering
* In OnTrack user Read packets from webrtc.RTPReceiver and save to disk

#### [Fan out WebRTC Input](examples/fanout)
Example of simple SFU

* User declares codec they wish to fan out
* Uploader connects and we assert that our chosen codec is supported, if so we start reading
* Downloader connects and we assert our chosen codec is supported
* We send our video packets and the developer doesn't have to worry about managing PayloadTypes/SSRCes etc..

We can also implement some basic receiver feedback. These need to be supported, but are not mandatory
* Determine the lowest support bitrate across all receivers, and forward it back to the sender
* Forward NACK/PLI messages from each receiver back to the sender
* Simulcast (see below)

#### [Simulcast Receive](examples/simulcast-send)
A user should be able to receive multiple feeds for a single Track

**TODO** (Works in master)

#### [Error Resilience Send](examples/error-resilience-send)
A user sending video should be able to receive NACKs and respond to them

**TODO**

#### [Error Resilience Receive](examples/error-resilience-receive)
A user receiving video should be able to send NACKs and receive retransmissions

**TODO**

#### [Congestion Control Send](examples/congestion-control-send)
A user sending video should be able to receive REMB/Receiver Reports/TWCC and adjust bitrate

**TODO**

#### [Congestion Control Receive](examples/congestion-control-receive)
A user receiving video should be able to send REMB/Receiver Reports/TWCC

**TODO**

#### [Simulcast Send](examples/simulcast-send)
A user should be able to send simulcast to another Pion instance. This isn't supported to the browser.

**TODO**

#### [Portable getUserMedia](examples/portable-getusermedia)
Users should be able to call getUserMedia and have it work in both their Go and WASM code.

Everything should be behind platform flags in the `mediadevices` repo so user doesn't need to write platform specific
code.
