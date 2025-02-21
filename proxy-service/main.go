package main

import (
	"comm/service"
	pbProxy "proto/proxy"
	"proxy/handler"

	"github.com/micro/plugins/v5/proxy/http"
)

var (
	homeUrl        = "/"
	listWeatherUrl = "https://aider.meizu.com/app/weather/listWeather"
)

func main() {
	r := http.NewSingleHostRouter()
	r.RegisterEndpoint("MicroService.Home", homeUrl)
	r.RegisterEndpoint("MeiZuService.ListWeather", listWeatherUrl)
	service := service.NewService(
		service.Name("proxy"),
		service.Router(r),
	)
	h := handler.Handler{}
	pbProxy.RegisterMicroServiceHandler(service.Server(), &h)
	pbProxy.RegisterMeiZuServiceHandler(service.Server(), &h)
	service.Run()
}
