package endpoints

import (
	"flutterexplorer/models"
	"flutterexplorer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/thoas/go-funk"
)

type GetWidgetParams struct {
	Query string `query:"q"`
}

func GetWidgets(c *fiber.Ctx) error {

	getWidgetParams := new(GetWidgetParams)
	var w []models.Widget

	if err := c.QueryParser(getWidgetParams); err != nil {
		return c.Status(400).JSON(models.FailResponse("The query is invalid", 400))
	}

	widgets, err := utils.WidgetScraper()

	if err != nil {
		return c.Status(500).JSON(models.FailResponse("Internal Server Error", 500))
	}

	if getWidgetParams.Query != "" {
		for _, widget := range widgets {
			if funk.Contains(widget.Name, getWidgetParams.Query) {
				w = append(w, widget)
			}
		}
		return c.Status(200).JSON(models.SuccessResponse(w))
	} else {
		return c.Status(200).JSON(models.SuccessResponse(widgets))
	}

}
