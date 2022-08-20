package controllers

import (
	"net/http"
	"winning_fish_backend/config"
)

func StartMainSever() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
