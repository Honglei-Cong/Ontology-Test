
package main

import(
        "fmt"
)

func main() {
    str := "abc"
    fmt.Println(perm(str, 0))
    fmt.Println(perm(str, 1))
    fmt.Println(perm(str, 2))
    fmt.Println(perm(str, 3))
    fmt.Println(perm(str, 4))
    fmt.Println(perm(str, 5))
}

func perm(str string, kth int) string {
        s, _ := perm_("", str, kth)
        return s
}

func perm_(prefix string, str string, kth int) (string, int) {
    n := len(str)
    if n == 0 || kth == 0{
        // fmt.Println(prefix)
        return prefix + str, 1
    } else {
        k := kth
        for i := 0; i < n; i++ {
                s, m := perm_(prefix+str[i:i+1], str[0:i]+str[(i+1):n], k)
                if k == m {
                        return s, kth
                }
                k -= m
        }
        return "", kth - k
    }
}

