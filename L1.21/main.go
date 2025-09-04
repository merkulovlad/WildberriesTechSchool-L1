package main

import "fmt"

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
    fmt.Println("Client inserts Lightning connector into computer.")
    com.InsertIntoLightningPort()
}

type Computer interface {
	InsertIntoLightningPort()
}

type MacBook struct {
}

func (m *MacBook) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is inserted into MacBook.")
}

type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is inserted into Windows.")
}

type WindowsAdapter struct {
    windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := Client{}
	
	mac := MacBook{}
	client.InsertLightningConnectorIntoComputer(&mac)
	win := Windows{}
	adapter := WindowsAdapter{windowMachine: &win}
	client.InsertLightningConnectorIntoComputer(&adapter)
}


// Пример использования паттерна "Адаптер" в Go.
// Например: скачиваем данные в формате XML, а работаем с JSON.
// В этом случае XML будет адаптироваться в JSON.
// + -  Отделяет и скрывает от клиента подробности преобразования различных интерфейсов.
// - -  Усложняет код программы за счёт введения дополнительных классов.
