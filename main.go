package main

import (
	. "./widgets"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
)

var edit *walk.TextEdit
var treeView *walk.TreeView

func main() {
	_, _ = MainWindow{
		Title:   "Tester",
		MinSize: Size{600, 400},
		Layout:  HBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					VSplitter{
						Children: []Widget{
							CreateTreeView(Size{Width: 50, Height: 390}),
							PushButton{
								Text: "Generate tests",
							},
						},
					},
					VSplitter{
						Children: []Widget{
							HSplitter{
								Children: []Widget{
									CreateDialectPair(Size{
										Width:  50,
										Height: 20,
									}),
								},
							},
							TextEdit{
								AssignTo: &edit,
							},
						},
					},
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
