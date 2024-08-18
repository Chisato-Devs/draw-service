package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
)

func drawNothing(coords [2]float64, rd context.Editor) {
	rd.DrawAlign(
		"NOTHING",
		33,
		coords,
		144,
		false,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)
}

func DrawEconomyProfile(w http.ResponseWriter, r *http.Request) error {
	userAvatar, err := context.GetParameterFromURL(r, "userAvatar", "Parameter userAvatar not found", w)
	if err != nil {
		return err
	}

	userName, err := context.GetParameterFromURL(r, "userName", "Parameter userName not found", w)
	if err != nil {
		return err
	}

	moneyOnHands, err := context.GetParameterFromURL(r, "moneyOnHands", "Parameter moneyOnHands not found", w)
	if err != nil {
		return err
	}

	topPosition, err := context.GetParameterFromURL(r, "topPosition", "Parameter topPosition not found", w)
	if err != nil {
		return err
	}

	petUrl := r.URL.Query().Get("petUrl")

	petStamina, err := context.GetParameterFromURL(r, "petStamina", "Parameter petStamina not found", w)
	if err != nil {
		return err
	}

	petMana, err := context.GetParameterFromURL(r, "petMana", "Parameter petMana not found", w)
	if err != nil {
		return err
	}

	petLevel, err := context.GetParameterFromURL(r, "petLevel", "Parameter petLevel not found", w)
	if err != nil {
		return err
	}

	userImage, _ := gg.LoadImage("./images/economy_card.png")
	rd := context.Editor{Context: gg.NewContextForImage(userImage)}
	colo := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	avatarImage, response, status := context.GetImageFromUrl(userAvatar)
	if status != 200 {
		context.CreateErrorResponse(w, response, status)
		return err
	}

	rd.DrawObject(
		avatarImage,
		[2]int{101, 92},
		[2]int{153, 151},
		true,
	)

	rd.DrawAlign(
		rd.TrimText(userName, 30, 165, context.Upped),
		30,
		[2]float64{101, 274.8},
		142,
		false,
		colo,
	)

	rd.DrawAlign(
		moneyOnHands,
		33,
		[2]float64{433, 106},
		110,
		false,
		colo,
	)

	rd.DrawAlign(
		topPosition,
		33,
		[2]float64{433, 215},
		110,
		false,
		color.RGBA{R: 255, G: 255, B: 255, A: 255},
	)

	if petUrl != "" {
		petImage, response, status := context.GetImageFromUrl(petUrl)
		if status != 200 {
			context.CreateErrorResponse(w, response, status)
			return err
		}
		rd.DrawObject(
			petImage,
			[2]int{129, 363},
			[2]int{104, 104},
			false,
		)
	} else {
		greenStars, err := gg.LoadImage("./images/green_stars.png")
		if err != nil {
			context.CreateErrorResponse(w, "green_stars not found", http.StatusBadRequest)
			return err
		}
		rd.DrawObject(
			greenStars,
			[2]int{129, 363},
			[2]int{104, 104},
			false,
		)
	}

	if petStamina != "0" {
		rd.DrawAlign(
			petStamina,
			33,
			[2]float64{413.57, 317.75},
			144,
			false,
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		)
	} else {
		drawNothing(
			[2]float64{413.57, 317.75},
			rd,
		)
	}

	if petMana != "0" {
		rd.DrawAlign(
			petMana,
			33,
			[2]float64{413.57, 379.43},
			144,
			false,
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		)
	} else {
		drawNothing(
			[2]float64{413.57, 379.43},
			rd,
		)
	}

	if petStamina != "0" {
		rd.DrawAlign(
			petLevel,
			33,
			[2]float64{413.57, 442.21},
			144,
			false,
			color.RGBA{R: 255, G: 255, B: 255, A: 255},
		)
	} else {
		drawNothing(
			[2]float64{413.57, 442.21},
			rd,
		)
	}

	response, status = rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}
	return nil
}
