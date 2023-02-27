package practical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInArray(t *testing.T) {
	actual := InArray(1, []int{3, 4, 1, 5})
	assert.Equal(t, true, actual)
}
func TestArrayCombine(t *testing.T) {
	expect := map[int]string{
		5: "ArrayCombine",
		7: "InArray",
	}
	actual := ArrayCombine[int, string]([]int{5, 7}, []string{"ArrayCombine", "InArray"})
	assert.Equal(t, expect, actual)
	expect2 := map[string]string{
		"key1": "ArrayCombine",
		"key2": "InArray",
	}
	actual2 := ArrayCombine[string, string]([]string{"key1", "key2"}, []string{"ArrayCombine", "InArray"})
	assert.Equal(t, expect2, actual2)
	expect3 := map[interface{}]string{
		5:      "ArrayCombine",
		"key2": "InArray",
	}
	actual3 := ArrayCombine[any, string]([]interface{}{5, "key2"}, []string{"ArrayCombine", "InArray"})
	assert.Equal(t, expect3, actual3)
}

func TestArrayMerge(t *testing.T) {
	expect := []string{
		"111",
		"222",
		"333",
		"444",
	}
	actual := ArrayMerge([]string{"111"}, []string{"222", "333"}, []string{"444"})
	assert.Equal(t, expect, actual)
}
func TestArraySlice(t *testing.T) {
	expect := []string{
		"333",
		"444",
		"555",
	}
	actual := ArraySlice([]string{"111", "222", "333", "444", "555"}, 2, 3)
	assert.Equal(t, expect, actual)
}
func TestArrayDiff(t *testing.T) {
	expect := []int{
		1, 2, 3, 6, 7, 8, 9, 10,
	}
	actual := ArrayDiff([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{4, 13}, []int{5, 20})
	assert.Equal(t, expect, actual)
}

func TestArrayIntersect(t *testing.T) {
	expect := []int{
		4, 5,
	}
	actual := ArrayIntersect([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{4, 13}, []int{5, 20})
	assert.Equal(t, expect, actual)
}
func TestArraySearch(t *testing.T) {
	expect := 2
	actual := ArraySearch([]int{1, 2, 3, 4, 5}, 3)
	assert.Equal(t, expect, actual)
}
func TestArraySum(t *testing.T) {
	var expect1 float64 = 10.0
	actual1 := ArraySum([]float64{1.0, 9.0})
	assert.Equal(t, expect1, actual1)
	var expect2 int16 = 10
	actual2 := ArraySum([]int16{1, 9})
	assert.Equal(t, expect2, actual2)
}
func TestArrayUnique(t *testing.T) {
	expect := []float64{
		1.0,
		2.0,
		10.0,
	}
	actual := ArrayUnique([]float64{1.0, 2.0, 1.0, 10.0, 2.0, 10.0})
	assert.Equal(t, expect, actual)
}

func TestArrayRand(t *testing.T) {
	expect := []int{
		1, 2, 3,
	}
	actual := ArrayRand([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
	assert.Equal(t, expect, actual)
}
func TestArrayCountValues(t *testing.T) {
	expect := map[int]int{
		1: 2,
		2: 3,
		5: 1,
	}
	actual := ArrayCountValues([]int{1, 1, 2, 2, 2, 5})
	assert.Equal(t, expect, actual)
}
func TestArraySplice(t *testing.T) {
	expect := []int{
		1, 2, 8, 9, 6, 7,
	}
	actual := ArraySplice([]int{1, 2, 4, 5, 6, 7}, 2, 2, []int{8, 9})
	assert.Equal(t, expect, actual)
}
func TestArrayShuffle(t *testing.T) {
	expect := []int{
		1, 6, 5, 3, 2, 10, 7, 8, 4, 9,
	}
	actual := ArrayShuffle([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, expect, actual)
}
func TestArrayPop(t *testing.T) {
	expect := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}
	actual := ArrayPop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, expect, actual)
}
func TestArrayPush(t *testing.T) {
	expect := []int{
		1, 2, 3,
	}
	actual := ArrayPush([]int{1}, 2, 3)
	assert.Equal(t, expect, actual)
}

func TestArrayUnshift(t *testing.T) {
	expect := []int{
		1, 2, 3,
	}
	actual := ArrayUnshift([]int{2, 3}, 1)
	assert.Equal(t, expect, actual)
}
func TestArrayDescSort(t *testing.T) {
	expect := []int{
		10, 9, 8, 7, 6, 5, 4, 3, 2, 1,
	}
	actual := ArrayDescSort([]int{3, 1, 2, 5, 10, 9, 4, 8, 7, 6})
	assert.Equal(t, expect, actual)
}
func TestArrayAscSort(t *testing.T) {
	expect := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	actual := ArrayAscSort([]int{3, 1, 2, 5, 10, 9, 4, 8, 7, 6})
	assert.Equal(t, expect, actual)
}
func TestArrayReverse(t *testing.T) {
	expect := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}
	actual := ArrayReverse([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	assert.Equal(t, expect, actual)
}
func TestArrayToString(t *testing.T) {
	expect := "1,2,3,4,5,6,7,8,9,10"
	actual := ArrayToString([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	assert.Equal(t, expect, actual)
}
