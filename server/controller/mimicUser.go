package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/Brightscout/mattermost-plugin-mimic-user/server/config"
)

var mimicUser = &Endpoint{
	Path:         "/create-post",
	Method:       http.MethodPost,
	Execute:      handleMimicUser,
	RequiresAuth: true,
}

func handleMimicUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	post := model.Post{}
	if err := decoder.Decode(&post); err != nil {
		config.Mattermost.LogError("Error decoding post params: ", err.Error())
		return
	}
	_, err := config.Mattermost.CreatePost(&post); if err != nil {
		config.Mattermost.LogError("Error creating post", err.Error())
	}
}
