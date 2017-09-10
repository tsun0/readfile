## Description

- flagパッケージの勉強用
- 特定のファイル・特定の行を読み込む

## Use

`git clone https://github.com/tsun0/readfile.git`

## Run on Mac localhost:8080

```
# readfileのディレクトリに移動
cd readfile

# バイナリファイルを作成
go build

# "readfile"と書かれたファイルが作成されるので、実行
./readfile -file [ファイル名] -line[指定の行数]
```