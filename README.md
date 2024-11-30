# Neovim default themes for Zed

This is a port of the default [Neovim](https://neovim.io/) themes made for the
[Zed](https://zed.dev/) editor.

It uses the original palette and highlight groups to the best of my ability for
almost everything. As it's a different editor there will never be a 1:1 mapping.

I am by no means good with design and colors. Zed has a lot of UI that Neovim
doesn't and I tried to fill in the gaps while being as true to the original
theme as I could. Input is greatly appreciated.

## Screenshots

On their way :)

## Development

- The original highlight groups and palette live in `default/`
- Some syntax groups have been modified slightly to better resemble what I get
  with Tree-sitter in Neovim. I plan on making those variants to keep in line
  with being faithful for the default themes.
- There is a go program that reads and parses those files and writes the Zed
  theme to STDOUT. The mapping is done in theme.go. Just `go run .`
- Run `make build` to build the theme and write it to
  `themes/neovim-default.json`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

[MIT](LICENSE)

## Acknowledgements

The original colorschemes and palette can be found
[here](https://github.com/neovim/neovim/blob/master/src/nvim/highlight_group.c)
