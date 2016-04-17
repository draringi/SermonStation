package audio

import (
	"errors"
	"github.com/gordonklaus/portaudio"
)

type Manager struct {
	device     *portaudio.DeviceInfo
	channels   int
	sampleSize int
	recording  *Recording
}

func (m *Manager) SetDevice(device *portaudio.DeviceInfo) {
	m.device = device
}

func (m *Manager) Device() *portaudio.DeviceInfo {
	return m.device
}

func (m *Manager) SetChannelCount(n int) {
	m.channels = n
}

func (m *Manager) Channels() int {
	return m.channels
}

func (m *Manager) SetSampleSize(size int) error {
	switch size {
	case 8, 16, 24, 32:
		m.sampleSize = size
	default:
		return errors.New("Invalid Sample Size")
	}
	return nil
}

func (m *Manager) SampleSize() int {
	return m.sampleSize
}

func NewManager() (m *Manager, err error) {
	m = new(Manager)
	m.device, err = portaudio.DefaultInputDevice()
	if err != nil {
		return nil, err
	}
	m.channels = m.device.MaxInputChannels
	m.sampleSize = 32
	// Future initialization code goes here...
	return m, nil
}

func (m *Manager) Parameters() portaudio.StreamParameters {
	return portaudio.StreamParameters{Input: portaudio.StreamDeviceParameters{Device: m.device, Channels: m.channels, Latency: m.device.DefaultHighInputLatency}, SampleRate: 44100, FramesPerBuffer: 128}
}

func (m *Manager) NewRecording(path string) (*Recording, error) {
	if m.recording != nil {
		if m.recording.status != STOPPED {
			return m.recording, errors.New("Already Active Recording exists")
		}
	}
	m.recording = NewRecording(path, m.Parameters(), m.channels, m.sampleSize)
	return m.recording, nil
}

func (m *Manager) Status() map[string]interface{} {
	var status map[string]interface{}
	if m.recording != nil {
		status["status"] = m.recording.Status()
		status["duration"] = m.recording.Duration()
	} else {
		status["status"] = "Stopped"
	}
	return status
}

func init() {
	portaudio.Initialize()
}
