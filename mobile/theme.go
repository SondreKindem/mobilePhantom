package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type myTheme struct{}

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	green := color.RGBA{
		R: 71,
		G: 122,
		B: 30,
		A: 255,
	}

	if name == theme.ColorNameButton {
		return green
	}

	if name == theme.ColorNamePrimary {
		return green
	}

	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	/*if name == theme.IconNameHome {
		fyne.NewStaticResource("myHome", homeBytes)
	}*/

	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name) * 2
}
