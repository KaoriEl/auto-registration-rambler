package imageCrop

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"runtime"
)

func Crop(name string) {
	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())

	// load original image
	img, err := imaging.Open(name)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// crop from center
	centercropimg := imaging.Crop(img, image.Rect(1000, 650, 1300, 800))

	// save cropped image
	err = imaging.Save(centercropimg, name)
}
