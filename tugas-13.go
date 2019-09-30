package main

import "fmt"
import "encoding/base64"
import "crypto/sha1"

func main(){
  var nama = "Mukhamat Jafar"

  var encodestring = base64.StdEncoding.EncodeToString([]byte(nama))
  fmt.Println("encoding string: ", encodestring)

  var sha = sha1.New()
  sha.Write([]byte(nama))
  var enkripsi = sha.Sum(nil)
  var stringenkripsi = fmt.Sprintf("%x", enkripsi)

  fmt.Println(stringenkripsi)

}
