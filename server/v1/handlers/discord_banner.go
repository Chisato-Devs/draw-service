package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
)

func DrawDiscordBanner(w http.ResponseWriter, r *http.Request) error {
	guildCount, err := context.GetParameterFromURL(r, "guildCount", "Parameter guildCount not found", w)
	if err != nil {
		return err
	}

	usersCount, err := context.GetParameterFromURL(r, "usersCount", "Parameter usersCount not found", w)
	if err != nil {
		return err
	}

	bannerImage, _ := gg.LoadImage("./images/discord_banner.png")
	rd := context.Editor{Context: gg.NewContextForImage(bannerImage)}
	coloText := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	rd.DrawAlign(
		guildCount,
		100,
		[2]float64{914, 90},
		260,
		false,
		coloText,
	)

	rd.DrawAlign(
		usersCount,
		100,
		[2]float64{914, 291},
		260,
		false,
		coloText,
	)

	response, status := rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}
	return nil
}
