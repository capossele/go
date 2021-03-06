package gossip

import (
    "github.com/iotadevelopment/go/packages/network"
    "github.com/iotadevelopment/go/packages/transaction"
    "reflect"
)

//region callbackEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type callbackEvent struct {
    callbacks map[uintptr]Callback
}

func (this *callbackEvent) Attach(callback Callback) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *callbackEvent) Detach(callback Callback) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *callbackEvent) Trigger() {
    for _, callback := range this.callbacks {
        callback()
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region dataEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type dataEvent struct {
    callbacks map[uintptr]DataConsumer
}

func (this *dataEvent) Attach(callback DataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *dataEvent) Detach(callback DataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *dataEvent) Trigger(data []byte) {
    for _, callback := range this.callbacks {
        callback(data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region transactionEvent //////////////////////////////////////////////////////////////////////////////////////////////

type transactionEvent struct {
    callbacks map[uintptr]TransactionConsumer
}

func (this *transactionEvent) Attach(callback TransactionConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *transactionEvent) Detach(callback TransactionConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *transactionEvent) Trigger(transaction *transaction.Transaction) {
    for _, callback := range this.callbacks {
        callback(transaction)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region neighborEvent /////////////////////////////////////////////////////////////////////////////////////////////////

type neighborEvent struct {
    callbacks map[uintptr]NeighborConsumer
}

func (this *neighborEvent) Attach(callback NeighborConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *neighborEvent) Detach(callback NeighborConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *neighborEvent) Trigger(neighbor *Neighbor) {
    for _, callback := range this.callbacks {
        callback(neighbor)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerEvent /////////////////////////////////////////////////////////////////////////////////////////////////////

type peerEvent struct {
    callbacks map[uintptr]PeerConsumer
}

func (this *peerEvent) Attach(callback PeerConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerEvent) Detach(callback PeerConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerEvent) Trigger(peer network.Connection) {
    for _, callback := range this.callbacks {
        callback(peer)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region errorEvent ////////////////////////////////////////////////////////////////////////////////////////////////////

type errorEvent struct {
    callbacks map[uintptr]ErrorConsumer
}

func (this *errorEvent) Attach(callback ErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *errorEvent) Detach(callback ErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *errorEvent) Trigger(err error) {
    for _, callback := range this.callbacks {
        callback(err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerDataEvent /////////////////////////////////////////////////////////////////////////////////////////////////

type peerDataEvent struct {
    callbacks map[uintptr]PeerDataConsumer
}

func (this *peerDataEvent) Attach(callback PeerDataConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerDataEvent) Detach(callback PeerDataConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerDataEvent) Trigger(peer network.Connection, data []byte) {
    for _, callback := range this.callbacks {
        callback(peer, data)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerErrorEvent ////////////////////////////////////////////////////////////////////////////////////////////////

type peerErrorEvent struct {
    callbacks map[uintptr]PeerErrorConsumer
}

func (this *peerErrorEvent) Attach(callback PeerErrorConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerErrorEvent) Detach(callback PeerErrorConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerErrorEvent) Trigger(peer network.Connection, err error) {
    for _, callback := range this.callbacks {
        callback(peer, err)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////

//region peerTransactionEvent //////////////////////////////////////////////////////////////////////////////////////////

type peerTransactionEvent struct {
    callbacks map[uintptr]PeerTransactionConsumer
}

func (this *peerTransactionEvent) Attach(callback PeerTransactionConsumer) {
    this.callbacks[reflect.ValueOf(callback).Pointer()] = callback
}

func (this *peerTransactionEvent) Detach(callback PeerTransactionConsumer) {
    delete(this.callbacks, reflect.ValueOf(callback).Pointer())
}

func (this *peerTransactionEvent) Trigger(peer network.Connection, transaction *transaction.Transaction) {
    for _, callback := range this.callbacks {
        callback(peer, transaction)
    }
}

//endregion ////////////////////////////////////////////////////////////////////////////////////////////////////////////
