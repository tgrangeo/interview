package main

import (
    "image"
    "image/jpeg"
    "os"
)

func rotateImage90(img image.Image) image.Image {
    bounds := img.Bounds()
    rotated := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))
    for x := 0; x < bounds.Dx(); x++ {
        for y := 0; y < bounds.Dy(); y++ {
            rotated.Set(bounds.Dy()-y-1, x, img.At(x, y))
        }
    }
    return rotated
}

func main() {
    file, err := os.Open("test.jpg")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    img, _, err := image.Decode(file)
    if err != nil {
        panic(err)
    }
    rotatedImg := rotateImage90(img)
    outFile, err := os.Create("output.jpg")
    if err != nil {
        panic(err)
    }
    defer outFile.Close()
    err = jpeg.Encode(outFile, rotatedImg, nil)
    if err != nil {
        panic(err)
    }
}

