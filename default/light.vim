hi Normal guifg=NvimDarkGrey2 guibg=NvimLightGrey2 ctermfg=NONE ctermbg=NONE

" UI
hi Added                guifg=NvimDarkGreen                                  ctermfg=2
hi Changed              guifg=NvimDarkCyan                                   ctermfg=6
hi ColorColumn                               guibg=NvimLightGrey4            cterm=reverse
hi Conceal              guifg=NvimLightGrey4
hi CurSearch            guifg=NvimLightGrey1 guibg=NvimDarkYellow            ctermfg=15 ctermbg=3
hi CursorColumn                              guibg=NvimLightGrey3
hi CursorLine                                guibg=NvimLightGrey3
hi DiffAdd              guifg=NvimDarkGrey1  guibg=NvimLightGreen            ctermfg=15 ctermbg=2
hi DiffChange           guifg=NvimDarkGrey1  guibg=NvimLightGrey4
hi DiffDelete           guifg=NvimDarkRed                          gui=bold  ctermfg=1 cterm=bold
hi DiffText             guifg=NvimDarkGrey1  guibg=NvimLightCyan             ctermfg=15 ctermbg=6
hi Directory            guifg=NvimDarkCyan                                   ctermfg=6
hi ErrorMsg             guifg=NvimDarkRed                                    ctermfg=1
hi FloatShadow                               guibg=NvimLightGrey4            ctermbg=0 blend=80
hi FloatShadowThrough                        guibg=NvimLightGrey4            ctermbg=0 blend=100
hi Folded               guifg=NvimDarkGrey4  guibg=NvimLightGrey3
hi LineNr               guifg=NvimLightGrey4
hi MatchParen                                guibg=NvimLightGrey4  gui=bold  cterm=bold,underline
hi ModeMsg              guifg=NvimDarkGreen                                  ctermfg=2
hi MoreMsg              guifg=NvimDarkCyan                                   ctermfg=6
hi NonText              guifg=NvimLightGrey4
hi NormalFloat                               guibg=NvimLightGrey1
hi Pmenu                                     guibg=NvimLightGrey3            cterm=reverse
hi PmenuThumb                                guibg=NvimLightGrey4
hi Question             guifg=NvimDarkCyan                                   ctermfg=6
hi QuickFixLine         guifg=NvimDarkCyan                                   ctermfg=6
hi RedrawDebugClear                          guibg=NvimLightYellow           ctermfg=15 ctermbg=3
hi RedrawDebugComposed                       guibg=NvimLightGreen            ctermfg=15 ctermbg=2
hi RedrawDebugRecompose                      guibg=NvimLightRed              ctermfg=15 ctermbg=1
hi Removed              guifg=NvimDarkRed                                    ctermfg=1
hi Search               guifg=NvimDarkGrey1  guibg=NvimLightYellow           ctermfg=15 ctermbg=3
hi SignColumn           guifg=NvimLightGrey4
hi SpecialKey           guifg=NvimLightGrey4
hi SpellBad             guisp=NvimDarkRed    gui=undercurl                   cterm=undercurl
hi SpellCap             guisp=NvimDarkYellow gui=undercurl                   cterm=undercurl
hi SpellLocal           guisp=NvimDarkGreen  gui=undercurl                   cterm=undercurl
hi SpellRare            guisp=NvimDarkCyan   gui=undercurl                   cterm=undercurl
hi StatusLine           guifg=NvimLightGrey3 guibg=NvimDarkGrey3             cterm=reverse
hi StatusLineNC         guifg=NvimDarkGrey3  guibg=NvimLightGrey3            cterm=bold,underline
hi Title                guifg=NvimDarkGrey2                        gui=bold  cterm=bold
hi Visual                                    guibg=NvimLightGrey4            ctermfg=15 ctermbg=0
hi WarningMsg           guifg=NvimDarkYellow                                 ctermfg=3
hi WinBar               guifg=NvimDarkGrey4  guibg=NvimLightGrey1  gui=bold  cterm=bold
hi WinBarNC             guifg=NvimDarkGrey4  guibg=NvimLightGrey1            cterm=bold

" Syntax
hi Constant   guifg=NvimDarkGrey2
hi Operator   guifg=NvimDarkGrey2
hi PreProc    guifg=NvimDarkGrey2
hi Type       guifg=NvimDarkGrey2
hi Delimiter  guifg=NvimDarkGrey2

hi Comment    guifg=NvimDarkGrey4
hi String     guifg=NvimDarkGreen                    ctermfg=2
hi Identifier guifg=NvimDarkBlue                     ctermfg=4
hi Function   guifg=NvimDarkCyan                     ctermfg=6
hi Statement  guifg=NvimDarkGrey2 gui=bold           cterm=bold
hi Special    guifg=NvimDarkCyan                     ctermfg=6
hi Error      guifg=NvimDarkGrey1 guibg=NvimLightRed ctermfg=15 ctermbg=1
hi Todo       guifg=NvimDarkGrey2 gui=bold           cterm=bold

" Diagnostic
hi DiagnosticError          guifg=NvimDarkRed                      ctermfg=1
hi DiagnosticWarn           guifg=NvimDarkYellow                   ctermfg=3
hi DiagnosticInfo           guifg=NvimDarkCyan                     ctermfg=6
hi DiagnosticHint           guifg=NvimDarkBlue                     ctermfg=4
hi DiagnosticOk             guifg=NvimDarkGreen                    ctermfg=2
hi DiagnosticUnderlineError guisp=NvimDarkRed    gui=underline     cterm=underline
hi DiagnosticUnderlineWarn  guisp=NvimDarkYellow gui=underline     cterm=underline
hi DiagnosticUnderlineInfo  guisp=NvimDarkCyan   gui=underline     cterm=underline
hi DiagnosticUnderlineHint  guisp=NvimDarkBlue   gui=underline     cterm=underline
hi DiagnosticUnderlineOk    guisp=NvimDarkGreen  gui=underline     cterm=underline
hi DiagnosticDeprecated     guisp=NvimDarkRed    gui=strikethrough cterm=strikethrough

" Treesitter standard groups
hi @variable guifg=NvimDarkGrey2
