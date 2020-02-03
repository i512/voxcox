package compressor

//MockCompressorImpl this is mock for interface Compressor. Hi does nothing.
type MockCompressorImpl struct {
}

//Compress just return same value
func (m MockCompressorImpl) Compress(notCompressedSound int) int {
	return notCompressedSound
}
