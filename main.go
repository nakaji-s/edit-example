package main

import (
	"net/http"
	"os"

	"flag"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"
)

func main() {
	// parse arguments
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Println("need to set target filePath")
		os.Exit(1)
	}
	filePath := os.Args[1]

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file not found")
			os.Exit(1)
		} else {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// start server
	e := echo.New()

	e.File("/", "index.html")

	e.GET("/data", func(c echo.Context) error {
		return c.String(http.StatusOK, string(b))
	})

	e.PUT("/data", func(c echo.Context) error {
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filePath, body, 0644)
		if err != nil {
			return err
		}
		b = body
		return c.String(http.StatusOK, string(body))
	})

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}
