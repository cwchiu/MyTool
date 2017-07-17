package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strconv"
	"time"
)

const AREA = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const SEX = "12"
const NUM = "0123456789"

func SetupTaiwanPidCommand(rootCmd *cobra.Command) {
	var area_code string
	var sex string

	POS_WEIGHT := []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	AREA_WEIGHT := map[string]string{
		"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15", "G": "16", "H": "17", "I": "34", "J": "18", "K": "19", "L": "20", "M": "21", "N": "22", "O": "35", "P": "23", "Q": "24", "R": "25", "S": "26", "T": "27", "U": "28", "V": "29", "W": "32", "X": "30", "Y": "31", "Z": "33",
	}
	cmd := &cobra.Command{
		Use:   "taiwan-pid",
		Short: "產生符合驗證規則的身分證字號",
		Run: func(cmd *cobra.Command, args []string) {
			rand.Seed(time.Now().UnixNano())
			sn := ""
			weight, exists := AREA_WEIGHT[area_code]
			if exists {
				sn += weight
			} else {
				area_code = string(AREA[rand.Intn(len(AREA))])
				sn += AREA_WEIGHT[area_code]
			}

			if sex == "1" || sex == "2" {
				sn += sex
			} else {
				sn += string(SEX[rand.Intn(len(SEX))])
			}
			fmt.Println(rand.Intn(len(NUM)))
			for i := 0; i < 7; i++ {
				sn = sn + string(NUM[rand.Intn(len(NUM))])
			}

			sum := 0
			for i, ch := range sn {
				v1, _ := strconv.Atoi(string(ch))
				sum = sum + POS_WEIGHT[i]*v1

			}
			mod := (10 - sum%10) % 10
			fmt.Printf("%s%s%v\n", area_code, sn[2:], mod)
		},
	}
	cmd.Flags().StringVarP(&area_code, "area", "a", "", "區域代碼")
	cmd.Flags().StringVarP(&sex, "sex", "s", "", "性別: 1(男), 2(女)")
	rootCmd.AddCommand(cmd)
}
