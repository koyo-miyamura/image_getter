## image_getter
CORS避けるためにURLをPOSTのパラメータにして画像を取得してbase64エンコードして返すくん。

## Usage
Run server

```bash
SECRET_KEY="hoge" go run main.go
```

And request
```bash
curl -X POST -d 'url=https://...&key=hoge' http://localhost:12345/
```
