package algo

import (
	"container/heap"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/emirpasic/gods/trees/binaryheap"
)

func Add(a int, b int) int {
	fmt.Println(a, b)
	return a + b
}

const (
	header_separator = ","
	lang_separator   = "-"
)

func parseHeader(clientLangs string) []string {
	var validClientLangs []byte
	for _, char := range clientLangs {
		if byte(char) == byte(' ') {
			continue
		}
		validClientLangs = append(validClientLangs, byte(char))
	}
	return strings.Split(string(validClientLangs), header_separator)
}

type supportLangs struct {
	langs []string
	dict  map[string]bool
}

func NewSupportLangs() *supportLangs {
	return &supportLangs{
		dict: make(map[string]bool),
	}
}

func (sl *supportLangs) append(lang string) {
	if _, ok := sl.dict[lang]; ok {
		return
	}
	sl.langs = append(sl.langs, lang)
	sl.dict[lang] = true
}

func ParseLang(clientLangs string, serverLangs []string) []string {
	fmt.Println("=======")
	langDict := make(map[string]struct{})
	tagDict := make(map[string][]string)
	for _, lang := range serverLangs {
		langDict[lang] = struct{}{}
		langStrs := strings.Split(lang, lang_separator)
		if len(langStrs) != 2 {
			continue
		}
		tagDict[langStrs[0]] = append(tagDict[langStrs[0]], lang)
	}

	clientLangArr := parseHeader(clientLangs)
	res := NewSupportLangs()

	for _, lang := range clientLangArr {
		if lang == "" {
			continue
		}
		if lang == "*" {
			fmt.Println("come here")
			for _, slang := range serverLangs {
				res.append(slang)
			}
			break
		}
		langStrs := strings.Split(lang, lang_separator)
		if len(langStrs) > 1 {
			if _, ok := langDict[lang]; ok {
				res.append(lang)
			}
		} else {
			// lang is tag
			for _, d := range tagDict[lang] {
				res.append(d)
			}
		}
	}
	return res.langs
}

type Py struct {
	Name  string `json:"name"`
	Val   int    `json:"bal"`
	index int
}

func (p *Py) String() string {
	return fmt.Sprintf("name:%v,bal:%v;", p.Name, p.Val)
}

type PriorityQ []*Py

func (p *PriorityQ) Len() int {
	return len(*p)
}

func (p PriorityQ) Less(i, j int) bool {
	arr := p
	return arr[i].Val > arr[j].Val
}

func (p PriorityQ) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *PriorityQ) Push(d interface{}) {
	sz := len(*p)
	item := d.(*Py)
	item.index = sz
	*p = append(*p, item)
}

func (p *PriorityQ) Pop() interface{} {
	arr := *p
	sz := len(arr)
	item := arr[sz-1]
	item.index = -1
	arr[sz-1] = nil
	*p = arr[0 : sz-1]
	return item
}

type CustomTimer struct {
	now time.Time
}

func (c *CustomTimer) SetNow(t time.Time) {
	c.now = t
}

func (c *CustomTimer) Now() time.Time {
	return c.now
}

// rate limit for each user
type limiter struct {
	tokens        int
	lastUpdatedTs time.Time

	// const
	refillCount int
	// const
	maxCount int
	// const
	refillInterval time.Duration
	timer          *CustomTimer
}

func newLimiter(initialTokens int, refillCount int, maxTokens int, refillInterval time.Duration,
	timer *CustomTimer) *limiter {
	return &limiter{
		tokens:         initialTokens,
		refillCount:    refillCount,
		maxCount:       maxTokens,
		lastUpdatedTs:  time.Now(),
		refillInterval: refillInterval,
		timer:          timer,
	}
}

// check whether need to refill, and do refill if needed
func (l *limiter) refill() {
	now := l.timer.Now()
	if now.Sub(l.lastUpdatedTs) < l.refillInterval {
		return
	}
	l.tokens += l.refillCount
	if l.tokens > l.maxCount {
		l.tokens = l.maxCount
	}
	l.lastUpdatedTs = now
}

func (l *limiter) execute() bool {
	l.refill()
	//fmt.Println("tokens", l.tokens)
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}

type RateLimiter struct {
	limiterDict    map[string]*limiter
	refillCount    int
	refillInterval time.Duration
	timer          *CustomTimer
}

func NewRateLimiter(refillInterval time.Duration, refillCount int, timer *CustomTimer) *RateLimiter {
	return &RateLimiter{
		limiterDict:    make(map[string]*limiter),
		timer:          timer,
		refillCount:    refillCount,
		refillInterval: refillInterval,
	}
}

func (r *RateLimiter) ExecuteEndpoint(customer string) bool {
	l, ok := r.limiterDict[customer]
	if !ok {
		r.limiterDict[customer] = newLimiter(
			r.refillCount, r.refillCount, r.refillCount, r.refillInterval, r.timer)
		l = r.limiterDict[customer]
	}
	//fmt.Println("come", customer)
	return l.execute()
}

