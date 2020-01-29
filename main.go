package main

import (
	"log"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	file, err := os.Open("./samples/file_example_WAV_1MG.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := wav.NewDecoder(file)
	buffer, err := decoder.FullPCMBuffer()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("sample rate: ", buffer.PCMFormat().SampleRate)
	log.Print("channels: ", buffer.PCMFormat().NumChannels)
	log.Print("SourceBitDepth: ", buffer.SourceBitDepth)
	log.Print("SampleBitDepth: ", decoder.SampleBitDepth())

	// output, err := os.OpenFile("result.wav", os.O_RDWR|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer output.Close()

	// steps := 2
	// for i := 0; i < len(buffer.Data)/steps; i++ {
	// 	for j := 0; j < steps; j++ {
	// 		buffer.Data[i*steps+j] = buffer.Data[i*steps]
	// 	}
	// }

	// for i := 0; i < len(buffer.Data)/2-1; i++ {
	// 	buffer.Data[i*2+1] = buffer.Data[i*2+2]
	// }

	// opusEncoder, err := opus.NewEncoder(buffer.PCMFormat().SampleRate, buffer.PCMFormat().NumChannels, opus.AppVoIP)
	// if err != nil {
	// 	log.Fatal("encoder fail ", err)
	// }
	// err = opusEncoder.SetBitrate(1024 * 8)
	// if err != nil {
	// 	log.Fatal("set bitrate fail ", err)
	// }

	// opusDecoder, err := opus.NewDecoder(buffer.PCMFormat().SampleRate, buffer.PCMFormat().NumChannels)
	// if err != nil {
	// 	log.Fatal("decoder fail", err)
	// }

	// bytesPerFrame := buffer.PCMFormat().SampleRate * 60 / 1000
	// pcmInt16Buffer := make([]int16, len(buffer.Data))
	// for i, a := range buffer.Data {
	// 	pcmInt16Buffer[i] = int16(a)
	// }
	// coded := make([]byte, 10000)
	// decoded := make([]int16, bytesPerFrame)
	// pcmRecoded := make([]int16, 0)

	// for frame := 0; frame < len(pcmInt16Buffer)/bytesPerFrame-1; frame++ {
	// 	start := frame * bytesPerFrame
	// 	pcmFrame := pcmInt16Buffer[start : start+bytesPerFrame]

	// 	log.Print("frame size: ", len(pcmFrame))

	// 	n, err := opusEncoder.Encode(pcmFrame, coded)
	// 	log.Print("encoded: ", n)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	n, err = opusDecoder.Decode(coded[:n], decoded)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Print("decoded: ", n)

	// 	pcmRecoded = append(pcmRecoded, decoded[:n]...)
	// }

	// pcmRecodedInt := make([]int, len(pcmRecoded))
	// for i, b := range pcmRecoded {
	// 	pcmRecodedInt[i] = int(b)
	// }

	// buffer.Data = pcmRecodedInt

	// encoder := wav.NewEncoder(
	// 	output,
	// 	buffer.PCMFormat().SampleRate,
	// 	buffer.SourceBitDepth,
	// 	buffer.PCMFormat().NumChannels,
	// 	1,
	// )
	// defer encoder.Close()
	// encoder.Write(buffer)
}
