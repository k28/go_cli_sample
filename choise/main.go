package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
)

func main() {
    // 乱数のタネを初期化
    rand.Seed(time.Now().UnixNano())

    // コマンドライン引数の数をカウント
    // 設定されていなければエラーにする
    c := len(os.Args) - 1
    if c < 1 {
        fmt.Println(os.Stderr, "[usage] %s choise1 choise2...", os.Args[0])
        os.Exit(1)
    }

    // 乱数で生成したものを表示
    fmt.Printf(os.Args[rand.Intn(c) + 1])
}

