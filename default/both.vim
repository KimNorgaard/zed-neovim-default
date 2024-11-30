hi Cursor            guifg=bg      guibg=fg
hi CursorLineNr      gui=bold      cterm=bold
hi PmenuMatch        gui=bold      cterm=bold
hi PmenuMatchSel     gui=bold      cterm=bold
hi PmenuSel          gui=reverse   cterm=reverse,underline blend=0
hi RedrawDebugNormal gui=reverse   cterm=reverse
hi TabLineSel        gui=bold      cterm=bold
hi TermCursor        gui=reverse   cterm=reverse
hi Underlined        gui=underline cterm=underline
hi lCursor           guifg=bg      guibg=fg

" UI
hi default link CursorIM         Cursor
hi default link CursorLineFold   FoldColumn
hi default link CursorLineSign   SignColumn
hi default link EndOfBuffer      NonText
hi default link FloatBorder      NormalFloat
hi default link FloatFooter      FloatTitle
hi default link FloatTitle       Title
hi default link FoldColumn       SignColumn
hi default link IncSearch        CurSearch
hi default link LineNrAbove      LineNr
hi default link LineNrBelow      LineNr
hi default link MsgSeparator     StatusLine
hi default link MsgArea          NONE
hi default link NormalNC         NONE
hi default link PmenuExtra       Pmenu
hi default link PmenuExtraSel    PmenuSel
hi default link PmenuKind        Pmenu
hi default link PmenuKindSel     PmenuSel
hi default link PmenuSbar        Pmenu
hi default link Substitute       Search
hi default link StatusLineTerm   StatusLine
hi default link StatusLineTermNC StatusLineNC
hi default link TabLine          StatusLineNC
hi default link TabLineFill      TabLine
hi default link TermCursorNC     NONE
hi default link VertSplit        WinSeparator
hi default link VisualNOS        Visual
hi default link Whitespace       NonText
hi default link WildMenu         PmenuSel
hi default link WinSeparator     Normal

" Syntax
hi default link Character      Constant
hi default link Number         Constant
hi default link Boolean        Constant
hi default link Float          Number
hi default link Conditional    Statement
hi default link Repeat         Statement
hi default link Label          Statement
hi default link Keyword        Statement
hi default link Exception      Statement
hi default link Include        PreProc
hi default link Define         PreProc
hi default link Macro          PreProc
hi default link PreCondit      PreProc
hi default link StorageClass   Type
hi default link Structure      Type
hi default link Typedef        Type
hi default link Tag            Special
hi default link SpecialChar    Special
hi default link SpecialComment Special
hi default link Debug          Special
hi default link Ignore         Normal

" Built-in LSP
hi default link LspCodeLens                 NonText
hi default link LspCodeLensSeparator        LspCodeLens
hi default link LspInlayHint                NonText
hi default link LspReferenceRead            LspReferenceText
hi default link LspReferenceText            Visual
hi default link LspReferenceWrite           LspReferenceText
hi default link LspReferenceTarget          LspReferenceText
hi default link LspSignatureActiveParameter Visual
hi default link SnippetTabstop              Visual

" Diagnostic
hi default link DiagnosticFloatingError    DiagnosticError
hi default link DiagnosticFloatingWarn     DiagnosticWarn
hi default link DiagnosticFloatingInfo     DiagnosticInfo
hi default link DiagnosticFloatingHint     DiagnosticHint
hi default link DiagnosticFloatingOk       DiagnosticOk
hi default link DiagnosticVirtualTextError DiagnosticError
hi default link DiagnosticVirtualTextWarn  DiagnosticWarn
hi default link DiagnosticVirtualTextInfo  DiagnosticInfo
hi default link DiagnosticVirtualTextHint  DiagnosticHint
hi default link DiagnosticVirtualTextOk    DiagnosticOk
hi default link DiagnosticSignError        DiagnosticError
hi default link DiagnosticSignWarn         DiagnosticWarn
hi default link DiagnosticSignInfo         DiagnosticInfo
hi default link DiagnosticSignHint         DiagnosticHint
hi default link DiagnosticSignOk           DiagnosticOk
hi default link DiagnosticUnnecessary      Comment

" Treesitter standard groups
hi default link @variable.builtin           Special
hi default link @variable.parameter.builtin Special

hi default link @constant         Constant
hi default link @constant.builtin Special

hi default link @module         Structure
hi default link @module.builtin Special
hi default link @label          Label

hi default link @string             String
hi default link @string.regexp      @string.special
hi default link @string.escape      @string.special
hi default link @string.special     SpecialChar
hi default link @string.special.url Underlined

hi default link @character         Character
hi default link @character.special SpecialChar

hi default link @boolean      Boolean
hi default link @number       Number
hi default link @number.float Float

hi default link @type         Type
hi default link @type.builtin Special

hi default link @attribute         Macro
hi default link @attribute.builtin Special
hi default link @property          Identifier

hi default link @function         Function
hi default link @function.builtin Special

hi default link @constructor Special
hi default link @operator    Operator

hi default link @keyword Keyword

hi default link @punctuation         Delimiter
hi default link @punctuation.special Special

hi default link @comment Comment

hi default link @comment.error   DiagnosticError
hi default link @comment.warning DiagnosticWarn
hi default link @comment.note    DiagnosticInfo
hi default link @comment.todo    Todo

hi @markup.strong        gui=bold          cterm=bold
hi @markup.italic        gui=italic        cterm=italic
hi @markup.strikethrough gui=strikethrough cterm=strikethrough
hi @markup.underline     gui=underline     cterm=underline

hi default link @markup         Special
hi default link @markup.heading Title
hi default link @markup.link    Underlined

hi default link @diff.plus  Added
hi default link @diff.minus Removed
hi default link @diff.delta Changed

hi default link @tag         Tag
hi default link @tag.builtin Special

" :help
" Highlight "===" and "---" heading delimiters specially.
hi default @markup.heading.1.delimiter.vimdoc guibg=bg guifg=bg guisp=fg gui=underdouble,nocombine ctermbg=NONE ctermfg=NONE cterm=underdouble,nocombine
hi default @markup.heading.2.delimiter.vimdoc guibg=bg guifg=bg guisp=fg gui=underline,nocombine ctermbg=NONE ctermfg=NONE cterm=underline,nocombine

" LSP semantic tokens
hi default link @lsp.type.class         @type
hi default link @lsp.type.comment       @comment
hi default link @lsp.type.decorator     @attribute
hi default link @lsp.type.enum          @type
hi default link @lsp.type.enumMember    @constant
hi default link @lsp.type.event         @type
hi default link @lsp.type.function      @function
hi default link @lsp.type.interface     @type
hi default link @lsp.type.keyword       @keyword
hi default link @lsp.type.macro         @constant.macro
hi default link @lsp.type.method        @function.method
hi default link @lsp.type.modifier      @type.qualifier
hi default link @lsp.type.namespace     @module
hi default link @lsp.type.number        @number
hi default link @lsp.type.operator      @operator
hi default link @lsp.type.parameter     @variable.parameter
hi default link @lsp.type.property      @property
hi default link @lsp.type.regexp        @string.regexp
hi default link @lsp.type.string        @string
hi default link @lsp.type.struct        @type
hi default link @lsp.type.type          @type
hi default link @lsp.type.typeParameter @type.definition
hi default link @lsp.type.variable      @variable

hi default link @lsp.mod.deprecated DiagnosticDeprecated
