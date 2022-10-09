package config

import "github.com/kataras/blocks"

func Views() *blocks.Blocks {
	views := blocks.New("./views")

	err := views.Load()

	if err != nil {
		panic("Template engine can not be loaded")
	}

	return views
}
