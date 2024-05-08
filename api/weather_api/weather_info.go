package weather_api

import (
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (WeatherApi) WeatherInfoView(c *gin.Context) {
	city := c.Query("city")
	logrus.Info(city)
	res.Ok(models.WeatherMap[city], "获取天气信息成功", c)
}
