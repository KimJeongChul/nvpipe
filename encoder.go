package nvpipe

/*
#include <NvPipe.h>
#cgo LDFLAGS: -L./ -lNvPipe
*/
import "C"
import (
	"unsafe"
)

// Encoder ...
type Encoder struct {
	// NvPipe Encoder
	enc    unsafe.Pointer
	width  int
	height int
}

// NewEncoder Creates a new encoder instance.
/**
 * @param format Format of input frame.
 * @param codec Possible codecs are H.264 and HEVC if available.
 * @param compression Lossy or lossless compression.
 * @param bitrate Bitrate in bit per second, e.g., 32 * 1000 * 1000 = 32 Mbps (for lossy compression only).
 * @param targetFrameRate At this frame rate the effective data rate approximately equals the bitrate (for lossy compression only).
 * @param width Initial width of the encoder.
 * @param height Initial height of the encoder.
 * @return NULL on error.
 */
func NewEncoder(format NvPipeFormat, codec NvPipeCodec, compression NvPipeCompression, bitrateMbps int, targetFPS int, width int, height int) *Encoder {
	var encoder Encoder
	enc := C.NvPipe_CreateEncoder(
		C.NvPipe_Format(format),
		C.NvPipe_Codec(codec),
		C.NvPipe_Compression(compression),
		C.ulong(bitrateMbps*1000*1000),
		C.uint32_t(targetFPS),
		C.uint32_t(width),
		C.uint32_t(height),
	)
	encoder.width = width
	encoder.height = height
	encoder.enc = enc
	return &encoder
}

// Encode a single frame from device or host memory.
/**
 * @param nvp Encoder instance.
 * @param src Device or host memory pointer.
 * @param srcPitch Pitch of source memory.
 * @param dst Host memory pointer for compressed output.
 * @param dstSize Available space for compressed output.
 * @param width Width of input frame in pixels.
 * @param height Height of input frame in pixels.
 * @param forceIFrame Enforces an I-frame instead of a P-frame.
 * @return Size of encoded data in bytes or 0 on error.
 */
func (encoder *Encoder) Encode(src []uint8, dst []byte) int {
	srcPitch := 4 * encoder.width
	forceIFrame := false
	dstSize := 4 * encoder.width * encoder.height

	//cstr := C.CString(string(src[:]))
	in := unsafe.Pointer(&src[0])
	//in := unsafe.Pointer(&cstr)                  // const void* src
	out := (*C.uint8_t)(unsafe.Pointer(&dst[0])) // uint8_t*

	size := C.NvPipe_Encode(
		encoder.enc,
		in,
		C.uint64_t(srcPitch),
		out,
		C.uint64_t(dstSize),
		C.uint32_t(encoder.width),
		C.uint32_t(encoder.height),
		C.bool(forceIFrame),
	)

	return int(size)
}

/*
func (encoder *Encoder) Encode(src []byte, dst []int8) int {
	srcPitch := 4 * encoder.width
	forceIFrame := false
	dstSize := 4 * encoder.width * encoder.height

	cstr := C.CString(string(src[:]))
	in := unsafe.Pointer(&cstr)                  // const void* src
	out := (*C.uint8_t)(unsafe.Pointer(&dst[0])) // uint8_t*

	size := C.NvPipe_Encode(
		encoder.enc,
		in,
		C.uint64_t(srcPitch),
		out,
		C.uint64_t(dstSize),
		C.uint32_t(encoder.width),
		C.uint32_t(encoder.height),
		C.bool(forceIFrame),
	)

	return int(size)
}

*/
