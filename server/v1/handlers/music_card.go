package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"errors"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
)

func DrawMusicCard(w http.ResponseWriter, r *http.Request) error {
	musicArtwork, err := context.GetParameterFromURL(r, "musicArtwork", "Parameter musicArtwork not found", w)
	if err != nil {
		return err
	}

	musicName, err := context.GetParameterFromURL(r, "musicName", "Parameter musicName not found", w)
	if err != nil {
		return err
	}

	musicArtistName, err := context.GetParameterFromURL(r, "musicArtistName", "Parameter musicArtist not found", w)
	if err != nil {
		return err
	}

	musicSource, err := context.GetParameterFromURL(r, "musicSource", "Parameter musicSource not found", w)
	if err != nil {
		return err
	}

	musicFilter, err := context.GetParameterFromURL(r, "musicFilter", "Parameter musicFilter not found", w)
	if err != nil {
		return err
	}

	bannerImage, err := gg.LoadImage("./images/music_cards/" + musicFilter + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "music_card not found", http.StatusBadRequest)
		return err
	}

	rd := context.Editor{Context: gg.NewContextForImage(bannerImage)}

	if musicArtwork != "None" {
		artworkImage, err, status := context.GetImageFromUrl(musicArtwork)
		if status == 400 && artworkImage == nil {
			context.CreateErrorResponse(w, "Music image not found", 400)
			return errors.New(err)
		} else {
			expandedObj := imaging.Resize(artworkImage, 145, 145, imaging.Lanczos)

			maskImage, err := gg.LoadImage("./images/masks/music_artwork.png")
			if err != nil {
				context.CreateErrorResponse(w, "maskImage not found", http.StatusBadRequest)
				return err
			}

			rd.DrawWithMask(expandedObj, maskImage, [2]int{61, 102})
		}
	} else {
		noteImage, err := gg.LoadImage("./images/music_note.png")
		if err != nil {
			context.CreateErrorResponse(w, "music_note not found", http.StatusBadRequest)
			return err
		}
		rd.DrawObject(
			noteImage,
			[2]int{103, 131},
			[2]int{0, 0},
			false,
		)
	}

	sourceImage, err := gg.LoadImage("./images/music_sources/" + musicSource + ".png")
	if err == nil {
		rd.DrawObject(
			sourceImage,
			[2]int{164, 205},
			[2]int{38, 38},
			false,
		)
	}

	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	rd.DrawAlign(
		rd.TrimText(musicName, 55, 372, context.Default),
		55,
		[2]float64{258, 113},
		372,
		false,
		white,
	)

	rd.DrawAlign(
		rd.TrimText(musicArtistName, 40, 270, context.Default),
		40,
		[2]float64{254, 217},
		270,
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
