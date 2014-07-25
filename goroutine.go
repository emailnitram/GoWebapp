  // A _goroutine_ is a lightweight thread of execution.
  
  package main
  
  import (
      "fmt"
      //"runtime"
  )
  
  func f(from string) {
      for i := 0; i < 33; i++ {
          fmt.Println(from, ":", i)
      }
  }
  
  func main() {
      //runtime.GOMAXPROCS(2)
  
      // Suppose we have a function call `f(s)`. Here's how
      // we'd call that in the usual way, running it
      // synchronously.
      f("direct")
  
      // To invoke this function in a goroutine, use
      // `go f(s)`. This new goroutine will execute
      // concurrently with the calling one.
      go f("goroutine")
      go f("thaigogo")
      go f("happy")
      go f("library")
      go f("banana")
  
      // You can also start a goroutine for an anonymous
      // function call.
  //    go func(msg string) {
  //        fmt.Println(msg)
  //    }("going")
  
      // Our two goroutines are running asynchronously in
      // separate goroutines now, so execution falls through
      // to here. This `Scanln` code requires we press a key
      // before the program exits.
      var input string
      fmt.Scanln(&input)
      fmt.Println("done")
  }
  

//  package main
//  
//  import (
//      "fmt"
//      "runtime"
//      "sync"
//  )
//  
//  func main() {
//      runtime.GOMAXPROCS(2)
//  
//      var wg sync.WaitGroup
//      wg.Add(2)
//  
//      fmt.Println("Starting Go Routines")
//      go func() {
//          defer wg.Done()
//  
//          for char := 'a'; char < 'a'+26; char++ {
//              fmt.Printf("%c ", char)
//          }
//      }()
//  
//      go func() {
//          defer wg.Done()
//  
//          for number := 1; number < 27; number++ {
//              fmt.Printf("%d ", number)
//          }
//      }()
//  
//      fmt.Println("Waiting To Finish")
//      wg.Wait()
//  
//      fmt.Println("\nTerminating Program")
//  }
//  
