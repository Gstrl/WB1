package main

import (
	"WB1/Tasks/Task_21/data"
	data_service "WB1/Tasks/Task_21/data-service"
)

//Представим что мы работаем только с форматом json, но сервис аналитики работает с xml
//JsonDocument не реализует SendXmlData(), чтобы не добавлять этот метод структруе
//создаем адаптер этой структуры JsonDocumentAdapter и NewJsonDocumentAdapter
//который принимает JsonDocument и создает JsonDocumentAdapter.
//Адаптер конвертирует json в xml и реализует SendXmlData(), а соответсвено
//и интерфейс AnalyticalDataService.
//Теперь не меняя ничего в пакете data, JsonDocument через адаптер реализует нужный нам интерфейc.

type JsonDocumentAdapter struct {
	*data.JsonDocument
}

func (doc JsonDocumentAdapter) ConvertToXml() string {
	return "<xml>...</xml>"
}

func (adapter *JsonDocumentAdapter) SendXmlData() {
	s := adapter.ConvertToXml()
	println("send json convert to xml", s)
}

func NewJsonDocumentAdapter(json *data.JsonDocument) *JsonDocumentAdapter {
	return &JsonDocumentAdapter{json}
}

func main() {
	jsonDoc := NewJsonDocumentAdapter(&data.JsonDocument{})
	data_service.AnalyticalDataService(jsonDoc).SendXmlData()
	xmlDoc := data_service.XmlDocument{}
	data_service.AnalyticalDataService(xmlDoc).SendXmlData()
}
