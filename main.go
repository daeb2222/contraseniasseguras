package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dlclark/regexp2"
)

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.DarkTheme())
	w := a.NewWindow("Contraseñas seguras")
	w.Resize(fyne.NewSize(300, 400))

	labelTitulo := widget.NewLabel("Verificador de contraseñas")
	labelTitulo.TextStyle.Bold = true
	labelTitulo.TextStyle.Italic = true

	labelCorreo := widget.NewLabel("Correo")
	labelContrasena := widget.NewLabel("Contraseña")
	labelConfirmar := widget.NewLabel("Confirmar contraseña")
	labelVerificar := widget.NewLabel("")

	inputCorreo := widget.NewEntry()
	inputCorreo.SetPlaceHolder("correo@ejemplo.com")
	inputContrasena := widget.NewPasswordEntry()
	inputContrasena.SetPlaceHolder("Contraseña")
	inputConfirmar := widget.NewPasswordEntry()
	inputConfirmar.SetPlaceHolder("Confirmar contraseña")

	botonConfirmar := widget.NewButton("Confirmar", func() {
		expression := regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[{},\[\]\-+.*,:;()]).{12,}$`, 0)

		if isMatch, _ := expression.MatchString(inputContrasena.Text); isMatch {
			if inputContrasena.Text == inputConfirmar.Text {
				labelVerificar.SetText("Usuario Creado")
			} else {
				labelVerificar.SetText("Los passwords no coinciden")
			}
		} else {
			labelVerificar.SetText("Mínimo doce caracteres de longitud\n Una letra mayúscula. \nUna letra minúscula. \nUn carácter especial de los siguientes: {,}, [,], -, +,. , (, ) \nLa cadena no puede contener espacios en blanco \n .")
		}
	})
	contenedor := container.NewVBox(
		labelTitulo,
		labelCorreo,
		inputCorreo,
		labelContrasena,
		inputContrasena,
		labelConfirmar,
		inputConfirmar,
		botonConfirmar,
		labelVerificar)
	w.SetContent(contenedor)

	w.ShowAndRun()

}
