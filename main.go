package main

import (
	"OLC2_Proyecto1_202111478/TswiftP"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

//VARIABLE GLOBAL DEL PATH
var rutaCompleta string

//FUNCION PARA CREAR LAS TABS DE CONSOLA, SIMBOLOS, ERRORES
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

//FUNCION PARA GUARDAR UNICAMENTE CON EL PATH
func saveToFile(filePath, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

//FUNCION PARA MOSTRAR LA VENTANA
func VentanaM() {
	//setup de la ventana
	myApp := app.New()
	myWindow := myApp.NewWindow("Be Swift")
	myWindow.Resize(fyne.NewSize(1920, 1080))
	rutaCompleta = ""

	//OBJETOS
	//Multi linea donde se ingresa el codigo
	Entrada := widget.NewMultiLineEntry()
	Entrada.Wrapping = fyne.TextTruncate
	Entrada.Resize(fyne.NewSize(1400, 550))
	Entrada.Move(fyne.NewPos(30, 30))

	//Tabs inferiores
	Consola := widget.NewMultiLineEntry()
	Consola.Wrapping = fyne.TextWrapBreak
	Consola.Disable()

	Simbolos := widget.NewMultiLineEntry()
	Simbolos.Wrapping = fyne.TextWrapBreak
	Simbolos.Disable()

	Errores := widget.NewMultiLineEntry()
	Errores.Wrapping = fyne.TextWrapBreak
	Errores.Disable()

	//botones
	//BOTON EJECUTAR
	btn1 := widget.NewButton("Ejecutar", func() {})
	btn1.Resize(fyne.NewSize(200, 50))
	btn1.Move(fyne.NewPos(1550, 230))
	//BOTON REPORTES
	btn2 := widget.NewButton("Mostrar Reportes", func() {})
	btn2.Resize(fyne.NewSize(200, 50))
	btn2.Move(fyne.NewPos(1550, 330))
	// label
	ruta := widget.NewLabel("")
	ruta.Move(fyne.NewPos(30, 0))

	//Menu de despliegue de Archivo
	ArchivoItem1 := fyne.NewMenuItem("Abrir", func() {
		fileDialog := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			data, _ := ioutil.ReadAll(uc)
			result := fyne.NewStaticResource("entrada", data)
			rutaCompleta = uc.URI().Path()
			ruta.SetText("RUTA: " + rutaCompleta)
			Entrada.SetText(string(result.StaticContent))
		}, myWindow)
		fileDialog.Show()
	})
	ArchivoItem2 := fyne.NewMenuItem("Guardar", func() {
		if rutaCompleta != "" {
			content := Entrada.Text
			err := saveToFile(rutaCompleta, content)
			if err != nil {
				dialog.ShowError(err, myWindow)
			} else {
				mensaje := "Archivo guardado exitosamente en:"
				mensaje += rutaCompleta
				dialog.ShowInformation("Información", mensaje, myWindow)
				ruta.SetText("RUTA: " + rutaCompleta)
			}
		} else {
			dialog.ShowInformation("Error", "No existe una ruta para guardar el archivo.", myWindow)
		}

	})
	ArchivoItem3 := fyne.NewMenuItem("Guardar Como", func() {
		fileDialog := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			data := []byte(Entrada.Text)
			uc.Write(data)
			rutaCompleta = uc.URI().Path()
			mensaje := "Archivo guardado exitosamente en:"
			mensaje += rutaCompleta
			dialog.ShowInformation("Información", mensaje, myWindow)
			ruta.SetText("RUTA: " + rutaCompleta)
			uc.Close()
		}, myWindow)
		fileDialog.SetFileName("")
		fileDialog.Show()
	})

	//Menu de despliegue de Archivo
	EditorItem1 := fyne.NewMenuItem("Limpiar entrada", func() {
		Entrada.SetText("")
	})
	EditorItem2 := fyne.NewMenuItem("Limpiar ruta", func() {
		rutaCompleta = ""
		ruta.SetText("")
	})
	EditorItem3 := fyne.NewMenuItem("Opcion 3", func() {})

	AcercaItem1 := fyne.NewMenuItem("Información", func() {})

	//Menu principal ARCHIVO
	ArchivoMenu := fyne.NewMenu("Archivo", ArchivoItem1, ArchivoItem2, ArchivoItem3)
	EditarMenu := fyne.NewMenu("Editar", EditorItem1, EditorItem2, EditorItem3)
	AcercaMenu := fyne.NewMenu("Acerca de", AcercaItem1)
	//Menu Principal
	Menu := fyne.NewMainMenu(ArchivoMenu, EditarMenu, AcercaMenu)
	myWindow.SetMainMenu(Menu)

	tabContainer := container.NewAppTabs(
		container.NewTabItem("Consola", createTabContent(Consola)),
		container.NewTabItem("Símbolos", createTabContent(Simbolos)),
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
			ruta,
			Entrada,
			content,
		),
	)

	myWindow.ShowAndRun()
}

//MOSTRAR MAIN
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
