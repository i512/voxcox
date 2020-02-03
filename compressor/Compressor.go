package compressor

//Compressor interface for sound compressor's
type Compressor interface {
	Compress(notCompressedSound int) int
}
