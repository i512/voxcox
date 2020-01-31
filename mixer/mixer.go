package mixer

import (
	"errors"

	"github.com/go-audio/audio"
)

//Mixer TODO: Write description
type Mixer struct {
	comapreBlockSize int
	audioCompressor  Compressor
}

func compare(voice1, voice2, joinedVoices *audio.IntBuffer) error {

	if voice1 == nil || voice2 == nil || joinedVoices == nil {
		return errors.New("buffer's must be not nil")
	}

	if len(voice1.Data) != len(voice2.Data) && len(voice1.Data) != len(joinedVoices.Data) {
		return errors.New("voice buffer's must be equivalent lenght")
	}

	for i := 0; i < len(voice1.Data); i++ {
		joinedVoices.Data[i] = voice1.Data[i] + voice2.Data[i]
	}

	return nil
}

func simpleDataCompare(voiceBlock1, voiceBlock2 []int) []int {
	maxLen := len(voiceBlock2)
	for i := 0; i < maxLen; i++ {
		voiceBlock1[i] += voiceBlock2[i]
	}

	return voiceBlock1
}
