package subtitle

import (
	"bytes"
	"io/ioutil"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func extReplace(file, new_ext string) string {
	ext := path.Ext(file)
	return file[0:len(file)-len(ext)] + new_ext
}

// http://angelonotes.blogspot.tw/2014/07/srt-ass-golang.html
func srt2ass(file string) {
	outfile := extReplace(file, ".ass")

	nl := regexp.MustCompile("\\d+ *(\\n|\\r\\n|\\r)\\d{2}:\\d{2}:\\d{2},\\d{3}")
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	str := string(bs)
	newline := nl.FindStringSubmatch(str)[1]
	var buffer bytes.Buffer
	buffer.WriteString("[Script Info]" + newline + "ScriptType: v4.00+" + newline + newline + "[V4+ Styles]" + newline + "Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding" + newline + "Style: Default,WenQuanYi Micro Hei,20,&H00FFFFFF,&HF0000000,&H00B36E1F,&HF0000000,0,0,0,0,100,100,0,0,1,1,1,2,0,0,10,1" + newline + newline + "[Events]" + newline + "Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text" + newline)
	re := regexp.MustCompile("\\d+" + newline + "(\\d{2}:\\d{2}:\\d{2},\\d{3}) --> (\\d{2}:\\d{2}:\\d{2},\\d{3})" + newline + "((.+" + newline + "){0,4})")
	subtitles := re.FindAllStringSubmatch(str, -1)
	for _, v := range subtitles {
		buffer.WriteString("Dialogue: 0," + v[1][:8] + "." + v[1][9:11] + "," + v[2][:8] + "." + v[2][9:11] + ",Default,,0,0,0,," + strings.Replace(strings.TrimSpace(v[3]), newline, "\\N", -1) + newline)
	}

	ioutil.WriteFile(outfile, buffer.Bytes(), 0777)
}

// http://angelonotes.blogspot.tw/2014/07/ass-srt-golang-re2.html
func ass2srt(file string) {
	outfile := extReplace(file, ".srt")

	re := regexp.MustCompile("\\[Events\\] *(\\n|\\r\\n|\\r)Format: .*Text *")
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	str := string(bs)
	reloc := re.FindStringSubmatchIndex(str)
	newline := string(str[reloc[2]:reloc[3]])
	subtitles := strings.Split(str[reloc[1]+len(newline):], newline)
	cols := strings.Count(str[reloc[0]:reloc[1]], ",") + 1
	for i, v := range subtitles {
		if v != "" {
			temp := strings.SplitN(v, ",", cols)
			if temp[cols-1] != "" {
				t1, t2 := strings.Replace(temp[1], ".", ",", -1)+"0", strings.Replace(temp[2], ".", ",", -1)+"0"
				if len(t1) == 11 {
					t1 = "0" + t1
				}
				if len(t2) == 11 {
					t2 = "0" + t2
				}
				subtitles[i] = t1 + " --> " + t2 + newline + strings.Replace(strings.TrimSpace(temp[cols-1]), "\\N", newline, -1) + newline + newline
			} else {
				subtitles[i] = ""
			}
		} else {
			subtitles[i] = ""
		}
	}
	sort.Strings(subtitles)
	j := 0
	var buffer bytes.Buffer
	for _, v := range subtitles {
		if v != "" {
			buffer.WriteString(strconv.Itoa(j+1) + newline + v)
			j++
		}
	}

	ioutil.WriteFile(outfile, buffer.Bytes(), 0777)
}
