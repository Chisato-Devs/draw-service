package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"errors"
	"github.com/fogleman/gg"
	"net/http"
)

func DrawLoveBanner(w http.ResponseWriter, r *http.Request) error {
	bannerName, err := context.GetParameterFromURL(r, "bannerName", "Parameter bannerName not found", w)
	if err != nil {
		return err
	}

	firstAvatar, err := context.GetParameterFromURL(r, "firstAvatar", "Parameter firstAvatar not found", w)
	if err != nil {
		return err
	}

	secondAvatar, err := context.GetParameterFromURL(r, "secondAvatar", "Parameter secondAvatar not found", w)
	if err != nil {
		return err
	}

	bannerImage, err := gg.LoadImage("./images/love/" + bannerName + ".png")
	if err != nil {
		context.CreateErrorResponse(w, "Banner image with "+bannerName+" not found", http.StatusBadRequest)
		return err
	}

	rd := context.Editor{Context: gg.NewContextForImage(bannerImage)}

	firstImage, _, status := context.GetImageFromUrl(firstAvatar)
	if status == 400 && firstImage == nil {
		context.CreateErrorResponse(w, "Not found member avatar", 400)
	}
	rd.DrawObject(
		firstImage,
		[2]int{86, 208},
		[2]int{240, 240},
		true,
	)

	secondImage, stringErr, status := context.GetImageFromUrl(secondAvatar)
	if status == 400 && secondImage == nil {
		context.CreateErrorResponse(w, "Not found member avatar", 400)
		return errors.New(stringErr)
	}
	rd.DrawObject(
		secondImage,
		[2]int{986, 208},
		[2]int{240, 240},
		true,
	)

	response, status := rd.Save()
	if status == 200 {
		context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	} else {
		context.CreateErrorResponse(w, response, status)
	}
	return nil
}
