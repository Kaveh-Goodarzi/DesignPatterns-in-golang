package main

import "fmt"

type Theme interface {
	TextColor() string
	BgColor()   string
}

type LightTheme struct {}

func (l *LightTheme) BgColor() string {
	return "light"
}

func (l *LightTheme) TextColor() string {
	return "light"
}

type DarkTheme struct {}

func (d *DarkTheme) BgColor() string {
	return "dark"
}

func (d *DarkTheme) TextColor() string {
	return "dark"
}

type WebPage interface {
	getContent() string
	SetTheme(theme Theme)
}

type AbstractWebPage struct {
	theme Theme
}

func (aw *AbstractWebPage) SetTheme(theme Theme) {
	aw.theme = theme
}

type HomePage struct {
	AbstractWebPage
}

func (h *HomePage) getContent() string {
	return fmt.Sprintf(
		"<div style='background-color:%s; color:%s;'>Welcome To Our Site</div>",
		h.theme.BgColor(),
		h.theme.TextColor(),
	 )
}

func main() {
	homePage := &HomePage{}
	homePage.SetTheme(&DarkTheme{})

	fmt.Println(homePage.getContent())
	fmt.Println(homePage.theme.TextColor())
	fmt.Println(homePage.theme.BgColor())

	homePage.SetTheme(&LightTheme{})
	fmt.Println(homePage.getContent())

	fmt.Println(homePage.theme.TextColor())
	fmt.Println(homePage.theme.BgColor())
}