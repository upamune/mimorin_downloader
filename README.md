# mimorin_downloader
[![Circle CI](https://circleci.com/gh/upamune/mimorin_downloader/tree/master.svg?style=svg)](https://circleci.com/gh/upamune/mimorin_downloader/tree/master)


## What is this ?
三森すずこさんの画像をBingから約50件取得してきます。CPU数に合わせて並列ダウンロードを行います。

## Support OS

- Windows 64-bit
- MacOSX  64-bit
- Linux   64-bit
- FreeBSD 64-bit
- Solaris 64-bit

## Install
### For User
実行ファイルを[ダウンロード](https://github.com/upamune/mimorin_downloader/releases)してパスが通っている場所に配置してください。

### For Developer

```bash
go get github.com/upamune/mimorin_downloader
```

## Uses

### Setup
環境変数 ```BING_API_KEY``` をセットする必要があります。 ```BING_API_KEY``` には [BingSearchAPI](https://datamarket.azure.com/dataset/bing/search) に登録して、 [Your Account](https://datamarket.azure.com/account) から取得できるプライマリキーをセットします。

```
export BING_API_KEY="YOUR_API_KEY"
```

### Run

```
mimorin_downloader
```
