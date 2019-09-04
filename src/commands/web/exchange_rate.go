package web

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"strings"
)

func SetupExchangeRateCommand(rootCmd *cobra.Command) {
	var source string
	var target string
	var list bool

	dictCurrency := map[string]string{
		"AED": "United Arab Emirates Dirham (AED)",
		"ANG": "Netherlands Antillean Guilder (ANG)",
		"ARS": "Argentine Peso (ARS)",
		"AUD": "Australian Dollar (AUD)",
		"BDT": "Bangladeshi Taka (BDT)",
		"BGN": "Bulgarian Lev (BGN)",
		"BHD": "Bahraini Dinar (BHD)",
		"BND": "Brunei Dollar (BND)",
		"BOB": "Bolivian Boliviano (BOB)",
		"BRL": "Brazilian Real (BRL)",
		"BWP": "Botswanan Pula (BWP)",
		"CAD": "Canadian Dollar (CAD)",
		"CHF": "Swiss Franc (CHF)",
		"CLP": "Chilean Peso (CLP)",
		"CNY": "Chinese Yuan (CNY)",
		"COP": "Colombian Peso (COP)",
		"CRC": "Costa Rican Colón (CRC)",
		"CZK": "Czech Republic Koruna (CZK)",
		"DKK": "Danish Krone (DKK)",
		"DOP": "Dominican Peso (DOP)",
		"DZD": "Algerian Dinar (DZD)",
		"EEK": "Estonian Kroon (EEK)",
		"EGP": "Egyptian Pound (EGP)",
		"EUR": "Euro (EUR)",
		"FJD": "Fijian Dollar (FJD)",
		"GBP": "British Pound Sterling (GBP)",
		"HKD": "Hong Kong Dollar (HKD)",
		"HNL": "Honduran Lempira (HNL)",
		"HRK": "Croatian Kuna (HRK)",
		"HUF": "Hungarian Forint (HUF)",
		"IDR": "Indonesian Rupiah (IDR)",
		"ILS": "Israeli New Sheqel (ILS)",
		"INR": "Indian Rupee (INR)",
		"JMD": "Jamaican Dollar (JMD)",
		"JOD": "Jordanian Dinar (JOD)",
		"JPY": "Japanese Yen (JPY)",
		"KES": "Kenyan Shilling (KES)",
		"KRW": "South Korean Won (KRW)",
		"KWD": "Kuwaiti Dinar (KWD)",
		"KYD": "Cayman Islands Dollar (KYD)",
		"KZT": "Kazakhstani Tenge (KZT)",
		"LBP": "Lebanese Pound (LBP)",
		"LKR": "Sri Lankan Rupee (LKR)",
		"LTL": "Lithuanian Litas (LTL)",
		"LVL": "Latvian Lats (LVL)",
		"MAD": "Moroccan Dirham (MAD)",
		"MDL": "Moldovan Leu (MDL)",
		"MKD": "Macedonian Denar (MKD)",
		"MUR": "Mauritian Rupee (MUR)",
		"MVR": "Maldivian Rufiyaa (MVR)",
		"MXN": "Mexican Peso (MXN)",
		"MYR": "Malaysian Ringgit (MYR)",
		"NAD": "Namibian Dollar (NAD)",
		"NGN": "Nigerian Naira (NGN)",
		"NIO": "Nicaraguan Córdoba (NIO)",
		"NOK": "Norwegian Krone (NOK)",
		"NPR": "Nepalese Rupee (NPR)",
		"NZD": "New Zealand Dollar (NZD)",
		"OMR": "Omani Rial (OMR)",
		"PEN": "Peruvian Nuevo Sol (PEN)",
		"PGK": "Papua New Guinean Kina (PGK)",
		"PHP": "Philippine Peso (PHP)",
		"PKR": "Pakistani Rupee (PKR)",
		"PLN": "Polish Zloty (PLN)",
		"PYG": "Paraguayan Guarani (PYG)",
		"QAR": "Qatari Rial (QAR)",
		"RON": "Romanian Leu (RON)",
		"RSD": "Serbian Dinar (RSD)",
		"RUB": "Russian Ruble (RUB)",
		"SAR": "Saudi Riyal (SAR)",
		"SCR": "Seychellois Rupee (SCR)",
		"SEK": "Swedish Krona (SEK)",
		"SGD": "Singapore Dollar (SGD)",
		"SKK": "Slovak Koruna (SKK)",
		"SLL": "Sierra Leonean Leone (SLL)",
		"SVC": "Salvadoran Colón (SVC)",
		"THB": "Thai Baht (THB)",
		"TND": "Tunisian Dinar (TND)",
		"TRY": "Turkish Lira (TRY)",
		"TTD": "Trinidad and Tobago Dollar (TTD)",
		"TWD": "New Taiwan Dollar (TWD)",
		"TZS": "Tanzanian Shilling (TZS)",
		"UAH": "Ukrainian Hryvnia (UAH)",
		"UGX": "Ugandan Shilling (UGX)",
		"USD": "US Dollar (USD)",
		"UYU": "Uruguayan Peso (UYU)",
		"UZS": "Uzbekistan Som (UZS)",
		"VEF": "Venezuelan Bolívar (VEF)",
		"VND": "Vietnamese Dong (VND)",
		"XOF": "CFA Franc BCEAO (XOF)",
		"YER": "Yemeni Rial (YER)",
		"ZAR": "South African Rand (ZAR)",
		"ZMK": "Zambian Kwacha (ZMK)",
	}

	cmd := &cobra.Command{
		Use:   "exchange-rate -s USD -t TWD",
		Short: "目前匯率",
		Long:  `目前匯率`,
		Run: func(cmd *cobra.Command, args []string) {
			if list {
				for k, v := range dictCurrency {
					fmt.Printf("%s : %s\n", k, v)
				}

				return
			}
			source_name, source_exists := dictCurrency[source]
			if !source_exists {
				fmt.Printf("%s not support\n", source)
				return
			}

			target_name, target_exists := dictCurrency[target]
			if !target_exists {
				fmt.Printf("%s not support\n", target)
				return
			}

			value := fmt.Sprintf("%s%s=X", source, target)

			request := gorequest.New()
			_, body, err := request.Get("http://finance.yahoo.com/d/quotes.csv").Param("e", ".csv").Param("f", "sl1d1t1").Param("s", value).Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)").SetDebug(false).End()
			if err != nil {
				panic(err)
			}
			// fmt.Println(body)
			parts := strings.Split(body, ",")
			fmt.Printf("%s => %s\n1 => %v\n%s %s\n", source_name, target_name, parts[1], strings.Replace(parts[2], `"`, "", -1), strings.Replace(parts[3], `"`, "", -1))
		},
	}
	cmd.Flags().StringVarP(&source, "source", "s", "USD", "來源貨幣")
	cmd.Flags().StringVarP(&target, "target", "t", "TWD", "目的貨幣")
	cmd.Flags().BoolVarP(&list, "list", "l", false, "列出支援貨幣")
	rootCmd.AddCommand(cmd)

}
