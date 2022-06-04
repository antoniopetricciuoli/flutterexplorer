package endpoints

import (
	"flutterexplorer/models"
	"flutterexplorer/utils"
	"github.com/gofiber/fiber/v2"
)

type GetPackageParams struct {
	Query string `query:"q"`
}

func GetPackages(c *fiber.Ctx) error {
	getPackageParams := new(GetPackageParams)
	var p []models.Package

	if err := c.QueryParser(getPackageParams); err != nil {
		return c.Status(400).JSON(models.FailResponse("The query is invalid", 400))
	}

	p, err := utils.PackagesScraper(getPackageParams.Query)
	if err != nil {
		return c.Status(500).JSON(models.FailResponse("Internal Server Error", 500))
	}

	return c.Status(200).JSON(models.SuccessResponse(p))
}
