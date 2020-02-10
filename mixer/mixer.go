package mixer

import (
	"github.com/NIHERASE/voxcox/compressor"
)

//Mixer compare voice streams
type Mixer struct {
	audioCompressor compressor.Compressor
	voice1          VoiceBit
	voice2          VoiceBit
	outputChanel    chan<- int
}

//VoiceBit use for storage value of voice in per time
type VoiceBit struct {
	volume int
}

//Compare two voices to one
func (m *Mixer) Compare(chanVoice1, chanVoice2 chan int, joinedVoices chan<- int) {

	m.outputChanel = joinedVoices

	if m.audioCompressor == nil {
		panic("audio compressor must be initialized")
	}

	for chanVoice1 != nil || chanVoice2 != nil {
		select {
		case bitChanVoice1, ok := <-chanVoice1:
			if ok {
				m.setRecivedVoice1(bitChanVoice1)
				m.waitAnotherChan(chanVoice2, &m.voice2)
			} else {
				chanVoice1 = nil
			}
		case bitChanVoice2, ok := <-chanVoice2:
			if ok {
				m.setRecivedVoice2(bitChanVoice2)
				m.waitAnotherChan(chanVoice1, &m.voice1)
			} else {
				chanVoice2 = nil
			}
		}

		m.sendComparedVoice()

	}
}

func (m *Mixer) waitAnotherChan(ch chan int, voice *VoiceBit) {
	if ch != nil {
		bitVoice := <-ch
		voice.volume = bitVoice
	}
}

func (m *Mixer) sendComparedVoice() {
	m.outputChanel <- m.compareWithCompressor(m.voice1.volume, m.voice2.volume)
}

func (m *Mixer) compareWithCompressor(bitVoice1, bitVoice2 int) int {
	return m.audioCompressor.Compress(bitVoice1 + bitVoice2)
}

func (m *Mixer) setRecivedVoice1(volume int) {
	m.voice1.volume = volume
}

func (m *Mixer) setRecivedVoice2(volume int) {
	m.voice2.volume = volume
}
