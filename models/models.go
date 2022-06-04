package models

type Widget struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
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
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Scores      Scores `json:"scores"`
	Metadata    any    `json:"metadata"`
	Badge       Badge  `json:"badge"`
}

type Fail struct {
	Ok           bool   `json:"ok"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
}

type Success struct {
	Ok     bool `json:"ok"`
	Result any  `json:"result"`
}

func SuccessResponse(result interface{}) *Success {
	return &Success{Ok: true, Result: result}
}

func FailResponse(err string, code int) *Fail {
	return &Fail{Ok: false, Code: code, ErrorMessage: err}
}
