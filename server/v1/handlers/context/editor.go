package context

import (
	"bytes"
	"chisato-draw-service/server/controllers"
	"encoding/base64"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
	"strconv"
	"strings"
)

type TextChars int

const (
	Upped TextChars = iota
	Lowered
	Default
)

type Editor struct {
	Context *gg.Context
}

func (rd *Editor) DrawAlign(
	name string,
	points float64,
	coords [2]float64,
	weight float64,
	shadow bool,
	rgba color.RGBA,
) {
	err := rd.Context.LoadFontFace("./Gilroy.ttf", points)
	if err != nil {
		fmt.Println("Error when loading font: ", err)
		return
	}

	textWeight, _ := rd.Context.MeasureString(name)
	x := (((coords[0] * 2) + weight) - textWeight) / 2
	if shadow {
		rd.Context.SetColor(color.RGBA{A: 255})
		rd.Context.DrawStringAnchored(name, x+2, coords[1]+2, 0, 1)
	}

	rd.Context.SetColor(rgba)
	rd.Context.DrawStringAnchored(name, x, coords[1], 0, 1.2)
}

func (rd *Editor) HexToRGBA(hex string) (color.RGBA, error) {
	var rgba color.RGBA

	r, err := strconv.ParseInt(hex[1:3], 16, 0)
	if err != nil {
		return rgba, err
	}
	g, err := strconv.ParseInt(hex[3:5], 16, 0)
	if err != nil {
		return rgba, err
	}
	b, err := strconv.ParseInt(hex[5:7], 16, 0)
	if err != nil {
		return rgba, err
	}

	rgba.R = uint8(r)
	rgba.G = uint8(g)
	rgba.B = uint8(b)
	rgba.A = 255

	return rgba, nil
}

func (rd *Editor) DrawSimple(
	name string,
	coords [2]float64,
	points float64,
	color color.RGBA,
) {
	err := rd.Context.LoadFontFace("./Gilroy.ttf", points)
	if err != nil {
		fmt.Println("Error when loading font: ", err)
		return
	}
	rd.Context.SetColor(color)
	rd.Context.DrawStringAnchored(name, coords[0], coords[1], 0, 1.2)
}

func (rd *Editor) DrawRight(
	name string,
	coords [2]float64,
	points float64,
	width float64,
	color color.RGBA,
) {
	err := rd.Context.LoadFontFace("./Gilroy.ttf", points)
	if err != nil {
		fmt.Println("Error when loading font: ", err)
		return
	}
	rd.Context.SetColor(color)
	rd.Context.DrawStringWrapped(
		name,
		coords[0],
		coords[1],
		0,
		0,
		width,
		0,
		gg.AlignRight,
	)
}

func (rd *Editor) trimTextToWidth(text string, maxWidth int) string {
	_points, _ := rd.Context.MeasureString("...")
	maxWidth -= int(_points)
	var trimmedText string
	for _, char := range text {
		charWidth, _ := rd.Context.MeasureString(string(char))
		if charWidth <= float64(maxWidth) {
			trimmedText += string(char)
			maxWidth -= int(charWidth)
		} else {
			trimmedText += "..."
			break
		}
	}
	return trimmedText
}

func (rd *Editor) TrimText(text string, points float64, maxWidth int, upper TextChars) string {
	err := rd.Context.LoadFontFace("./Gilroy.ttf", points)
	if err != nil {
		controllers.Logger().Warnf("Failed to truncate text: %v", err)
		return ""
	}
	trimmedText := rd.trimTextToWidth(text, maxWidth)

	switch upper {
	case Upped:
		return strings.ToUpper(trimmedText)
	case Lowered:
		return strings.ToLower(trimmedText)
	default:
		return trimmedText
	}
}

func (rd *Editor) DrawObject(obj image.Image, coords [2]int, size [2]int, circle bool) {
	var expandedObj image.Image
	if size[0] != 0 && size[1] != 0 {
		expandedObj = imaging.Resize(obj, size[0], size[1], imaging.Lanczos)
	} else {
		expandedObj = obj
	}

	if !circle {
		rd.Context.DrawImage(expandedObj, coords[0], coords[1])
	} else {
		mask := gg.NewContextForRGBA(
			image.NewRGBA(image.Rect(0, 0, size[0], size[1])),
		)

		mask.SetColor(color.White)
		mask.DrawEllipse(
			float64(size[0]/2),
			float64(size[1]/2),
			float64(size[0]/2),
			float64(size[1]/2),
		)
		mask.Fill()

		rgba := image.NewRGBA(image.Rect(0, 0, size[0], size[1]))
		draw.DrawMask(rgba, rgba.Bounds(), expandedObj, image.Point{}, mask.AsMask(), image.Point{}, draw.Over)
		rd.Context.DrawImage(rgba, coords[0], coords[1])
	}
}

func (rd *Editor) DrawWithMask(obj image.Image, maskImage image.Image, coords [2]int) {
	size := obj.Bounds().Size()
	mask := gg.NewContextForImage(imaging.Resize(maskImage, size.X, size.Y, imaging.Lanczos))

	rgba := image.NewRGBA(obj.Bounds())
	draw.DrawMask(rgba, rgba.Bounds(), obj, image.Point{}, mask.AsMask(), image.Point{}, draw.Over)
	rd.Context.DrawImage(rgba, coords[0], coords[1])
}

func (rd *Editor) Save() (string, int) {
	var buf bytes.Buffer

	err := png.Encode(&buf, rd.Context.Image())
	if err != nil {
		return "", http.StatusInternalServerError
	}

	encodedStr := base64.StdEncoding.EncodeToString(buf.Bytes())
	return encodedStr, http.StatusOK
}
