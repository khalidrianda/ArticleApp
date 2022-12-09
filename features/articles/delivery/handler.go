package delivery

import (
	"net/http"

	"github.com/khalidrianda/ArticleApp/features/articles/domain"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {

}

func (ah *articleHandler) Insert() echo.HandlerFunc {
	return func(c echo.Context) error {
		var inputData PostFormat
		if err := c.Bind(&inputData); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request."))
		}

		cnv := ToDomain(inputData)
		res, err := ah.srv.Create(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is a problem on server"))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success post the article", ToResponse(res, "add")))
	}
}

func (ah *articleHandler) GetData() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("query")
		author := c.QueryParam("author")

		res, err := ah.srv.Show(query, author)
		if err != nil {
			return c.JSON(http.StatusOK, SuccessResponse("no data", res))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get data", ToResponse(res, "get")))
	}
}
