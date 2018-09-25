package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func showMatch(m [][]byte) {
	for i := 0; i < len(m); i++ {
		fmt.Printf("%d: %s \n", i, m[i])
	}
}

type SelectStmt struct {
	Fields  []string `json:"fields"`
	Tables  []string `json:"tables"`
	StartID int      `json:"start_id"`
	EndID   int      `json:"end_id"`
}

func ParseSelectStmt(stmt string) (*SelectStmt, error) {
	sel := &SelectStmt{nil, nil, -1, -1}

	src := []byte(stmt)
	pat := `(?i:select)\s+((?:[\w*]+,\s+)*[\w*]+)\s+(?i:from)\s+(\w+(?:,\s*\w+)*)(\s*(?i:where))*([[:ascii:]]*)`
	reg := regexp.MustCompile(pat)

	fields := []byte{}
	tables := []byte{}
	conds := []byte{}

	if reg.Match(src) {
		m1 := reg.FindSubmatch(src)
		//showMatch(m1)
		fields = m1[1]
		tables = m1[2]
		conds = m1[4]
	} else {
		return sel, errors.New("Failed to find select/from")
	}

    for _, s := range strings.Split(string(fields), ",") {
            sel.Fields = append(sel.Fields, strings.TrimSpace(s))
    }
    for _, s := range strings.Split(string(tables), ",") {
            sel.Tables = append(sel.Tables, strings.TrimSpace(s))
    }

	reg3 := regexp.MustCompile(`\s*(?i:id)\s*>\s*(\d+)`)
	if reg3.Match(conds) {
		m := reg3.FindSubmatch(conds)
		val, err := strconv.Atoi(string(m[1]))
		if err == nil {
			sel.StartID = val
		} else {
			return sel, fmt.Errorf("invalid gt %s", string(conds))
		}
	} else {
		fmt.Println("no gt matched")
	}

	reg4 := regexp.MustCompile(`\s*(?i:id)\s*<\s*(\d+)`)
	if reg4.Match(conds) {
		m := reg4.FindSubmatch(conds)
		val, err := strconv.Atoi(string(m[1]))
		if err == nil {
			sel.EndID = val
		} else {
			return sel, fmt.Errorf("invalid lt %s", string(conds))
		}
	} else {
		fmt.Println("no lt matched")
	}

	return sel, nil
}

func main() {
	stmt := "select a from b"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

	stmt = "select * from b"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

	stmt = "select   aa, ab, ac from  xx,yy"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

	stmt = "select   a from  xx, yy where id > 10"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

	stmt = "select   aa, ab, ac from  xx, yy where id > 10"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

	stmt = "select   aa, ab, ac from  xx, yy where id > 10 and ID < 20"
	fmt.Println(stmt)
	fmt.Println(ParseSelectStmt(stmt))
	fmt.Printf("=================\n\n")

}
