package main

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main()  {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})
	e.GET("/api/v2/project/obs/:id/getProjectObsBind", uni)

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
	in := []byte(`{"code":4,"msg":null,"data":{"projectId":3855,"obsProductId":11,"obsPlanProductId":539,"yjnTiUni":"4444","id":5,"createAt":1111111,"updateAt":11111111}}`)
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