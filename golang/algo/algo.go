package algo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strings"
)

func Add(a int, b int) int {
	fmt.Println(a, b)
	return a + b
}

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
	sz := len(arr) + 1
	newArr = make([]string, sz)
	j := 0
	for i := 0; i < sz; i++ {
		if i == index {
			newArr[i] = str
		} else {
			newArr[i] = arr[j]
			j++
		}
	}
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

type edge struct {
	name string
	val  float64
}

type pair struct {
	src string
	tgt string
}

func dfs(m map[string][]edge, dict map[pair]float64, cur string, tgt string) float64 {
	if cur == tgt {
		return 1
	}
	p := pair{src: cur, tgt: tgt}
	if v, ok := dict[p]; ok {
		return v
	}
	dict[p] = -1

	for _, e := range m[cur] {
		nv := dfs(m, dict, e.name, tgt)
		if nv == -1 {
			continue
		}
		dict[p] = nv * e.val
		return dict[p]
	}
	return -1
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string][]edge)
	for i, eq := range equations {
		m[eq[0]] = append(m[eq[0]], edge{
			name: eq[1],
			val:  values[i],
		})
		m[eq[1]] = append(m[eq[1]], edge{
			name: eq[0],
			val:  1 / values[i],
		})
	}
	dict := make(map[pair]float64)

	var res []float64
	for _, q := range queries {
		if _, ok := m[q[0]]; !ok {
			res = append(res, -1)
			continue
		}
		res = append(res, dfs(m, dict, q[0], q[1]))
	}
	return res
}

type player int

const (
	noone player = iota
	p1
	p2
	tie
)

const (
	rowNum    = 7
	colNum    = 6
	threshold = 4
)

type FourConn struct {
	board  [][]player
	indexs []int
	moves  [][]int
}

func NewFourConn() *FourConn {
	f := &FourConn{
		board:  make([][]player, rowNum),
		indexs: make([]int, rowNum),
	}
	for i := range f.board {
		f.board[i] = make([]player, colNum)
	}
	return f
}

func (f *FourConn) PlaceAndCheck(p player, row int) player {
	if !f.Place(p, row) {
		return tie
	}
	res := f.CheckWin(p, row, f.indexs[row]-1)
	fmt.Println(res)
	return res
}

func (f *FourConn) Place(p player, row int) bool {
	if row >= len(f.indexs) {
		return false
	}
	if f.indexs[row] >= colNum {
		return false
	}
	f.board[row][f.indexs[row]] = p
	f.moves = append(f.moves, []int{row, f.indexs[row]})
	f.indexs[row]++
	return true
}

func (f *FourConn) getVal(x int, y int) (player, bool) {
	if x >= 0 && x < rowNum && y >= 0 && y < colNum {
		return f.board[x][y], true
	}
	return noone, false
}

func (f *FourConn) undo() {
	if len(f.moves) == 0 {
		return
	}
	lastMove := f.moves[len(f.moves)-1]
	f.moves = f.moves[:len(f.moves)-1]
	f.board[lastMove[0]][lastMove[1]] = 0
	f.indexs[lastMove[0]]--
}

func (f *FourConn) anotherPlayer(p player) player {
	if p == p1 {
		return p2
	}
	return p1
}

func (f *FourConn) minimax(p player, expectP player) (int, int) {
var bestMove int
	if p == expectP {
		ret := math.MinInt32
		for i := 0; i < rowNum; i++ {
			res := f.PlaceAndCheck(p, i)
			if res == noone {
				res, _ = minimax(f.anotherPlayer(p), expectP)
			}
			f.undo()
			if res > ret {
				res = ret
				bestMoves = i
			}
			if res == 1 {
				return 1, move
			}
		}  
	} else {
		ret := math.MaxInt32
		for i := 0; i < rowNum; i++ {
			res := f.PlaceAndCheck(p, i)
			if res == noone {
				res, _ = minimax(f.anotherPlayer(p), expectP)
			}
			f.undo()
			if res < ret {
				res = ret
				bestMoves = i
			}
			if res == -1 {
				return 1, move
			}
		}  
	}
}

func (f *FourConn) CheckWin(p player, x int, y int) player {
	moves := [][]int{
		{-1, 0, 1, 0},
		{0, -1, 0, 1},
		{-1, -1, 1, 1},
		{1, -1, -1, 1},
	}

	for _, move := range moves {
		l, r := 0, 0
		for true {
			curp, ok := f.getVal(x+move[0]*l, y+move[1]*l)
			if !ok || curp != p {
				break
			}
			l++
		}
		for true {
			curp, ok := f.getVal(x+move[2]*r, y+move[3]*r)
			if !ok || curp != p {
				break
			}
			r++
		}
		count := l + r - 1
		if count >= threshold {
			return p
		}
	}

	for _, index := range f.indexs {
		if index != colNum {
			return noone
		}
	}
	return tie
}

func Run() error {
	fmt.Println("start")
	f := NewFourConn()
	f.PlaceAndCheck(p1, 0)

	f.PlaceAndCheck(p2, 1)
	f.PlaceAndCheck(p1, 1)

	f.PlaceAndCheck(p2, 2)
	f.PlaceAndCheck(p2, 2)
	f.PlaceAndCheck(p1, 2)

	f.PlaceAndCheck(p2, 3)
	f.PlaceAndCheck(p2, 3)
	f.PlaceAndCheck(p2, 3)
	f.PlaceAndCheck(p1, 3)

	client := &http.Client{}
	resp, err := client.Get("http://192.168.55.197:9000/echo")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := struct {
		Err bool   `json:"error"`
		Msg string `json:"message"`
	}{}
	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}
	fmt.Println(res)

	fmt.Println(calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))
	return nil
}
