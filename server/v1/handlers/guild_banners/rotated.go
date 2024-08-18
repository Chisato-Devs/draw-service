package guild_banners

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/structs"
	"errors"
	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"image/color"
	"net/http"
)

func DrawRotated(
	w http.ResponseWriter,
	bannerName, guildLanguage, guildMembers, voiceMembers, activityMemberAvatarUrl, activityMemberName, activityMemberStatus string,
) error {
	userImage, _ := gg.LoadImage("./images/banners/" + guildLanguage + "/" + bannerName + ".png")
	rd := context.Editor{Context: gg.NewContextForImage(userImage)}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	rd.DrawAlign(
		guildMembers,
		35,
		[2]float64{398, 278.79},
		110,
		false,
		white,
	)

	rd.DrawSimple(
		voiceMembers,
		[2]float64{342.5, 163.5},
		100,
		white,
	)

	avatarImage, err, status := context.GetImageFromUrl(activityMemberAvatarUrl)
	if status == 400 && avatarImage == nil {
		context.CreateErrorResponse(w, "Avatar image not found or discord threw an error", 500)
		return errors.New(err)
	} else {
		obj := imaging.Resize(avatarImage, 150, 150, imaging.Lanczos)
		maskImage, err := gg.LoadImage("./images/banners/masks/rotated.png")
		if err != nil {
			context.CreateErrorResponse(w, "maskImage not found", http.StatusBadRequest)
			return err
		}

		rd.DrawWithMask(obj, maskImage, [2]int{740, 329})
	}

	rd.DrawRight(
		rd.TrimText(activityMemberName, 60, 395, context.Upped),
		[2]float64{312.5, 360.79},
		60,
		395,
		white,
	)

	rd.DrawRight(
		rd.TrimText(activityMemberStatus, 25, 361, context.Upped),
		[2]float64{345.5, 416.29},
		25,
		361,
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
