package nvpipe

/*
#include <NvPipe.h>
#cgo LDFLAGS: -L./ -lNvPipe
*/
import "C"
import (
	"unsafe"
)

// Decoder ...
type Decoder struct {
	// NvPipe Decoder
	dec    unsafe.Pointer
	width  int
	height int
}

// NewDecoder Create a new decoder instance.
/**
 * @param format Format of output frame.
 * @param codec Possible codecs are H.264 and HEVC if available.
 * @param width Initial width of the decoder.
 * @param height Initial height of the decoder.
 * @return NULL on error.
 */
func NewDecoder(format NvPipeFormat, codec NvPipeCodec, width int, height int) *Decoder {
	var decoder Decoder
	dec := C.NvPipe_CreateDecoder(
		C.NvPipe_Format(format),
		C.NvPipe_Codec(codec),
		C.uint32_t(width),
		C.uint32_t(height),
	)
	decoder.dec = dec
	decoder.width = width
	decoder.height = height
	return &decoder
}

// Decode a single frame to device or host memory.
/**
 * @param nvp Decoder instance.
 * @param src Compressed frame data in host memory.
 * @param srcSize Size of compressed data.
 * @param dst Device or host memory pointer.
 * @param width Width of frame in pixels.
 * @param height Height of frame in pixels.
 * @return Size of decoded data in bytes or 0 on error.
 */
func (decoder *Decoder) Decode(src []byte, srcSize int, dst []uint8) int {

	in := (*C.uint8_t)(unsafe.Pointer(&src[0])) // uint8_t*
	out := unsafe.Pointer(&dst[0])

	size := C.NvPipe_Decode(
		decoder.dec,
		in,
		C.uint64_t(srcSize),
		out,
		C.uint32_t(decoder.width),
		C.uint32_t(decoder.height),
	)

	return int(size)
}

// Destroy Cleans up an decoder instance.
func (decoder *Decoder) Destroy() {
	C.NvPipe_Destroy(decoder.dec)
	decoder = nil
}

// GetError Returns an error message for the last error that occured.
/**
 * @return Returned string must not be deleted.
 */
func (decoder *Decoder) GetError() string {
	err := C.NvPipe_GetError(decoder.dec)
	return C.GoString(err)
}
