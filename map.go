package practical

import (
	"fmt"
)

// MapKeyExists 判断某个map中是否存在指定的key 存在返回true 否则返回false
func MapKeyExists[T comparable, E any](m map[T]E, key T) bool {
	_, ok := m[key]
	return ok
}

// MapKeys 返回包含map中所有键名的一个数组
func MapKeys[T comparable, E any](m map[T]E) []T {
	array := make([]T, 0)
	for key := range m {
		array = append(array, key)
	}
	return array
}

// MapValues 返回一个包含给定map中所有键值的数组
func MapValues[T comparable, E any](m map[T]E) []E {
	array := make([]E, 0)
	for _, value := range m {
		array = append(array, value)
	}
	return array
}

// MapFlip 反转/交换map中的键名和对应关联的键值
func MapFlip[T comparable, E comparable](m map[T]E) map[E]T {
	m1 := make(map[E]T)
	for k, v := range m {
		m1[v] = k
	}
	return m1
}

// MapShift 删除map中第一个元素，并返回被删除元素的值
func MapShift[T comparable, E any](m map[T]E, key T) E {
	if v, ok := m[key]; ok {
		delete(m, key)
		return v
	} else {
		panic(fmt.Sprintf("key %v not exists", key))
	}

}
func MapKeyDescSort() {

}
func MapKeyAscSort() {

}
func MapValueDescSort() {

}
func MapValueAscSort() {

}
