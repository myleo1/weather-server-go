package weather

import (
	"github.com/mizuki1412/go-core-kit/class"
	"github.com/mizuki1412/go-core-kit/service/restkit/context"
	"github.com/mizuki1412/go-core-kit/service/restkit/router"
	"weather-server-go/mod/weather/dao/weatherdao"
)

func Init(router *router.Router) {
	tag := "weather"
	r := router.Group("/rest/weather")
	{
		r.Post("/getCity", getCity).Tag(tag).Summary("获取地级市列表")
		r.Post("/getWeather", getWeather).Tag(tag).Summary("获取天气").Param(getWeatherParams{})
	}
}

func getCity(ctx *context.Context) {
	ret1 := weatherdao.CityRet()
	ctx.JsonSuccess(ret1)
}

type getWeatherParams struct {
	GetLive     class.Bool `description:"是否获取实时天气" default:"true"`
	GetForecast class.Bool `description:"是否获取预报天气" default:"false"`
	Location    string     `validate:"required" description:"城市编码"`
}

func getWeather(ctx *context.Context) {
	params := getWeatherParams{}
	ctx.BindForm(&params)
	var ret1, ret2 interface{}
	if params.GetLive.Bool {
		ret1 = weatherdao.GetWeather(params.Location + "00")
	}
	if params.GetForecast.Bool {
		ret2 = weatherdao.GetWeatherCast(params.Location + "00")
	}
	ctx.JsonSuccess(map[string]interface{}{
		"live":     ret1,
		"forecast": ret2,
	})
}
