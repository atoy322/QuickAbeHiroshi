package main

import (
    "os/exec"
)

func main() {
    com := exec.Command("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe", "http://abehiroshi.la.coocan.jp")
    e := com.Run()

    if e != nil {
        panic(e)
    }
}
