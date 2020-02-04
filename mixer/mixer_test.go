package mixer

import (
	"os"
	"testing"

	"github.com/NIHERASE/voxcox/compressor"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func TestOneChanelInput(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	exit := make(chan int)
	exit2 := make(chan int)

	output := make(chan int, 1000)

	m := new(Mixer)
	compressor := new(compressor.MockCompressorImpl)
	m.audioCompressor = compressor

	voice1, err1 := os.Open("../samples/Spanish vocal phrase.wav")
	defer voice1.Close()

	if err1 != nil {
		t.Errorf("file not exist's")
	}

	decoder1 := wav.NewDecoder(voice1)
	buffer1, err1 := decoder1.FullPCMBuffer()

	if err1 != nil {
		t.Errorf("i can't extract buffer from file")
	}

	go saveChanToFile(output, exit2, buffer1, t)
	go m.Compare(ch1, ch2, exit, output)

	i := 0
	for ; i < len(buffer1.Data); i++ {
		ch1 <- buffer1.Data[i]
	}

	exit <- 1
	exit2 <- 1

}

func saveChanToFile(input, exit chan int, buffer *audio.IntBuffer, t *testing.T) {
	work := true
	var data []int

	for work {
		select {
		case in := <-input:
			data = append(data, in)
		case <-exit:
			work = false
		}
	}
	buffer.Data = data

	simpleSaveToFile(buffer, t)

}

func TestMixTwoFile(t *testing.T) {
	voice1, err1 := os.Open("../samples/Spanish vocal phrase dollars-.wav")
	defer voice1.Close()

	voice2, err2 := os.Open("../samples/Telephone prompt poss.wav")
	defer voice2.Close()

	if err1 != nil || err2 != nil {
		t.Errorf("file not exist's")
	}

	decoder1 := wav.NewDecoder(voice1)
	decoder2 := wav.NewDecoder(voice2)
	buffer1, err1 := decoder1.FullPCMBuffer()
	buffer2, err2 := decoder2.FullPCMBuffer()

	if err1 != nil || err2 != nil {
		t.Errorf("i can't extract buffer from file")
	}

	bufferIndex := 0
	if len(buffer1.Data) < len(buffer2.Data) {
		bufferIndex = 1
	}

	if bufferIndex == 0 {
		simpleCompareBuffer(buffer1, buffer2)
		simpleSaveToFile(buffer1, t)
	} else {
		simpleCompareBuffer(buffer2, buffer1)
		simpleSaveToFile(buffer2, t)
	}

}
func simpleCompareBuffer(voice1, voice2 *audio.IntBuffer) {
	for i := 0; i < len(voice2.Data); i++ {
		voice1.Data[i] += voice2.Data[i]
	}

}
func simpleSaveToFile(buffer *audio.IntBuffer, t *testing.T) {
	output, err := os.OpenFile("result.wav", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		t.Errorf("i can't save result file")
	}
	defer output.Close()

	encoder := wav.NewEncoder(
		output,
		buffer.PCMFormat().SampleRate,
		buffer.SourceBitDepth,
		buffer.PCMFormat().NumChannels,
		1,
	)
	defer encoder.Close()
	encoder.Write(buffer)
}
