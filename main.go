package main

import (
	interprete "OLC2_Proyecto1_202111478/Interprete"
	"OLC2_Proyecto1_202111478/TswiftP"
	"fmt"
	"io/ioutil"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/antlr4-go/antlr/v4"
)

//VARIABLE GLOBAL DEL PATH-----------------------------------------------------------------------------------------------------------------------------------------------------
var rutaCompleta string
var Consola = widget.NewMultiLineEntry()
var Simbolos = widget.NewMultiLineEntry()
var Errores = widget.NewMultiLineEntry()

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

//FUNCION PARA GUARDAR UNICAMENTE CON EL PATH-----------------------------------------------------------------------------------------------------------------------------------
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

//ERRORES----------------------------------------------------------------------------------------------------------------------------------------------------------------------
type Error struct {
	Descripcion string
	Tipo        string
	Linea       int
	Columna     int
}
type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errores []Error
}

func NewErrorListener() *CustomErrorListener {
	return &CustomErrorListener{Errores: make([]Error, 0)}
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	error_ := Error{
		Descripcion: msg,
		Tipo:        "Sintactico",
		Linea:       line,
		Columna:     column,
	}
	Errores.SetText(Errores.Text + "Error: " + error_.Tipo + " " + error_.Descripcion + " en la linea " + fmt.Sprint(error_.Linea) + " y columna " + fmt.Sprint(error_.Columna) + "\n")
}

//EJECUTAR VISITOR INTERPRETE-------------------------------------------------------------------------------------------------------------------------------------------------
func ejecutarInterprete(entrada string) {
	is := antlr.NewInputStream(entrada)
	lexer := TswiftP.NewTswift_GrammarLexer(is)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := TswiftP.NewTswift_GrammarParser(tokens)
	//agregar el error listener
	ErrorListener := NewErrorListener()
	parser.RemoveErrorListeners()
	parser.AddErrorListener(ErrorListener)

	//Se captura cualquier error
	visitor := NewVisitorInterprete()
	arbol := parser.S()
	raiz := visitor.Visit(arbol).(interprete.AbstractExpression)

	//Hacemos contexto
	context := interprete.NewContexto()
	//se inicializa la memoria global
	context.ZGlobal = append(context.ZGlobal, context.Memoria)

	raiz.Interpretar(context)

	//igualar la caja de texto a Consola
	Consola.SetText(string(context.Consola))
	//agregar cada error a la caja de texto de errores
	for _, element := range context.Errores {
		Errores.SetText(Errores.Text + element + "\n")
	}

	//agregar cada simbolo a la caja de texto de simbolos
	for _, memoria := range context.ZGlobal {
		for clave, simbolo := range memoria.Variables {
			Simbolos.SetText(Simbolos.Text + "\n Tipo simbolo: " + simbolo.Categoria + "\n")
			Simbolos.SetText(Simbolos.Text + "------------------------------------------\n")
			Simbolos.SetText(Simbolos.Text + "Ambito: " + simbolo.Ambito + "\n")
			Simbolos.SetText(Simbolos.Text + "Identificador: " + clave + "\n")
			Simbolos.SetText(Simbolos.Text + "Tipo: " + simbolo.Tipo.String() + "\n")
			Simbolos.SetText(Simbolos.Text + "Valor: " + simbolo.Resultado.ValorS + "\n")
			Simbolos.SetText(Simbolos.Text + "Linea: " + fmt.Sprint(simbolo.Linea) + "| Columna: " + fmt.Sprint(simbolo.Columna) + "\n")
			Simbolos.SetText(Simbolos.Text + "------------------------------------------\n")
		}
	}

}

//FUNCION PARA MOSTRAR LA VENTANA----------------------------------------------------------------------------------------------------------------------------------------------
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
	Consola = widget.NewMultiLineEntry()
	Consola.Wrapping = fyne.TextWrapBreak
	Consola.TextStyle = fyne.TextStyle{}
	//Consola.Disable()

	Simbolos = widget.NewMultiLineEntry()
	Simbolos.Wrapping = fyne.TextWrapBreak
	//Simbolos.Disable()

	Errores = widget.NewMultiLineEntry()
	Errores.Wrapping = fyne.TextWrapBreak
	//Errores.Disable()

	//botones
	//BOTON EJECUTAR
	btn1 := widget.NewButton("Ejecutar", func() {
		Consola.SetText("")
		Simbolos.SetText("")
		Errores.SetText("")
		if Entrada.Text != "" {
			ejecutarInterprete(Entrada.Text)
		} else {
			dialog.ShowInformation("Error", "La entrada está vacía, no puede interpretar nada", myWindow)
		}
	})
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
			if err == nil {
				fmt.Print("")
				//No se hace nada
			}
			if uc != nil {
				data, _ := ioutil.ReadAll(uc)
				result := fyne.NewStaticResource("entrada", data)
				rutaCompleta = uc.URI().Path()
				ruta.SetText("RUTA: " + rutaCompleta)
				Entrada.SetText(string(result.StaticContent))
				uc.Close()
			}
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
			if err == nil {
				fmt.Print("")
				//No se hace nada
			}
			if uc != nil {
				data := []byte(Entrada.Text)
				uc.Write(data)
				rutaCompleta = uc.URI().Path()
				mensaje := "Archivo guardado exitosamente en:"
				mensaje += rutaCompleta
				dialog.ShowInformation("Información", mensaje, myWindow)
				ruta.SetText("RUTA: " + rutaCompleta)
				uc.Close()
			}
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

//MOSTRAR MAIN-----------------------------------------------------------------------------------------------------------------------------------------------------------------
func main() {
	VentanaM()
}
