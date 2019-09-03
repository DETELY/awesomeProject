package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"log"
)

var treeView *walk.TreeView
var treeViewDeclaration TreeView

type Directory struct {
	name     string
	children []*Directory
}

type DirectoryTreeModel struct {
	walk.TreeModelBase
	roots []*Directory
}

func (d *DirectoryTreeModel) RootCount() int {
	panic("implement me")
}

func (d *DirectoryTreeModel) RootAt(index int) walk.TreeItem {
	panic("implement me")
}

var _ walk.TreeModel = new(DirectoryTreeModel)

func CreateTreeView(size Size) TreeView {
	if treeView == nil {
		treeModel, err := NewDirectoryTreeModel()
		if err != nil {
			log.Fatal(err)
		}
		CurrentDialectPair = dialectPairs[0]
		treeViewDeclaration = TreeView{
			AssignTo: &treeView,
			Model:    treeModel,
		}
	}
	return treeViewDeclaration
}

func NewDirectoryTreeModel() (*DirectoryTreeModel, error) {
	model := new(DirectoryTreeModel)

	drives, err := walk.DriveNames()
	if err != nil {
		return nil, err
	}

	for _, drive := range drives {
		switch drive {
		case "A:\\", "B:\\":
			continue
		}

		model.roots = append(model.roots, NewDirectory(drive, nil))
	}

	return model, nil
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name}
}
