## image_getter
URLをPOSTのパラメータとして画像を取得しbase64エンコードして返すくん。
（CORS回避用）

## Usage
Run server

```bash
SECRET_KEY="hoge" go run main.go
```

And request
```bash
curl -X POST -d 'url=https://...&key=hoge' http://localhost:12345/
```

## Other
You can change default port
```bash
go run main.go 8080
```
