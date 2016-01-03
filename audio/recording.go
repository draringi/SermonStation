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
}

func (r *Recording) Start() error {
	f, err := os.Create(r.path)
	if err != nil {
		return err
	}
	r.startedAt = time.Now()
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
	err = binary.Write(f, binary.BigEndian, int32(0))
	if err != nil {
		return err
	}
	frameCount := 0
	if err != nil {
		return err
	}
	return nil
}
