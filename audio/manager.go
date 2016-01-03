package audio

import (
	"code.google.com/p/portaudio-go/portaudio"
)

type Manager struct {
	stream *portaudio.DeviceInfo
}

func init() {
	portaudio.Initialize()
}
