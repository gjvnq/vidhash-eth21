// test different hash variants for images
package main

// TODO

// 1. pHash every frame
// 2. Calcualte first derivative (hamming distance)
// 2. Round low values to zero?
// 3. Calculate second derivative
// 4. Mark all inflexion points (zero second deriv) without near candidates
// 5. For the remaining candidates, select the "winners" by making the neighbouring segments as long as possible
// 6. Join very similar neighbouring segments (shortest first)
// 7. Compute the "average pHash" of each segment (will be a sequence of values between 0 and 1)
// 8. Make a list of triples (average pHash, start frame, end frame)


import "os"
import "fmt"
import "strconv"
import "strings"
import "path/filepath"
import "github.com/corona10/goimagehash/transforms"
import "image"
import "image/color"
import "image/draw"
import "image/png"
import _ "image/jpeg"
import "github.com/disintegration/imaging"

func panicIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}

func mess_img_up(srcImg image.Image, name string) ([]image.Image, []string) {
	ansImgs := make([]image.Image, 0)
	ansNames := make([]string, 0)
	
	// Blur
	ansImgs = append(ansImgs, imaging.Blur(srcImg, 0.1))
	ansNames = append(ansNames, name+"-blur-0.1")
	ansImgs = append(ansImgs, imaging.Blur(srcImg, 0.5))
	ansNames = append(ansNames, name+"-blur-0.5")
	ansImgs = append(ansImgs, imaging.Blur(srcImg, 0.9))
	ansNames = append(ansNames, name+"-blur-0.9")
	ansImgs = append(ansImgs, imaging.Blur(srcImg, 2.0))
	ansNames = append(ansNames, name+"-blur-2.0")
	ansImgs = append(ansImgs, imaging.Blur(srcImg, 7.0))
	ansNames = append(ansNames, name+"-blur-7.0")

	// Gamma
	ansImgs = append(ansImgs, imaging.AdjustGamma(srcImg, 0.75))
	ansNames = append(ansNames, name+"-gamma-0.75")
	ansImgs = append(ansImgs, imaging.AdjustGamma(srcImg, 1.25))
	ansNames = append(ansNames, name+"-gamma-1.25")

	// Contrast
	ansImgs = append(ansImgs, imaging.AdjustContrast(srcImg, 15))
	ansNames = append(ansNames, name+"-contrast-pos")
	ansImgs = append(ansImgs, imaging.AdjustContrast(srcImg, -15))
	ansNames = append(ansNames, name+"-contrast-neg")

	// Rotate
	ansImgs = append(ansImgs, imaging.Rotate90(srcImg))
	ansNames = append(ansNames, name+"-rot-90")
	ansImgs = append(ansImgs, imaging.Rotate180(srcImg))
	ansNames = append(ansNames, name+"-rot-280")
	ansImgs = append(ansImgs, imaging.Rotate270(srcImg))
	ansNames = append(ansNames, name+"-rot-270")

	// Flip
	ansImgs = append(ansImgs, imaging.FlipH(srcImg))
	ansNames = append(ansNames, name+"-flip-h")
	ansImgs = append(ansImgs, imaging.FlipV(srcImg))
	ansNames = append(ansNames, name+"-flip-v")
	ansImgs = append(ansImgs, imaging.FlipH(imaging.FlipV(srcImg)))
	ansNames = append(ansNames, name+"-flip-hv")

	// Border
	w := srcImg.Bounds().Dx()
	h := srcImg.Bounds().Dy()
	bordered := imaging.New(w, h, color.Black)
	r := image.Rect(w/10, h/10, 9*w/10, 9*h/10)
	draw.Draw(bordered, r, srcImg, image.Point{0,0}, draw.Src)
	ansImgs = append(ansImgs, bordered)
	ansNames = append(ansNames, name+"-bordered")

	return ansImgs, ansNames
}

func hash_img(fullImg image.Image, name string) []byte {
    // Reduce resolution
    smallImg := imaging.Resize(fullImg, 64, 64, imaging.Lanczos)
	
	// Remove colour
	baseImg := transforms.Rgb2Gray(smallImg)

    // Turn into signal space
    freqMap := transforms.DCT2D(baseImg, 16, 16)

    hash_val := make([]byte, )

    fmt.Println(freqMap)

	// file, err := os.OpenFile("gray.png", os.O_RDWR|os.O_CREATE, 0755)
	// panicIfNotNil(err)
	// png.Encode(file, baseImg)

    return nil
}

func main() {
	fmt.Println("Started")
	if len(os.Args) != 3 {
		fmt.Printf("USAGE: %s [SRC PATH] [OUT PATH]\n", os.Args[0])
		return
	}
	imgSrcPath := os.Args[1] 
	imgOutPath := os.Args[2]

	name := filepath.Base(imgSrcPath)
	name = strings.Split(name, ".")[0]
	imgSrcFile, err := os.Open(imgSrcPath)
	panicIfNotNil(err)

	imgSrc, _, err := image.Decode(imgSrcFile)
	panicIfNotNil(err)
	fmt.Println("Decoded "+name)
	newImgs, _ := mess_img_up(imgSrc, name)
	hash_img(imgSrc, name)
	for i, img := range newImgs {
		path := imgOutPath+"/"+strconv.Itoa(i)+".png"
		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
		panicIfNotNil(err)
		png.Encode(file, img)
		file.Close()
	}

	// Plot 2D: (distance from base img, distance from nearest other img)
}