package main

import "zed-neovim/zed"

func addTheme(theme *zed.Theme, name string, appearance zed.AppearanceContent, style *zed.ThemeStyleContent) {
	theme.Themes = append(theme.Themes, zed.ThemeContent{
		Name:       name,
		Appearance: appearance,
		Style:      *style,
	})
}

const (
	Opacity50 = "50"
	Opacity70 = "70"
	Opacity80 = "80"
)

func toZedStyle(colors map[string]Color) zed.ThemeStyleContent {
	style := zed.ThemeStyleContent{}
	style.Background = ptr(colors["PmenuThumb"].Background)
	style.Border = ptr(colors["PmenuThumb"].Background)
	style.BorderFocused = ptr(colors["Normal"].Foreground)
	style.BorderSelected = ptr(colors["Normal"].Foreground)
	style.BorderTransparent = ptr("#00000000")
	style.BorderVariant = ptr(colors["Pmenu"].Background)
	style.BorderDisabled = ptr(colors["PmenuThumb"].Foreground + Opacity70)
	style.Conflict = ptr(colors["WarningMsg"].Foreground)
	style.ConflictBackground = ptr(colors["Pmenu"].Background)
	style.ConflictBorder = ptr(colors["WarningMsg"].Foreground)
	style.Created = ptr(colors["DiagnosticOk"].Foreground)
	style.CreatedBackground = ptr(colors["Pmenu"].Background)
	style.CreatedBorder = ptr(colors["DiagnosticOk"].Foreground)
	style.Deleted = ptr(colors["ErrorMsg"].Foreground)
	style.DeletedBackground = ptr(colors["Pmenu"].Background)
	style.DeletedBorder = ptr(colors["ErrorMsg"].Foreground)
	style.DropTargetBackground = ptr(colors["NonText"].Foreground + Opacity80)
	style.EditorActiveLineNumber = ptr(colors["Normal"].Foreground)
	style.EditorBackground = ptr(colors["Normal"].Background)
	style.EditorActiveWrapGuide = ptr(colors["PmenuThumb"].Background + Opacity70)
	style.EditorDocumentHighlightReadBackground = ptr(colors["MatchParen"].Background)
	style.EditorDocumentHighlightWriteBackground = ptr(colors["MatchParen"].Background)
	style.EditorForeground = ptr(colors["Normal"].Foreground)
	style.EditorGutterBackground = ptr(colors["Normal"].Background)
	style.EditorHighlightedLineBackground = ptr(colors["CursorLine"].Background)
	style.EditorInvisible = ptr(colors["NonText"].Foreground)
	style.EditorLineNumber = ptr(colors["LineNr"].Foreground)
	style.EditorSubheaderBackground = ptr(colors["Normal"].Background)
	style.EditorWrapGuide = ptr(colors["PmenuThumb"].Background + Opacity50)
	style.ElementActive = ptr(colors["PmenuThumb"].Background)
	style.ElementBackground = ptr(colors["Normal"].Background)
	style.ElementDisabled = ptr(colors["Normal"].Background)
	style.ElementHover = ptr(colors["Pmenu"].Background)
	style.ElementSelected = ptr(colors["PmenuThumb"].Background)
	style.ElevatedSurfaceBackground = ptr(colors["Normal"].Background)
	style.Error = ptr(colors["ErrorMsg"].Foreground)
	style.ErrorBackground = ptr(colors["Pmenu"].Background)
	style.ErrorBorder = ptr(colors["ErrorMsg"].Foreground)
	style.GhostElementActive = ptr(colors["PmenuThumb"].Background)
	style.GhostElementBackground = ptr("#00000000")
	style.GhostElementDisabled = ptr(colors["Normal"].Background)
	style.GhostElementHover = ptr(colors["Pmenu"].Background)
	style.GhostElementSelected = ptr(colors["PmenuThumb"].Background)
	style.Hidden = ptr(colors["Normal"].Foreground + Opacity70)
	style.HiddenBackground = ptr(colors["Pmenu"].Background)
	style.HiddenBorder = ptr(colors["Normal"].Foreground)
	style.Hint = ptr(colors["PmenuThumb"].Background)
	style.HintBackground = ptr(colors["Pmenu"].Background)
	style.HintBorder = ptr(colors["PmenuThumb"].Foreground)
	style.Icon = ptr(colors["Normal"].Foreground)
	style.IconAccent = ptr(colors["Special"].Foreground)
	style.IconDisabled = ptr(colors["Normal"].Foreground + Opacity70)
	style.IconMuted = ptr(colors["Comment"].Foreground)
	style.IconPlaceholder = ptr(colors["Comment"].Foreground)
	style.Ignored = ptr(colors["Normal"].Foreground + Opacity70)
	style.IgnoredBackground = ptr(colors["Pmenu"].Background)
	style.IgnoredBorder = ptr(colors["Normal"].Foreground)
	style.Info = ptr(colors["Special"].Foreground)
	style.InfoBackground = ptr(colors["Pmenu"].Background)
	style.InfoBorder = ptr(colors["Special"].Foreground)
	style.LinkTextHover = ptr(colors["Special"].Foreground)
	style.Modified = ptr(colors["WarningMsg"].Foreground)
	style.ModifiedBackground = ptr(colors["Pmenu"].Background)
	style.ModifiedBorder = ptr(colors["WarningMsg"].Foreground)
	style.PaneFocusedBorder = nil
	style.PanelBackground = ptr(colors["Normal"].Background)
	style.PanelFocusedBorder = nil
	style.Predictive = ptr(colors["Comment"].Foreground)
	style.PredictiveBackground = ptr(colors["Pmenu"].Background)
	style.PredictiveBorder = ptr(colors["Comment"].Foreground)
	style.Renamed = ptr(colors["Special"].Foreground)
	style.RenamedBackground = ptr(colors["Pmenu"].Background)
	style.RenamedBorder = ptr(colors["Special"].Foreground)
	style.ScrollbarThumbBorder = ptr(colors["Pmenu"].Background)
	style.ScrollbarThumbBackground = ptr(colors["Pmenu"].Background + Opacity80)
	style.ScrollbarThumbHoverBackground = ptr(colors["Pmenu"].Background)
	style.ScrollbarTrackBorder = ptr(colors["Normal"].Background)
	style.ScrollbarTrackBackground = ptr("#00000000")
	style.SearchMatchBackground = ptr(colors["Search"].Background)
	style.StatusBarBackground = ptr(colors["Pmenu"].Background)
	style.Success = ptr(colors["DiagnosticOk"].Foreground)
	style.SuccessBackground = ptr(colors["Pmenu"].Background)
	style.SuccessBorder = ptr(colors["DiagnosticOk"].Foreground)
	style.SurfaceBackground = ptr(colors["Normal"].Background)
	style.TabActiveBackground = ptr(colors["Normal"].Background)
	style.TabInactiveBackground = ptr(colors["TabLine"].Background)
	style.TabBarBackground = ptr(colors["Pmenu"].Background)
	// Terminal*
	style.Text = ptr(colors["Normal"].Foreground)
	style.TextAccent = ptr(colors["Special"].Foreground)
	style.TextDisabled = ptr(colors["Normal"].Foreground + Opacity70)
	style.TextMuted = ptr(colors["Comment"].Foreground)
	style.TextPlaceholder = ptr(colors["Normal"].Foreground + Opacity70)
	style.TitleBarBackground = ptr(colors["Pmenu"].Background)
	style.ToolbarBackground = ptr(colors["Normal"].Background)
	style.Unreachable = ptr(colors["Comment"].Foreground)
	style.UnreachableBackground = ptr(colors["Pmenu"].Background)
	style.UnreachableBorder = ptr(colors["Comment"].Foreground)
	style.Warning = ptr(colors["WarningMsg"].Foreground)
	style.WarningBackground = ptr(colors["Pmenu"].Background)
	style.WarningBorder = ptr(colors["WarningMsg"].Foreground)
	style.Players = []zed.PlayerColorContent{
		{
			Background: ptr(colors["Normal"].Foreground),
			Cursor:     ptr(colors["Normal"].Foreground),
		},
	}
	style.Syntax = zed.ThemeStyleContentSyntax{
		"attribute": zed.HighlightStyleContent{
			Color: ptr(colors["@attribute"].Foreground),
		},
		"boolean": zed.HighlightStyleContent{
			Color: ptr(colors["@boolean"].Foreground),
		},
		"character": zed.HighlightStyleContent{
			Color: ptr(colors["@character"].Foreground),
		},
		"class": zed.HighlightStyleContent{
			Color: ptr(colors["@lsp.type.class"].Foreground),
		},
		"comment": zed.HighlightStyleContent{
			Color: ptr(colors["@lsp.type.comment"].Foreground),
		},
		"comment.doc": zed.HighlightStyleContent{
			Color: ptr(colors["@lsp.type.comment"].Foreground),
		},
		"constant": {
			Color: ptr(colors["@constant"].Foreground),
		},
		"constant.builtin": {
			Color: ptr(colors["zed.syntax.constant.builtin"].Foreground),
		},
		"constructor": {
			Color: ptr(colors["@constructor"].Foreground),
		},
		"decorator": {
			Color: ptr(colors["@lsp.type.decorator"].Foreground),
		},
		"enum": {
			Color: ptr(colors["@lsp.type.enum"].Foreground),
		},
		"enum.member": {
			Color: ptr(colors["@lsp.type.enumMember"].Foreground),
		},
		"function": {
			Color: ptr(colors["@lsp.type.function"].Foreground),
		},
		"hint": {
			Color: ptr(colors["LspInlayHint"].Foreground),
		},
		"interface": {
			Color: ptr(colors["@lsp.type.interface"].Foreground),
		},
		"keyword": {
			Color:      ptr(colors["@lsp.type.keyword"].Foreground),
			FontWeight: ptr(900),
		},
		"label": {
			Color:      ptr(colors["@label"].Foreground),
			FontWeight: ptr(900),
		},
		"markup": {
			Color: ptr(colors["@markup"].Foreground),
		},
		"markup.heading": {
			Color: ptr(colors["@markup.heading"].Foreground),
		},
		"markup.link": {
			FontStyle: zed.FontStyleContentItalic,
		},
		"macro": {
			Color: ptr(colors["Macro"].Foreground),
		},
		"method": {
			Color: ptr(colors["zed.syntax.method"].Foreground),
		},
		"modifier": {
			Color: ptr(colors["Operator"].Foreground),
		},
		"module": {
			Color: ptr(colors["@module"].Foreground),
		},
		"namespace": {
			Color: ptr(colors["@lsp.type.namespace"].Foreground),
		},
		"number": {
			Color: ptr(colors["@lsp.type.number"].Foreground),
		},
		"operator": {
			Color: ptr(colors["@lsp.type.operator"].Foreground),
		},
		"parameter": {
			Color: ptr(colors["Special"].Foreground),
		},
		"preproc": {
			Color: ptr(colors["PreProc"].Foreground),
		},
		"property": {
			Color: ptr(colors["zed.syntax.property"].Foreground),
		},
		"punctuation": {
			Color: ptr(colors["@punctuation"].Foreground),
		},
		"punctuation.bracket": {
			Color: ptr(colors["@punctuation"].Foreground),
		},
		"punctuation.delimiter": {
			Color: ptr(colors["@punctuation"].Foreground),
		},
		"punctuation.list_marker": {
			Color: ptr(colors["@punctuation"].Foreground),
		},
		"punctuation.special": {
			Color: ptr(colors["@punctuation.special"].Foreground),
		},
		"string": {
			Color: ptr(colors["@string"].Foreground),
		},
		"string.escape": {
			Color: ptr(colors["@string.escape"].Foreground),
		},
		"string.regexp": {
			Color: ptr(colors["@string.regexp"].Foreground),
		},
		"string.special": {
			Color: ptr(colors["@string.special"].Foreground),
		},
		"string.special.symbol": {
			Color: ptr(colors["@string.special"].Foreground),
		},
		"struct": {
			Color: ptr(colors["@lsp.type.struct"].Foreground),
		},
		"tag": {
			Color: ptr(colors["@tag"].Foreground),
		},
		"text.literal": {
			Color: ptr(colors["String"].Foreground),
		},
		"title": {
			Color: ptr(colors["Title"].Foreground),
		},
		"type": {
			Color: ptr(colors["zed.syntax.type"].Foreground),
		},
		"type.builtin": {
			Color: ptr(colors["zed.syntax.type.builtin"].Foreground),
		},
		"variable": {
			Color: ptr(colors["zed.syntax.variable"].Foreground),
		},
		"variable.member": {
			Color: ptr(colors["zed.syntax.variable.member"].Foreground),
		},
		"variable.builtin": {
			Color: ptr(colors["@variable.builtin"].Foreground),
		},
		"variable.special": {
			Color: ptr(colors["SpecialChar"].Foreground),
		},
	}

	return style
}

func ptr[T any](v T) *T {
	return &v
}
