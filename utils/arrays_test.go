package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MarshinalTest struct {
	ID   int
	Name string
}

func TestWrapArray(t *testing.T) {
	arrObj := []MarshinalTest{
		MarshinalTest{ID: 1, Name: "User1"},
		MarshinalTest{ID: 2, Name: "User2"},
	}
	responseJSONObj := WrapArray(arrObj)
	fmt.Println(responseJSONObj)
	assert.Equal(t, responseJSONObj["data"] != nil, true)
}

func TestWrapArrayWithElemName(t *testing.T) {
	arrObj := []MarshinalTest{
		MarshinalTest{ID: 1, Name: "User1"},
		MarshinalTest{ID: 2, Name: "User2"},
	}
	responseJSONObjCase1 := WrapArrayWithElemName(arrObj, "myTagName")
	assert.Equal(t, responseJSONObjCase1["myTagName"] != nil, true)
	assert.Equal(t, responseJSONObjCase1["wrongName"] != nil, false)

	responseJSONObjCase2 := WrapArrayWithElemName(nil, "myTagName")
	assert.Equal(t, responseJSONObjCase2 != nil, true)
	assert.Equal(t, responseJSONObjCase2["myTagName"] != nil, true)

	responseJSONObjCase3 := WrapArrayWithElemName([]MarshinalTest{}, "myTagName")
	assert.Equal(t, responseJSONObjCase3["myTagName"] != nil, true)

	responseJSONObjCase4 := WrapArrayWithElemName([]MarshinalTest{}, "")
	assert.Equal(t, responseJSONObjCase4[""] != nil, true)

	responseJSONObjCase5 := WrapArrayWithElemName("wrong type", "wrongTypeTag")
	assert.Equal(t, responseJSONObjCase5["wrongTypeTag"] != nil, true)
}

func TestIndexOfInt(t *testing.T) {
	inArr := []int{1, 2, 3, 4, 5, 6}
	result := IndexOfInt(inArr, 3)
	assert.Equal(t, result, 2)
	assert.NotEqual(t, result, 10)
}

func TestExistsInt(t *testing.T) {
	inArr := []string{"apple", "ball", "word", "speech", "angle", "happy"}
	result1 := ExistsString(inArr, "apple")
	assert.Equal(t, result1, true)

	result2 := ExistsString(inArr, "noExisits")
	assert.Equal(t, result2, false)
}

func TestUniqInt(t *testing.T) {
	inArr := []int{1, 1, 3, 2, 3, 5, 6, 7, 8}
	expectResult := []int{1, 3, 2, 5, 6, 7, 8}
	result := UniqInt(inArr)

	for index, ele := range result {
		fmt.Println(ele)
		assert.Equal(t, expectResult[index], ele)
	}
}

func TestIntArrayToString(t *testing.T) {
	inArr := []int{1, 1, 3, 2, 3, 5, 6, 7, 8}
	expectResult1 := "1,1,3,2,3,5,6,7,8"
	expectResult2 := "1-1-3-2-3-5-6-7-8"
	expectResult3 := "113235678"

	result1 := IntArrayToString(inArr, ",")
	assert.Equal(t, result1, expectResult1)

	result2 := IntArrayToString(inArr, "-")
	assert.Equal(t, result2, expectResult2)

	result3 := IntArrayToString(inArr, "")
	assert.Equal(t, result3, expectResult3)
}

func TestStringToIntArray(t *testing.T) {
	input1 := "1,1,3,2,3,5,6,7,8"
	expectResult1 := []int{1, 1, 3, 2, 3, 5, 6, 7, 8}

	result1, err1 := StringToIntArray(input1)
	for index, ele := range result1 {
		assert.Equal(t, expectResult1[index], ele)
	}
	assert.Equal(t, err1, nil)

	input2 := "1,1,3,2,3,5,6,7,x"
	result2, err2 := StringToIntArray(input2)
	for index, ele := range result2 {
		if index < len(result2)-1 {
			assert.Equal(t, expectResult1[index], ele)
		} else {
			if assert.Error(t, err2) {
				assert.Equal(t, 1, 1)
			}
		}
	}

	input3 := ""
	result3, err3 := StringToIntArray(input3)
	assert.Equal(t, len(result3), 0)
	assert.Equal(t, err3, nil)
}

func TestRemoveNumInArray(t *testing.T) {
	input1 := []int{1, 1, 3, 2, 3, 5, 6, 7, 8}
	expectResult1 := []int{1, 1, 2, 5, 6, 7, 8}
	result1 := RemoveNumInArray(3, input1)

	for index, ele := range result1 {
		assert.Equal(t, expectResult1[index], ele)
	}
}

func TestSliceInt(t *testing.T) {
	input := []int{1, 1, 3, 2, 3, 5, 6, 7, 8, 2, 3, 4, 2, 6, 4}
	expectResult1 := []int{3, 2, 3, 5, 6}
	expectResult3 := []int{3, 2, 3, 5, 6, 7, 8, 2, 3, 4, 2, 6, 4}

	result1 := SliceInt(2, 5, input)
	for index, ele := range result1 {
		assert.Equal(t, expectResult1[index], ele)
	}

	result2 := SliceInt(20, 5, input)
	assert.Equal(t, len(result2), 0)

	result3 := SliceInt(2, 50, input)
	for index, ele := range result3 {
		assert.Equal(t, expectResult3[index], ele)
	}
}

func TestSortMapByValueForTime(t *testing.T) {
	// var input map[int]time.Time
	// input[0] = time.Time{}
}

func TestSortMapByValueForFloat(t *testing.T) {
	var input map[int]float64
	input = make(map[int]float64, 4)
	input[0] = 3.8
	input[1] = 2.77
	input[2] = 7.77
	input[3] = 4.21
	expectResult1 := []int{2, 3, 0, 1}
	result1 := SortMapByValueForFloat(input)

	for index, ele := range result1 {
		assert.Equal(t, expectResult1[index], ele)
	}

}
