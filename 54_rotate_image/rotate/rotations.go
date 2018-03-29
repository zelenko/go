package rotate

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

// FlipH flips the image horizontally (from left to right) and returns the transformed image.
func FlipH(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.w
	dstH := src.h
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcY := dstY
			src.scan(0, srcY, src.w, srcY+1, dst.Pix[i:i+rowSize])
			reverse(dst.Pix[i : i+rowSize])
		}
	})
	return dst
}

// FlipV flips the image vertically (from top to bottom) and returns the transformed image.
func FlipV(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.w
	dstH := src.h
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcY := dstH - dstY - 1
			src.scan(0, srcY, src.w, srcY+1, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}

// Transpose flips the image horizontally and rotates 90 degrees counter-clockwise.
func Transpose(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.h
	dstH := src.w
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstY
			src.scan(srcX, 0, srcX+1, src.h, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}

// Transverse flips the image vertically and rotates 90 degrees counter-clockwise.
func Transverse(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.h
	dstH := src.w
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstH - dstY - 1
			src.scan(srcX, 0, srcX+1, src.h, dst.Pix[i:i+rowSize])
			reverse(dst.Pix[i : i+rowSize])
		}
	})
	return dst
}

// Rotate90 rotates the image 90 degrees counter-clockwise and returns the transformed image.
func Rotate90(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.h
	dstH := src.w
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstH - dstY - 1
			src.scan(srcX, 0, srcX+1, src.h, dst.Pix[i:i+rowSize])
		}
	})
	return dst
}

// Rotate180 rotates the image 180 degrees counter-clockwise and returns the transformed image.
func Rotate180(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.w
	dstH := src.h
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcY := dstH - dstY - 1
			src.scan(0, srcY, src.w, srcY+1, dst.Pix[i:i+rowSize])
			reverse(dst.Pix[i : i+rowSize])
		}
	})
	return dst
}

// Rotate270 rotates the image 270 degrees counter-clockwise and returns the transformed image.
func Rotate270(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dstW := src.h
	dstH := src.w
	rowSize := dstW * 4
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))
	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			i := dstY * dst.Stride
			srcX := dstY
			src.scan(srcX, 0, srcX+1, src.h, dst.Pix[i:i+rowSize])
			reverse(dst.Pix[i : i+rowSize])
		}
	})
	return dst
}

// Rotate rotates an image by the given angle counter-clockwise .
// The angle parameter is the rotation angle in degrees.
// The bgColor parameter specifies the color of the uncovered zone after the rotation.
func Rotate(img image.Image, angle float64, bgColor color.Color) *image.NRGBA {
	angle = angle - math.Floor(angle/360)*360

	switch angle {
	case 0:
	return Clone(img)
	case 90:
		return Rotate90(img)
	case 180:
		return Rotate180(img)
	case 270:
		return Rotate270(img)
	}

	src := toNRGBA(img)
	//src := img
	srcW := src.Bounds().Max.X
	srcH := src.Bounds().Max.Y
	dstW, dstH := rotatedSize(srcW, srcH, angle)
	dst := image.NewNRGBA(image.Rect(0, 0, dstW, dstH))

	if dstW <= 0 || dstH <= 0 {
		return dst
	}

	srcXOff := float64(srcW)/2 - 0.5
	srcYOff := float64(srcH)/2 - 0.5
	dstXOff := float64(dstW)/2 - 0.5
	dstYOff := float64(dstH)/2 - 0.5

	bgColorNRGBA := color.NRGBAModel.Convert(bgColor).(color.NRGBA)
	sin, cos := math.Sincos(math.Pi * angle / 180)

	parallel(0, dstH, func(ys <-chan int) {
		for dstY := range ys {
			for dstX := 0; dstX < dstW; dstX++ {
				xf, yf := rotatePoint(float64(dstX)-dstXOff, float64(dstY)-dstYOff, sin, cos)
				xf, yf = xf+srcXOff, yf+srcYOff
				interpolatePoint(dst, dstX, dstY, src, xf, yf, bgColorNRGBA)
			}
		}
	})

	return dst
}

////

// parallel processes the data in separate goroutines.
func parallel(start, stop int, fn func(<-chan int)) {
	count := stop - start
	if count < 1 {
		return
	}

	procs := runtime.GOMAXPROCS(0)
	if procs > count {
		procs = count
	}

	c := make(chan int, count)
	for i := start; i < stop; i++ {
		c <- i
	}
	close(c)

	var wg sync.WaitGroup
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(c)
		}()
	}
	wg.Wait()
}

func rotatePoint(x, y, sin, cos float64) (float64, float64) {
	return x*cos - y*sin, x*sin + y*cos
}

