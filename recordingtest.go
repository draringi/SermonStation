package main

import (
	"github.com/draringi/SermonStation/audio"
	"github.com/gordonklaus/portaudio"
	"fmt"
	"os"
	"os/signal"
)

func getDevice() *portaudio.DeviceInfo {
	hs, err := portaudio.HostApis()
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("Please Select an API")
		for index, api := range hs {
			fmt.Printf("%d: %s\n", index, api.Name)
		}
		var selection int
		fmt.Printf("API #: ")
		fmt.Scanf("%d", &selection)
		api := hs[selection]
		fmt.Println("Please select an input device")
		for index, dev := range api.Devices {
			fmt.Printf("%d: %s (%d channels)\n", index, dev.Name, dev.MaxInputChannels)
		}
		fmt.Printf("Device #: ")
		fmt.Scanf("%d", &selection)
		dev := api.Devices[selection]
		return dev
	}
}

func getSampleSize() int {
	for {
		var sampleSize int
		fmt.Printf("How big should a sample be? {32, 24, 16, 8} ")
		fmt.Scanf("%d", &sampleSize)
		switch sampleSize {
		case 8, 16, 24, 32:
			return sampleSize
		default:
			fmt.Println("Invalid Sample Size. Must be 32, 24, 16 or 8")
		}
	}
}

func main() {
	device := getDevice()
	fmt.Printf("Recording with %s\n", device.Name)
	var chanCount, sampleSize int
	for {
		fmt.Printf("How many channels? [1,%d] ", device.MaxInputChannels)
		fmt.Scanf("%d", &chanCount)
		if chanCount > 0 && chanCount <= device.MaxInputChannels {
			break
		} else {
			fmt.Println("Invalid number of channels")
		}
	}
	sampleSize = getSampleSize()
	var path string
	fmt.Println("Where should the audio be saved to?")
	fmt.Scanf("%s", &path)
	var params portaudio.StreamParameters
	var devParams portaudio.StreamDeviceParameters
	devParams.Channels = chanCount
	devParams.Device = device
	devParams.Latency = device.DefaultLowInputLatency
	params.SampleRate = device.DefaultSampleRate
	params.FramesPerBuffer = 128
	params.Input = devParams
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	recorder := audio.NewRecording(path, params, chanCount, sampleSize)
	fmt.Println("Starting Recording")
	recorder.Start()
	fmt.Println("Recording")
	for {
		switch {
		case stopCmd := <-sig:
			recorder.Stop()
		case recorder.Status() == audio.STOPPED:
			fmt.Println("Recording Stopped")
			return
		default:
		}
	}
}
