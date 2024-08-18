package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"errors"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
	"strconv"
)

func DrawLevelCard(w http.ResponseWriter, r *http.Request) error {
	levelValue, err := context.GetParameterFromURL(r, "levelValue", "Parameter levelValue not found", w)
	if err != nil {
		return err
	}

	prestigeValue, err := context.GetParameterFromURL(r, "prestigeValue", "Parameter prestigeValue not found", w)
	if err != nil {
		return err
	}

	userName, err := context.GetParameterFromURL(r, "userName", "Parameter userName not found", w)
	if err != nil {
		return err
	}

	userAvatar, err := context.GetParameterFromURL(r, "userAvatar", "Parameter userAvatar not found", w)
	if err != nil {
		return err
	}

	nowExp, err := context.GetParameterFromURL(r, "nowExp", "Parameter nowExp not found", w)
	if err != nil {
		return err
	}
	if nowExp == "0" {
		nowExp = "1"
	}

	needExp, err := context.GetParameterFromURL(r, "needExp", "Parameter needExp not found", w)
	if err != nil {
		return err
	}

	userImage, _ := gg.LoadImage("./images/level_card.png")
	rd := context.Editor{Context: gg.NewContextForImage(userImage)}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	userImage, stringErr, status := context.GetImageFromUrl(userAvatar)
	if status == 400 && userImage == nil {
		context.CreateErrorResponse(w, "Not found member avatar", 400)
		return errors.New(stringErr)
	}
	rd.DrawObject(
		userImage,
		[2]int{72, 98},
		[2]int{126, 125},
		true,
	)

	rd.DrawAlign(
		prestigeValue,
		27.23,
		[2]float64{107, 252.5},
		57,
		false,
		white,
	)

	rd.DrawAlign(
		rd.TrimText(userName, 61.72, 334, context.Upped),
		61.72,
		[2]float64{277, 109},
		321,
		false,
		white,
	)

	rgba, _ := rd.HexToRGBA("#373653")
	rd.DrawAlign(
		levelValue,
		27.53,
		[2]float64{557, 226.5},
		46,
		false,
		rgba,
	)

	nowExpInt, _ := strconv.Atoi(nowExp)
	needExpInt, _ := strconv.Atoi(needExp)

	rd.DrawRight(
		nowExp+" xp / "+needExp+" xp",
		[2]float64{416, 230.5},
		14,
		126,
		white,
	)

	if float64(needExpInt/nowExpInt) < 24.75 {
		hexToRGBA, err := rd.HexToRGBA("#E6E5FF")
		if err != nil {
			return err
		}
		rd.Context.SetColor(hexToRGBA)
		rd.Context.DrawRoundedRectangle(
			float64(283),
			250.5,
			float64(256*nowExpInt/needExpInt),
			10.4,
			5,
		)
		rd.Context.Fill()
	}

	response, status := rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}
	return nil
}
