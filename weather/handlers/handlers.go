package handlers

import (
	"log"
	"time"

	"example.com/labwork_8/weather/cfg"
	"example.com/labwork_8/weather/db"
	"example.com/labwork_8/weather/model"
	owm "github.com/briandowns/openweathermap"
	"github.com/gofiber/fiber/v2"
)

func getForecastByGeo(lat, long float64, key string) *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "en", key)

	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByCoordinates(&owm.Coordinates{
		Longitude: lat,
		Latitude:  long,
	})

	return w
}

func ForecastNowHandler(c *fiber.Ctx) error {
	key := cfg.GetProperty("WEATHER_API_KEY")

	latitude := c.Locals("latitude").(float64)
	longtitude := c.Locals("longtitude").(float64)

	w := getForecastByGeo(latitude, longtitude, key)

	return c.JSON(
		model.ForecastData{
			Temperature: w.Main.Temp,
			WindSpeed:   w.Wind.Speed,
			Humidity:    w.Main.Humidity,
			Clouds:      w.Clouds.All,
			Latitude:    w.GeoPos.Latitude,
			Longtitude:  w.GeoPos.Longitude,
		})
}

func ForecastHistoryHandler(c *fiber.Ctx) error {
	db := db.DB

	dateFrom := c.Locals("dateFrom").(string)
	dateTo := c.Locals("dateTo").(string)

	var forecasts []model.ForecastDB

	if err := db.Where(
		"created_at BETWEEN TO_TIMESTAMP(? || ' 00:00:00', 'YYYY-MM-DD HH24:MI:SS')"+
			"AND TO_TIMESTAMP(? || ' 23:59:59', 'YYYY-MM-DD HH24:MI:SS')",
		dateFrom,
		dateTo).Find(&forecasts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to fetch data from db!"})
	}

	return c.JSON(forecasts)
}

func ForecastSheduler(duration time.Duration) {
	ticker := time.NewTicker(duration * time.Second)
	apiKey := cfg.GetProperty("WEATHER_API_KEY")

	defer ticker.Stop()

	for range ticker.C {
		for _, value := range cities {
			w := getForecastByGeo(value.Latitude, value.Longitude, apiKey)

			forecast := model.ForecastDB{
				ForecastData: model.ForecastData{
					Temperature: w.Main.Temp,
					WindSpeed:   w.Wind.Speed,
					Humidity:    w.Main.Humidity,
					Clouds:      w.Clouds.All,
					Latitude:    value.Latitude,
					Longtitude:  value.Longitude,
				},
				CreatedAt: time.Now().UTC(),
			}

			resp := db.DB.Create(&forecast)

			if resp.Error != nil {
				log.Println("Error sheduler error while fetching data!")
			} else {
				log.Println("Forecast sheduler saved data into db successfully!")
			}

		}
	}
}

var cities = map[string]*owm.Coordinates{
	"Lviv": {
		Latitude:  49.838300,
		Longitude: 24.023200,
	},
	"London": {
		Latitude:  51.508500,
		Longitude: -0.125700,
	},
	"Dnipro": {
		Latitude:  48.45,
		Longitude: 34.9833,
	},
}
