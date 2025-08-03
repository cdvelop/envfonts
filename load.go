package envfonts

import (
	. "github.com/cdvelop/tinystring"
)

// FontStyle defines the available font styles
const (
	FontRegular = "regular"
	FontBold    = "bold"
	FontItalic  = "italic"
)

const fontPublicPath = "fonts/"

// Font represents font files for different styles
type Font struct {
	Regular string // Regular font file name eg: "regular.ttf"
	Bold    string // Bold font file name eg: "bold.ttf"
	Italic  string // Italic font file name eg: "italic.ttf"
	Path    string // Base path for fonts eg: "fonts/"
}

// FontConfig represents different font configurations for document sections
type FontConfig struct {
	Family Font
}

// loadFonts loads the fonts from the Font struct
func (f *Envfonts) loadFonts() error {
	fontPath := f.fontConfig.Family.Path

	// add regular font
	if err := f.AddTTFFont(FontRegular, fontPath+f.fontConfig.Family.Regular); err != nil {
		return err
	}
	// add bold font
	if f.fontConfig.Family.Bold == "" {
		f.fontConfig.Family.Bold = f.fontConfig.Family.Regular
	} else {
		if err := f.AddTTFFont(FontBold, fontPath+f.fontConfig.Family.Bold); err != nil {
			return err
		}
	}
	// add italic font
	if f.fontConfig.Family.Italic == "" {
		f.fontConfig.Family.Italic = f.fontConfig.Family.Regular
	} else {
		if err := f.AddTTFFont(FontItalic, fontPath+f.fontConfig.Family.Italic); err != nil {
			return err
		}
	}

	if f.fontConfig.Family.Path == "" {
		f.fontConfig.Family.Path = fontPublicPath
	}

	return nil
}

// extracts the font name from the font path eg: "public/fonts/regular.ttf" => "regular"
func extractNameFromPath(path string) string {
	if path == "" {
		return ""
	}
	// normalize path separators to forward slash
	path = Convert(path).Replace("\\", "/").String()

	// split the path by "/"
	parts := Convert(path).Split("/")

	// get the last part (filename)
	filename := parts[len(parts)-1]

	// split by dot to get all parts
	nameParts := Convert(filename).Split(".")

	// remove the last part if it's an extension (ttf, otf, etc)
	if len(nameParts) > 1 {
		nameParts = nameParts[:len(nameParts)-1]
	}

	// join all parts without dots
	return Convert(nameParts).Join("").String()
}

// defaultFontConfig returns word-processor like defaults
func defaultFontConfig() FontConfig {
	return FontConfig{
		Family: Font{
			Regular: "regular.ttf",
			Bold:    "bold.ttf",
			Italic:  "italic.ttf",
			Path:    fontPublicPath,
		},
	}
}
