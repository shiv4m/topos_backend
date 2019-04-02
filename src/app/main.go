package main

import (
	_ "awesomeProject/src/github.com/lib/pq"
	"database/sql"
	_ "database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type buildings []struct {
	BaseBbl    string    `json:"base_bbl"`
	Bin        string    `json:"bin"`
	CnstrctYr  string    `json:"cnstrct_yr"`
	FeatCode   string    `json:"feat_code"`
	Groundelev string    `json:"groundelev"`
	Heightroof string    `json:"heightroof"`
	Lststatype string    `json:"lststatype,omitempty"`
	ShapeArea  string    `json:"shape_area"`
	ShapeLen   string    `json:"shape_len"`
}


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, string("Note: All the analysis is based on 1000 records in the database - https://data.cityofnewyork.us/resource/mtik-6c5q.json\n\n"))
	borCount(w)
}
func avg(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Note: All the analysis is based on 1000 records in the database - https://data.cityofnewyork.us/resource/mtik-6c5q.json\n\n")
	average(w)
}
func avglen(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Note: All the analysis is based on 1000 records in the database - https://data.cityofnewyork.us/resource/mtik-6c5q.json\n\n")
	averagelen(w)
}
func featcode(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Note: All the analysis is based on 1000 records in the database - https://data.cityofnewyork.us/resource/mtik-6c5q.json\n\n")
	feat_code(w)
}
func feat_code(w http.ResponseWriter) {
	connStr := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		fmt.Println(err)
	}
	boroughs := map[string]string{
		"2100":"Building Constructed",
		"5100":"Building Under Construction",
		"5110":"Garage",
		"2110":"Skybridge",
		"1001":"Gas Station Canopy",
		"1002":"Storage Tank",
		"1003":"Placeholders",
		"1004":"Auxiliary Structure",
		"1005":"Temporary Structure",
	}

	m := map[string]int{
	"Building Constructed" : 0,
	"Building Under Construction" : 0,
	"Garage": 0,
	"Skybridge" : 0,
	"Gas Station Canopy" : 0,
	"Storage Tank" : 0,
	"Placeholders" : 0,
	"Auxiliary Structure" : 0,
	"Temporary Structure" : 0,
	}
	roww := "SELECT (featcode) FROM buildDb LIMIT $1"
	var c string

	r, _:= db.Query(roww, 1000) //querying 1000 rows
	for r.Next(){
	_= r.Scan(&c)
	m[boroughs[c]] += 1
	}

	fmt.Fprintf(w, "Buildings Constructed : %d\n", m["Building Constructed"]) //printing on http
	fmt.Fprintf(w, "Buildings under construction: %d\n", m["Building Under Construction"])
	fmt.Fprintf(w, "Garages: %d\n", m["Garage"])
	fmt.Fprintf(w, "Skybridges: %d\n", m["Skybridge"])
	fmt.Fprintf(w, "Gas Stations : %d\n", m["Gas Station Canopy"])
	fmt.Fprintf(w, "Storage Tanks : %d\n", m["Storage Tank"])
	fmt.Fprintf(w, "Placeholder : %d\n", m["Placeholders"])
	fmt.Fprintf(w, "Auxiliary Structures : %d\n", m["Auxiliary Structure"])
	fmt.Fprintf(w, "Temporary Structures : %d\n", m["Temporary Structure"])
	db.Close()
}
func averagelen(w http.ResponseWriter){
	connStr := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		fmt.Println(err)
	}
	roww := "SELECT (heightroof) FROM buildDb LIMIT $1"
	var c float64
	a := 0.0
	r, _:= db.Query(roww, 1000) //querying 1000 rows
	for r.Next(){
		_ = r.Scan(&c)
		a += c
	}
	a = a/1000
	fmt.Fprintf(w, "Average length of the buildings in NYC %f\n", a)
}
func average(w http.ResponseWriter){
	connStr := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		fmt.Println(err)
	}
	roww := "SELECT (shapearea) FROM buildDb LIMIT $1"
	var c float64
	a := 0.0
	r, _:= db.Query(roww, 1000) //querying 1000 rows
	for r.Next(){
		_ = r.Scan(&c)
		a += c
	}
	a = a/1000

	fmt.Fprintf(w, "Average area of the buildings in NYC %f\n", a)
}

func borCount(w http.ResponseWriter) {
	connStr := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil{
		fmt.Println(err)
	}
	boroughs := map[string]string{
		"1": "Manhattan",
		"2": "The Bronx",
		"3": "Brooklyn",
		"4": "Queens",
		"5": "Staten Island",
	}

	m := map[string]int{
		"Manhattan" : 0,
		"The Bronx" : 0,
		"Brooklyn": 0,
		"Queens" : 0,
		"Staten Island" : 0,
	}
	roww := "SELECT (bin) FROM buildDb LIMIT $1"
	var c string

	r, _:= db.Query(roww, 1000) //querying 1000 rows
	for r.Next(){
		_= r.Scan(&c)
		m[boroughs[string([]rune(c)[0])]] += 1 //count buildings based on the first number of the bin column where 1 is Manhattan, 2 is bronx and so on
	}

	fmt.Fprintf(w, "Number of buildings in Manhattan borough : %d\n", m["Manhattan"]) //printing on http
	fmt.Fprintf(w, "Number of buildings in The Bronx borough: %d\n", m["The Bronx"])
	fmt.Fprintf(w, "Number of buildings in Brooklyn borough: %d\n", m["Brooklyn"])
	fmt.Fprintf(w, "Number of buildings in Queens borough: %d\n", m["Queens"])
	fmt.Fprintf(w, "Number of buildings in Staten Island : %d\n", m["Staten Island"])
	db.Close()
}

func main() {
	connStr := "user=postgres dbname=postgres password=123456 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("success")
	stemt := "DROP TABLE IF EXISTS buildDb;" //drop table if it already exists.
	ls1, err := db.Prepare(stemt)
	ls1.Exec()
	stem1 := "CREATE TABLE buildDb(BaseBbl varchar(255),Bin varchar(255),CnstrctYr varchar(255),FeatCode varchar(255),Groundelev varchar(255),Heightroof varchar(255),Lststatype varchar(255),ShapeArea varchar(255),ShapeLen varchar(255))"
	ls2, err := db.Prepare(stem1)
	ls2.Exec()
	statement := "INSERT INTO buildDb(BaseBbl, Bin, CnstrctYr, FeatCode, Groundelev, Heightroof, Lststatype, ShapeArea, ShapeLen) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err = db.Prepare(statement)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Printf(os.Getwd())
	jsonData, err := os.Open("building.json") //all the footprint data stored in a json file
	if err != nil{
		fmt.Println(err)
	}
	defer jsonData.Close()
	byteValue, _ := ioutil.ReadAll(jsonData)
	var result buildings
	json.Unmarshal([]byte(byteValue), &result)
	for i:=0; i < len(result); i++ {
		_, err := db.Exec(statement, result[i].BaseBbl, result[i].Bin, result[i].CnstrctYr, result[i].FeatCode, result[i].Groundelev, result[i].Heightroof, result[i].Lststatype, result[i].ShapeArea, result[i].ShapeLen)
		if err !=nil{
			panic(err)
		}
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/avgarea/", avg)
	http.HandleFunc("/avglen/", avglen)
	http.HandleFunc("/type/", featcode)
	http.ListenAndServe(":8080", nil)
	db.Close()
}
