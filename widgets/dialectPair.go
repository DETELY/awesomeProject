package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var dialectPairs = []string{
	"OracleDw_To_Glue",
	"Oracle_To_Postgre",
	"MsSql_To_MySql",
}

var CurrentDialectPair string
var comboBoxDeclaration ComboBox
var comboBox *walk.ComboBox

func CreateDialectPair(size Size) ComboBox {
	if comboBox == nil {
		CurrentDialectPair = dialectPairs[0]
		comboBoxDeclaration = ComboBox{
			AssignTo:              &comboBox,
			Model:                 dialectPairs,
			CurrentIndex:          0,
			OnCurrentIndexChanged: onCurrentIndexChanged,
		}
	}
	return comboBoxDeclaration
}

func onCurrentIndexChanged() {
	//TODO need add refresh project
	CurrentDialectPair = dialectPairs[comboBox.CurrentIndex()]
}
