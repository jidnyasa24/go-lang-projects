// Question: Show explicit type conversion.
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)

    fmt.Println("Int:", i, "Float:", f, "Uint:", u)
}
