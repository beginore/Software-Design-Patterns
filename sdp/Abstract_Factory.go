package main

type Button interface {
	Render()
}

type DarkThemeButton struct{}

func (b *DarkThemeButton) Render() {
	println("Rendering Dark Theme Button")
}

type LightThemeButton struct{}

func (b *LightThemeButton) Render() {
	println("Rendering Light Theme Button")
}

type UIFactory interface {
	CreateButton() Button
}

type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button {
	return &DarkThemeButton{}
}

type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button {
	return &LightThemeButton{}
}
