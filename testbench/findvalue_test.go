package main

import (
  "fmt"
  "testing"
  "../golang"
  "time"
)

func TestFindValue(t *testing.T){
  start := d7024e.StartNode(8500, "none", "2111111400000000000000000000000000000000")
  time.Sleep(50 * time.Millisecond)
  d7024e.StartNode(8501, "localhost:8500", "2111111400000000000000000000000000000001")
  d7024e.StartNode(8502, "localhost:8500", "2111111400000000000000000000000000000002")
  d7024e.StartNode(8503, "localhost:8500", "2111111400000000000000000000000000000003")
  fmt.Println("All nodes connected")
  contact := start.SendStoreMessage(d7024e.NewKademliaID("2111111400000000000000000000000000000000"), []byte("Test store"))
  time.Sleep(50 * time.Millisecond)
  if string(contact) != "stored"{
    t.Error("Value not stored!")
  }else{
    fmt.Println("Complete store.")
    //After storing an item on node 8401, look it up.
    find := start.SendFindValueMessage(d7024e.NewKademliaID("2111111400000000000000000000000000000000"))
    time.Sleep(50 * time.Millisecond)
    matchTo := d7024e.NewKademliaID("2111111400000000000000000000000000000000")
    if (find.Key == *matchTo) {
      fmt.Println("Successful find : ", string(find.Value))
    }else{
      t.Error("Did not find item!", find)
    }
  }
}
