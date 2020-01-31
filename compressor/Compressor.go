package compressor

//Compressor interface for sound compressor's
type Compressor interface {
	compress(notCompressedSound []int) []int
}
