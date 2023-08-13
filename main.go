package main

import (
	"OLC2_Proyecto1_202111478/TswiftP"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

func createTabContent(entry fyne.Widget) *fyne.Container {
	// Crear una caja de texto para el contenido de una tab
	tabEntry := entry
	tabEntry.Resize(fyne.NewSize(1850, 250))
	// Envolver la caja de texto en un contenedor
	tabContent := container.NewWithoutLayout(
		tabEntry,
	)
	return tabContent
}

func VentanaM() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Be Swift")
	myWindow.Resize(fyne.NewSize(1920, 1080))

	//Menu de despliegue de Archivo
	ArchivoItem1 := fyne.NewMenuItem("Abrir", func() {})
	ArchivoItem2 := fyne.NewMenuItem("Guardar", func() {})
	ArchivoItem3 := fyne.NewMenuItem("Guardar Como", func() {})

	//Menu de despliegue de Archivo
	EditorItem1 := fyne.NewMenuItem("Crear", nil)
	EditorItem2 := fyne.NewMenuItem("Borrar", nil)
	EditorItem3 := fyne.NewMenuItem("Opcion 3", nil)

	AcercaItem1 := fyne.NewMenuItem("Informacion", nil)

	//Menu principal ARCHIVO
	ArchivoMenu := fyne.NewMenu("Archivo", ArchivoItem1, ArchivoItem2, ArchivoItem3)
	EditarMenu := fyne.NewMenu("Editar", EditorItem1, EditorItem2, EditorItem3)
	AcercaMenu := fyne.NewMenu("Acerca de", AcercaItem1)
	//Menu Principal
	Menu := fyne.NewMainMenu(ArchivoMenu, EditarMenu, AcercaMenu)
	myWindow.SetMainMenu(Menu)
	//botones
	//BOTON EJECUTAR
	btn1 := widget.NewButton("Ejecutar", func() {})
	btn1.Resize(fyne.NewSize(200, 50))
	btn1.Move(fyne.NewPos(1550, 230))
	//BOTON REPORTES
	btn2 := widget.NewButton("Mostrar Reportes", func() {})
	btn2.Resize(fyne.NewSize(200, 50))
	btn2.Move(fyne.NewPos(1550, 330))

	//Multi linea donde se ingresa el codigo
	Entrada := widget.NewMultiLineEntry()
	Entrada.Wrapping = fyne.TextTruncate
	Entrada.Resize(fyne.NewSize(1400, 550))
	Entrada.Move(fyne.NewPos(30, 30))

	//Tabs inferiores
	Consola := widget.NewMultiLineEntry()
	Consola.Wrapping = fyne.TextWrapBreak

	Simbolos := widget.NewMultiLineEntry()
	Simbolos.Wrapping = fyne.TextWrapBreak

	Errores := widget.NewMultiLineEntry()
	Errores.Wrapping = fyne.TextWrapBreak

	tabContainer := container.NewAppTabs(
		container.NewTabItem("Consola", createTabContent(Consola)),
		container.NewTabItem("Simbolos", createTabContent(Simbolos)),
		container.NewTabItem("Errores", createTabContent(Errores)),
	)
	//se coloca en un contenedor las tabs
	content := container.NewVBox(
		tabContainer,
	)
	content.Move(fyne.NewPos(30, 600))
	content.Resize(fyne.NewSize(1850, 400))

	//Colocar contenido
	myWindow.SetContent(
		container.NewWithoutLayout(
			btn1,
			btn2,
			Entrada,
			content,
		),
	)

	myWindow.ShowAndRun()
}

func main() {
	VentanaM()
	fmt.Println("hola mundo")
	//abrir archivos
	archivo, err := antlr.NewFileStream("entrada.txt")
	if err != nil {
		log.Fatal(err)
	}
	//se ejecuta el lexer
	lexer := TswiftP.NewTswift_Lexer(archivo)
	//se ven los tokens
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	//se ve el parser
	parser := TswiftP.NewTswift_GrammarParser(tokens)
	//se visita por medio de visitor
	visitor := NewTswiftPVisitorImp()
	//comenzar el arbol
	arbol := parser.S() // S es la produccion inicial

	fmt.Println(visitor.Visit(arbol))
}
