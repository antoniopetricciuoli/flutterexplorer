package utils

import (
	"encoding/json"
	"flutterexplorerbot/models"
	"fmt"
	"net/http"
)

const URL = "https://flutterexplorer.antoniopetricciuoli.com/api"

func GetWidgets(widget string) ([]models.Widget, error) {
	resp, err := http.Get(fmt.Sprintf("%s/widget?q=%s", URL, widget))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var widgetResponse models.WidgetResponse
	if err := json.NewDecoder(resp.Body).Decode(&widgetResponse); err != nil {
		return nil, err
	}
	return widgetResponse.Result, nil
}

func GetPackages(name string) ([]models.Package, error) {
	resp, err := http.Get(fmt.Sprintf("%s/package?q=%s", URL, name))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var packageResponse models.PackageResponse
	if err := json.NewDecoder(resp.Body).Decode(&packageResponse); err != nil {
		return nil, err
	}

	return packageResponse.Result, nil
}
