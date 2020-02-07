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
	volume   int
	received bool
}

//Compare two voices to one
func (m *Mixer) Compare(chanVoice1, chanVoice2, exit chan int, joinedVoices chan<- int) {

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
	voice.received = true
	if ch != nil {
		bitVoice := <-ch
		voice.volume = bitVoice
	}
}

func (m *Mixer) sendOneChanelIfAlreadyRecived() {
	m.emptyVoiceFiller()
	m.sendComparedVoice()
}

func (m *Mixer) sendComparedVoice() {
	m.outputChanel <- m.compareWithCompressor(m.voice1.volume, m.voice2.volume)
	m.setAllToUnrecived()
}
func (m *Mixer) isReadyToCompare() bool {
	return m.voice1.received && m.voice2.received
}

func (m *Mixer) compareWithCompressor(bitVoice1, bitVoice2 int) int {
	return m.audioCompressor.Compress(bitVoice1 + bitVoice2)
}

func (m *Mixer) emptyVoiceFiller() {
	if !m.voice1.received {
		m.setRecivedVoice1(0)
	}

	if !m.voice2.received {
		m.setRecivedVoice2(0)
	}
}

func (m *Mixer) setRecivedVoice1(volume int) {
	m.voice1.volume = volume
	m.voice1.received = true
}

func (m *Mixer) setRecivedVoice2(volume int) {
	m.voice2.volume = volume
	m.voice2.received = true
}
func (m *Mixer) setAllToUnrecived() {
	m.setUnrecivedVoice1()
	m.setUnrecivedVoice2()
}
func (m *Mixer) setUnrecivedVoice1() {
	m.voice1.received = false
}

func (m *Mixer) setUnrecivedVoice2() {
	m.voice2.received = false
}
