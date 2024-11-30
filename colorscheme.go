package main

import (
	"log/slog"
	"os"

	"zed-neovim/vim"
)

type colorSchemes map[string]vim.HlGroups

func loadColorSchemesFromFiles(schemes ...map[string]string) (colorSchemes, error) {
	colorSchemes := make(colorSchemes)

	for _, scheme := range schemes {
		for name, fileName := range scheme {
			slog.Info("Loading color scheme", slog.Any("name", name), slog.Any("file", fileName))
			f, err := os.Open(fileName)
			defer func() {
				if err := f.Close(); err != nil {
					slog.Error("closing file", slog.Any("error", err))
				}
			}()

			if err != nil {
				return nil, err
			}
			hlGroups, err := vim.LoadHlGroups(f)
			if err != nil {
				return nil, err
			}
			colorSchemes[name] = hlGroups
		}
	}

	return colorSchemes, nil
}
