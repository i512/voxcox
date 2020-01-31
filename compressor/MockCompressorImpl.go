package compressor

//MockCompressorImpl this is mock for interface Compressor. Hi does nothing.
type MockCompressorImpl struct {
}

//Just return same value
func (m MockCompressorImpl) compress(notCompressedSound []int) []int {
	return notCompressedSound
}
