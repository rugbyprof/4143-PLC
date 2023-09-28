## Image Pixel Colors

- Go has a standard library package called `image` that you can use for basic image manipulation tasks, including reading pixel data. 
- If you need more advanced functionality, there's also a third-party package called `github.com/disintegration/imaging`, which extends the standard `image` package.

Here's a simple example using the standard `image` package to read the color of each pixel in an image:

```go
package main

import (
	"fmt"
	"image"
	_ "image/png" // import this package to decode PNGs
	"os"
)

func main() {
	reader, err := os.Open("image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
```

- Replace the image file named `image.png` with your own file path.
- You'll notice the line `_ "image/png"`. The underscore is used to import a package solely for its initialization side-effects. 
- This particular import enables your code to decode PNG files. 
- If you want to support JPEG files, you'd add `_ "image/jpeg"` in a similar manner.
