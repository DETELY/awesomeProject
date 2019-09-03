package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
)

var edit *walk.TextEdit
var treeView *walk.TreeView

func main() {
	_, _ = MainWindow{
		Title:   "GoPad",
		MinSize: Size{600, 400},
		Layout:  HBox{},
		Children: []Widget{
			VSplitter{
				MinSize: Size{Width: 50, Height: 400},
				Children: []Widget{
					TreeView{AssignTo: &treeView},
					PushButton{
						MinSize: Size{Width: 50, Height: 10},
						Text:    "Generate tests",
					},
				},
			},
			VSplitter{
				MinSize: Size{Width: 550, Height: 400},
				Children: []Widget{
					HSplitter{
						MinSize: Size{Width: 550, Height: 20},
						Children: []Widget{
							ComboBox{

							},
						},
					},
					TextEdit{
						AssignTo: &edit,
						MinSize:  Size{Width: 550, Height: 380}},
				},
			},
		},
	}.Run()
}

func copy() {
	walk.Clipboard().SetText(edit.Text())
}

func paste() {
	if text, err := walk.Clipboard().Text(); err == nil {
		edit.SetText(edit.Text() + text)
	}
}

func save() {
	ioutil.WriteFile("file.txt", []byte(edit.Text()), 0x777)
}

func load() {
	contents, _ := ioutil.ReadFile("file.txt")
	edit.SetText(string(contents))
}
