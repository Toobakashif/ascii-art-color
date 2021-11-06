package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {

	// Reading  Arguments from cmd
	files := os.Args

	var indexArray []int

	//
	if len(files) >= 4 {
		match, _ := regexp.MatchString(`^(\-\-COLOR\=(black|red|green|yellow|blue|purple|cyan|white|lightgray|darkgray|lightred|lightgreen|lightyellow|lightblue|lightmagent|lightcyan|\#000000|\#FF0000|\#00FF00|\#FFFF00|\#0000FF|#800080|\#00FFFF|\#FFFFFF|\#D3D3D3\#A9A9A9|\#FF7F7F|\#90EE90|\#FFFFE0|\#ADD8E6|\#E78BE7|\#004C54|rgb\(0,0,0\)|rgb\(255,0,0\)|rgb\(0,255,0\)|rgb\(255,255,0\)|rgb\(0,0,255\)|rgb\(128,0,128\)|rgb\(0,255,255\)|rgb\(255,255,255\)|rgb\(211,211,211\)|rgb\(169,169,169\)|rgb\(255,127,127\)|rgb\(144,238,144\)|rgb\(255,255,224\)|rgb\(173,216,230\)|rgb\(231,139,231\)|rgb\(224,255,255\)))`, files[3])
		if !match && len(files) == 4 {
			errorHandler("errorcolor")

		}

		// here color
		if len(files) == 5 {

			matchP, _ := regexp.MatchString(`^(\-\-POSITION\=((\d+)|(\d\:\d+))|(\-\-LETTER\=\w)|(\-\-WORD\=\w+)|(\-\-RANDOM))$`, files[4])
			matchPos, _ := regexp.MatchString(`^\-\-POSITION\=((\d+)|(\d\:\d+))$`, files[4])

			match2Letter, _ := regexp.MatchString(`^\-\-LETTER\=\w$`, files[4])
			match2Word, _ := regexp.MatchString(`^\-\-WORD\=\w+$`, files[4])
			match2Random, _ := regexp.MatchString(`^\-\-RANDOM$`, files[4])

			if !(matchP && match) {
				errorHandler("errorpositionandcolor")
			}
			//here we are splite values with  "="
			lineArray := strings.Split(files[4], "=")

			if matchPos {
				if strings.Contains(lineArray[1], ":") {
					nArray := strings.Split(string(lineArray[1]), ":")
					aNumber, _ := strconv.Atoi(nArray[0])
					bNumber, _ := strconv.Atoi(nArray[1])
					if aNumber > bNumber {
						errorHandler("errorbiggernumber")
					}
					indexArray = []int{aNumber - 1, bNumber}

				} else {
					aNumber, _ := strconv.Atoi(lineArray[1])
					indexArray = []int{aNumber - 1, aNumber}
				}

			} else if match2Letter {
				letter := lineArray[1]
				re := regexp.MustCompile(letter)
				preindexArray := re.FindAllStringIndex(string(files[1]), -1)
				indexArray = colorArray(preindexArray)

			} else if match2Word {
				word := lineArray[1]
				re := regexp.MustCompile(word)
				preindexArray := re.FindAllStringIndex(string(files[1]), -1)
				indexArray = colorArray(preindexArray)

			} else if match2Random {
				rand.Seed(time.Now().Unix())
				number1 := rand.Intn(len(files[1]))
				number2 := rand.Intn(len(files[1]))
				indexArray = []int{number1, number2}
				sort.Ints(indexArray)

			}

		} else {
			aNumber := 0
			bNumber := len(files[1])
			indexArray = []int{aNumber, bNumber}
		}
		color := strings.Split(files[3], "=")

		// Color dictionary
		colorDict := map[string]string{
			"black":            "\033[30m",
			"red":              "\033[31m",
			"green":            "\033[32m",
			"yellow":           "\033[33m",
			"blue":             "\033[0;34m",
			"purple":           "\033[35m",
			"cyan":             "\033[36m",
			"white":            "\033[97m",
			"lightgray":        "\033[37m",
			"darkgray":         "\033[90m",
			"lightred":         "\033[1;31m",
			"lightgreen":       "\033[1;32m",
			"lightyellow":      "\033[93m",
			"lightblue":        "\033[1;36m",
			"lightmagenta":     "\033[95m",
			"lightcyan":        "\033[96m",
			"#000000":          "\033[30m",
			"#FF0000":          "\033[31m",
			"#00FF00":          "\033[32m",
			"#FFFF00":          "\033[33m",
			"#0000FF":          "\033[34m",
			"#800080":          "\033[35m",
			"#00FFFF":          "\033[36m",
			"#FFFFFF":          "\033[97m",
			"#D3D3D3":          "\033[37m",
			"#A9A9A9":          "\033[90m",
			"#FF7F7F":          "\033[1;31m",
			"#90EE90":          "\033[1;32m",
			"#FFFFE0":          "\033[93m",
			"#ADD8E6":          "\033[1;36m",
			"#E78BE7":          "\033[95m",
			"#E0FFFF":          "\033[96m",
			"rgb(0,0,0)":       "\033[30m",
			"rgb(255,0,0)":     "\033[31m",
			"rgb(0,255,0)":     "\033[32m",
			"rgb(255,255,0)":   "\033[33m",
			"rgb(0,0,255)":     "\033[34m",
			"rgb(128,0,128)":   "\033[35m",
			"rgb(0,255,255)":   "\033[36m",
			"rgb(255,255,255)": "\033[97m",
			"rgb(211,211,211)": "\033[37m",
			"rgb(169,169,169)": "\033[90m",
			"rgb(255,127,127)": "\033[1;31m",
			"rgb(144,238,144)": "\033[1;32m",
			"rgb(255,255,224)": "\033[93m",
			"rgb(173,216,230)": "\033[1;36m",
			"rgb(231,139,231)": "\033[95m",
			"rgb(224,255,255)": "\033[96m",
			"colorreset":       "\033[0m",
		}

		line, err := ioutil.ReadFile(files[2] + ".txt")
		if err != nil {
			panic(err)
		}
		//It will split readed lines
		label := strings.Split(string(line), "\n")
		word := strings.Split(string(files[1]), "\\n")
		for i := 0; i < len(word); i++ {
			counter := 0
			if word[i] == "" {
				break
			}
			if i > 0 {
				counter = len(word[i-1])
			}
			//Print out Ascii lines
			for j := 1; j < 9; j++ {

				for k := 0; k < len(word[i]); k++ {
					if control(indexArray, k, counter) {
						fmt.Print(colorDict[string(color[1])], label[(int(word[i][k])-32)*9+j])

					} else {
						fmt.Print(colorDict["colorreset"], label[(int(word[i][k])-32)*9+j])
					}

				}
				fmt.Println()
			}

		}

	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	}
}

