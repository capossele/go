package network

import "net"

type Connection interface {
    GetProtocol() string
    GetConnection() net.Conn
    Write(data []byte)
    OnReceiveData(callback DataConsumer) Connection
    OnDisconnect(callback Callback) Connection
    OnError(callback ErrorConsumer) Connection
    TriggerReceiveData(data []byte) Connection
    TriggerDisconnect() Connection
    TriggerError(err error) Connection
    HandleConnection()
}

type Callback func()

type ErrorConsumer func(err error)

type DataConsumer func(data []byte)
