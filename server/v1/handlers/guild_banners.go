package handlers

import (
	"chisato-draw-service/server/v1/handlers/context"
	"chisato-draw-service/server/v1/handlers/guild_banners"
	"net/http"
)

func DrawGuildBanner(w http.ResponseWriter, r *http.Request) error {
	bannerName, err := context.GetParameterFromURL(r, "bannerName", "Parameter bannerName not found", w)
	if err != nil {
		return err
	}

	guildLanguage, err := context.GetParameterFromURL(r, "guildLanguage", "Parameter guildLanguage not found", w)
	if err != nil {
		return err
	}

	guildMembers, err := context.GetParameterFromURL(r, "guildMembers", "Parameter guildMembers not found", w)
	if err != nil {
		return err
	}

	voiceMembers, err := context.GetParameterFromURL(r, "voiceMembers", "Parameter voiceMembers not found", w)
	if err != nil {
		return err
	}

	activityMemberAvatarUrl, err := context.GetParameterFromURL(r, "activityMemberAvatarUrl", "Parameter activityMemberAvatarUrl not found", w)
	if err != nil {
		return err
	}

	activityMemberName, err := context.GetParameterFromURL(r, "activityMemberName", "Parameter activityMemberName not found", w)
	if err != nil {
		return err
	}

	activityMemberStatus, err := context.GetParameterFromURL(r, "activityMemberStatus", "Parameter activityMemberStatus not found", w)
	if err != nil {
		return err
	}

	switch bannerName {
	case "yellow":
		return guild_banners.DrawRotated(
			w,
			bannerName,
			guildLanguage,
			guildMembers,
			voiceMembers,
			activityMemberAvatarUrl,
			activityMemberName,
			activityMemberStatus,
		)
	case "blue":
		return guild_banners.DrawRotated(
			w,
			bannerName,
			guildLanguage,
			guildMembers,
			voiceMembers,
			activityMemberAvatarUrl,
			activityMemberName,
			activityMemberStatus,
		)
	case "green":
		return guild_banners.DrawNoRotated(
			w,
			bannerName,
			guildLanguage,
			guildMembers,
			voiceMembers,
			activityMemberAvatarUrl,
			activityMemberName,
			activityMemberStatus,
		)
	case "pink":
		return guild_banners.DrawNoRotated(
			w,
			bannerName,
			guildLanguage,
			guildMembers,
			voiceMembers,
			activityMemberAvatarUrl,
			activityMemberName,
			activityMemberStatus,
		)
	default:
		context.CreateErrorResponse(w, "Got unknown banner name...", http.StatusBadRequest)
	}

	//response, status := rd.Save()
	//if status == 200 {
	//	context.CreateSuccessResponse(w, structs.OKResponse{Encode: response})
	//} else {
	//	context.CreateErrorResponse(w, response, status)
	//}
	return nil
}
