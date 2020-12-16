package cstr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDim(t *testing.T) {
	_, err := fmt.Println(
		Dim("This is a Black string...", "Black"),
		Dim("This is a Red string.....", "Red"),
		Dim("This is a Green string...", "Green"),
		Dim("This is a Yellow string..", "Yellow"),
	)
	assert.Nil(t, err)
	_, err = fmt.Println(
		Dim("This is a Blue string....", "Blue"),
		Dim("This is a Pink string....", "Pink"),
		Dim("This is a Skyblue string.", "Skyblue"),
		Dim("This is a White string...", "White"),
	)
	fmt.Println()
	ShowAllColor()
	fmt.Println()
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("THIS IS A COLORFUL STRING...........\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("这是一个彩色的字符串................\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("這是一個彩色的字串..................\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("これは色付きの文字列です............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("이것은 컬러 문자열입니다............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("C’est une corde colorée.............\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("Il s’agit d’une chaîne de couleur...\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("Dies ist eine farbige Zeichenfolge..\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("Questa è una stringa colorata.......\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("Dit is een gekleurde tekenreeks.....\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)
	_, err = fmt.Print(DimColorful("Esta es una cadena de color.........\n", []string{"Red", "Green", "Yellow", "Blue", "Pink", "Skyblue", "White"}))
	assert.Nil(t, err)

}
