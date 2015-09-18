package crypto

import (
  "fmt"
)

func ExampleSha() {
  someString := "Hello Hash"
  hashed := Sha(someString, 10)
  fmt.Println(hashed)
  // Output: 7102552d18
}
