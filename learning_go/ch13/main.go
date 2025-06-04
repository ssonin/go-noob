package main

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type DateTime struct {
	DayOfWeek  time.Weekday `json:"day_of_week"`
	DayOfMonth int          `json:"day_of_month"`
	Month      time.Month   `json:"month"`
	Year       int          `json:"year"`
	Hour       int          `json:"hour"`
	Minute     int          `json:"minute"`
	Second     int          `json:"second"`
}

func NewDateTime(t time.Time) *DateTime {
	return &DateTime{
		DayOfWeek:  t.Weekday(),
		DayOfMonth: t.Day(),
		Month:      t.Month(),
		Year:       t.Year(),
		Hour:       t.Hour(),
		Minute:     t.Minute(),
		Second:     t.Second(),
	}
}

func Logger(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.LogAttrs(c, slog.LevelInfo, "Incoming request",
			slog.String("ip", c.ClientIP()))
		c.Next()
	}
}

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(Logger(logger))

	r.GET("/datetime", func(c *gin.Context) {
		c.JSON(http.StatusOK, NewDateTime(time.Now()))
	})
	r.Run()
}
