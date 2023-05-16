package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
)

func main() {
	e := echo.New()

	// remarsh, err := json.Marshal(mod)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	e.GET("/", func(c echo.Context) error {
		jsonstr := `
		{
			"name": "laravel/lara*vel",
			"description": "The Laravel Framework.",
			"license": "MIT",
			"type": "åäöÅÄÖ*-_.<>!@#$%^&*_+"
		}
		`

		mod := &ComposerJson{}
		err := json.Unmarshal([]byte(jsonstr), mod)

		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, err.Error())
		}

		if errs := validator.Validate(mod); errs != nil {
			fmt.Println(errs)
			return c.JSON(http.StatusInternalServerError, errs)
			// c.String(http.StatusInternalServerError, errs.Error())
		}

		return c.JSON(http.StatusOK, mod)
		// return c.String(http.StatusOK, json.Marshal(mod)[0])
	})

	e.Logger.Fatal(e.Start(":8123"))
}
