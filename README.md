# tuneweathergo

## 概要

Golangの練習のため、簡単なCLIツールを作りました

Rustでも作っています。

https://github.com/tunemage/tuneweather

![image](https://user-images.githubusercontent.com/911649/167286835-59cf937b-03dd-44fc-ae4e-34d0b738b629.png)


上画像の例のように、パラメータに都市名を指定して（tokyo,osaka,nagoyaのみ対応）直近5日分の天気を取得します。

データ元は、ログインなしで使用できるWebAPIをお借りしています（ありがとうございます）

https://open-meteo.com/en

```
## 実行例
go install
~/go/bin/tuneweather tuneweather -c tokyo
```

## 注意点

* 安全性・保守性等、色々と問題があると思います。

## 初めてGolangを触ってみた感想

非常に簡単な機能しかないですが、

* コマンドパラメータの受け取り（cobra)
* WebAPIからのデータ取得(net/http)
* JSONから構造体への変換(encoding/json)

など、色々勉強になりました。

Golangは書きやすさにかなり比重が置かれてる言語のような印象を持ちました。

また、標準ライブラリの範囲で色々できるのも魅力に感じました。

次は簡単でいいので実用性のあるアプリを作ってみようと思います
