package web

import (
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geckoboard/cli-table"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type proxyData struct {
	IpPort   string
	Type     string
	Speed    string
	Category string
	Country  string
	City     string
}

type proxyDatas []*proxyData

func (s proxyDatas) Len() int      { return len(s) }
func (s proxyDatas) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type BySpeed struct{ proxyDatas }

func (s BySpeed) Less(i, j int) bool {
	a, err := strconv.ParseFloat(s.proxyDatas[i].Speed, 64)
	if err != nil {
		a = 0
	}

	b, err := strconv.ParseFloat(s.proxyDatas[j].Speed, 64)
	if err != nil {
		b = 0
	}

	return a > b
}

type TypeDatas []string

func (s TypeDatas) Conv(v string) string {
	if v == "HTTP" {
		return "no"
	}

	if v == "HTTPS" {
		return "yes"
	}

	return "any"
}

var VALID_TYPE = TypeDatas{
	"HTTP", "HTTPS",
}

type CountryDatas map[string]string

func (s CountryDatas) IsValid(v string) bool {
	_, exists := s[v]
	return exists
}

func (s CountryDatas) List() {
	t := table.New(2)
	t.SetHeader(0, "縮寫", table.AlignLeft)
	t.SetHeader(1, "全名", table.AlignLeft)
	for k, v := range s {
		t.Append([]string{k, v})
	}
	t.Write(os.Stdout, table.StripAnsi)
}

var VALID_COUNTRY = CountryDatas{
	"usa-and-canada": "USA and Canada",
	"western-europe": "Western Europe",
	"estern-europe":  "Estern Europe",
	"arab-world":     "Arab World",
	"western-asia":   "Western Asia",
	"estern-asia":    "Eastern Asia",
	"AF":             "Afghanistan",
	"AL":             "Albania",
	"DZ":             "Algeria",
	"AO":             "Angola",
	"AR":             "Argentina",
	"AU":             "Australia",
	"AT":             "Austria",
	"AZ":             "Azerbaijan",
	"BD":             "Bangladesh",
	"BY":             "Belarus",
	"BE":             "Belgium",
	"BO":             "Bolivia",
	"BA":             "Bosnia",
	"BW":             "Botswana",
	"BR":             "Brazil",
	"BG":             "Bulgaria",
	"KH":             "Cambodia",
	"CM":             "Cameroon",
	"CA":             "Canada",
	"KY":             "Cayman&nbsp;Islands",
	"CL":             "Chile",
	"CN":             "China",
	"CO":             "Colombia",
	"CG":             "Congo",
	"CR":             "Costa Rica",
	"CI":             "Cote D'Ivoire",
	"HR":             "Croatia",
	"CZ":             "Czech Republic",
	"DK":             "Denmark",
	"DJ":             "Djibouti",
	"EC":             "Ecuador",
	"EG":             "Egypt",
	"EU":             "Europe",
	"FR":             "France",
	"GA":             "Gabon",
	"GM":             "Gambia",
	"GE":             "Georgia",
	"DE":             "Germany",
	"GH":             "Ghana",
	"GR":             "Greece",
	"GT":             "Guatemala",
	"GN":             "Guinea",
	"HN":             "Honduras",
	"HK":             "Hong Kong",
	"HU":             "Hungary",
	"IN":             "India",
	"ID":             "Indonesia",
	"IQ":             "Iraq",
	"IE":             "Ireland",
	"IL":             "Israel3",
	"IT":             "Italy",
	"JM":             "Jamaica",
	"JP":             "Japan",
	"KZ":             "Kazakhstan",
	"KE":             "Kenya",
	"KG":             "Kyrgyzstan",
	"LV":             "Latvia",
	"LB":             "Lebanon",
	"LS":             "Lesotho",
	"LT":             "Lithuania",
	"MO":             "Macau",
	"MK":             "Macedonia",
	"MW":             "Malawi",
	"MY":             "Malaysia",
	"ML":             "Mali",
	"MX":             "Mexico",
	"MN":             "Mongolia",
	"ME":             "Montenegro",
	"MZ":             "Mozambique",
	"NP":             "Nepal",
	"NL":             "Netherlands",
	"NZ":             "New Zealand",
	"NG":             "Nigeria",
	"NO":             "Norway",
	"PK":             "Pakistan",
	"PS":             "Palestinian&nbsp;Territory",
	"PA":             "Panama",
	"PY":             "Paraguay",
	"PE":             "Peru",
	"PH":             "Philippines",
	"PL":             "Poland",
	"PT":             "Portugal",
	"PR":             "Puerto Rico",
	"QA":             "Qatar",
	"RO":             "Romania",
	"RU":             "Russia",
	"RW":             "Rwanda",
	"WS":             "Samoa",
	"A2":             "Satellite Provider",
	"SA":             "Saudi Arabia",
	"RS":             "Serbia",
	"SC":             "Seychelles",
	"SG":             "Singapore",
	"SX":             "Sint&nbsp;Maarten&nbsp;(Dutch&nbsp;part)",
	"SK":             "Slovakia",
	"ZA":             "South Africa",
	"ES":             "Spain",
	"SE":             "Sweden",
	"CH":             "Switzerland",
	"TW":             "Taiwan",
	"TJ":             "Tajikistan",
	"TH":             "Thailand",
	"TL":             "Timor-Leste",
	"TR":             "Turkey",
	"UG":             "Uganda",
	"UA":             "Ukrain",
	"AE":             "United Arab Emirates",
	"GB":             "United Kingdom",
	"US":             "United States",
	"UC":             "Unknown&nbsp;Country",
	"UY":             "Uruguay",
	"VE":             "Venezuela",
	"VN":             "Vietnam",
	"YE":             "Yemen",
	"ZM":             "Zambia",
	"ZW":             "Zimbabwe",
}

