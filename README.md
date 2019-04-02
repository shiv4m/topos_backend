# topos_backend
A Go API to insert and query with PostgreSQL


Instructions to run the code - 

download the "topos_backend" folder from github.

Assumptions are - 
1. Postgresql is installed and configured
2. Go is installed with proper GOROOT and GOPATH

Place the topos_backend folder in the directory_to_go/src

Running the code - 
1. Change the configurations of postgresql database accordingly on lines 44, 93, 110, 129, 169.
2. run the go file from "topos_backend/src/app/main.go
3. open browser and type localhost:8080/
4. Browse with localhost:8080/,
               localhost:8080/avgarea,
               localhost:8080/avglen,
               localhost:8080/type
   for different analysis. 
   
A demo.mp4 is also attached.


https://www.github.com/lib/pq is used as postgresql driver for go.
