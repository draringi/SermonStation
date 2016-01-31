package audio

import (
	"github.com/gordonklaus/portaudio"
)

type Manager struct {
	stream *portaudio.DeviceInfo
}

func init() {
	portaudio.Initialize()
}
