package events

import (
	"sync"
	"time"
)

// EVENTOS: carrega dados
type EventInterface interface {
	GetName() string
	GetDateName() time.Time
	GetPayload() interface{} //INTERFACE VAZIA POIS QUALQUER COISA PODE IMPLEMENTAR O PAYLOAD
}

// OPERAÇÕES QUANDO EVENTO É CHAMADO
type EventHandlerInterface interface {
	Handler(event EventInterface, wg *sync.WaitGroup) //EXECUTA A OPERAÇÃO, POR ISSO PRECISA DO EventInterface
}

// GERENCIADOR
type EventDispatcherInterface interface {
	Register(eventName, string, handler EventHandlerInterface) error //MÉTODO REGISTER REGISTRA O EVENTO e PASSA AS OPERAÇÕES
	Dispatch(event EventInterface) error                             //FAZ COM QUE O EVENTO ACONTEÇA E QUE OS EVENTOS SEJAM EXECUTADOS
	Remove(eventName string, handler EventHandlerInterface) error    // REMOVER O EVENTO DA LISTA
	Has(eventName string, handler EventHandlerInterface) bool        //VERIFICA SE O TEM UM EVENTO COM HANDLER ESTÁ REGISTRADO
	Clear() error                                                    //LIMPA O EVENTDISPACHER
}
