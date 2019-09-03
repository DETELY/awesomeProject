package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
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

var root DirectoryTreeModel

func CreateTreeView(size Size) TreeView {
	if treeView == nil {
		CurrentDialectPair = dialectPairs[0]
		treeViewDeclaration = TreeView{
			AssignTo: &treeView,
			//Model:    root,
		}
	}
	return treeViewDeclaration
}
