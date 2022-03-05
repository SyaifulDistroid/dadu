package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var Pemain, Dadu, TotalDadu, Eval int
	var ListmapPemain []map[string]interface{}
	giliran := 0

	fmt.Print("Masukan Jumlah Pemain = ")
	fmt.Scan(&Pemain)
	fmt.Print("Masukan Jumlah Dadu = ")
	fmt.Scan(&Dadu)

	for i := 1; i <= Pemain; i++ {
		mapPemain := make(map[string]interface{})
		mapPemain["pemain"] = i
		mapPemain["poin"] = 0
		mapPemain["dadu"] = Dadu
		fmt.Printf("Pemain ke #%v Masukan %v Angka Dadu = ", mapPemain["pemain"], mapPemain["dadu"])
		fmt.Scan(&TotalDadu)
		mapPemain["listdadu"] = TotalDadu
		ListmapPemain = append(ListmapPemain, mapPemain)
	}

	for {
		fmt.Println("=============================")
		giliran++
		fmt.Printf("Giliran %v Lempar Dadu", giliran)
		fmt.Println()

		for _, v := range ListmapPemain {
			fmt.Printf("Pemain ke #%v (%v) = %v", v["pemain"], v["poin"], v["listdadu"])
			fmt.Println()
		}

		fmt.Println("=============================")
		fmt.Print("[1] Evaluasi")
		fmt.Println()
		fmt.Print("[2] Keluar")
		fmt.Println()
		fmt.Scan(&Eval)
		if Eval == 1 {
			ListmapPemain = Evaluasi(ListmapPemain)
			fmt.Println("====== Setelah Evaluasi =====")
			for _, v := range ListmapPemain {
				fmt.Printf("Pemain ke #%v (%v) = %v", v["pemain"], v["poin"], v["listdadu"])
				fmt.Println()
			}
			fmt.Println("=============================")

			var NewList []map[string]interface{}
			for _, v := range ListmapPemain {
				mapPemain := make(map[string]interface{})
				mapPemain["pemain"] = v["pemain"]
				mapPemain["poin"] = v["poin"]
				mapPemain["dadu"] = v["dadu"]
				fmt.Printf("Pemain ke #%v Masukan %v Angka Dadu = ", v["pemain"], v["dadu"])
				fmt.Scan(&TotalDadu)
				mapPemain["listdadu"] = TotalDadu
				NewList = append(NewList, mapPemain)
			}
			ListmapPemain = NewList
		} else {
			break
		}
	}
}

func Evaluasi(ListMap []map[string]interface{}) []map[string]interface{} {
	var NewList []map[string]interface{}
	var List []map[string]interface{}

	// Evaluasi if 6
	for _, v := range ListMap {
		valListDadu, _ := v["listdadu"].(int)
		valDadu, _ := v["dadu"].(int)
		valPoin, _ := v["poin"].(int)
		NewMap := make(map[string]interface{})
		if strings.ContainsAny(strconv.Itoa(valListDadu), "6") == true {
			NewMap["pemain"] = v["pemain"]
			NewMap["poin"] = valPoin + strings.Count(strconv.Itoa(valListDadu), "6")
			NewMap["listdadu"] = strings.ReplaceAll(strconv.Itoa(valListDadu), "6", "")
			NewMap["dadu"] = valDadu - 1
		} else {
			NewMap["pemain"] = v["pemain"]
			NewMap["poin"] = v["poin"]
			NewMap["dadu"] = v["dadu"]
			NewMap["listdadu"] = v["listdadu"]
		}
		NewList = append(NewList, NewMap)
	}

	//Evaluasi if 1
	for _, v := range NewList {
		valListDadu, _ := v["listdadu"].(int)
		valDadu, _ := v["dadu"].(int)
		NewMap := make(map[string]interface{})
		if strings.ContainsAny(strconv.Itoa(valListDadu), "1") == true {
			NewMap["pemain"] = v["pemain"]
			NewMap["poin"] = v["poin"]
			NewMap["listdadu"] = strings.ReplaceAll(strconv.Itoa(valListDadu), "1", "")
			NewMap["dadu"] = valDadu - 1
		} else {
			NewMap["pemain"] = v["pemain"]
			NewMap["poin"] = v["poin"]
			NewMap["dadu"] = v["dadu"]
			NewMap["listdadu"] = v["listdadu"]
		}
		List = append(List, NewMap)
	}

	return List
}
