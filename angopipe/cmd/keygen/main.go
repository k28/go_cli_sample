package main

import (
    "crypto/rand"
    "encoding/base64"
    "fmt"
    "io"
)

func main() {
    // 32バイトの配列を作成して乱数データで埋める
    key := make([]byte, 32)
    io.ReadFull(rand.Reader, key)
    // バイト列をBase64に変換する
    readableKey := base64.StdEncoding.EncodeToString(key)
    fmt.Println(readableKey)
}

