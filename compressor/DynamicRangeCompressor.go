package compressor

//DynamicRangeCompressor can process operation that reduces the volume of loud sounds
//Treadshold it is a volume-level above which it is necessary to begin compression
//Attack it is a delay for begin compression after exceeding the threshold
//release it is a time while compression will work
//compretionRatio it is a coefficient by which the signal will be multiplied
type DynamicRangeCompressor struct {
	threshold       float64
	attack          int
	release         int
	compretionRatio float64

	attackTimer  int
	releaseTimer int

	isInitialize            bool
	thresholHasBeenExceeded bool
}

//NewDynamicRangeCompressor creates new DynamicRangeCompressor to process sounds
func NewDynamicRangeCompressor(threshold, compretionRatio float64, attack, release int) *DynamicRangeCompressor {
	return &DynamicRangeCompressor{
		threshold:       threshold,
		attack:          attack,
		attackTimer:     attack,
		release:         release,
		compretionRatio: compretionRatio,
		isInitialize:    true,
	}
}

//Compress sound
func (d *DynamicRangeCompressor) Compress(notCompressedSound float64) float64 {

	if d == nil || !d.isInitialize {
		panic("dynamicRangeCOmpressor must be initialized. Use NewCompressor function")
	}

	compressedSound := notCompressedSound

	if notCompressedSound > d.threshold {
		d.releaseTimer = d.release
		d.thresholHasBeenExceeded = true
	}

	if d.attackTimer == 0 && d.releaseTimer == 0 && d.thresholHasBeenExceeded {
		d.attackTimer = d.attack
		d.thresholHasBeenExceeded = false
		compressedSound = d.compressFunction(notCompressedSound)

	} else if d.attackTimer == 0 && d.releaseTimer > 0 {
		compressedSound = d.compressFunction(notCompressedSound)
		d.releaseTimer--

	} else if d.thresholHasBeenExceeded {
		d.attackTimer--
	}

	return compressedSound
}

func (d *DynamicRangeCompressor) compressFunction(notCompressedSound float64) float64 {
	return notCompressedSound * d.compretionRatio
}
