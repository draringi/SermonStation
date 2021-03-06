package audio

import (
	"encoding/binary"
	"errors"
	"github.com/gordonklaus/portaudio"
	"log"
	"os"
	"time"
)

type action int
type Status int

type Recording struct {
	path        string
	streamInfo  portaudio.StreamParameters
	stream      *portaudio.Stream
	startedAt   time.Time
	file        *os.File
	err         error
	channels    int
	sampleSize  int
	buffer      portaudio.Buffer
	actionQueue chan action
	status      Status
	frameCount  int
}

// States for the recording
const (
	RECORDING Status = iota
	STOPPED
	PENDING
)

func (s Status) String() string {
	switch s {
	case RECORDING:
		return "Recording"
	case STOPPED:
		return "Stopped"
	case PENDING:
		return "Pending"
	default:
		return "Unknown"
	}
}

const (
	stop action = iota
	pause
)

const (
	aiffFORMSize       = 4
	aiffCOMMSize       = 8 + 18
	aiffSSNDHeaderSize = 16
	paBufferSize       = 128
)

func (r *Recording) Start() error {
	log.Printf("Starting Recording to %s (c: %d, b: %d)\n", r.path, r.channels, r.sampleSize)
	r.file, r.err = os.Create(r.path)
	f := r.file
	if r.err != nil {
		return r.err
	}
	// Form Chunk
	_, r.err = f.WriteString("FORM")
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(0))
	if r.err != nil {
		return r.err
	}
	_, r.err = f.WriteString("AIFF")
	if r.err != nil {
		return r.err
	}
	// Common Chunk
	_, r.err = f.WriteString("COMM")
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(18))
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int16(r.channels))
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(0))
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int16(r.sampleSize))
	if r.err != nil {
		return r.err
	}
	_, r.err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0})
	if r.err != nil {
		return r.err
	}
	// Sound Data Chunk
	_, r.err = f.WriteString("SSND")
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(0))
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(0))
	if r.err != nil {
		return r.err
	}
	r.err = binary.Write(f, binary.BigEndian, int32(0))
	if r.err != nil {
		return r.err
	}
	r.startedAt = time.Now()
	switch r.sampleSize {
	case 32:
		tmpBuffer := make([]int32, r.channels*paBufferSize)
		r.buffer = tmpBuffer
	case 24:
		tmpBuffer := make([]portaudio.Int24, r.channels*paBufferSize)
		r.buffer = tmpBuffer
	case 16:
		tmpBuffer := make([]int16, r.channels*paBufferSize)
		r.buffer = tmpBuffer
	case 8:
		tmpBuffer := make([]int8, r.channels*paBufferSize)
		r.buffer = tmpBuffer
	default:
		r.err = errors.New("Invalid sample size")
		return r.err
	}
	go r.run()
	return nil
}

func (r *Recording) run() {
	r.frameCount = 0
	f := r.file
	defer func() {
		if r.err != nil {
			return
		}
		bytesPerSample := r.sampleSize / 8
		audioSize := r.frameCount * r.channels * bytesPerSample
		totalSize := aiffCOMMSize + aiffSSNDHeaderSize + audioSize + aiffFORMSize
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
		r.err = binary.Write(f, binary.BigEndian, int32(r.frameCount))
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
		r.stream.Close()
		r.status = STOPPED
		log.Printf("Done recording to %s (%v)\n", r.path, r.Duration())
	}()
	r.stream, r.err = portaudio.OpenStream(r.streamInfo, r.buffer)
	if r.err != nil {
		return
	}
	r.status = RECORDING
	r.err = r.stream.Start()
	if r.err != nil {
		return
	}
	for {
		r.stream.Read()
		switch r.sampleSize {
		case 32:
			tmpBuffer := r.buffer.([]int32)
			l := len(tmpBuffer) / r.channels
			r.err = binary.Write(f, binary.BigEndian, tmpBuffer)
			if r.err != nil {
				return
			}
			r.frameCount += l
		case 24:
			tmpBuffer := r.buffer.([]portaudio.Int24)
			l := len(tmpBuffer) / r.channels
			r.err = binary.Write(f, binary.BigEndian, tmpBuffer)
			if r.err != nil {
				return
			}
			r.frameCount += l
		case 16:
			tmpBuffer := r.buffer.([]int16)
			l := len(tmpBuffer) / r.channels
			r.err = binary.Write(f, binary.BigEndian, tmpBuffer)
			if r.err != nil {
				return
			}
			r.frameCount += l
		case 8:
			tmpBuffer := r.buffer.([]int8)
			l := len(tmpBuffer) / r.channels
			r.err = binary.Write(f, binary.BigEndian, tmpBuffer)
			if r.err != nil {
				return
			}
			r.frameCount += l
		default:
			r.err = errors.New("Invalid sample size")
			return
		}
		select {
		case <-r.actionQueue:
			return
		default:

		}
	}
}

func (r *Recording) Stop() {
	r.actionQueue <- stop
}

func (r *Recording) Status() Status {
	return r.status
}

func (r *Recording) Duration() time.Duration {
	nanoseconds := (float64(r.frameCount) / 44100) * float64(time.Second)
	return time.Duration(nanoseconds)
}

func NewRecording(path string, params portaudio.StreamParameters, channels, sampleSize int) *Recording {
	r := new(Recording)
	r.path = path
	r.actionQueue = make(chan action, 1)
	r.channels = channels
	r.sampleSize = sampleSize
	r.status = PENDING
	r.streamInfo = params
	return r
}

func (r *Recording) Error() error {
	return r.err
}
