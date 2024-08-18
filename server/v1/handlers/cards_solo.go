package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"github.com/fogleman/gg"
	"image"
	"net/http"
)

func DrawCardSolo(w http.ResponseWriter, r *http.Request) error {
	cardImageName, err := context.GetParameterFromURL(r, "cardImageName", "Parameter cardImageName not found", w)
	if err != nil {
		return err
	}

	cardRarity, err := context.GetParameterFromURL(r, "cardRarity", "Parameter cardRarity not found", w)
	if err != nil {
		return err
	}

	cardImage, err := gg.LoadImage("./images/cards/" + cardImageName + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "Unknown card image for "+cardImageName, 400)
		return err
	}

	rarityFrame, err := gg.LoadImage("./images/cards_rarity/" + cardRarity + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "Unknown rarity", http.StatusBadRequest)
		return err
	}

	rd := context.Editor{Context: gg.NewContextForRGBA(image.NewRGBA(image.Rect(0, 0, 1050, 740)))}

	rd.DrawObject(
		cardImage,
		[2]int{285, 10},
		[2]int{0, 0},
		false,
	)
	rd.DrawObject(
		rarityFrame,
		[2]int{285, 10},
		[2]int{0, 0},
		false,
	)

	response, status := rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}
	return nil
}
