package main

import (
	"fmt"
	"github.com/draringi/SermonStation/audio"
	"github.com/draringi/SermonStation/web"
	"github.com/gordonklaus/portaudio"
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
	defer portaudio.Terminate()
	audioManager, err := audio.NewManager()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	web.StartServer(audioManager)
	device := getDevice()
	fmt.Printf("Recording with %s\n", device.Name)
	audioManager.SetDevice(device)
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
	audioManager.SetChannelCount(chanCount)
	sampleSize = getSampleSize()
	audioManager.SetSampleSize(sampleSize)
	var path string
	fmt.Println("Where should the audio be saved to?")
	fmt.Scanf("%s", &path)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	recorder, err := audioManager.NewRecording(path)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Println("Starting Recording")
	recorder.Start()
	fmt.Println("Recording")
	for {
		select {
		case <-sig:
			fmt.Println("Sending Stop Signal")
			recorder.Stop()
		default:
		}
		if recorder.Status() == audio.STOPPED {
			fmt.Println("Recording Stopped")
			break
		} else if recorder.Error() != nil {
			fmt.Printf("ERROR: %s\n", recorder.Error())
			break
		}
	}
}
