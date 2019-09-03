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
	parent   *Directory
	name     string
	children []*Directory
}

func (d Directory) Text() string {
	return d.name
}

func (d Directory) Parent() walk.TreeItem {
	if d.parent == nil {
		// We can't simply return d.parent in this case, because the interface
		// value then would not be nil.
		return nil
	}
	return d.parent
}

func (d Directory) ChildCount() int {
	return len(d.children)
}

func (d Directory) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

type DirectoryTreeModel struct {
	walk.TreeModelBase
	roots []*Directory
}

func (m *DirectoryTreeModel) RootCount() int {
	return len(m.roots)
}

func (m *DirectoryTreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

var _ walk.TreeModel = new(DirectoryTreeModel)

func CreateTreeView(size Size) TreeView {
	if treeView == nil {
		treeModel, err := newDirectoryTreeModel()
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

func newDirectoryTreeModel() (*DirectoryTreeModel, error) {
	model := new(DirectoryTreeModel)
	dir := newDirectory("a0", nil)
	model.roots = append(model.roots, dir)

	_ = append(dir.children, newDirectory("a1", dir))
	_ = append(dir.children, newDirectory("a2", dir))

	return model, nil
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name}
}
