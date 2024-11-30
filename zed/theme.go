package zed

import "encoding/json"

type Theme struct {
	Schema string `json:"$schema,omitempty"`
	*V010Json
}

func NewTheme(name string, author string) *Theme {
	theme := &Theme{
		Schema: "https://zed.dev/schema/themes/v0.1.0.json",
		V010Json: &V010Json{
			Author: author,
			Name:   name,
			Themes: make([]ThemeContent, 0),
		},
	}
	return theme
}

func MarshalJSON(theme *Theme) ([]byte, error) {
	return json.MarshalIndent(theme, "", "  ")
}
