package algo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	// can not make too large, because coinbase api has rate limiter
	parallel_count = 20
)

type CurrencyPair struct {
	ID  string `json:"id"`
	Src string `json:"base_currency"`
	Tgt string `json:"quote_currency"`
}

func sendReq(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getPrice(productUrl string, src string, tgt string) (Ticker, error) {
	url := fmt.Sprintf("%v/%v-%v/ticker", productUrl, src, tgt)
	quoteData, err := sendReq(url)
	if err != nil {
		return Ticker{}, err
	}
	quoteJson := struct {
		Bid string `json:"bid"`
		Ask string `json:"ask"`
	}{}
	if err := json.Unmarshal(quoteData, &quoteJson); err != nil {
		return Ticker{}, err
	}
	// fmt.Printf("data for %v is %v\n", url, quoteJson)
	bidPrice, err := strconv.ParseFloat(quoteJson.Bid, 64)
	if err != nil {
		return Ticker{}, err
	}
	askPrice, err := strconv.ParseFloat(quoteJson.Ask, 64)
	if err != nil {
		return Ticker{}, err
	}
	t := Ticker{
		src: src,
		tgt: tgt,
		bid: bidPrice,
		ask: askPrice,
	}
	//fmt.Println("tickr", t)
	return t, nil
}

type Ticker struct {
	src string
	tgt string
	ask float64 // tgt->src
	bid float64 // src->tgt
}

type Edge struct {
	tgt string
	val float64
}

// a->b b->c c->a a->d

func BestCoversion(tickers []Ticker, base string, quote string, amount float64) (float64, []string) {
	graph := make(map[string][]Edge)
	for _, t := range tickers {
		graph[t.src] = append(graph[t.src], Edge{
			tgt: t.tgt,
			val: t.bid,
		})
		graph[t.tgt] = append(graph[t.tgt], Edge{
			tgt: t.src,
			val: 1 / t.ask,
		})
	}

	var maxPaths []string
	dict := make(map[pair]cachePath)
	cp := dfs(graph, base, quote, dict, maxPaths)
	return cp.val * amount, cp.paths
}

type pair struct {
	src string
	tgt string
}

type cachePath struct {
	paths []string
	val   float64
}

func dfs(graph map[string][]Edge, cur string, expect string, dict map[pair]cachePath, maxPaths []string) cachePath {
	if cur == expect {
		return cachePath{val: 1, paths: []string{expect}}
	}
	p := pair{
		src: cur,
		tgt: expect,
	}
	if v, ok := dict[p]; ok {
		return v
	}
	// marking loop
	dict[p] = cachePath{
		val: -1,
	}

	maxV := float64(-1)
	for _, e := range graph[cur] {
		childPaths := dfs(graph, e.tgt, expect, dict, maxPaths)
		childVal := childPaths.val
		if childVal == -1 {
			continue
		}
		if childVal*e.val > maxV {
			var newPath []string
			newPath = append(newPath, cur)
			newPath = append(newPath, childPaths.paths...)
			maxPaths = newPath
			maxV = childVal * e.val
		}
	}
	// change it to only record value
	if maxV == -1 {
		delete(dict, p)
		return cachePath{val: -1, paths: []string{expect}}
	}
	dict[p] = cachePath{
		val:   maxV,
		paths: maxPaths,
	}
	return dict[p]
}
