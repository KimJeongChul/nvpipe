# Go wrapper for NvPipe

This package provides Go bindings for the [NVIDIA Nvpipe](https://github.com/NVIDIA/NvPipe) libraries which is convenience wrapper around the low-level NVENC/NVDEC APIS in the official [NVIDIA Video Codec SDK](https://developer.nvidia.com/nvidia-video-codec-sdk).

## Import
```go
import "github.com/KimJeongChul/nvpipe"
```

## Encoding
```go
const (
  codec = nvpipe.NvPipeH264
  format = nvpipe.NvPipeRGBA32
  compression = nvpipe.NvPipeLossless
  bitrateMbps = 32
  targetFPS = 30
  width = 1280
  height = 720
)

encoder := nvpipe.NewEncoder(format, codec, compression, bitrateMbps, targetFPS, width, height)
encodeSize := width * height
encodeData := make([]byte, encodeSize)
n := encoder.Encode(rgba.ToBytes(), encodeData)
if n == 0 {
  fmt.Println("[ERROR] nvenc error : ", encoder.GetError())
}

log.Println(encodeData[:n])

encoder.Destroy()
```

## Decoding
```go
const (
  codec = nvpipe.NvPipeH264
  format = nvpipe.NvPipeRGBA32
  width = 1280
  height = 720
  rgbaChannel = 4
)

decoder := nvpipe.NewDecoder(format, codec, width, height)
decodeSize := width*height*rgbaChannel
decodeData := make([]uint8, decodeSize)
n := decoder.Decode(encodeData, len(encodeData), decodeData)
if n == 0 {
  fmt.Println("[ERROR] nvdnc error : ", decoder.GetError())
}

log.Println(decodeData[:n]) 

decoder.Destroy()
```

## Build & Installation
This package requires NVIDIA CUDA.
