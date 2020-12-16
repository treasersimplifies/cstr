package cstr

import (
	"fmt"
	"strings"
)

// BasicColors supported
var BasicColors []string = []string{
	"Black", "Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White",
}

// AllColors supported
var AllColors []string = []string{
	"Black", "Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White",
	"UnderlineRed", "UnderlineGreen", "UnderlineBlue",
	"ShiningYellow", "ShiningSkyblue", "ShiningSkyblue",
	"BlackBeforeBlack", "RedBeforeBlack", "GreenBeforeBlack", "YellowBeforeBlack", "BlueBeforeBlack", "PinkBeforeBlack", "SkyblueBeforeBlack", "WhiteBeforeBlack",
	"BlackBeforeRed", "RedBeforeRed", "GreenBeforeRed", "YellowBeforeRed", "BlueBeforeRed", "PinkBeforeRed", "SkyblueBeforeRed", "WhiteBeforeRed",
	"BlackBeforeGreen", "RedBeforeGreen", "GreenBeforeGreen", "YellowBeforeGreen", "BlueBeforeGreen", "PinkBeforeGreen", "SkyblueBeforeGreen", "WhiteBeforeGreen",
	"BlackBeforeYellow", "RedBeforeYellow", "GreenBeforeYellow", "YellowBeforeYellow", "BlueBeforeYellow", "PinkBeforeYellow", "SkyblueBeforeYellow", "WhiteBeforeYellow",
	"BlackBeforeBlue", "RedBeforeBlue", "GreenBeforeBlue", "YellowBeforeBlue", "BlueBeforeBlue", "PinkBeforeBlue", "SkyblueBeforeBlue", "WhiteBeforeBlue",
	"BlackBeforePink", "RedBeforePink", "GreenBeforePink", "YellowBeforePink", "BlueBeforePink", "PinkBeforePink", "SkyblueBeforePink", "WhiteBeforePink",
	"BlackBeforeSkyblue", "RedBeforeSkyblue", "GreenBeforeSkyblue", "YellowBeforeSkyblue", "BlueBeforeSkyblue", "PinkBeforeSkyblue", "SkyblueBeforeSkyblue", "WhiteBeforeSkyblue",
	"BlackBeforeSkyWhite", "RedBeforeSkyWhite", "GreenBeforeSkyWhite", "YellowBeforeSkyWhite", "BlueBeforeSkyWhite", "PinkBeforeSkyWhite", "SkyblueBeforeSkyWhite", "WhiteBeforeSkyWhite",
}

// Dim dims string with single color
func Dim(str string, style string) string {
	color := strings.Split(style, "Before")
	if len(color) == 1 {

		switch style {
		case "Black":
			return "\033[30m" + str + "\033[0m"
		case "Red":
			return "\033[31m" + str + "\033[0m"
		case "Green":
			return "\033[32m" + str + "\033[0m"
		case "Yellow":
			return "\033[33m" + str + "\033[0m"
		case "Blue":
			return "\033[34m" + str + "\033[0m"
		case "Pink":
			return "\033[35m" + str + "\033[0m"
		case "Skyblue":
			return "\033[36m" + str + "\033[0m"
		case "White":
			return "\033[37m" + str + "\033[0m"

		// TODO: format underline colors
		case "UnderlineRed":
			return "\033[4;31m" + str + "\033[0m"
		case "UnderlineGreen":
			return "\033[4;32m" + str + "\033[0m"
		case "UnderlineBlue":
			return "\033[4;34m" + str + "\033[0m"

		// TODO: format shining colors
		case "ShiningYellow":
			return "\033[5;33m" + str + "\033[0m"
		case "ShiningPink":
			return "\033[5;35m" + str + "\033[0m"
		case "ShiningSkyblue":
			return "\033[5;36m" + str + "\033[0m"

		case "Invisible":
			return "\033[8m" + str + "\033[0m"
		default:
			return str
		}
	}

	dimLeft := ""
	dimRight := "\033[0m"
	switch color[1] {
	case "Black":
		dimLeft = "\033[40;"
	case "Red":
		dimLeft = "\033[41;"
	case "Green":
		dimLeft = "\033[42;"
	case "Yellow":
		dimLeft = "\033[43;"
	case "Blue":
		dimLeft = "\033[44;"
	case "Pink":
		dimLeft = "\033[45;"
	case "Skyblue":
		dimLeft = "\033[46;"
	case "White":
		dimLeft = "\033[47;"
	default:
		return str
	}

	switch color[0] {
	case "Black":
		dimLeft += "30m"
	case "Red":
		dimLeft += "31m"
	case "Green":
		dimLeft += "32m"
	case "Yellow":
		dimLeft += "33m"
	case "Blue":
		dimLeft += "34m"
	case "Pink":
		dimLeft += "35m"
	case "Skyblue":
		dimLeft += "36m"
	case "White":
		dimLeft += "37m"
	default:
		return str
	}

	return dimLeft + str + dimRight
}

// DimColorful .
func DimColorful(str string, style []string) string {
	l := len(style)
	strc := ""
	for i, ch := range str {
		strc += Dim(string(ch), style[i%l])
	}
	return strc
}

// ShowAllColor shows all colors
func ShowAllColor() {
	fmt.Printf("show all colors(currently support %d types): \n", len(AllColors))
	str := ""
	for _, v := range AllColors {
		str += Dim(v+"==>", v)
	}
	fmt.Println(str)
}
