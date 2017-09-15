package main

import (
  "fmt"
  "testing"
  "../golang"
  "time"
)

/*
Connection pattern : nodes B,C,D are connected with A.
A request a store on node all other nodes and then searches for it.
*/
func TestFindValue(t *testing.T){
  A := d7024e.NewKademlia(8500, "5111111400000000000000000000000000000000")
  A.Start(8500)
  time.Sleep(50 * time.Millisecond)

  B := d7024e.NewKademlia(8501, "5111111400000000000000000000000000000001")
  B.Start(8501)
  B.Ping("localhost:8500")
  time.Sleep(50 * time.Millisecond)

  C := d7024e.NewKademlia(8502, "5111111400000000000000000000000000000002")
  C.Start(8502)
  C.Ping("localhost:8500")
  time.Sleep(50 * time.Millisecond)

  D := d7024e.NewKademlia(8503, "5111111400000000000000000000000000000003")
  D.Start(8503)
  D.Ping("localhost:8500")
  time.Sleep(50 * time.Millisecond)

  fmt.Println("All nodes connected")
  contact := A.SendStoreMessage(d7024e.NewKademliaID("5111111400000000000000000000000000000000"), []byte("Test store"))
  time.Sleep(50 * time.Millisecond)

  if string(contact) != "stored"{
    t.Error("Value not stored!")
  }else{
    fmt.Println("Complete store.")
    //After storing an item on node 8401, look it up.
    find := A.SendFindValueMessage(d7024e.NewKademliaID("5111111400000000000000000000000000000000"))
    time.Sleep(50 * time.Millisecond)
    matchTo := d7024e.NewKademliaID("5111111400000000000000000000000000000000")
    if (find.Key == *matchTo) {
      fmt.Println("Successful find : ", string(find.Value))
    }else{
      t.Error("Did not find item!", find)
    }
  }
}
