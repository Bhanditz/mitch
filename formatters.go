package mitch

func FormatUser(user *User) Any {
	res := Any{
		"id":           user.ID,
		"gamer":        user.Gamer,
		"developer":    user.Developer,
		"press_user":   user.PressUser,
		"display_name": user.DisplayName,
		"username":     user.Username,
		"url":          "http://example.org",
		"cover_url":    "http://example.org",
	}
	if user.AllowTelemetry {
		res["allow_telemetry"] = true
	}
	return res
}

func FormatGame(game *Game) Any {
	res := Any{
		"id":             game.ID,
		"user_id":        game.UserID,
		"title":          game.Title,
		"min_price":      game.MinPrice,
		"type":           game.Type,
		"classification": game.Classification,
	}
	return res
}

func FormatUpload(upload *Upload) Any {
	res := Any{
		"id":       upload.ID,
		"game_id":  upload.GameID,
		"type":     upload.Type,
		"storage":  upload.Storage,
		"size":     upload.Size,
		"filename": upload.Filename,
		"url":      upload.URL,
	}
	platforms := Any{}
	if upload.PlatformLinux {
		platforms["linux"] = "all"
	}
	if upload.PlatformWindows {
		platforms["windows"] = "all"
	}
	if upload.PlatformMac {
		platforms["osx"] = "all"
	}
	res["platforms"] = platforms
	return res
}

func FormatUploads(uploads []*Upload) []Any {
	var res []Any
	for _, u := range uploads {
		res = append(res, FormatUpload(u))
	}
	return res
}
