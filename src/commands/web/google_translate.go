package web

import (
	"fmt"
	// "encoding/json"
	// "net/url"
	"github.com/parnurzeal/gorequest"
	"github.com/robertkrimen/otto"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

// https://github.com/Neoyyy/google-CommandLine-Translation-Tool

func getTKKCode() string {
	request := gorequest.New()
	_, body, _ := request.Get("https://translate.google.cn").End()
	// fmt.Println(body)
	re := regexp.MustCompile(`TKK=(eval.*?\(\)\)'\));`)
	ret := re.FindStringSubmatch(body)
	return ret[1]
}

func parseResponse(vm *otto.Otto, response string) otto.Value {
	code := `
        function dealresult(obj){
            var result = {
                orginal : '',
                pronunciation : '',
                translate : '',
                intro : '',
                other : '',
                
            }
            
            if (obj[1] == null) {
                if(obj[0][0][0]){
                    result.translate = obj[0][0][0];
                }else{
                    result = 'translate error';
                }
            }else{
                cleannull(obj);
                if (obj[0][0][1] !== undefined) {
                    result.orginal = obj[0][0][1];
                }
                if (obj[0][1][0] !== undefined) {
                    result.pronunciation = (obj[2] == 'en')? obj[0][1][1] : obj[0][1][0];

                }
                if (obj[0][0][0] !== undefined) {
                    result.translate = obj[0][0][0];
                }
                if (obj[1][0][0] !== undefined && obj[1][0][2][0][1] !== undefined) {
                    result.intro = obj[1][0][0]+','+obj[1][0][2][0][1] ;
                }else if (obj[1][0][0] !== undefined) {
                    result.intro = obj[1][0][0] ;
                }else if (obj[1][0][2][0][1] !== undefined) {
                    result.intro = obj[1][0][2][0][1] ;
                }
                if (obj[6][0] !== undefined) {

                    result.other = (obj[2] == 'en')? (obj[6][0][1][0]+obj[6][0][1][1]) : obj[6][0];

                }
            }


            return result;


        }


        function cleannull(arr){



            if ( arr == null) {
                return;
            }else{

                var i = arr.length;

                while(i--){
                    if (arr[i] == null) {
                        arr.splice(i,1);
                    
                    }else{
                        if (Array.isArray(arr[i])) {
                                cleannull(arr[i]);

                        }
                    }
                }


                return arr;

            }




        }
        
    `
	// fmt.Println(response)
	code += "dealresult( " + response + `)`
	// fmt.Println(code)
	value, err := vm.Run(code)
	if err != nil {
		panic(err)
	}

	return value
}

func trans(vm *otto.Otto, tk string, text string, to string) otto.Value {
	request := gorequest.New()
	request.Get("https://translate.google.cn/translate_a/single?client=t&dt=at&dt=bd&dt=ex&dt=ld&dt=md&dt=qca&dt=rw&dt=rm&dt=ss&dt=t&ie=UTF-8&oe=UTF-8&otf=2&ssel=0&tsel=0&kc=2")
	request.Param("sl", "auto")
	request.Param("tl", to)
	request.Param("hl", to)
	request.Param("tk", tk)
	request.Param("q", text)
	_, body, err := request.End()
	if err != nil {
		panic(err)
	}
	return parseResponse(vm, body)
}

func vmGetTk(vm *otto.Otto, tkk_code string, text string) string {
	text_escape := strings.Replace(text, "'", `\'`, 0)
	js_code := `
        function sM(a) {
            var b;
            if (null !== yr)
                b = yr;
            else {
                b = wr(String.fromCharCode(84));
                var c = wr(String.fromCharCode(75));
                b = [b(), b()];
                b[1] = c();
                b = (yr = window[b.join(c())] || "") || ""
            }
            var d = wr(String.fromCharCode(116))
                , c = wr(String.fromCharCode(107))
                , d = [d(), d()];
            d[1] = c();
            c = "&" + d.join("") + "=";
            d = b.split(".");
            b = Number(d[0]) || 0;
            for (var e = [], f = 0, g = 0; g < a.length; g++) {
                var l = a.charCodeAt(g);
                128 > l ? e[f++] = l : (2048 > l ? e[f++] = l >> 6 | 192 : (55296 == (l & 64512) && g + 1 < a.length && 56320 == (a.charCodeAt(g + 1) & 64512) ? (l = 65536 + ((l & 1023) << 10) + (a.charCodeAt(++g) & 1023),
                    e[f++] = l >> 18 | 240,
                    e[f++] = l >> 12 & 63 | 128) : e[f++] = l >> 12 | 224,
                    e[f++] = l >> 6 & 63 | 128),
                    e[f++] = l & 63 | 128)
            }
            a = b;
            for (f = 0; f < e.length; f++)
                a += e[f],
                    a = xr(a, "+-a^+6");
            a = xr(a, "+-3^+b+-f");
            a ^= Number(d[1]) || 0;
            0 > a && (a = (a & 2147483647) + 2147483648);
            a %= 1E6;
            return c + (a.toString() + "." + (a ^ b))
        }

        var yr = null;
        var wr = function(a) {
            return function() {
                return a
            }
        }
        , xr = function(a, b) {
        for (var c = 0; c < b.length - 2; c += 3) {
            var d = b.charAt(c + 2)
                , d = "a" <= d ? d.charCodeAt(0) - 87 : Number(d)
                , d = "+" == b.charAt(c + 1) ? a >>> d : a << d;
            a = "+" == b.charAt(c) ? a + d & 4294967295 : a ^ d
        }
        return a
    };

    function getToken(text){
        var tk = sM(text);
        tk = tk.replace('&tk=', '');
        //return {name: 'tk', value: tk};
        return tk;
    }
    
    ` + `window = {TKK:` + tkk_code + `};` + `
    
    getToken('` + text_escape + `');`

	// fmt.Println(js_code)
	tk, err := vm.Run(js_code)
	if err != nil {
		panic(err)
	}

	return tk.String()
}

func SetupGoogleTranslateCommand(rootCmd *cobra.Command) {
	var lang string
	cmd := &cobra.Command{
		Use:   "google-trans <text>",
		Short: "google 翻譯",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("required <text>")
			}

			tkk_code := getTKKCode()

			vm := otto.New()

			text := args[0]

			tk := vmGetTk(vm, tkk_code, text)
			// fmt.Println(tk)
			value := trans(vm, tk, text, lang)
			if value.IsObject() {

				obj := value.Object()
				translate, err := obj.Get("translate")
				if err != nil {
					panic(err)
				}

				fmt.Println(translate)
				intro, err := obj.Get("intro")
				if err != nil {
					panic(err)
				}
				fmt.Println(intro)
			} else {
				fmt.Println(value)
			}

			// fmt.Println(obj.Get("intro"), err)
		},
	}

	cmd.Flags().StringVarP(&lang, "lang", "l", "zh-TW", "翻譯後的語言")

	rootCmd.AddCommand(cmd)

}
