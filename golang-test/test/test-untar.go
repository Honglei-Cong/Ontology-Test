
package main

import (
        "archive/tar"
        "compress/gzip"
        "os"
        "io"
        "path"
        "fmt"
)

func main() {
        mpath := "a.tgz"
        f, err := os.Open(mpath)
        if err != nil {
                panic(err)
        }
        defer f.Close()

        g, err := gzip.NewReader(f)
        if err != nil {
                panic(err)
        }
        defer g.Close()

        t := tar.NewReader(g)
        for {
                hdr, err := t.Next()
                if err == io.EOF {
                        break
                }
                if err != nil {
                        panic(err)
                }

                if hdr.Typeflag != tar.TypeDir {
                        p := path.Dir(hdr.Name)
                        if err := os.MkdirAll(p, os.ModePerm); err != nil {
                                panic(err)
                        }
                        fmt.Println(hdr.Name)
                        ow, err := os.Create(hdr.Name)
                        if err != nil {
                                panic(err)
                        }
                        if _, err := io.Copy(ow, t); err != nil {
                                panic(err)
                        }
                        ow.Close()
                }
        }
}
