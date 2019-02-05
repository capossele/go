package tangle

import (
    "fmt"
    "github.com/iotadevelopment/go/modules/gossip"
    "github.com/iotadevelopment/go/packages/database"
    "github.com/iotadevelopment/go/packages/ixi"
    "github.com/iotadevelopment/go/packages/transaction"
    "github.com/iotadevelopment/go/packages/network"
    "strconv"
)

var MODULE = ixi.NewIXIModule().OnRun(func() {
    transactionsDatabase := database.Get("transactions")

    counter := 0

    gossip.IXI().OnReceiveTransaction(func(peer network.Peer, transaction transaction.Transaction) {
        counter++

        err := transactionsDatabase.Set([]byte(transaction.GetHash().ToString() + strconv.Itoa(counter)), transaction.GetBytes())
        if err != nil {
            fmt.Println(err)
        }
    })
})
