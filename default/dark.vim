hi Normal guifg=NvimLightGrey2 guibg=NvimDarkGrey2 ctermfg=NONE ctermbg=NONE

" UI
hi Added                guifg=NvimLightGreen                                ctermfg=10
hi Changed              guifg=NvimLightCyan                                 ctermfg=14
hi ColorColumn                                guibg=NvimDarkGrey4           cterm=reverse
hi Conceal              guifg=NvimDarkGrey4
hi CurSearch            guifg=NvimDarkGrey1   guibg=NvimLightYellow         ctermfg=0 ctermbg=11
hi CursorColumn                               guibg=NvimDarkGrey3
hi CursorLine                                 guibg=NvimDarkGrey3
hi DiffAdd              guifg=NvimLightGrey1  guibg=NvimDarkGreen           ctermfg=0 ctermbg=10
hi DiffChange           guifg=NvimLightGrey1  guibg=NvimDarkGrey4
hi DiffDelete           guifg=NvimLightRed                         gui=bold ctermfg=9 cterm=bold
hi DiffText             guifg=NvimLightGrey1  guibg=NvimDarkCyan            ctermfg=0 ctermbg=14
hi Directory            guifg=NvimLightCyan                                 ctermfg=14
hi ErrorMsg             guifg=NvimLightRed                                  ctermfg=9
hi FloatShadow                                guibg=NvimDarkGrey4           ctermbg=0 blend=80
hi FloatShadowThrough                         guibg=NvimDarkGrey4           ctermbg=0 blend=100
hi Folded               guifg=NvimLightGrey4  guibg=NvimDarkGrey3
hi LineNr               guifg=NvimDarkGrey4
hi MatchParen                                 guibg=NvimDarkGrey4  gui=bold cterm=bold,underline
hi ModeMsg              guifg=NvimLightGreen                                ctermfg=10
hi MoreMsg              guifg=NvimLightCyan                                 ctermfg=14
hi NonText              guifg=NvimDarkGrey4
hi NormalFloat                                guibg=NvimDarkGrey1
hi Pmenu                                      guibg=NvimDarkGrey3           cterm=reverse
hi PmenuThumb                                 guibg=NvimDarkGrey4
hi Question             guifg=NvimLightCyan                                 ctermfg=14
hi QuickFixLine         guifg=NvimLightCyan                                 ctermfg=14
hi RedrawDebugClear                           guibg=NvimDarkYellow          ctermfg=0 ctermbg=11
hi RedrawDebugComposed                        guibg=NvimDarkGreen           ctermfg=0 ctermbg=10
hi RedrawDebugRecompose                       guibg=NvimDarkRed             ctermfg=0 ctermbg=9
hi Removed              guifg=NvimLightRed                                  ctermfg=9
hi Search               guifg=NvimLightGrey1  guibg=NvimDarkYellow          ctermfg=0 ctermbg=11
hi SignColumn           guifg=NvimDarkGrey4
hi SpecialKey           guifg=NvimDarkGrey4
hi SpellBad             guisp=NvimLightRed    gui=undercurl                 cterm=undercurl
hi SpellCap             guisp=NvimLightYellow gui=undercurl                 cterm=undercurl
hi SpellLocal           guisp=NvimLightGreen  gui=undercurl                 cterm=undercurl
hi SpellRare            guisp=NvimLightCyan   gui=undercurl                 cterm=undercurl
hi StatusLine           guifg=NvimDarkGrey3   guibg=NvimLightGrey3          cterm=reverse
hi StatusLineNC         guifg=NvimLightGrey3  guibg=NvimDarkGrey3           cterm=bold,underline
hi Title                guifg=NvimLightGrey2                       gui=bold cterm=bold
hi Visual                                     guibg=NvimDarkGrey4           ctermfg=0 ctermbg=15
hi WarningMsg           guifg=NvimLightYellow                               ctermfg=11
hi WinBar               guifg=NvimLightGrey4  guibg=NvimDarkGrey1  gui=bold cterm=bold
hi WinBarNC             guifg=NvimLightGrey4  guibg=NvimDarkGrey1           cterm=bold

" Syntax
hi Constant   guifg=NvimLightGrey2"
hi Operator   guifg=NvimLightGrey2
hi PreProc    guifg=NvimLightGrey2
hi Type       guifg=NvimLightGrey2
hi Delimiter  guifg=NvimLightGrey2

hi Comment    guifg=NvimLightGrey4
hi String     guifg=NvimLightGreen                   ctermfg=10
hi Identifier guifg=NvimLightBlue                    ctermfg=12
hi Function   guifg=NvimLightCyan                    ctermfg=14
hi Statement  guifg=NvimLightGrey2 gui=bold          cterm=bold
hi Special    guifg=NvimLightCyan                    ctermfg=14
hi Error      guifg=NvimLightGrey1 guibg=NvimDarkRed ctermfg=0 ctermbg=9
hi Todo       guifg=NvimLightGrey2 gui=bold          cterm=bold

" Diagnostic
hi DiagnosticError          guifg=NvimLightRed                      ctermfg=9
hi DiagnosticWarn           guifg=NvimLightYellow                   ctermfg=11
hi DiagnosticInfo           guifg=NvimLightCyan                     ctermfg=14
hi DiagnosticHint           guifg=NvimLightBlue                     ctermfg=12
hi DiagnosticOk             guifg=NvimLightGreen                    ctermfg=10
hi DiagnosticUnderlineError guisp=NvimLightRed    gui=underline     cterm=underline
hi DiagnosticUnderlineWarn  guisp=NvimLightYellow gui=underline     cterm=underline
hi DiagnosticUnderlineInfo  guisp=NvimLightCyan   gui=underline     cterm=underline
hi DiagnosticUnderlineHint  guisp=NvimLightBlue   gui=underline     cterm=underline
hi DiagnosticUnderlineOk    guisp=NvimLightGreen  gui=underline     cterm=underline
hi DiagnosticDeprecated     guisp=NvimLightRed    gui=strikethrough cterm=strikethrough

" Treesitter standard groups
hi @variable guifg=NvimLightGrey2
