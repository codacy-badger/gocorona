package main

import (
	"github.com/gizak/termui/v3/widgets"
)

// Loading represents the loading widget
// displayed while data is being fetched
type Loading struct {
	Widget *widgets.Paragraph
}

// Construct creates the loading widget
func (l *Loading) Construct() {
	widget := widgets.NewParagraph()
	widget.Text = `
	 ____
  / ___| ___   ___ ___  _ __ ___  _ __   __ _
	| |  _ / _ \ / __/ _ \| '__/ _ \| '_ \ / _' |
	| |_| | (_) | (_| (_) | | | (_) | | | | (_| |
	 \____|\___/ \___\___/|_|  \___/|_| |_|\__,_|

 [gocorona](fg:black,bg:yellow) by Ayooluwa Isaiah

 Worldwide Coronavirus (COVID-19) Statistics for your terminal
	`
	widget.SetRect(0, 0, 100, 50)
	widget.Border = false

	l.Widget = widget
}