func rotatedSize(w, h int, angle float64) (int, int) {
	if w <= 0 || h <= 0 {
		return 0, 0
	}

	sin, cos := math.Sincos(math.Pi * angle / 180)
	x1, y1 := rotatePoint(float64(w-1), 0, sin, cos)
	x2, y2 := rotatePoint(float64(w-1), float64(h-1), sin, cos)
	x3, y3 := rotatePoint(0, float64(h-1), sin, cos)

	minx := math.Min(x1, math.Min(x2, math.Min(x3, 0)))
	maxx := math.Max(x1, math.Max(x2, math.Max(x3, 0)))
	miny := math.Min(y1, math.Min(y2, math.Min(y3, 0)))
	maxy := math.Max(y1, math.Max(y2, math.Max(y3, 0)))

	neww := maxx - minx + 1
	if neww-math.Floor(neww) > 0.1 {
		neww++
	}
	newh := maxy - miny + 1
	if newh-math.Floor(newh) > 0.1 {
		newh++
	}

	return int(neww), int(newh)
}

func interpolatePoint(dst *image.NRGBA, dstX, dstY int, src *image.NRGBA, xf, yf float64, bgColor color.NRGBA) {
	dstIndex := dstY*dst.Stride + dstX*4

	x0 := int(math.Floor(xf))
	y0 := int(math.Floor(yf))
	bounds := src.Bounds()
	if !image.Pt(x0, y0).In(image.Rect(bounds.Min.X-1, bounds.Min.Y-1, bounds.Max.X, bounds.Max.Y)) {
		dst.Pix[dstIndex+0] = bgColor.R
		dst.Pix[dstIndex+1] = bgColor.G
		dst.Pix[dstIndex+2] = bgColor.B
		dst.Pix[dstIndex+3] = bgColor.A
		return
	}

	xq := xf - float64(x0)
	yq := yf - float64(y0)

	var pxs [4]color.NRGBA
	var cfs [4]float64

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			k := i*2 + j
			pt := image.Pt(x0+j, y0+i)
			if pt.In(bounds) {
				l := pt.Y*src.Stride + pt.X*4
				pxs[k].R = src.Pix[l+0]
				pxs[k].G = src.Pix[l+1]
				pxs[k].B = src.Pix[l+2]
				pxs[k].A = src.Pix[l+3]
			} else {
				pxs[k] = bgColor
			}
		}
	}

	cfs[0] = (1 - xq) * (1 - yq)
	cfs[1] = xq * (1 - yq)
	cfs[2] = (1 - xq) * yq
	cfs[3] = xq * yq

	var r, g, b, a float64
	for i := range pxs {
		wa := float64(pxs[i].A) * cfs[i]
		r += float64(pxs[i].R) * wa
		g += float64(pxs[i].G) * wa
		b += float64(pxs[i].B) * wa
		a += wa
	}

	if a != 0 {
		r /= a
		g /= a
		b /= a
	}

	dst.Pix[dstIndex+0] = clamp(r)
	dst.Pix[dstIndex+1] = clamp(g)
	dst.Pix[dstIndex+2] = clamp(b)
	dst.Pix[dstIndex+3] = clamp(a)
}

// clamp rounds and clamps float64 value to fit into uint8.
func clamp(x float64) uint8 {
	v := int64(x + 0.5)
	if v > 255 {
		return 255
	}
	if v > 0 {
		return uint8(v)
	}
	return 0
}

func toNRGBA(img image.Image) *image.NRGBA {
	if img, ok := img.(*image.NRGBA); ok {
		return &image.NRGBA{
			Pix:    img.Pix,
			Stride: img.Stride,
			Rect:   img.Rect.Sub(img.Rect.Min),
		}
	}
	return Clone(img)
}

func reverse(pix []uint8) {
	if len(pix) <= 4 {
		return
	}
	i := 0
	j := len(pix) - 4
	for i < j {
		pix[i+0], pix[j+0] = pix[j+0], pix[i+0]
		pix[i+1], pix[j+1] = pix[j+1], pix[i+1]
		pix[i+2], pix[j+2] = pix[j+2], pix[i+2]
		pix[i+3], pix[j+3] = pix[j+3], pix[i+3]
		i += 4
		j -= 4
	}
}

// Clone returns a copy of the given image.
func Clone(img image.Image) *image.NRGBA {
	src := newScanner(img)
	dst := image.NewNRGBA(image.Rect(0, 0, src.w, src.h))
	size := src.w * 4
	parallel(0, src.h, func(ys <-chan int) {
		for y := range ys {
			i := y * dst.Stride
			src.scan(0, y, src.w, y+1, dst.Pix[i:i+size])
		}
	})
	return dst
}
