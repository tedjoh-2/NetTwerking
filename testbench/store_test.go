package main

import (
  "fmt"
  "testing"
  "../golang"
  "time"
)
/*
Connection pattern : (A) connect with (B) connect with both (C) & (D)
*/
func TestStore(t *testing.T){
  start := d7024e.StartNode(8400, "none", "2111111400000000000000000000000000000000")
  time.Sleep(50 * time.Millisecond)
  d7024e.StartNode(8401, "localhost:8400", "2111111400000000000000000000000000000001")
  d7024e.StartNode(8402, "localhost:8401", "2111111400000000000000000000000000000002")
  d7024e.StartNode(8403, "localhost:8401", "2111111400000000000000000000000000000003")
  fmt.Println("All nodes connected", start)
  contact := start.SendStoreMessage(d7024e.NewKademliaID("2111111400000000000000000000000000000000"), []byte("Test store"))
  if string(contact) != "stored"{
    t.Error("Value not stored!", contact)
  }else{
    fmt.Println("Value is", string(contact))
  }
}
