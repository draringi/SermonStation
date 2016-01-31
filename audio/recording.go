package audio

import (
	"code.google.com/p/portaudio-go/portaudio"
	"encoding/binary"
	"os"
	"time"
)

type Recording struct {
	path       string
	streamInfo *portaudio.StreamParameters
	stream     *portaudio.Stream
	startedAt  time.Time
	file       *os.File
	err        error
	channels   int
	sampleSize int
	buffer     portaudio.Buffer
}

const (
	aiffFORMSize       = 4
	aiffCOMMSize       = 8 + 18
	aiffSSNDHeaderSize = 16
	paBufferSize       = 128
)

func (r *Recording) Start() error {
	r.file, err = os.Create(r.path)
	f := r.file
	if err != nil {
		return err
	}
	// Form Chunk
	_, err = f.WriteString("FORM")
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	_, err = f.WriteString("AIFF")
	if err != nil {
		return err
	}
	// Common Chunk
	_, err = f.WriteString("COMM")
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(18))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int16(r.streamInfo.Input.Channels))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int16(32))
	if err != nil {
		return err
	}
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return err
	}
	// Sound Data Chunk
	_, err = f.WriteString("SSND")
y	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	r.startedAt = time.Now()
	switch sampleSize {
	case 32:
		r.buffer = make([][]int32, r.channels)
		for c := 0; c < r.channels; c++ {
			r.buffer[c] = make([]int32, paBufferSize)
		}
	case 24:
		r.buffer = make([][]Int24, r.channels)
		for _, c := range(r.buffer) {
			c = make([]Int24, paBufferSize)
		}
	case 16:
		r.buffer = make([][]int16, r.channels)
		for _, c := range(r.buffer) {
			c = make([]int16, paBufferSize)
		}
	case 8:
		r.buffer = make([][]int8, r.channels)
		for _, c := range(r.buffer) {
			c = make([]int8, paBufferSize)
		}
	default:
		r.err = error("Invalid sample size")
		return r.err
	}
	go r.run()
	return nil
}

func (r *Recording) run() {
	frameCount := 0
	f := r.file
	defer func() {
		if r.err != nil {
			return
		}
		bytesPerSample = r.sampleSize / 8
		audioSize = framecount * r.channels * bytesPerSample
		totalSize = aiffCOMMSize + aiffSSNDHeaderSize + audioSize + aiffFORMSize
		_, r.err = f.Seek(4, 0)
		if r.err != nil {
			return
		}
		r.err = binary.Write(f, binary.BigEndian, int32(totalSize))
		if r.err != nil {
			return
		}
		_, r.err = f.Seek(22, 0)
		if r.err != nil {
			return
		}
		r.err = binary.Write(f, binary.BigEndian, int32(frameCount))
		if r.err != nil {
			return
		}
		_, r.err = f.Seek(42, 0)
		if r.err != nil {
			return
		}
		r.err = binary.Write(f, binary.BigEndian, int32(audioSize+8))
		if r.err != nil {
			return
		}
		r.err = f.Close()
	}()

}
