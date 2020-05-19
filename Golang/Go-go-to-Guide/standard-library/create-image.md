# Create a New Image

* Use the `image`, `image/color`, and `image/png` packages to create a PNG image

```go
width := 200
height := 100

upLeft := image.Point{0, 0}
lowRight := image.Point{width, height}

img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
cyan := color.RGBA{100, 200, 200, 0xff}

// Set color for each pixel.
for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
        switch {
        case x < width/2 && y < height/2: // upper left quadrant
            img.Set(x, y, cyan)
        case x >= width/2 && y >= height/2: // lower right quadrant
            img.Set(x, y, color.White)
        default:
            // Use zero value.
        }
    }
}

// Encode as PNG.
f, _ := os.Create("image.png")
png.Encode(f, img)
```

## Go Image Support

* The `image` package implements a basic 2-D image library without painting or drawing functionality
* Additionally, the `image/draw` package provides image composition functions that can be used to perform a number of common image manipulation tasks
