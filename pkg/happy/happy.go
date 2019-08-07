package happy

import (
    "net/http"

    "github.com/labstack/echo"
)

func HeppyWorld(c echo.Context) error {
    return c.JSON(http.StatusOK, "happy world !!!")
}

