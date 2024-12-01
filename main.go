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
	darkColors["zed.syntax.constant.builtin"] = darkColors["@constant"]
	darkColors["zed.syntax.method"] = darkColors["@lsp.type.method"]
	darkColors["zed.syntax.property"] = darkColors["@lsp.type.property"]
	darkColors["zed.syntax.type"] = darkColors["@lsp.type.type"]
	darkColors["zed.syntax.type.builtin"] = darkColors["Special"]
	darkColors["zed.syntax.variable"] = darkColors["@lsp.type.variable"]
	darkColors["zed.syntax.variable.member"] = darkColors["@lsp.type.variable"]

	lightColors := make(map[string]Color)
	resolve(colorSchemes["light"], palette, lightColors)
	resolve(colorSchemes["both"], palette, lightColors)
	resolve(colorSchemes["cmdline"], palette, lightColors)
	lightColors["zed.syntax.constant.builtin"] = lightColors["@constant"]
	lightColors["zed.syntax.method"] = lightColors["@lsp.type.method"]
	lightColors["zed.syntax.property"] = lightColors["@lsp.type.property"]
	lightColors["zed.syntax.type"] = lightColors["@lsp.type.type"]
	lightColors["zed.syntax.type.builtin"] = lightColors["Special"]
	lightColors["zed.syntax.variable"] = lightColors["@lsp.type.variable"]
	lightColors["zed.syntax.variable.member"] = lightColors["@lsp.type.variable"]

	darkColorsAccented := make(map[string]Color)
	for k := range darkColors {
		darkColorsAccented[k] = darkColors[k]
	}
	darkColorsAccented["zed.syntax.constant.builtin"] = darkColors["Special"]
	darkColorsAccented["zed.syntax.method"] = darkColors["@lsp.type.function"]
	darkColorsAccented["zed.syntax.property"] = darkColors["Identifier"]
	darkColorsAccented["zed.syntax.type"] = darkColors["Special"]
	darkColorsAccented["zed.syntax.type.builtin"] = darkColors["Special"]
	darkColorsAccented["zed.syntax.variable.member"] = darkColors["Identifier"]

	lightColorsAccented := make(map[string]Color)
	for k := range lightColors {
		lightColorsAccented[k] = lightColors[k]
	}
	lightColorsAccented["zed.syntax.constant.builtin"] = lightColors["Special"]
	lightColorsAccented["zed.syntax.method"] = lightColors["@lsp.type.function"]
	lightColorsAccented["zed.syntax.property"] = lightColors["Identifier"]
	lightColorsAccented["zed.syntax.type"] = lightColors["Special"]
	lightColorsAccented["zed.syntax.type.builtin"] = lightColors["Special"]
	lightColorsAccented["zed.syntax.variable.member"] = lightColors["Identifier"]

	// for _, color := range colors {
	// 	fmt.Printf("%s: %s %s %s %d %v\n", color.Name, color.Foreground, color.Background, color.Special, color.Blend, color.Attr)
	// }

	theme := zed.NewTheme("Neovim default", "Kim Nørgaard")
	addTheme(theme, "Neovim default dark", zed.AppearanceContentDark, toZedStyle(darkColors))
	addTheme(theme, "Neovim default light", zed.AppearanceContentLight, toZedStyle(lightColors))
	addTheme(theme, "Neovim default accented dark", zed.AppearanceContentDark, toZedStyle(darkColorsAccented))
	addTheme(theme, "Neovim default accented light", zed.AppearanceContentLight, toZedStyle(lightColorsAccented))

	if data, err := zed.MarshalJSON(theme); err == nil {
		fmt.Println(string(data))
	}
}
