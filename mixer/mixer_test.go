package mixer

import (
	"os"
	"testing"

	"github.com/go-audio/audio"
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

func TestMixOneFile(t *testing.T) {
	file, err := os.Open("../samples/file_example_WAV_1MG.wav")

	if err != nil {
		t.Errorf("test file not exists")
	}

	defer file.Close()

}
