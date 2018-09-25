
package main

import (
        "encoding/json"
        "fmt"
)

type productID struct {
        id      int          `json:"ID"`
}

type hashNode struct {
        Data    []byte       `json:"data"`
        SubNode *hashNode    `json:"subNode"`
}

type nodeFlag struct {
        Hash    *hashNode   `json:"hash"`
        Gen     uint16      `json:"gen"`
        Dirty   bool        `json:"dirty"`
}

type byteArray struct {
        Data    []byte  `json:"data"`
}

type stringMapStruct struct {
        Str1    int     `json:"str1"`
        Str2    float64 `json:"str2"`
        Str3    string  `json:"str3"`
}

type CreateProductParam struct {
	ProductCode   string `json:"product_code"`
	ProductName   string `json:"product_name"`
	RiskRate      string `json:"risk_rate"`
	Period        string `json:"period"`
	ExpAnnualRate uint64 `json:"exp_annual_rate"`
	Issuer        string `json:"issuer"`
	IssueScale    uint64 `json:"issue_scale"`
	Username      string `json:"username"`
}

func main() {
        subNode := &hashNode {
                Data:   []byte("subNodeData"),
                SubNode: nil,
        }
        x := &nodeFlag {
                Hash:   &hashNode {
                        Data:   []byte("testdata"),
                        SubNode: subNode,
                },
                Gen: 123,
                Dirty: false,
        }

        bytes, err := json.Marshal(x)
        if err != nil {
                fmt.Println(err.Error())
        }
        fmt.Println(string(bytes))

        x2 := &nodeFlag{}
        if err := json.Unmarshal(bytes, x2); err != nil {
                fmt.Println(err.Error())
        }

        fmt.Println(string(x2.Hash.Data))
        fmt.Println(string(x2.Hash.SubNode.Data))

        x3 := &byteArray{
                Data: []byte("test-string"),
        }
        bytes, err = json.Marshal(x3)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Println(string(bytes))

        x4 := &byteArray{}
        if err := json.Unmarshal(bytes, x4); err != nil {
                fmt.Println(err)
        }
        fmt.Println(string(x4.Data))

        fmt.Println("test ID marshal")
        // idStr := "{\"test-ID\":1}"
        idStr := "{}"
        id := &productID{}
        if err := json.Unmarshal([]byte(idStr), id); err != nil {
                fmt.Println(err)
        }
        fmt.Printf("unmarshaled ID: %d \n", id.id)


        strMap := &stringMapStruct {
                Str1: 1,
                Str3: "testStr3",
        }
        bytes, err = json.Marshal(strMap)
        if err != nil {
                fmt.Println(err)
        }
        strMap2 := make(map[string]interface{})
        if err := json.Unmarshal(bytes, &strMap2); err != nil {
                fmt.Println(err)
        }
        fmt.Println(strMap2)
        bytes, _ = json.Marshal(strMap2)
        fmt.Println(bytes)


        xxx := "{\"product_code\":\"b\",\"product_name\":\"b\",\"risk_rate\":\"b\",\"period\":\"100\",\"exp_annual_rate\":100,\"issuer\":\"b\",\"issue_scale\":10000,\"username\":\"b\"}"
        param := &CreateProductParam{}
        err = json.Unmarshal([]byte(xxx), param)
        if err != nil {
                fmt.Println("unmarshal failed", err.Error())
        }
}

