package main

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.GET("/api/v2/project/obs/:id/getProjectObsBind", uni)
	e.POST("/apis/v1/GetTempSecret", tempSecret)

	e.Logger.Fatal(e.Start(":1333"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, id)
}

func uni(c echo.Context) error {
	id := c.Param("id")
	if id == "1" {
		return errors.New("id error")
	}

	return commonResult(c, `{"code":0,"msg":null,"data":{"projectId":3855,"obsProductId":11,"obsPlanProductId":539,"yjnTiUni":"4444","id":5,"createAt":1111111,"updateAt":11111111}}`)
}

func tempSecret(c echo.Context) error {
	return commonResult(c, `{"Code":0,"Data":{"SecretId":"sec","SecretKey":"xxxx","Token":"12345","ExpiredTime":3223}}`)
}

func commonResult(c echo.Context, result string) error {
	in := []byte(result)
	var raw map[string]interface{}
	if err := json.Unmarshal(in, &raw); err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, raw)
}


func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	if team == "" {
		team = "aa"
	}
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}