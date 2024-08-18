package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
)

func DrawPlaylistCard(w http.ResponseWriter, r *http.Request) error {
	playlistName, err := context.GetParameterFromURL(r, "playlistName", "Parameter playlistName not found", w)
	if err != nil {
		return err
	}

	ownerAvatar, err := context.GetParameterFromURL(r, "ownerAvatar", "Parameter ownerAvatar not found", w)
	if err != nil {
		return err
	}

	ownerName, err := context.GetParameterFromURL(r, "ownerName", "Parameter ownerName not found", w)
	if err != nil {
		return err
	}

	tracksCount, err := context.GetParameterFromURL(r, "tracksCount", "Parameter tracksCount not found", w)
	if err != nil {
		return err
	}

	listenedCount, err := context.GetParameterFromURL(r, "listenedCount", "Parameter listenedCount not found", w)
	if err != nil {
		return err
	}

	cardName, err := context.GetParameterFromURL(r, "cardName", "Parameter cardName not found", w)
	if err != nil {
		return err
	}

	playlistImageBg, err := gg.LoadImage("./images/playlist_cards/" + cardName + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "playlist_card not found", http.StatusBadRequest)
		return err
	}
	rd := context.Editor{Context: gg.NewContextForImage(playlistImageBg)}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	ownerImage, _, status := context.GetImageFromUrl(ownerAvatar)
	if status == 400 && ownerImage == nil {
		context.CreateErrorResponse(w, "Not found member avatar", 400)
	}
	rd.DrawObject(
		ownerImage,
		[2]int{409, 270},
		[2]int{140, 140},
		true,
	)

	rd.DrawAlign(
		rd.TrimText(playlistName, 50, 541, context.Default),
		50,
		[2]float64{362, 97},
		541,
		false,
		white,
	)

	rd.DrawAlign(
		rd.TrimText(ownerName, 50, 169, context.Upped),
		25,
		[2]float64{393, 446},
		169,
		false,
		white,
	)

	rd.DrawAlign(
		tracksCount,
		30,
		[2]float64{734, 299},
		150,
		false,
		white,
	)

	rd.DrawAlign(
		listenedCount,
		30,
		[2]float64{734, 425},
		150,
		false,
		white,
	)

	response, status := rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}

	return nil
}
