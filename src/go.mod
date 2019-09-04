module github.com/cwchiu/MyTool

go 1.13

require (
	github.com/PuerkitoBio/goquery v1.5.0
	github.com/abbot/go-http-auth v0.4.0
	github.com/anmitsu/go-shlex v0.0.0-20161002113705-648efa622239 // indirect
	github.com/atotto/clipboard v0.1.2
	github.com/bogem/id3v2 v1.1.1
	github.com/boltdb/bolt v1.3.1
	github.com/boombuler/barcode v1.0.0
	github.com/bradfitz/go-smtpd v0.0.0-20170404230938-deb6d6237625
	github.com/cavaliercoder/grab v2.0.0+incompatible
	github.com/cwchiu/go-pastebin v0.0.0-20180221142448-e19cd4f8d4a1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elazarl/goproxy v0.0.0-20190711103511-473e67f1d7d2
	github.com/fclairamb/ftpserver v0.0.0-20181215165500-058ccd38e144
	github.com/geckoboard/cli-table v0.0.0-20161124114750-ba78d7928542
	github.com/gliderlabs/ssh v0.2.2
	github.com/go-ole/go-ole v1.2.4
	github.com/gorilla/websocket v1.4.1
	github.com/hajimehoshi/go-mp3 v0.2.1
	github.com/hajimehoshi/oto v0.4.0
	github.com/jlaffaye/ftp v0.0.0-20190828173736-6aaa91c7796e
	github.com/kapmahc/epub v0.1.1
	github.com/mattn/go-runewidth v0.0.4 // indirect
	github.com/mattn/go-sqlite3 v1.11.0
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/miekg/dns v1.1.16
	github.com/mojocn/base64Captcha v0.0.0-20190801020520-752b1cd608b2
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/mozillazg/go-pinyin v0.15.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nosixtools/solarlunar v0.0.0-20170821083946-832f6f39411c
	github.com/olekukonko/tablewriter v0.0.1
	github.com/parnurzeal/gorequest v0.2.15
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/sftp v1.10.1
	github.com/qianlnk/guerrillamail v0.0.0-20170509071238-356e8e64acc3
	github.com/rakyll/statik v0.1.6
	github.com/ritchie46/GOPHY v0.0.0-20170315173114-9b8a7f05cfa1
	github.com/robertkrimen/otto v0.0.0-20180617131154-15f95af6e78d
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/github_flavored_markdown v0.0.0-20181002035957-2122de532470
	github.com/shurcooL/highlight_diff v0.0.0-20181222201841-111da2e7d480 // indirect
	github.com/shurcooL/highlight_go v0.0.0-20181215221002-9d8641ddf2e1 // indirect
	github.com/shurcooL/octicon v0.0.0-20181222203144-9ff1a4cf27f4 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/skratchdot/open-golang v0.0.0-20190402232053-79abb63cd66e
	github.com/sourcegraph/annotate v0.0.0-20160123013949-f4cad6c6324d // indirect
	github.com/sourcegraph/syntaxhighlight v0.0.0-20170531221838-bd320f5d308e // indirect
	github.com/spf13/cobra v0.0.5
	github.com/tuotoo/qrcode v0.0.0-20190222102259-ac9c44189bf2
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/image v0.0.0-20190902063713-cb417be4ba39
	golang.org/x/text v0.3.0
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/cheggaaa/pb.v1 v1.0.28
	gopkg.in/ldap.v2 v2.5.1
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/redis.v5 v5.2.9
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
)

replace github.com/qianlnk/guerrillamail => ./3rd/guerrillamail
