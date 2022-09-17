package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {

	recebe := "91468384066"
	agrupa(recebe)

	troca := make([]int, 11)
	copy(troca, cpfint)

	cpfgiro = troca
	calc(cpfgiro)

	valido = compara(cpfint, cpfOk)
	fmt.Println(valido)

	if valido == true {
		formata(cpfOk)
	} else {
		formata(cpfint)
	}

	Resp = Resposta{valido, formatado}

	fmt.Println(Resp)
	fmt.Println("Iniciando o servidor Rest com Go")

	HandleRequest()
}

var cpfint []int
var cpfgiro []int
var cpfOk []int
var cpfForm []string

var valido bool
var formatado string

type Resposta struct {
	Valido    bool   `json:"Valido"`
	Formatado string `json:"Formatado"`
}

var Resp Resposta

func agrupa(cpfSt string) []int {
	for _, num := range cpfSt {
		intnum, err := strconv.Atoi(string(num))
		if err != nil {
			fmt.Println("Erro:", nil)
		}
		cpfint = append(cpfint, intnum)
	}
	return cpfint
}

func calc(x []int) []int {
	total := 0
	mult := 10

	cpfOk = append(x[:9])

	for _, v := range cpfOk {
		total += v * mult
		mult--
	}

	res1 := 11 - (total % 11)
	switch {
	case res1 >= 10:
		d1 := 0
		cpfOk = append(cpfOk, d1)
	case res1 < 10:
		d1 := res1
		cpfOk = append(cpfOk, d1)
	}

	mult = 11
	total = 0

	for _, v := range cpfOk {
		total += v * mult
		mult--
	}

	res2 := 11 - (total % 11)
	switch {
	case res2 >= 10:
		d1 := 0
		cpfOk = append(cpfOk, d1)
	case res2 < 10:
		d1 := res2
		cpfOk = append(cpfOk, d1)
	}

	return cpfOk
}

func compara(cpfi []int, cpfOki []int) bool {

	cpfs := make([]string, len(cpfi))
	for i, x := range cpfi {
		cpfs[i] = strconv.Itoa(x)
	}

	cpfOks := make([]string, len(cpfOki))
	for j, k := range cpfOki {
		cpfOks[j] = strconv.Itoa(k)
	}

	cpfs2 := strings.Join(cpfs, " ")
	cpfOks2 := strings.Join(cpfOks, " ")

	if cpfs2 != cpfOks2 {
		return false //fmt.Println("CPF INVÁLIDO!")
	} else {
		return true //fmt.Println("CPF OK!")
	}
}

func formata(cpfOki []int) {

	cpfOks := make([]string, len(cpfOki))
	for j, k := range cpfOki {
		cpfOks[j] = strconv.Itoa(k)
	}
	cpfOks = append(cpfOks, ".", "-", ".")

	cpfForm = append(cpfForm, cpfOks[:3]...)
	cpfForm = append(cpfForm, cpfOks[11:12]...)
	cpfForm = append(cpfForm, cpfOks[3:6]...)
	cpfForm = append(cpfForm, cpfOks[11:12]...)
	cpfForm = append(cpfForm, cpfOks[6:9]...)
	cpfForm = append(cpfForm, cpfOks[12:13]...)
	cpfForm = append(cpfForm, cpfOks[9:11]...)

	formatado = strings.Join(cpfForm, " ")
	fmt.Println(formatado)
}

func HandleRequest() {
	http.HandleFunc("/valida-cpf", Retorna)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Retorna(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Resp)
}
