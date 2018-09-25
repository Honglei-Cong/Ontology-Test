
package main

import (
        "fmt"
        "sort"
        "strconv"
        "bytes"
        "encoding/binary"
)

type Trans struct {
        Amount  int
        Date    string
}

type ByAmount []Trans
func (a ByAmount) Len() int { return len(a) }
func (a ByAmount) Less(i, j int) bool { return a[i].Amount < a[j].Amount }
func (a ByAmount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByDate []Trans
func (a ByDate) Len() int { return len(a) }
func (a ByDate) Less (i, j int) bool { return a[i].Date > a[j].Date }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func main() {

        count := 20
        trans := make([]Trans, 0, count)
        for i := count; i > 0; i-- {
                trans = append(trans, Trans{ Amount: i, Date: strconv.Itoa(100 + i)})
        }

        sort.Sort(ByAmount(trans))
        fmt.Println(trans)

        sort.Sort(ByDate(trans))
        fmt.Println(trans)

        nums := []int{100, 500, 400, 300, 200}
        sort.Ints(nums)
        fmt.Println(nums)

        numBytes := make([]string, 0)
        for _, n := range nums {
                buf := new(bytes.Buffer)
                if err := binary.Write(buf, binary.BigEndian, uint64(n)); err != nil {
			fmt.Println("err:", err)
		}
                fmt.Printf("%d, %v \n", n, buf.Bytes())
                numBytes = append(numBytes, strconv.Itoa(n) + buf.String())
        }
        sort.Strings(numBytes)
        for _, x := range numBytes {
                fmt.Printf("%v \n", x)
        }

        b := []byte{0x0, 0x0, 0x03, 0xe8}
        buf := bytes.NewBuffer(b)
        var x int32
        binary.Read(buf, binary.BigEndian, &x)
        fmt.Println(x)
	fmt.Println(b)
}