func colorArray(array [][]int) []int {
	var mergeArray []int
	for _, value := range array {
		mergeArray = append(mergeArray, value...)
	}
	return mergeArray
}
func control(numbers []int, index int, counter int) bool {
	length := len(numbers)
	control := false
	for v := 0; v < length; v += 2 {
		if (numbers[v]) <= index+counter && index+counter < numbers[v+1] {
			control = true
			break
		}
	}
	return control
}

func errorHandler(errorNo string) {
	switch errorNo {
	case "errorcolor":
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard  --COLOR=<color>\n\nCOLOR=red (green, yellow, blue, purple, cyan, white, lightgray, darkgray, lightred, lightgreen, lightyellow, lightblue, lightmagent, lightcyan)")

		os.Exit(0)
	case "errorpositionandcolor":
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION][OPTION2]\n\nEX: go run . something standard  --COLOR=<color> --POSITION=<begin:end>\n\nCOLOR=red (green, yellow, blue, purple, cyan, white, lightgray, darkgray, lightred, lightgreen, lightyellow, lightblue, lightmagent, lightcyan)\n\nOPTION2 --POSITION=<begin:end> or --WORD=<number> or --LETTER=<letter> or --RANDOM")
		os.Exit(0)
	case "errorbiggernumber":
		fmt.Println("Wrong!: Begin number is bigger than end. It's impossible to use! Usage: OPTION2 --POSITION=<begin:end> begin")
		os.Exit(0)

	}
}
