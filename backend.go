package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readshowme() string {
	dat, err := ioutil.ReadFile("/home/ec2-user/environment/showme.txt")
	check(err)
	// 	fmt.Print(string(dat))
	dat2, err := ioutil.ReadFile("/home/ec2-user/environment/pong.txt")
	check(err)
	return string(dat) + string(dat2)
}

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, readshowme())
	})

	// Start server
	fmt.Println("Start Server from port 9936")
	e.Logger.Fatal(e.Start(":9936"))
}
