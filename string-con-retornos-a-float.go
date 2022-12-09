package main

import (
        "fmt"
)

func main() {

        str1 := "1.2345678\n2.2345678\n3.2345678"

        var float64Array [7]float64

        fmt.Println("str1:\n", str1)
        fmt.Sscanf(str1, "%g\n%g\n%g", &float64Array[0], &float64Array[1], &float64Array[2])
        fmt.Println("float64:")
        for i := 0; i < 3; i++ {
                fmt.Println(float64Array[i])
        }
        //NO FUNCIONA, imprime bytes
        for _, element := range str1 {
                fmt.Println("range:", element)
        }

        //arr := [3]string{"Jack", "James", "Jim"}
        str1 = "Jack\nJames\nJim" //imprime bytes
        arr := str1
        for index, element := range arr {
                //NO FUNCIONA: cannot use "\n" (type untyped string) as type rune
                //if element == "\n" {fmt.Println("<<< Es un retorno")}
                fmt.Println("Index", index, "has a value of", element)
        }

}

