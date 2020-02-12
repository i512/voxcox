package compressor

//Compressor interface for sound compressor's
type Compressor interface {
	Compress(notCompressedSound float64) float64
}
