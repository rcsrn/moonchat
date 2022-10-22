package view

import (
	"fmt"
)

type Printer struct {
}

const (
	WELCOME = "Bienvenido a Moonchat!"
	IDENTIFY_INSTRUCTION = "/username <NombreDeUsuario>"
	CLOSE_INSTRUCTION = "/close"
	STATUS_INSTRUCTION = "/status <AWAY/ACTIVE/BUSY>"
	NEW_ROOM_INSTRUCTION = "/newroom <NombreDeCuarto>"
	JOIN_ROOM_INSTRUCTION = "/joinroom <NombreDeCuarto>"
	INVITE_INSTRUCTION = "/invite <NombreDeCuarto> <Usuario(s)>"
	LEAVE_ROOM_INSTRUCTION = "/leaveroom <NombreDeCuarto>"
	USER_LIST_INSTRUCTION = "/userlist"
	ROOM_USER_LIST_INSTRUCTION = "/roomlist <NombreDelCuarto"
	ROOM_MESSAGE_INSTRUNCTION = "/roommessage <Mensaje>"
	PRIVATE_MESSAGE_INSTRUCTION = "/private <UsuarioDestinatario> <Mensaje>"
	CONNECT_ERROR = "Algo ha salido mal al tratar de conectar con el servidor. :("
)

func GetPrinterInstance() *Printer {
	return nil
}

func (printer *Printer) WarnConnectError() {
	fmt.Println(CONNECT_ERROR)
}

func (printer *Printer) PrintInstructions() {
	fmt.Println(WELCOME + "\n")
	fmt.Println("Escribe para enviar un mensaje publico.")
	fmt.Println("O puedes realizar las siguientes operaciones: \n")
	fmt.Println("*Para ingresar nombre de usuario: " + IDENTIFY_INSTRUCTION)
	fmt.Println("*Para cambiar de estado: " + STATUS_INSTRUCTION)
	fmt.Println("*Para obtener la lista de usuarios: " + USER_LIST_INSTRUCTION)
	fmt.Println("*Para enviar un mensaje privado a algun usuario: " + PRIVATE_MESSAGE_INSTRUCTION)
	fmt.Println("*Para crear una sala: " + NEW_ROOM_INSTRUCTION)	
	fmt.Println("*Para invitar a usuarios a una sala: " + INVITE_INSTRUCTION)
	fmt.Println("*Para enviar un mensaje a una sala: " + ROOM_MESSAGE_INSTRUNCTION)
	fmt.Println("*Para unirse a una sala: " + JOIN_ROOM_INSTRUCTION)
	fmt.Println("*Para salir de una sala: " + LEAVE_ROOM_INSTRUCTION)
	fmt.Println("*Para desconectarse del chat: " + CLOSE_INSTRUCTION + "\n")
}

func (printer *Printer) RequestUserName() {
	fmt.Println("Ingrese nombre de usuario:")
}

func (printer *Printer) Use() {
	fmt.Println("Uso: ./clien_controller <host> <port>\n")
}


