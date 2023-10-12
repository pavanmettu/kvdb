package main

/*
 README:
 Commands to Use:
 SET KEY VALUE // Sets Key to value. If part of Txn, till COMMIT
               // is called, its in mem. If ROLLBACK, then deleted.
 GET KEY  // Gets value first from Pending Txns and then Global DB
 BEGIN   // Starts a transaction of multi keys
 ROLLBACK // Takes back all members of the transacted items
 COMMIT // Adds key-value pairs from transactions to DB
 END   // Dont commit the transactions
 EXIT  // Exit from the program
*/

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    nst := &KVMemStk{}
    for {
        fmt.Printf(">> ")
        txt, _ := reader.ReadString('\n')
        arg := strings.Fields(txt)
        switch arg[0] {
        case "SET":
            nst.Set(arg[1], arg[2], nst)
        case "GET":
            fmt.Println(nst.Get(arg[1], nst))
        case "DELETE":
            nst.Delete(arg[1], nst)
        case "BEGIN":
            nst.NewStkTx()
        case "ROLLBACK":
            nst.RollbkTx()
        case "COMMIT":
            nst.Commit()
        case "END":
            nst.RmTx()
        case "EXIT":
            os.Exit(0)
        default:
            fmt.Println("ERROR: Invalid Method")
        }
    }
}

var KVDB = make(map[string]string)

type Txn struct {
    tmap map[string]string
    next *Txn
}

type KVMemStk struct {
    head *Txn
}

func (T *KVMemStk) Get(key string) string {
    PendingTx := T.LatestTx()
    // Not part of Transaction -- Go directly to DB
    if PendingTx == nil {
        if v, ok := KVDB[key]; ok {
            return v
        }
        return ""
    }
    // go through the pending txs to check key, if not DB
    for PendingTx != nil {
        if v, ok := PendingTx.tmap[key]; ok {
            return v
        }
        PendingTx = PendingTx.next
    }
    // Check in database 
    if v, ok := KVDB[key]; ok {
        return v
    }
    return ""
}

func (T *KVMemStk) Set(key string, value string) {
    PendingTx := T.LatestTx()
    if PendingTx == nil {
        KVDB[key] = value
        return
    }
    PendingTx.tmap[key] = value
}

func (T *KVMemStk) Delete(key string) {
    PendingTx := T.LatestTx()
    if PendingTx == nil {
        delete(KVDB, key)
        return
    }
    for PendingTx != nil {
        if _, ok := PendingTx.tmap[key]; ok {
            delete(PendingTx.tmap, key)
        }
        PendingTx = PendingTx.next
    }
}

func (T *KVMemStk) NewStkTx() {
    tmpts := Txn{tmap: make(map[string]string)}
    tmpts.next = T.head
    T.head = &tmpts
}

func (T *KVMemStk) RmTx() {
    T.head.tmap = nil
    T.head.next = nil
}

func (T *KVMemStk) LatestTx() *Txn {
    return T.head
}
func (T *KVMemStk) RollbkTx() {
    tmptx := T.LatestTx()
    if tmptx != nil {
        for k, _ := range tmptx.tmap {
            delete(tmptx.tmap, k)
        }
        tmptx = tmptx.next
    }
    ts.head.tmap = nil
    ts.head.next = nil
}

func (T *KVMemStk) Commit() {
    PendingTx := T.LatestTx()
    kvmap := map[string]string{}
    for PendingTx != nil {
        for k, v := range PendingTx.tmap {
            _, ok := kvmap[k]
            if !ok {
                kvmap[k] = v
            }
        }
        PendingTx = PendingTx.next
    }
    for k, v := range kvmap {
        KVDB[k] = v
    }
    T.head.tmap = map[string]string{}
    T.head.next = nil
}
