package events

import "time"

//EVENTOS: carrega dados
type EventInterface interface {
	GetName() string
	GetDateName() time.Time
	GetPayload() interface{} //INTERFACE VAZIA POIS QUALQUER COISA PODE IMPLEMENTAR O PAYLOAD
}

//OPERAÇÕES QUANDO EVENTO É CHAMADO
type EventHandlerInterface interface {
	Handler(event EventInterface) //EXECUTA A OPERAÇÃO, POR ISSO PRECISA DO EventInterface
}

//GERENCIADOR
type EventDispatcherInterface interface {
	Register(eventName, string, handler EventHandlerInterface) error //QUANDO ESTE ESTE EVENTO FOR EXECUTADO, EXECUTA O Handler
	Dispatch(event EventInterface) error                             //FAZ COM QUE O EVENTO ACONTEÇA E QUE OS EVENTOS SEJAM EXECUTADOS
	Remove(eventName string, handler EventHandlerInterface) error    // REMOVER O EVENTO DA LISTA
	Has(eventName string, handler EventHandlerInterface) bool        //VERICAA SE O TEM UM EENTNAME COM ESTE NOME
	Clear() error                                                    //LIMPA O EVENTDISPACHER
}
