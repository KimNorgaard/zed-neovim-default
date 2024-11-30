package vim

import (
	"bufio"
	"io"
	"log/slog"
	"strconv"
	"strings"
)

type HlGroups map[string]HlGroup

type HlGroup struct {
	Name      string
	IsLink    bool
	IsDefault bool
	FromGroup string
	ToGroup   string
	Attrs     *HlAttrs
	Raw       string
}

type HlAttrs struct {
	Cterm   HlAttr
	Start   string
	Stop    string
	Ctermfg int
	Ctermbg int
	Gui     HlAttr
	Font    string
	Guifg   string
	Guibg   string
	Guisp   string
	Blend   int
}

type HlAttr struct {
	Bold          bool
	Underline     bool
	Underdouble   bool
	Underdotted   bool
	Underdashed   bool
	Strikethrough bool
	Reverse       bool
	Inverse       bool
	Italic        bool
	Standout      bool
	Altfont       bool
	Nocombine     bool
}

func LoadHlGroups(r io.Reader) (HlGroups, error) {
	hlGroups := HlGroups{}
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		hlGroup, err := ParseHlGroup(line)
		if err != nil {
			slog.Warn("Failed to parse hi group", slog.Any("line", line), slog.Any("error", err))
		}
		if hlGroup.Name == "" {
			continue
		}
		hlGroups[hlGroup.Name] = hlGroup
	}
	return hlGroups, nil
}

func ParseHlGroup(hlCmd string) (HlGroup, error) {
	hlGroup := HlGroup{}
	if !strings.HasPrefix(hlCmd, "hi") {
		return hlGroup, nil
	}

	hlCmd = removeComments(hlCmd)

	tokens := strings.Fields(hlCmd)[1:]
	// invalid - there must be at least a group name or a link
	if len(tokens) < 2 {
		return hlGroup, nil
	}
	// ignore clear
	if tokens[0] == "clear" {
		return hlGroup, nil
	}
	// invalid - link must follow default
	if tokens[0] == "default" && tokens[1] != "link" {
		// TODO: normal groups can have a default
		return hlGroup, nil
	}

	hlGroup.Raw = hlCmd
	hlGroup.Attrs = &HlAttrs{}

	// Handle links
	var idx int
	if tokens[0] == "default" && tokens[1] == "link" {
		hlGroup.IsLink = true
		hlGroup.IsDefault = true
		if len(tokens) != 4 {
			slog.Warn("Invalid default link", slog.Any("line", hlCmd))
			return hlGroup, nil
		}
		idx = 2
	} else if tokens[0] == "link" {
		hlGroup.IsLink = true
		if len(tokens) != 3 {
			slog.Warn("Invalid default link", slog.Any("line", hlCmd))
			return hlGroup, nil
		}
		idx = 1
	}
	if hlGroup.IsLink {
		if tokens[:len(tokens)-1][0] == "NONE" {
			return hlGroup, nil
		}

		hlGroup.Name = tokens[idx]
		hlGroup.FromGroup = tokens[idx]
		hlGroup.ToGroup = tokens[idx+1]

		return hlGroup, nil
	}

	// Handle normal groups
	hlGroup.Name = tokens[0]
	if len(tokens) < 2 {
		slog.Warn("Invalid group", slog.Any("line", hlCmd))
		return hlGroup, nil
	}
	if tokens[1] == "NONE" {
		return hlGroup, nil
	}

	hlArgs := tokens[1:]
	ParseHlArgs(hlGroup, hlArgs)

	return hlGroup, nil
}

func ParseHlArgs(hlGroup HlGroup, hlArgs []string) {
	for _, arg := range hlArgs {
		kv := strings.Split(arg, "=")
		if len(kv) != 2 {
			slog.Warn("Invalid arg", slog.Any("arg", arg))
			continue
		}
		val := kv[1]
		switch kv[0] {
		// term
		case "cterm":
			hlGroup.Attrs.Cterm = ParseHlAttrs(val)
		case "start":
			hlGroup.Attrs.Start = val
		case "stop":
			hlGroup.Attrs.Stop = val
		case "ctermfg":
			hlGroup.Attrs.Ctermfg = ParseCtermColor(val)
		case "ctermbg":
			hlGroup.Attrs.Ctermbg = ParseCtermColor(val)
		// gui
		case "gui":
			hlGroup.Attrs.Gui = ParseHlAttrs(val)
		case "font":
			hlGroup.Attrs.Font = val
		case "guifg":
			hlGroup.Attrs.Guifg = val
		case "guibg":
			hlGroup.Attrs.Guibg = val
		case "guisp":
			hlGroup.Attrs.Guisp = val
		case "blend":
			if blend, err := strconv.Atoi(val); err != nil {
				hlGroup.Attrs.Blend = blend
			}
		}
	}
}

func ParseHlAttrs(val string) HlAttr {
	hlAttr := HlAttr{}
	attrs := strings.Split(val, ",")
	for _, attr := range attrs {
		switch attr {
		case "NONE":
			return hlAttr
		case "bold":
			hlAttr.Bold = true
		case "underline":
			hlAttr.Underline = true
		case "underdouble":
			hlAttr.Underdouble = true
		case "underdotted":
			hlAttr.Underdotted = true
		case "underdashed":
			hlAttr.Underdashed = true
		case "strikethrough":
			hlAttr.Strikethrough = true
		case "reverse":
			hlAttr.Reverse = true
		case "inverse":
			hlAttr.Inverse = true
		case "italic":
			hlAttr.Italic = true
		case "standout":
			hlAttr.Standout = true
		case "altfont":
			hlAttr.Altfont = true
		case "nocombine":
			hlAttr.Nocombine = true
		}
	}
	return hlAttr
}

// returns -1 if val is "NONE" or an invalid color name or number
// assumes 16+ color terminal
func ParseCtermColor(val string) int {
	if val == "NONE" {
		return -1
	}
	colors := map[string]int{
		"Black":        0,
		"DarkBlue":     1,
		"DarkGreen":    2,
		"DarkCyan":     3,
		"DarkRed":      4,
		"DarkMagenta":  5,
		"Brown":        6,
		"DarkYellow":   6,
		"LightGray":    7,
		"LightGrey":    7,
		"Gray":         7,
		"Grey":         7,
		"DarkGray":     8,
		"DarkGrey":     8,
		"Blue":         9,
		"LightBlue":    9,
		"Green":        10,
		"LightGreen":   10,
		"Cyan":         11,
		"LightCyan":    11,
		"Red":          12,
		"LightRed":     12,
		"Magenta":      13,
		"LightMagenta": 13,
		"Yellow":       14,
		"LightYellow":  14,
		"White":        15,
	}
	if _, ok := colors[val]; ok {
		return colors[val]
	}
	color, err := strconv.Atoi(val)
	if err != nil {
		slog.Warn("Failed to convert cterm to color name", slog.Any("val", val), slog.Any("error", err))
		return -1
	}
	if color < 0 {
		return -1
	}
	return color
}

func removeComments(line string) string {
	commentIndex := strings.Index(line, `"`)
	if commentIndex != -1 {
		line = line[:commentIndex]
	}
	return strings.TrimSpace(line)
}
