package compressor

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DynamicRangeCompressor", func() {

	It("Should return same value", func() {

		compressor := NewDynamicRangeCompressor(2.0, 0.5, 1, 2)

		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
		Expect(0.5).To(Equal(compressor.Compress(0.5)))
	})

	It("Should return compresed value with attack delay", func() {
		compressor := NewDynamicRangeCompressor(1.9, 0.5, 1, 2)

		Expect(2.0).To(Equal(compressor.Compress(2.0)))
		Expect(1.0).To(Equal(compressor.Compress(2.0)))
		Expect(0.5).To(Equal(compressor.Compress(1.0)))
		Expect(0.5).To(Equal(compressor.Compress(1.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
	})

	It("Should return compresed value without attack delay", func() {
		compressor := NewDynamicRangeCompressor(1.9, 0.5, 0, 2)

		Expect(1.0).To(Equal(compressor.Compress(2.0)))
		Expect(0.5).To(Equal(compressor.Compress(1.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
		Expect(5.0).To(Equal(compressor.Compress(10.0)))
		Expect(0.5).To(Equal(compressor.Compress(1.0)))
	})

	It("Should return compresed value without attack delay and witout release delay", func() {
		compressor := NewDynamicRangeCompressor(1.9, 0.5, 0, 0)

		Expect(1.0).To(Equal(compressor.Compress(2.0)))
		Expect(1.0).To(Equal(compressor.Compress(2.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
		Expect(5.0).To(Equal(compressor.Compress(10.0)))
		Expect(1.0).To(Equal(compressor.Compress(1.0)))
	})
})

func TestCompressor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Books Suite")
}
