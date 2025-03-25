package main

import (
	"comm/service"
	pbProxy "proto/proxy"

	"github.com/micro/plugins/v5/proxy/http"
)

var (
	homeUrl        = "/"
	listWeatherUrl = "https://aider.meizu.com/app/weather/listWeather"
)

func main() {
	router := http.NewSingleHostRouter()
	router.RegisterEndpoint("MicroService.Home", homeUrl)
	router.RegisterEndpoint("MeiZuService.ListWeather", listWeatherUrl)

	service := service.NewService(
		service.Name("proxy"),
		service.Router(router),
	)
	pbProxy.RegisterMicroServiceProxy(service.Server())
	pbProxy.RegisterMeiZuServiceProxy(service.Server())
	service.Run()
}
