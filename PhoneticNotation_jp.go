package main

import "fmt"
import "io/ioutil"
import "strings"
import "unicode/utf8"
import "regexp"
import "os"

func main() {
	f, err := ioutil.ReadFile("./japanese.txt")
	m := map[string]string{}
	if err == nil {
		allString := string(f)
		restpart := allString
		for strings.Index(restpart, "(") != -1 {
			word1_end_idx := strings.Index(restpart, "(")
			beginpart := restpart[0:word1_end_idx]
			restpart = restpart[word1_end_idx+1:]
			word1_begin_idx := strings.LastIndex(beginpart, " ")
			key := beginpart[word1_begin_idx+1 : word1_end_idx]
			roma1_end_idx := strings.Index(restpart, ")")
			val := restpart[:roma1_end_idx]
			m[key] = val
		}
	}

	f, err = ioutil.ReadFile("./source.txt")
	if err == nil {
		allString := string(f)
		restpart := strings.TrimSpace(allString)
		var output, announce string
		for len(restpart) > 0 {
			_, length1 := utf8.DecodeRuneInString(restpart)
			_, length2 := utf8.DecodeRuneInString(restpart[length1:])
			length2 += length1
			fmt.Println("length1:", length1)
			fmt.Println("length2:", length2)
			if m[restpart[:length2]] != "" {
				output = strings.Join([]string{output, restpart[:length2], m[restpart[:length2]]}, "")
				announce = strings.Join([]string{announce, m[restpart[:length2]]}, " ")
				restpart = restpart[length2:]
			} else if m[restpart[:length1]] != "" {
				output = strings.Join([]string{output, restpart[:length1], m[restpart[:length1]]}, "")
				announce = strings.Join([]string{announce, m[restpart[:length1]]}, " ")
				restpart = restpart[length1:]
			} else if restpart[:length1] == "\n" {
				//fmt.Println(1111)
				announce = strings.Join([]string{announce, "\r\n"}, "")
				restpart = restpart[length1:]
			} else if restpart[:length1] == " " {
				announce = strings.Join([]string{announce, " "}, "")
				output = output + " "
				restpart = restpart[length1:]
			} else {
				output = strings.Join([]string{output, restpart[:length1]}, "")
				engletterreg := regexp.MustCompile("[a-zA-Z0-9]")
				announce = strings.Join([]string{announce, engletterreg.FindString(restpart[:length1])}, "")
				restpart = restpart[length1:]
			}
		}
		// for key, val := range m {
		// 	fmt.Println(1, key, 2, val)
		// }
		fmt.Println(m)
		output = strings.Join([]string{output, announce}, "\r\n")
		ioutil.WriteFile("./output.txt", []byte(output), os.ModeAppend)
	}
}
