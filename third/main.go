package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"sort"
)

type (
	Input struct {
		ID   string  `json:"id"`
		Time int64   `json:"timestamp"`
		Temp float64 `json:"temperature"`
	}
	Output struct {
		ID   string    `json:"id"`
		Avg  float64   `json:"average"`
		Med  float64   `json:"median"`
		Mode []float64 `json:"mode"`
	}
)

func calculateData(data *[]Input) []Output {
	a, b, c := 0, 0, 0
	suma, sumb, sumc := 0.0, 0.0, 0.0
	adata, bdata, cdata := []float64{}, []float64{}, []float64{}
	ca, cb, cc := map[float64]int{}, map[float64]int{}, map[float64]int{}
	na, nb, nc := 0, 0, 0
	ma, mb, mc := []float64{}, []float64{}, []float64{}
	for _, v := range *data {
		switch v.ID {
		case "a":
			adata = append(adata, v.Temp)
			suma += v.Temp
			a++
			ca[v.Temp]++
			if na < ca[v.Temp] {
				ma = ma[:0]
				ma = append(ma, v.Temp)
				na = ca[v.Temp]
			} else if na == ca[v.Temp] {
				ma = append(ma, v.Temp)
			}
		case "b":
			bdata = append(bdata, v.Temp)
			sumb += v.Temp
			b++
			cb[v.Temp]++
			if nb < cb[v.Temp] {
				mb = mb[:0]
				mb = append(mb, v.Temp)
				nb = cb[v.Temp]
			} else if nb == cb[v.Temp] {
				mb = append(mb, v.Temp)
			}

		case "c":
			cdata = append(cdata, v.Temp)
			sumc += v.Temp
			c++
			cc[v.Temp]++
			if nc < cc[v.Temp] {
				mc = mc[:0]
				mc = append(mc, v.Temp)
				nc = cc[v.Temp]
			} else if nc == cc[v.Temp] {
				mc = append(mc, v.Temp)
			}

		}
	}

	sort.Float64s(adata)
	sort.Float64s(bdata)
	sort.Float64s(cdata)

	// Average
	avga, avgb, avgc := suma/float64(a), sumb/float64(b), sumc/float64(c)
	avga, avgb, avgc = math.Round(avga*100)/100, math.Round(avgb*100)/100, math.Round(avgc*100)/100

	// Median
	meda, medb, medc := 0.0, 0.0, 0.0
	if a%2 == 1 {
		meda = adata[(a+1)/2-1]
	} else {
		meda = (adata[a/2-1] + adata[a/2]) / 2
	}
	if b%2 == 1 {
		medb = bdata[(b+1)/2-1]
	} else {
		medb = (bdata[b/2-1] + bdata[b/2]) / 2
	}
	if c%2 == 1 {
		medc = cdata[(c+1)/2-1]
	} else {
		medc = (cdata[c/2-1] + cdata[c/2]) / 2
	}
	meda, medb, medc = math.Round(meda*100)/100, math.Round(medb*100)/100, math.Round(medc*100)/100

	// Output
	var out = []Output{
		Output{
			ID:   "a",
			Avg:  avga,
			Med:  meda,
			Mode: ma,
		},
		Output{
			ID:   "b",
			Avg:  avgb,
			Med:  medb,
			Mode: mb,
		},
		Output{
			ID:   "c",
			Avg:  avgc,
			Med:  medc,
			Mode: mc,
		},
	}

	return out

}

func main() {
	data := []Input{}

	file, err := os.Open("input.json")
	if err != nil {
		log.Panic(err)
	}

	parser := json.NewDecoder(file)
	if err = parser.Decode(&data); err != nil {
		log.Panic(err)
	}

	out := calculateData(&data)
	jsonOut, err := json.Marshal(out)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsonOut))
	err = ioutil.WriteFile("output.json", jsonOut, 0644)
}
