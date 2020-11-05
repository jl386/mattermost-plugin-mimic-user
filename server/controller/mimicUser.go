package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/v5/model"

	"github.com/jl386/mattermost-plugin-mimic-user/server/config"
)

var mimicUser = &Endpoint{
	Path:         "/create-post",
	Method:       http.MethodPost,
	Execute:      handleMimicUser,
	RequiresAuth: true,
}

func handleMimicUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	boaUsername := r.Header.Get("X-BOA-HEADER")
	config.Mattermost.LogInfo(boaUsername)
	user, err := config.Mattermost.GetUserByUsername(boaUsername)
	if err != nil {
		config.Mattermost.LogError("Error converting username", err.Error())
	}
	config.Mattermost.LogInfo(user.Id)

	decoder := json.NewDecoder(r.Body)
	post := model.Post{}
	if err := decoder.Decode(&post); err != nil {
		config.Mattermost.LogError("Error decoding post params: ", err.Error())
		return
	}
	post.UserId = user.Id

	if config.Mattermost.HasPermissionToChannel(post.UserId, post.ChannelId, model.PERMISSION_CREATE_POST) {
		savedPost, err := config.Mattermost.CreatePost(&post)
		if err != nil {
			config.Mattermost.LogError("Error creating post", err.Error())
			json, _ := json.Marshal(err)
			w.Write(json)
		}
		json, _ := json.Marshal(savedPost)
		w.Write(json)
		return
	}
	w.WriteHeader(http.StatusForbidden)

}
