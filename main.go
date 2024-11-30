package main

import (
	"fmt"
	"log/slog"
	"os"

	"zed-neovim/vim"
	"zed-neovim/zed"
)

const paletteFile = "default/palette.json"

type Color struct {
	Name       string
	Foreground string
	Background string
	Special    string
	Blend      int
	Attr       vim.HlAttr
}

// Flatten resolves the links on the scheme and makes sure all highlight groups
// are explicitly defined.
// It writes the colors to the given map. This makes so colors from different
// schemes can be merged in the chosen order by calling the function multiple
// times.
func resolve(groups vim.HlGroups, palette vim.Palette, colors map[string]Color) {
	linkGroups := make(vim.HlGroups)
	defaultGroups := make(vim.HlGroups)
	for name, group := range groups {
		// handle links in second pass
		if group.IsLink {
			if group.IsDefault {
				defaultGroups[group.Name] = group
			} else {
				linkGroups[group.Name] = group
			}
			continue
		}
		colors[name] = Color{
			Name:       name,
			Foreground: palette[group.Attrs.Guifg],
			Background: palette[group.Attrs.Guibg],
			Special:    palette[group.Attrs.Guisp],
			Blend:      group.Attrs.Blend,
			Attr:       group.Attrs.Gui,
		}
	}

	for _, group := range defaultGroups {
		resolvedGroup, ok := resolveLinks(colors, defaultGroups, group, make(map[string]bool))
		if !ok {
			continue
		}
		if newColors, ok := colors[resolvedGroup.ToGroup]; ok {
			newColors.Name = group.Name
			colors[group.Name] = newColors
		}
	}

	for _, group := range linkGroups {
		resolvedGroup, ok := resolveLinks(colors, linkGroups, group, make(map[string]bool))
		if !ok {
			continue
		}
		if newColors, ok := colors[resolvedGroup.ToGroup]; ok {
			newColors.Name = group.Name
			colors[group.Name] = newColors
		}
	}
}

func resolveLinks(colors map[string]Color, groups vim.HlGroups, group vim.HlGroup, seen map[string]bool) (vim.HlGroup, bool) {
	if group.ToGroup == "" {
		return group, true
	}
	if seen[group.ToGroup] {
		return group, false
	}
	seen[group.ToGroup] = true
	nextGroup, ok := groups[group.ToGroup]
	if ok {
		return resolveLinks(colors, groups, nextGroup, seen)
	}
	return group, true
}

func main() {
	slog.Info("Loading palette", slog.String("file", paletteFile))
	paletteF, err := os.Open(paletteFile)
	if err != nil {
		slog.Error("Failed to open palette file", slog.Any("error", err))
		os.Exit(1)
	}

	palette, err := vim.LoadPalette(paletteF)
	if err != nil {
		slog.Error("Failed to load palette", slog.Any("error", err))
	}

	schemes := map[string]string{
		"both":    "default/both.vim",
		"dark":    "default/dark.vim",
		"light":   "default/light.vim",
		"cmdline": "default/cmdline.vim",
	}
	colorSchemes, err := loadColorSchemesFromFiles(schemes)
	if err != nil {
		slog.Error("Failed to load colors", slog.Any("error", err))
	}
	darkColors := make(map[string]Color)
	resolve(colorSchemes["dark"], palette, darkColors)
	resolve(colorSchemes["both"], palette, darkColors)
	resolve(colorSchemes["cmdline"], palette, darkColors)
	lightColors := make(map[string]Color)
	resolve(colorSchemes["light"], palette, lightColors)
	resolve(colorSchemes["both"], palette, lightColors)
	resolve(colorSchemes["cmdline"], palette, lightColors)
	// for _, color := range colors {
	// 	fmt.Printf("%s: %s %s %s %d %v\n", color.Name, color.Foreground, color.Background, color.Special, color.Blend, color.Attr)
	// }

	theme := zed.NewTheme("Neovim default", "Kim Nørgaard")
	addTheme(theme, "Neovim default dark", zed.AppearanceContentDark, toZedStyle(darkColors))
	addTheme(theme, "Neovim default light", zed.AppearanceContentLight, toZedStyle(lightColors))

	if data, err := zed.MarshalJSON(theme); err == nil {
		fmt.Println(string(data))
	}
}
