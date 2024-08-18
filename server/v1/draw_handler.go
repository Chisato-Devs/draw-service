package v1

import (
	"chisato-draw-service/server/v1/handlers"
	"chisato-draw-service/server/v1/handlers/context"
	"net/http"
)

func DrawGetRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name, err := context.GetParameterFromURL(r, "name", "You need to specify the name of the image", w)
		if err != nil {
			return
		}

		switch name {
		case "level_card":
			err := handlers.DrawLevelCard(w, r)
			if err != nil {
				return
			}
		case "economy_profile":
			err := handlers.DrawEconomyProfile(w, r)
			if err != nil {
				return
			}
		case "love_banner":
			err := handlers.DrawLoveBanner(w, r)
			if err != nil {
				return
			}
		case "guild_banner":
			err := handlers.DrawGuildBanner(w, r)
			if err != nil {
				return
			}
		case "cards_solo":
			err := handlers.DrawCardSolo(w, r)
			if err != nil {
				return
			}
		case "cards_trade_frame":
			err := handlers.DrawCardTradeFrame(w, r)
			if err != nil {
				return
			}
		case "cards_trio":
			err := handlers.DrawCardTrio(w, r)
			if err != nil {
				return
			}
		case "discord_banner":
			err := handlers.DrawDiscordBanner(w, r)
			if err != nil {
				return
			}
		case "music_card":
			err := handlers.DrawMusicCard(w, r)
			if err != nil {
				return
			}
		case "playlist_card":
			err := handlers.DrawPlaylistCard(w, r)
			if err != nil {
				return
			}
		default:
			context.CreateErrorResponse(w, "An unknown card name is listed", http.StatusBadRequest)
		}
	} else {
		context.CreateErrorResponse(w, "Unknown method used, please use GET method", http.StatusBadRequest)
		return
	}
}
