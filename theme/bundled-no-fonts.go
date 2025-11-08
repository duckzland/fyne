//go:build no_fonts
// +build no_fonts

package theme

import "fyne.io/fyne/v2"

var regular = &fyne.StaticResource{
	StaticName:    "NotoSans-Regular.ttf",
	StaticContent: []byte{},
}

var bold = &fyne.StaticResource{
	StaticName:    "NotoSans-Bold.ttf",
	StaticContent: []byte{},
}

var italic = &fyne.StaticResource{
	StaticName:    "NotoSans-Italic.ttf",
	StaticContent: []byte{},
}

var bolditalic = &fyne.StaticResource{
	StaticName:    "NotoSans-BoldItalic.ttf",
	StaticContent: []byte{},
}

var monospace = &fyne.StaticResource{
	StaticName:    "DejaVuSansMono-Powerline.ttf",
	StaticContent: []byte{},
}

var symbol = &fyne.StaticResource{
	StaticName:    "InterSymbols-Regular.ttf",
	StaticContent: []byte{},
}
