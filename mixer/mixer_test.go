package mixer

import (
	"os"
	"testing"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func TestOkCase(t *testing.T) {

	voice1 := new(audio.IntBuffer)
	voice2 := new(audio.IntBuffer)
	joinedVoices := new(audio.IntBuffer)

	err := compare(voice1, voice2, joinedVoices)

	if err != nil {
		t.Errorf("Test for Ok failed")
	}
}

func TestMixTwoFile(t *testing.T) {
	voice1, err1 := os.Open("../samples/Spanish vocal phrase dollars-.wav")
	defer voice1.Close()

	voice2, err2 := os.Open("../samples/Spanish vocal phrase.wav")
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

	// maxSoundLen := 0
	bufferIndex := 0
	if len(buffer1.Data) > len(buffer2.Data) {
		// maxSoundLen = len(buffer1.Data)
	} else {
		// maxSoundLen = len(buffer2.Data)
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
