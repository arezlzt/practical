package practical

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// ArrayCombine
// 合并两个数组来创建一个map，其中的一个数组是键名，另一个数组的值为键值
func ArrayCombine[T comparable, E any](keys []T, values []E) map[T]E {
	valueLength := len(values)
	if valueLength != len(keys) {
		panic("the number of elements for each slice isn't equal")
	}
	maps := make(map[T]E, len(keys))
	for i := 0; i < len(keys); i++ {
		maps[keys[i]] = values[i]
	}
	return maps
}

// InArray
// 判断元素是否在数组里面
func InArray[T comparable](val T, array []T) bool {
	if len(array) == 0 {
		return false
	}
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return true
		}
	}
	return false
}

// ArrayMerge
// 数组合并
func ArrayMerge[T any](arr []T, options ...[]T) []T {
	if len(options) <= 0 {
		return arr
	}
	for _, value := range options {
		for _, val := range value {
			arr = append(arr, val)
		}
	}
	return arr
}

// ArraySlice
// 自定义位置截取
func ArraySlice[T any](arr []T, start, length int) []T {
	if start > len(arr) {
		panic("start:the start is less than the length of arr")
	}
	end := start + length
	if end < len(arr) {
		return arr[start:end]
	} else {
		return arr[start:]
	}
}

// ArrayDiff 返回差集
func ArrayDiff[T comparable](array []T, otherArrays ...[]T) []T {
	existMap := make(map[T]bool)
	for i := 0; i < len(array); i++ {
		existMap[array[i]] = false
	}
	for i := 0; i < len(otherArrays); i++ {
		for j := 0; j < len(otherArrays[i]); j++ {
			if _, flag := existMap[otherArrays[i][j]]; flag {
				existMap[otherArrays[i][j]] = true
			} else {
				existMap[otherArrays[i][j]] = false
			}
		}
	}
	result := make([]T, 0)
	for key, value := range existMap {
		if !value {
			result = append(result, key)
		}
	}
	return result
}

// ArrayIntersect 返回交集
func ArrayIntersect[T comparable](array []T, otherArrays ...[]T) []T {
	existMap := make(map[T]bool)
	for i := 0; i < len(array); i++ {
		existMap[array[i]] = false
	}
	for i := 0; i < len(otherArrays); i++ {
		for j := 0; j < len(otherArrays[i]); j++ {
			if _, flag := existMap[otherArrays[i][j]]; flag {
				existMap[otherArrays[i][j]] = true
			} else {
				existMap[otherArrays[i][j]] = false
			}
		}
	}
	result := make([]T, 0)
	for key, value := range existMap {
		if value {
			result = append(result, key)
		}
	}
	return result
}

// ArraySearch 搜索
func ArraySearch[T comparable](array []T, search T) int {
	index := -1
	for key, value := range array {
		if value == search {
			return key
		}
	}
	return index
}

// ArraySplice 从数组中移除选定的元素,并用新元素取代它
func ArraySplice[T any](array []T, start, length int, replaceArray []T) []T {
	startArray := array[:start]
	endArray := ArraySlice(array, start+length, len(array)-start+length)
	startArray = append(startArray, replaceArray...)
	startArray = append(startArray, endArray...)
	return startArray
}

// ArraySum 返回所有值的和
func ArraySum[T constraints.Float | constraints.Integer](array []T) T {
	var sum T
	for _, value := range array {
		sum += value
	}
	return sum
}

// ArrayUnique 移除数组中重复的值
func ArrayUnique[T comparable](array []T) []T {
	existMap := make(map[T]struct{})
	for _, value := range array {
		existMap[value] = struct{}{}
	}
	arr := make([]T, 0)
	for key := range existMap {
		arr = append(arr, key)
	}
	return arr
}

// ArrayRand 从数组中随机抽取length个元素,返回新的数组
func ArrayRand[T comparable](array []T, length int) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if length > len(array) {
		length = len(array)
	}
	n := make([]T, 0)
	for i, v := range r.Perm(len(array)) {
		if i >= length {
			break
		}
		n = append(n, array[v])
	}
	return n
}

// ArrayCountValues 统计数组中所有值出现的次数
func ArrayCountValues[T constraints.Ordered](array []T) map[T]int {
	countMap := make(map[T]int)
	for _, value := range array {
		if _, ok := countMap[value]; !ok {
			countMap[value] = 0
		}
		countMap[value]++
	}
	return countMap
}

// ArrayShuffle 将数组中的元素打乱，重新随机排列
func ArrayShuffle[T comparable](array []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := make([]T, 0)
	for _, v := range r.Perm(len(array)) {
		n = append(n, array[v])
	}
	return n
}

// ArrayUnshift 在数组开头插入一个或多个元素
func ArrayUnshift[T any](array []T, element ...T) []T {
	array = append(element, array...)
	return array
}

// ArrayPop 删除数组中最后一个元素。
func ArrayPop[T any](array []T) []T {
	if len(array) == 0 {
		return array
	}
	return array[0 : len(array)-1]
}

// ArrayPush 在数组结尾插入一个或多个元素
func ArrayPush[T any](array []T, element ...T) []T {
	array = append(array, element...)
	return array
}

// ArrayAscSort 升序排列
func ArrayAscSort[T constraints.Ordered](array []T) []T {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	return array
}

// ArrayDescSort 降序排列
func ArrayDescSort[T constraints.Ordered](array []T) []T {
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
	return array
}

// ArrayReverse 返回一个顺序相反的数组
func ArrayReverse[T comparable](array []T) []T {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
func ArrayToString[T any](array []T) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
