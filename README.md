# mp3 播放

無窮循環播放單首 

tool.exe mp3 play "d:\test.mp3" -r -1

播放多首音樂

tool.exe fs ls d:\my_music\*.mp3 | tool mp3 play

# 傳送檔案

tool.exe netcat server -p 1234 > test.txt

cat data.txt | tool.exe netcat client -t 127.0.0.1 -p 1234

# ftp server

tool server ftp -p 1021 -u arick -w 1234

# http proxy

tool server proxy -p 28080 

# http file server

tool server static -p 800 -r d:\

# 產生 "hello world" QR Code 輸出成 hello.png

tool barcode qr "hello world" -n hello

# 產生 Web 用的 base64 圖檔

cat hello.png | tool base64 encode > a.base64

# base64 圖檔還原

cat a.base64 | tool base64 decode  > a.png

# 目前系統 timestamp

tool date now -t

# json 格式化輸出

cat test.json | tool json pretty > pretty.json

# 瀏覽器開啟 markdown(.md) 檔案

tool md2html README.md -o

# 執行簡單的數學運算(javascript 語法, 不支援 ES6)

echo console.log(365*20^4) | tool js 

# HEX 方式檢視檔案

cat tool.go | tool hex

# 計算檔案的 hash

cat tool.exe | tool hash md5

cat tool.exe | tool hash crc32

cat tool.exe | tool hash sha1

cat tool.exe | tool hash sha256

cat tool.exe | tool hash sha384

cat tool.exe | tool hash sha512

#  Guerrillamail 臨時電子信箱

取得 Guerrillamail 臨時電子信箱

tool guerrillamail new

```
Email: xrirodem@guerrillamailblock.com
Timestamp: 1500457552
Alias: 8cbf80+1t1r8mbkrjzzw
Token: 4bumfh53ob4jl1l0ctnuag10c0
```

取得 Guerrillamail 郵件列表

tool guerrillamail list -k 4bumfh53ob4jl1l0ctnuag10c0

取得 Guerrillamail 指定郵件內文

tool guerrillamail fetch -k 4bumfh53ob4jl1l0ctnuag10c0 -i 1

# 目前對外 IP

tool web myip

# DNS 解析

tool web dns-resolve chuiwenchiu.wordpress.com

# 1 塊美金兌換多少台幣

tool web exchange-rate -s USD -t TWD

