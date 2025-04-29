package main

import "fmt"

func add(m map[string]map[string]int, path, country string) {
    mm, ok := m[path]
    if !ok {
        mm = make(map[string]int)
        m[path] = mm
    }
    mm[country]++
}

type Key struct {
    Path, Country string
}

func main(){
	hits := make(map[string]map[string]int)
	add(hits,"/docs/","br")
	add(hits,"/docs/","br")
	add(hits,"/docs/","br")
	add(hits,"/docs/","us")
	add(hits,"/docs/","us")
	add(hits,"/docs/","au")
	add(hits,"/health/","au")
	add(hits,"/health/","au")
	add(hits,"/health/","au")
	fmt.Println(hits)
	hits1 := make(map[Key]int)
	hits1[Key{Path: "/",Country: "BR"}]++
	hits1[Key{Path: "/",Country: "BR"}]++
	hits1[Key{Path: "/",Country: "BR"}]++
	hits1[Key{Path: "/",Country: "AU"}]++
	hits1[Key{Path: "/",Country: "AU"}]++
	hits1[Key{Path: "/",Country: "AU"}]++
	hits1[Key{Path: "/",Country: "AU"}]++
	hits1[Key{Path: "/",Country: "AU"}]++
	hits1[Key{Path: "/health",Country: "BR"}]++
	hits1[Key{Path: "/health",Country: "BR"}]++
	hits1[Key{Path: "/health",Country: "BR"}]++
	hits1[Key{Path: "/health",Country: "AU"}]++
	hits1[Key{Path: "/health",Country: "AU"}]++
	hits1[Key{Path: "/health",Country: "AU"}]++
	n := hits1[Key{"/ref/spec", "ch"}]
	fmt.Println("aqui",n)
}