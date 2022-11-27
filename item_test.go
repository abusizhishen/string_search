package string_search

import (
	"fmt"
	"log"
	"sort"
	"testing"
	"time"
)

func TestNewItem(t *testing.T) {
	var start = time.Now()
	var id = 1
	var str = "name"

	item := NewItem(id, str)
	item.Insert(2, "ame")
	item.Insert(3, "aeme")
	item.Insert(4, "张三")

	t.Logf("cost:%v", time.Since(start))
	log.Println(item)
}

func TestItem_Search(t *testing.T) {
	i := NewItem(1, "nabc")
	i.Insert(2, "abc")
	i.Insert(3, "aba")
	i.Insert(4, "张三")
	i.Insert(5, "abcd")
	i.Insert(6, "abc")

	var keywords = []string{
		"b",
		"aa",
		"bcd",
		"三",
	}

	var start = time.Now()
	for _, words := range keywords {
		result := i.Search(words)
		t.Logf("key: %s", words)
		t.Logf("result: %v", result)
	}
	t.Logf("cost: %v", time.Since(start))
}

func checkIds(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Ints(a)
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestName(t *testing.T) {
	var str = "123"
	for _, s := range str {
		fmt.Println(s)
	}

	log.Println(len(str), "---")
	for i := 0; i < len(str); i++ {
		log.Println(str[i])
	}
}
