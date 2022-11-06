# Go CLI tool sample

## CLIツールの作成方法

フォルダを作成して、次のコマンドをタイプ
```
$ go mod init <実行ファイル名>
```

go.modファイルが作成されて、その場所が作業フォルダのルートになる
main.goを作成して、パッケージはmainとする
```
import main

func main() {
  コードをかく
}

```

Makefile, testについてはhelloworldフォルダを参照.



## Software Design 2019-05を参考に...
- choise 複数の引数からランダムに選び出す
- processlog プロセス実行結果のログを表示

