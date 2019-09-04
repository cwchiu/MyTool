package web

import (
	"encoding/json"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

//ConfigJsonBody json request body.
type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

func CaptchaDemoHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(`<div>

    <img id="captch" src="" style="width:200px;height:200px">
    <div>
        <div><input type="text" id="txtCode" value=""></div>
        <button onclick="onVerify()">Verify</button>
    </div>
    <script>
        let captcha = null;
        
        function onVerify(){
            captchaVerify(document.getElementById('txtCode').value)
                .then( jd => {
                    document.getElementById('txtCode').value = '';
                    if(jd.code == "success"){
                        alert("pass");
                        return captchaGet();
                    }else{
                        alert("fail");
                    }
                })
                .catch(console.error);
        }
        
        
        async function captchaVerify(code) {
            const resp = await fetch('/captcha/verify', {
                method: 'post',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({"Id":captcha.captchaId,"VerifyValue":code})
            });
            
            const data = await resp.json();
            return data;
        }
        
        async function captchaGet() {
            const resp = await fetch('/captcha/get', {
                method: 'post',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({"CaptchaType":"digit","VerifyValue":"","ConfigAudio":{"CaptchaLen":6,"Language":"zh"},"ConfigCharacter":{"Height":60,"Width":240,"Mode":2,"ComplexOfNoiseText":0,"ComplexOfNoiseDot":0,"IsUseSimpleFont":true,"IsShowHollowLine":false,"IsShowNoiseDot":false,"IsShowNoiseText":false,"IsShowSlimeLine":false,"IsShowSineLine":false,"CaptchaLen":6},"ConfigDigit":{"Height":80,"Width":240,"CaptchaLen":5,"MaxSkew":0.7,"DotCount":80}})
            });
            
            captcha = await resp.json();
            document.getElementById("captch").src = captcha.data;
        }
        
        captchaGet().catch( console.error );
        
        ( async() =>{
            
        })();
    </script>
    </div>`))
}

// base64Captcha create http handler
func GetHandler(w http.ResponseWriter, r *http.Request) {
	//parse request parameters
	//接收客戶端發送來的請求參數
	decoder := json.NewDecoder(r.Body)
	var postParameters ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
		http.Error(w, "bad request", 400)
		return
	}
	defer r.Body.Close()

	//create base64 encoding captcha
	//創建base64圖像驗證碼

	var config interface{}
	switch postParameters.CaptchaType {
	case "audio":
		config = postParameters.ConfigAudio
	case "character":
		config = postParameters.ConfigCharacter
	default:
		config = postParameters.ConfigDigit
	}
	//GenerateCaptcha 第一個參數為空字符串,包會自動在服務器一個隨機種子給你產生隨機uiid.
	captchaId, digitCap := base64Captcha.GenerateCaptcha(postParameters.Id, config)
	base64Png := base64Captcha.CaptchaWriteToBase64Encoding(digitCap)

	//你也可以是用默認參數 生成圖像驗證碼
	//base64Png := captcha.GenerateCaptchaPngBase64StringDefault(captchaId)

	//set json response
	//設置json響應

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": 1, "data": base64Png, "captchaId": captchaId, "msg": "success"}
	json.NewEncoder(w).Encode(body)
}

// base64Captcha verify http handler
func VerifyHandle(w http.ResponseWriter, r *http.Request) {

	//parse request parameters
	//接收客戶端發送來的請求參數
	decoder := json.NewDecoder(r.Body)
	var postParameters ConfigJsonBody
	err := decoder.Decode(&postParameters)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	//verify the captcha
	//比較圖像驗證碼
	verifyResult := base64Captcha.VerifyCaptcha(postParameters.Id, postParameters.VerifyValue)

	//set json response
	//設置json響應
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	body := map[string]interface{}{"code": "error", "data": "驗證失敗", "msg": "captcha failed"}
	if verifyResult {
		body = map[string]interface{}{"code": "success", "data": "驗證通過", "msg": "captcha verified"}
	}
	json.NewEncoder(w).Encode(body)
}

func RegisterCaptacha() {
	http.HandleFunc("/captcha/get", GetHandler)
	http.HandleFunc("/captcha/verify", VerifyHandle)
	http.HandleFunc("/captcha/demo", CaptchaDemoHandle)
}
