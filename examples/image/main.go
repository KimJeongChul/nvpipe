package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/KimJeongChul/nvpipe"
)

func main() {
	codec := nvpipe.NvPipeH264
	if codec == nvpipe.NvPipeH264 {
		fmt.Println("NvPipe_Codec : H.264")
	} else {
		fmt.Println("NvPipe_Codec : HEVC")
	}

	compression := nvpipe.NvPipeLossless

	width := 1280
	height := 720
	rgbaChannel := 4

	fmt.Println("width : ", width, "height : ", height)

	bitrateMbps := 32
	targetFPS := 30

	fmt.Println("bitrate : ", bitrateMbps, " targetFPS : ", targetFPS)

	// NVDEC
	decoder := nvpipe.NewDecoder(nvpipe.NvPipeRGBA32, nvpipe.NvPipeH264, width, height)

	// NVENC
	encoder := nvpipe.NewEncoder(nvpipe.NvPipeRGBA32, nvpipe.NvPipeH264, nvpipe.NvPipeLossless, bitrateMbps, targetFPS, width, height)

	// Input jpg file
	imageFile, err := os.Open("input.jpg")
	if err != nil {
		fmt.Println("[ERROR] image file open : ", err)
	}
	defer imageFile.Close()

	// Image Decode
	src, _, err := image.Decode(imageFile)
	if err != nil {
		fmt.Println("[ERROR] image decode error : ", err)
	}

	// h264 Encoder
	h264Buf := bytes.NewBuffer(make([]byte, 0))

	opts := &x264.Options{
		Width:     width,
		Height:    height,
		FrameRate: targetFPS,
		Preset:    "veryfast",
		Tune:      "zerolatency",
		Profile:   "baseline",
	}

	enc, err := x264.NewEncoder(h264Buf, opts)
	err = enc.Encode(src)
	if err != nil {
		fmt.Println("[ERROR] x264 encode error : ", err)
	}

	// Image Data with 4 channel rgba
	decodeData := make([]uint8, width*height*rgbaChannel) // 1280 * 720 * 4 = 3686400

	// NvPipe Decode
	n = decoder.Decode(h264Buf.Bytes(), h264Buf.Len(), decodeData)

	// Save to Decode Data
	jpgFile, err := os.Create("output.jpg")
	if err != nil {
		fmt.Println("[ERROR] jpeg file open error : ", err)
	}

	// Preprocessing using image.RGBA
	start := image.Point{0, 0}
	end := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{start, end})

	// RGBA
	rgbaSlice := make([]color.RGBA, 0)
	for i := 0; i < width*height; i++ {
		r := output[4*i+0]
		g := output[4*i+1]
		b := output[4*i+2]
		a := output[4*i+3]
		rgbaSlice = append(rgbaSlice, color.RGBA{r, g, b, a})
	}

	// Draw RGBA
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, rgbaSlice[y*width+x])
		}
	}

	// File Write
	jpeg.Encode(jpgFile, img, nil)
}
