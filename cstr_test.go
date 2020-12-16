package cstr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDye(t *testing.T) {
	_, err := fmt.Println(
		Dye("This is a Black string...", "Black"),
		Dye("This is a Red string.....", "Red"),
		Dye("This is a Green string...", "Green"),
		Dye("This is a Yellow string..", "Yellow"),
	)
	assert.Nil(t, err)
	_, err = fmt.Println(
		Dye("This is a Blue string....", "Blue"),
		Dye("This is a Pink string....", "Pink"),
		Dye("This is a Skyblue string.", "Skyblue"),
		Dye("This is a White string...", "White"),
	)
	fmt.Println()
	ShowAllColor()
	fmt.Println()
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("THIS IS A COLORFUL STRING...........\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("这是一个彩色的字符串................\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("這是一個彩色的字串..................\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("これは色付きの文字列です............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("이것은 컬러 문자열입니다............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("C’est une corde colorée.............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("Il s’agit d’une chaîne de couleur...\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("Dies ist eine farbige Zeichenfolge..\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("Questa è una stringa colorata.......\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("Dit is een gekleurde tekenreeks.....\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DyeColorfully("Esta es una cadena de color.........\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)

}
