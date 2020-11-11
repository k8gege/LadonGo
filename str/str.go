package str
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
import (
	"strings"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func GetBetween(str, starting, ending string) string {

s := strings.Index(str, starting)

if s < 0 {

return ""

}

s += len(starting)

e := strings.Index(str[s:], ending)

if e < 0 {

return ""

}

return str[s : s+e]

}

func IsGBK(data []byte) bool {
    length := len(data)
    var i int = 0
    for i < length {
        //fmt.Printf("for %x\n", data[i])
        if data[i] <= 0xff {
            //编码小于等于127,只有一个字节的编码，兼容ASCII吗
            i++
            continue
        } else {
            //大于127的使用双字节编码
            if  data[i] >= 0x81 &&
                data[i] <= 0xfe &&
                data[i + 1] >= 0x40 &&
                data[i + 1] <= 0xfe &&
                data[i + 1] != 0xf7 {
                i += 2
                continue
            } else {
                return false
            }
        }
    }
    return true
}

func PreNUm(data byte) int {
    str := fmt.Sprintf("%b", data)
    var i int = 0
    for i < len(str) {
        if str[i] != '1' {
            break
        }
        i++
    }
    return i
}
func IsUtf8(data []byte) bool {
    for i := 0; i < len(data);  {
        if data[i] & 0x80 == 0x00 {
            // 0XXX_XXXX
            i++
            continue
        } else if num := PreNUm(data[i]); num > 2 {
            // 110X_XXXX 10XX_XXXX
            // 1110_XXXX 10XX_XXXX 10XX_XXXX
            // 1111_0XXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
            // 1111_10XX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
            // 1111_110X 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX 10XX_XXXX
            // PreNUm() 返回首个字节的8个bits中首个0bit前面1bit的个数，该数量也是该字符所使用的字节数
            i++
            for j := 0; j < num - 1; j++ {
                //判断后面的 num - 1 个字节是不是都是10开头
                if data[i] & 0xc0 != 0x80 {
                    return false
                }
                i++
            }
        } else  {
            //其他情况说明不是utf-8
            return false
        }
    }
    return true
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}