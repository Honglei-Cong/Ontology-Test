
package main

import (
        "fmt"
        "reflect"
        "encoding/json"
)

type XStruct struct {
        A string    `json:"a"`
        B int       `json:"b"`
        C float64   `json:"c"`
}

func printTypeInfo(x interface{}) {
        typ := reflect.TypeOf(x)
        if typ.Kind() == reflect.Ptr {
                typ = typ.Elem()
        } else if typ.Kind() != reflect.Struct {
                fmt.Println("Invalid type")
        }

        fmt.Println(typ.Name())
        for i := 0; i < typ.NumField(); i++ {
                fmt.Printf("%s, %s, %d \n", typ.Field(i).Name, typ.Field(i).Tag.Get("json"), typ.Field(i).Type.Kind())
        }

        bytes, err := json.Marshal(&x)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(string(bytes))
}

func updateX(x *XStruct) {
        x.B = 100

        y := *x
        y.A = "testA"
}

func main() {
        x := XStruct {
                A: "a", 
                B: 1,
                C: 1.2,
        }

        printTypeInfo(x)
        printTypeInfo(&x)

        updateX(&x)
        fmt.Printf("%v\n", x)
}
