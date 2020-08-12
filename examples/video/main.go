package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"strconv"

	"github.com/KimJeongChul/nvpipe"
	"gocv.io/x/gocv"
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

	video, err := gocv.OpenVideoCapture("input.mp4") // h264 video file
	if err != nil {
		fmt.Println("[ERROR] gocv open video capture error : ", err)
	}
	defer video.Close()

	mat := gocv.NewMat()
	defer mat.Close()

	fileIdx := 0
	for {
		if ok := video.Read(&img); !ok {
			return
		}

		if img.Empty() {
			break
		}

		img, err := mat.ToImage() // RGB 3 channel
		rect := img.Bounds()
		rgba := image.NewRGBA(rect)

		// NvPipe Encoding
		encodeSize := width * height
		encodeData := make([]byte, encodeSize)
		n := encoder.Encode(rgba.Pix, encodeData)
		if n == 0 {
			fmt.Println("[ERROR] nvdec error")
		}

		// Nvpipe Decoding
		decodeData := make([]uint8, width*height*rgbaChannel) // 1280 * 720 * 4 = 3686400
		n := decoder.Decode(encodeData[:n], n, output)

		jpgFile, err := os.Create("output" + strconv.Itoa(fileIdx) + ".jpg")
		if err != nil {
			fmt.Println("[ERROR] image file create : ", err)
		}

		// Preprocessing using image.RGBA
		start := image.Point{0, 0}
		end := image.Point{width, height}
		img := image.NewRGBA(image.Rectangle{start, end})

		// RGBA
		rgbaSlice := make([]color.RGBA, 0)
		for i := 0; i < width*height; i++ {
			r := decodeData[4*i+0]
			g := decodeData[4*i+1]
			b := decodeData[4*i+2]
			a := decodeData[4*i+3]
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

		fileIdx++
	}

}
