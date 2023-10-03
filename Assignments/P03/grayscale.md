## GrayScaling

- Before converting an image to `GrayScale`, you need to gain access to each of its color pixels. You can see that done [HERE](./colors.md)
- Converting an image to `GrayScale` in Go can be done using the same techniques as the [color example](./colors.md) by using the standard `image` and `image/color` libraries. 
- The only difference is converting each RGB value to a `GrayScale` value. 
- This can be done by looping through each pixel, extract its RGB values, and then set it to a gray value based on some formula. 

### Formula

I guess averaging 3 numbers is a "formula", but its the easiest way to obtain the "gray" value:

```cpp
Gray = (R + G + B) / 3
```

And then you set that pixel to its new gray color like so:


```cpp
RGB(Gray,Gray,Gray)
```

However a better way to choose the gray value for each channel can be seen below: 


```cpp
Gray = (0.3 * R) + (0.59 * G) + (0.11 * B)
```

Where we take 30% of the red channel value + 59% of the green channel and 11% of the blue channel to calculate the GrayScale value.

Here's a simple code snippet that demonstrates how to convert an image to GrayScale:

```go
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	// Open the original image
	reader, err := os.Open("image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image
	grayFile, err := os.Create("gray_image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer grayFile.Close()
	png.Encode(grayFile, grayImg)

	fmt.Println("Grayscale image saved.")
}
```

- In this example, we first open and decode the original image. 
- Then, we create a new `Gray` image using `image.NewGray(bounds)`. 
- After that, we loop through each pixel of the original image, calculate the grayscale value, and set it in the new image. 
- Finally, we save the new grayscale image as a PNG file.

Just run the code, and you'll get a grayscale version of your original image saved as `gray_image.png`.