func remoteGetVal(key string) string {
	return fmt.Sprintf("hello_%v", key)
}

func RemoteGetVal(key string) string {
	time.Sleep(time.Second)
	return remoteGetVal(key)
}

func RemoteGetVals(keys []string) (vals []string) {
	for _, key := range keys {
		vals = append(vals, remoteGetVal(key))
	}
	time.Sleep(time.Second)
	return vals
}

func PrintRes(key string, val string) {
	fmt.Printf("The val for key %v: %v\n", key, val)
}

type KeyFetcher struct {
	maxBatchSize     int
	maxFetchInterval time.Duration
	queue            chan *KV
	dispatchSignal   chan struct{}
	stopChan         chan struct{}
	callbackWg       sync.WaitGroup
}

func NewKeyFetcher(maxQueueSz int, maxBatchSize int, maxFetchInterval time.Duration) *KeyFetcher {
	return &KeyFetcher{
		queue:            make(chan *KV, maxQueueSz),
		dispatchSignal:   make(chan struct{}),
		maxBatchSize:     maxBatchSize,
		maxFetchInterval: maxFetchInterval,
		stopChan:         make(chan struct{}),
	}
}
func (k *KeyFetcher) AsyncStart() {
	go func() {
		k.Start()
	}()
}

func (k *KeyFetcher) Start() {
	for {
		select {
		case <-time.After(k.maxFetchInterval):
		case <-k.dispatchSignal:
			if len(k.queue) < k.maxBatchSize {
				continue
			}
		case <-k.stopChan:
			return
		}

		qSize := len(k.queue)
		if qSize > k.maxBatchSize {
			qSize = k.maxBatchSize
		}
		var kvs []*KV
		var keys []string
		for i := 0; i < qSize; i++ {
			kv := <-k.queue
			kvs = append(kvs, kv)
			keys = append(keys, kv.key)
		}
		vals := RemoteGetVals(keys)
		for i, kv := range kvs {
			kv.val <- vals[i]
		}
	}
}

func (k *KeyFetcher) Stop() {
	k.callbackWg.Wait()
	close(k.stopChan)
}

type KV struct {
	key string
	val chan string
}

func (k *KeyFetcher) GetKey(key string, callback func(key string, val string)) {
	kv := &KV{
		key: key,
		val: make(chan string, 1),
	}
	k.queue <- kv
	k.callbackWg.Add(1)
	go func() {
		defer k.callbackWg.Done()
		k.dispatchSignal <- struct{}{}
		callback(key, <-kv.val)
	}()
}

func dfsFilePath(rootPath string, fileSizeDict map[int64][]string) {
	finfo, err := os.Stat(rootPath)
	if err != nil {
		return
	}
	if !finfo.IsDir() {
		sz := finfo.Size()
		fileSizeDict[sz] = append(fileSizeDict[sz], rootPath)
		return
	}
	finfos, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return
	}
	for _, finfo := range finfos {
		fpath := path.Join(rootPath, finfo.Name())
		dfsFilePath(fpath, fileSizeDict)
	}
}

func FindDupFiles(rootPath string) (res [][]string) {
	fileSizeDict := make(map[int64][]string)
	dfsFilePath(rootPath, fileSizeDict)
	for _, fpaths := range fileSizeDict {
		if len(fpaths) == 1 {
			continue
		}
		md5sumDict := make(map[string][]string)
		for _, fpath := range fpaths {
			f, err := os.Open(fpath)
			if err != nil {
				continue
			}
			defer f.Close()

			h := md5.New()
			if _, err := io.Copy(h, f); err != nil {
				continue
			}
			sum := string(h.Sum(nil))
			md5sumDict[sum] = append(md5sumDict[sum], fpath)
		}
		for _, fpaths := range md5sumDict {
			if len(fpaths) == 1 {
				continue
			}
			res = append(res, fpaths)
		}
	}
	return res
}

func Run() error {
	t := time.Now()
	fmt.Println("start run", t)
	pys := []*Py{
		{Name: "hello", Val: 10},
		{Name: "world", Val: 11},
		{Name: "end", Val: 12},
		{Name: "again", Val: 13},
	}

	pq := PriorityQ(pys)
	heap.Init(&pq)

	comp := func(a interface{}, b interface{}) int {
		return a.(*Py).Val - b.(*Py).Val
	}
	m := binaryheap.NewWith(comp)
	for _, p := range pys {
		m.Push(p)
	}

	kvf := NewKeyFetcher(100, 3, time.Second*3)
	kvf.AsyncStart()
	fmt.Println(time.Now())
	kvf.GetKey("Py", PrintRes)
	kvf.GetKey("Py", PrintRes)
	kvf.GetKey("Py", PrintRes)
	kvf.GetKey("Py", PrintRes)
	kvf.GetKey("Py", PrintRes)
	kvf.GetKey("Py", PrintRes)
	kvf.Stop()
	fmt.Println(time.Now())
	// fmt.Println(FindDupFiles("/tmp/pytest"))
	return nil
}
