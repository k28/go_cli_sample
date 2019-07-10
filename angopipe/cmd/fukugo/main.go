/*
angoで暗号化した物を復号化する
    cat ango.txt | ./fukugo
*/
package main

import (
	"angopipe"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	gcm, err := angopipe.Prepare()
	if err != nil {
		fmt.Fprintf(os.Stdout, "")
	}

	// 先頭に12バイトのnonceが入っているので取り除く
	nonce := make([]byte, 12)
	n, err := io.ReadFull(os.Stdin, nonce)
	if n != 12 || err != nil {
		fmt.Fprintf(os.Stderr, "Can't read nonce: %v\n", err)
		os.Exit(1)
	}

	encrepted, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %v\n", err)
		os.Exit(1)
	}

	// 復号化の処理
	result, err := gcm.Open(nil, nonce, encrepted, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption Error: %v\n", err)
		os.Exit(1)
	}

	os.Stdout.Write(result)
}
