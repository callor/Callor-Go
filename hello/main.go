package main

import "fmt"

/*
Go 기본타입:

정수형들 (int, int8, int16, int32, rune, int64, uint, uintptr, uint8, uint16, uint64)
intergers를 나타내는 다양한 유형이 있으며
대부분 int를 사용하며 최적화를 위해 보다 전문화된 유형을 선택할 수 있다.

int 유형은 64비트 시스템에서 사용될 때 기본적으로 64비트, 32비트 시스템에서 32비트 등으로 사용된다.

uint는 부호 없는 정수이며,
숫자가 음수가 아닐 것이라는 것을 알고 있다면 저장할 수 있는 값의 양을 두 배로 늘리는 데 사용할 수 있다.

실수형들 (float32, float64), useful to represent decimals
Complex types (complex64, complex128), useful in math
Byte (byte), represents a single ASCII character

문자열 (string), a set of byte
Go의 문자열은 바이트 값의 시퀀스이다
문자열은 큰따옴표("")를 사용한다
문자열의 길이를 얻으려면 내장 len() 함수를 사용한다.

대괄호를 사용하여 개별 문자에 액세스하고 얻으려는 문자의 인덱스를 전달할 수 있습니다.
name := "홍길동"
name[0] //"홍" (indexes start at 0)
name[1] //"길"


Booleans (bool), true 또는 false

*/
func main() {
	fmt.Println("대한민국 만세")

	// age = 10, name = "Roger"
	var age1, name1 = 10, "Roger"
	fmt.Println(age1, name1)

	// var 키워드 없이 변수를 선언할때는 := 를 사용
	age2, name2 := 10, "Roger"
	fmt.Println(age2, name2)

	var nation = "Korea"
	fmt.Println(nation[0])
	fmt.Println(nation[1])

}
