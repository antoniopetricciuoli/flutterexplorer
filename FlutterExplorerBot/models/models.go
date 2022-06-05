package models

type Widget struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type PackageResponse struct {
	Ok     bool      `json:"ok"`
	Result []Package `json:"result"`
}

type WidgetResponse struct {
	Ok     bool     `json:"ok"`
	Result []Widget `json:"result"`
}

type Scores struct {
	Likes      string `json:"likes"`
	PubPoints  string `json:"pub_points"`
	Popularity string `json:"popularity"`
}

type Metadata struct {
	Version string `json:"version"`
}

type Badge struct {
	Platform []string `json:"platform"`
}

type Package struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Scores      Scores   `json:"scores"`
	Metadata    Metadata `json:"metadata"`
	Badge       Badge    `json:"badge"`
}
