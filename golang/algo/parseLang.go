package algo

import (
	"fmt"
	"sort"
	"strings"
)

func ParseLang(header string, supportList []string) (avbLan []string, err error) {
	acceptLans := strings.Split(header, ", ")
	dict := make(map[string]struct{})
	for _, lan := range supportList {
		dict[lan] = struct{}{}
	}
	for _, lan := range acceptLans {
		if _, ok := dict[lan]; ok {
			avbLan = append(avbLan, lan)
		}
	}
	return avbLan, nil
}

func ParseLang2(header string, supportList []string) ([]string, error) {
	headerLans := strings.Split(header, ", ")
	tagDict := make(map[string][]string)
	lanDict := make(map[string]struct{})
	for _, lan := range supportList {
		strs := strings.Split(lan, "-")
		if len(strs) != 2 {
			return nil, fmt.Errorf("The input %v in supportList is not valid", strs)
		}
		tag := strs[0]
		tagDict[tag] = append(tagDict[tag], lan)
		lanDict[lan] = struct{}{}
	}
	l := newLangCollect()

	for _, lan := range headerLans {
		strs := strings.Split(lan, "-")
		if len(strs) > 1 {
			if _, ok := lanDict[lan]; ok {
				l.append(lan, true)
			}
		} else {
			// it is tag only
			for _, d := range tagDict[lan] {
				l.append(d, false)
			}
		}
	}
	return l.list, nil
}

func insertStr(arr []string, str string, index int) (newArr []string) {
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, str)
	newArr = append(newArr, arr[index:]...)
	return newArr
}

type langCollect struct {
	list    []string
	checker map[string]bool
}

func newLangCollect() *langCollect {
	return &langCollect{
		checker: make(map[string]bool),
	}
}

func (lc *langCollect) insert(lan string, index int) bool {
	if _, ok := lc.checker[lan]; ok {
		return false
	}
	lc.checker[lan] = false
	lc.list = insertStr(lc.list, lan, index)
	fmt.Println("to insert", lan, "in", index, "res:", lc.list)
	return true
}

func (lc *langCollect) append(lan string, allowDup bool) bool {
	_, ok := lc.checker[lan]
	if ok && !allowDup {
		return false
	}
	lc.list = append(lc.list, lan)
	lc.checker[lan] = false
	return true
}

func ParseLang3(header string, supportList []string) ([]string, error) {
	headerLans := strings.Split(header, ", ")
	tagDict := make(map[string][]string)
	lanDict := make(map[string]struct{})
	for _, lan := range supportList {
		strs := strings.Split(lan, "-")
		if len(strs) != 2 {
			return nil, fmt.Errorf("The input %v in supportList is not valid", strs)
		}
		tag := strs[0]
		tagDict[tag] = append(tagDict[tag], lan)
		lanDict[lan] = struct{}{}
	}
	l := newLangCollect()

	for _, lan := range headerLans {
		if lan == "*" {
			for _, slan := range supportList {
				l.append(slan, false)
			}
			return l.list, nil
		}
		strs := strings.Split(lan, "-")
		if len(strs) > 1 {
			if _, ok := lanDict[lan]; ok {
				l.append(lan, true)
			}
		} else {
			// it is tag only
			for _, d := range tagDict[lan] {
				l.append(d, false)
			}
		}
	}
	return l.list, nil
}

func ParseLang4(header string, supportList []string) ([]string, error) {
	acceptLans := strings.Split(header, ", ")
	tagDict := make(map[string][]string)
	lanDict := make(map[string]struct{})
	for _, lan := range supportList {
		strs := strings.Split(lan, "-")
		if len(strs) != 2 {
			return nil, fmt.Errorf("The input %v in supportList is not valid", strs)
		}
		tag := strs[0]
		tagDict[tag] = append(tagDict[tag], lan)
		lanDict[lan] = struct{}{}
	}

	sort.Slice(acceptLans, func(i, j int) bool {
		arr1, arr2 := strings.Split(acceptLans[i], ";"), strings.Split(acceptLans[j], ";")
		return arr1[1] > arr2[1]
	})

	l := newLangCollect()
	startPos := -1

	for _, lanCm := range acceptLans {
		lan := strings.Split(lanCm, ";")[0]
		if lan == "*" {
			startPos = len(l.list)
			continue
		}

		strs := strings.Split(lan, "-")
		if len(strs) > 1 {
			if _, ok := lanDict[lan]; ok {
				l.append(lan, true)
			}
		} else {
			// it is tag only
			for _, d := range tagDict[lan] {
				l.append(d, false)
			}
		}
	}

	fmt.Println("before ins:", l.list)
	if startPos >= 0 {
		for _, supLan := range supportList {
			if l.insert(supLan, startPos) {
				startPos++
			}
		}
	}
	fmt.Println("after ins:", l.list)
	var res []string
	for i := len(l.list) - 1; i >= 0; i-- {
		if !l.checker[l.list[i]] {
			res = append(res, l.list[i])
			l.checker[l.list[i]] = true
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res, nil
}
