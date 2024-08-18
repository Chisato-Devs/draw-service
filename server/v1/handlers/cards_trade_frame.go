package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"fmt"
	"github.com/fogleman/gg"
	"net/http"
)

func DrawCardTradeFrame(w http.ResponseWriter, r *http.Request) error {
	firstCardImageName, err := context.GetParameterFromURL(r, "firstCardImageName", "Parameter firstCardImageName not found", w)
	if err != nil {
		return err
	}
	secondCardImageName, err := context.GetParameterFromURL(r, "secondCardImageUrl", "Parameter secondCardImageName not found", w)
	if err != nil {
		return err
	}

	firstCardRarity, err := context.GetParameterFromURL(r, "firstCardRarity", "Parameter firstCardRarity not found", w)
	if err != nil {
		return err
	}
	secondCardRarity, err := context.GetParameterFromURL(r, "secondCardRarity", "Parameter secondCardRarity not found", w)
	if err != nil {
		return err
	}

	firstCardImage, err := gg.LoadImage("./images/cards/" + firstCardImageName + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "Unknown card image for "+firstCardImageName, 400)
		return err
	}
	secondCardImage, err := gg.LoadImage("./images/cards/" + secondCardImageName + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "Unknown card image for "+secondCardImageName, 400)
		return err
	}

	firstRarityFrame, err := gg.LoadImage("./images/cards_rarity/" + firstCardRarity + ".png")
	if err != nil {
		context.CreateErrorResponse(w, fmt.Sprintf("Unknown rarity `%s`", firstCardRarity), http.StatusBadRequest)
		return err
	}
	secondRarityFrame, err := gg.LoadImage("./images/cards_rarity/" + secondCardRarity + ".png")
	if err != nil {
		context.CreateErrorResponse(w, fmt.Sprintf("Unknown rarity `%s`", secondCardRarity), http.StatusBadRequest)
		return err
	}

	tradeFrame, _ := gg.LoadImage("./images/trade_frame.png")
	rd := context.Editor{Context: gg.NewContextForImage(tradeFrame)}

	rd.DrawObject(
		firstCardImage,
		[2]int{0, 0},
		[2]int{0, 0},
		false,
	)
	rd.DrawObject(
		firstRarityFrame,
		[2]int{0, 0},
		[2]int{0, 0},
		false,
	)

	rd.DrawObject(
		secondCardImage,
		[2]int{917, 0},
		[2]int{0, 0},
		false,
	)
	rd.DrawObject(
		secondRarityFrame,
		[2]int{917, 0},
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
