package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app.io/config"
	"app.io/data/model"
	"app.io/lib/db"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		// error handling
	}

	e := echo.New()

	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "PONG",
		})
	})

	e.POST("/log/submit", func(c echo.Context) error {
		submitData := new(model.SubmitNewData)
		if err = c.Bind(submitData); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// if err = c.Validate(request); err != nil {
		// 	return err
		// }
		jsonSubmitData, errJson := json.Marshal(submitData.Data)
		if errJson != nil {
			panic(errJson)
		}
		s := fmt.Sprintf("{\"id\":\"%s\",\"level\":\"INFO\",\"message\":\"%s\"}", uuid.New(), string(jsonSubmitData))
		err := db.Store(s, config.DefaultProject, submitData.Environment, submitData.Service)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"request": submitData,
				"success": false,
				"message": "Unable to submit query.",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"log":         submitData.Data,
			"project":     submitData.Project,
			"environment": submitData.Environment,
			"service":     submitData.Service,
			"request":     submitData,
			"success":     true,
			"message":     "New log submitted",
		})
	})

	// example query
	//
	// 1. Simple open regex          .........  ^f1-.*d3.*
	// 2. Both criteria should met   .........  ^(?=.*4e45)(?=.*INFO).*
	// 3. Three criteria should met  .........  ^(?=.*INFO)(?=.*d5c1)(?=.*test).*
	// 4. Line matching              .........  ^(?=.*e88.*ce).*
	e.POST("/log/query", func(c echo.Context) error {
		request := new(model.Request)
		if err = c.Bind(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		// if err = c.Validate(request); err != nil {
		// 	return err
		// }
		result := []interface{}{}
		data, errOfRead := db.Read(request.Query, config.DefaultProject, request.Environment, request.Service)
		if errOfRead != nil {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"request": request,
				"success": false,
				"message": "Unable to query.",
			})
		}
		for _, val := range data {
			// fmt.Printf("[INFO] %s \n", val)
			result = append(result, val)
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"query":       request.Query,
			"project":     request.Project,
			"environment": request.Environment,
			"service":     request.Service,
			"result":      result,
			"Total":       len(result),
		})
	})

	e.Logger.Fatal(e.Start(config.Host + ":" + config.Port))
}
