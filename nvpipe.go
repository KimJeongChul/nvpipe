package nvpipe

/*
#include <NvPipe.h>
#cgo LDFLAGS: -L./ -lNvPipe
*/
import "C"

// NvPipeCodec Available video codecs in NvPipe.
type NvPipeCodec int

const (
	// NvPipeH264 H264 Codec
	NvPipeH264 = NvPipeCodec(C.NVPIPE_H264)
	// NvPipeHEVC HEVC/H265 Codec
	NvPipeHEVC = NvPipeCodec(C.NVPIPE_HEVC)
)

// NvPipeCompression Compression type used for encoding.
type NvPipeCompression int

const (
	// NvPipeLossy ...
	NvPipeLossy = NvPipeCompression(C.NVPIPE_LOSSY)
	// NvPipeLossless produces larger output.
	NvPipeLossless = NvPipeCompression(C.NVPIPE_LOSSLESS)
)

// NvPipeFormat Format of the input frame.
type NvPipeFormat int

const (
	NvPipeRGBA32 = NvPipeFormat(C.NVPIPE_RGBA32)
	NvPipeUInt4  = NvPipeFormat(C.NVPIPE_UINT4)
	NvPipeUInt8  = NvPipeFormat(C.NVPIPE_UINT8)
	NvPipeUInt16 = NvPipeFormat(C.NVPIPE_UINT16)
	NvPipeUInt32 = NvPipeFormat(C.NVPIPE_UINT32)
)
