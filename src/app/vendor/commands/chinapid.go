package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"time"
)

/*
18位的身份證號碼： 510104196307170239
（1）1~6位為地區代碼，其中1、2位數為各省級政府的代碼（四川省為51），3、4位數為地、市級政府的代碼（成都市為01），5、6位數為縣、區級政府代碼（錦江區為04）。

（2）7~10位為出生年份(4位)，如1963

（3）11~12位為出生月份，如07

（4）13~14位為出生日期，如17

（5）第15~17位為順序號，為縣、區級政府所轄派出所的分配碼，每個派出所分配碼位10個連續號碼，例如「020—029」，其中單數為男性分配碼，雙數為女性分配碼，如遇同年同月同日有兩人以上時順延第二、第三、第四、第五個分配碼。

（4）18位為效驗位（識別碼），通過複雜公式算出，普遍採用計算機自動生成。

15位的身份證號碼：

（1）1~6位為地區代碼

（2）7~8位為出生年份(2位)，9~10位為出生月份，11~12位為出生日期

（3）第13~15位為順序號，並能夠判斷性別，奇數為男，偶數為女。
*/
// http://www.welefen.com/lab/identify/?
// https://shenfenzheng.51240.com/110101199710121308__shenfenzheng/
// 11010119990101293x
func SetupChinaPidCommand(rootCmd *cobra.Command) {
	var year string
	const NUM = "0123456789"
	POS_WEIGHT := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	CRC_CODE := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	AREA := []string{
		"120101", // 天津 市轄區 和平區
		"110101", // 北京
		"130101", //河北 石家莊市 市轄區
		"140211", //山西 大同市 南郊區
	}
	cmd := &cobra.Command{
		Use:   "china-pid",
		Short: "產生符合驗證規則的身分證字號",
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
			area_code := AREA[rand.Intn(len(AREA))]
			dt := time.Unix(rand.Int63n(time.Now().Unix()-31536000)+2592000, 0)

			birthday := dt.Format("20060102")
			if year != "" && len(year) == 4 {
				birthday = year + birthday[4:]
			}

			serial := ""
			for i := 0; i < 3; i++ {
				serial += string(NUM[rand.Intn(len(NUM))])
			}
			sn := area_code + birthday + serial

			sum := 0
			for i, ch := range sn {
				v1, _ := strconv.Atoi(string(ch))
				sum = sum + POS_WEIGHT[i]*v1

			}
			fmt.Printf("%s%v\n", sn, CRC_CODE[sum%11])
		},
	}
	cmd.Flags().StringVarP(&year, "year", "y", "", "年份, ex:1900")
	rootCmd.AddCommand(cmd)
}