func SetupProxyCommand(rootCmd *cobra.Command) {
	var filter_country string
	var filter_ssl string
	var list bool

	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Proxy 列表",
		Run: func(cmd *cobra.Command, args []string) {
			if list {
				VALID_COUNTRY.List()
				return
			}
			request := gorequest.New()
			is_valid := VALID_COUNTRY.IsValid(filter_country)

			ssl := VALID_TYPE.Conv(filter_ssl)

			search := ""
			country := "any"
			if is_valid {
				search = filter_country
				country = filter_country
			}

			url := fmt.Sprintf("http://proxy-list.org/english/search.php?search=%s&country=%s&type=any&port=any&ssl=%s&p=1", search, country, ssl)
			resp, _, err := request.Get(fmt.Sprintf(url)).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			doc, err2 := goquery.NewDocumentFromResponse(resp)
			if err2 != nil {
				panic(err2)
			}
			result := []*proxyData{}
			re_proxy := regexp.MustCompile("Proxy\\('(.*)'\\)")
			re_speed := regexp.MustCompile("(.*)kbit")

			doc.Find("#proxy-table ul:not(.header-row)").Each(func(i int, s *goquery.Selection) {

				ret := re_proxy.FindStringSubmatch(s.Find("li.proxy").Text())
				if len(ret) < 2 {
					return
				}

				proxy_data, err := base64.StdEncoding.DecodeString(ret[1])
				if err != nil {
					return
				}

				speed := s.Find("li.speed").Text()
				if speed == "-" {
					return
				}

				// fmt.Println(s.Find("li.proxy").Text())
				// fmt.Println(s.Find("li.https").Text())
				// fmt.Println(speed)
				// fmt.Println(s.Find("li.type").Text())

				speed_data := re_speed.FindStringSubmatch(speed)
				// fmt.Println(speed_data)
				// fmt.Println(len(speed_data))
				if len(speed_data) < 2 {
					return
				}
				result = append(result, &proxyData{
					IpPort:   string(proxy_data),
					Type:     s.Find("li.type").Text(),
					Category: s.Find("li.https").Text(),
					City:     s.Find("li.country-city span.city").Text(),
					Speed:    speed_data[1],
					Country:  s.Find("li.country-city span.country").AttrOr("title", ""),
				})
			})
			// fmt.Println(result)
			sort.Sort(BySpeed{result})

			t := table.New(4)
			t.SetHeader(0, "Proxy", table.AlignLeft)
			t.SetHeader(1, "kbit", table.AlignRight)
			t.SetHeader(2, "Type", table.AlignLeft)
			t.SetHeader(3, "Country/City", table.AlignLeft)

			for _, item := range result {
				t.Append([]string{item.IpPort, item.Speed, item.Category + "," + item.Type, item.Country + "/" + item.City})
			}
			t.Write(os.Stdout, table.StripAnsi)
		},
	}
	cmd.Flags().StringVarP(&filter_country, "country", "c", "", "過濾指定國家")
	cmd.Flags().StringVarP(&filter_ssl, "ssl", "s", "", "ANY, HTTP or HTTPS")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "列出過濾國家")
	rootCmd.AddCommand(cmd)

}
