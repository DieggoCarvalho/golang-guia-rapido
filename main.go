package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Carro struct {
	Modelo string
	Ano    int
}

type Veiculo interface {
	Andar()
}

func (c Carro) Andar() {
	fmt.Println("O carro", c.Modelo, " está andando")
}

func VemAndarComigo(v Veiculo) {
	v.Andar()
}

func main() { // thread 1
	//println("Hello, Wolrd!")
	// println(evento.Soma(1,2))
	//a := 1 //inferir que o 1 é um int -- inferência de tipos
	//b := "String"
	//var c int
	//c = 1
	//http.HandleFunc("/", HelloWorld)
	//http.ListenAndServe(":8888", nil)
	go contador(10) // thread 2
	go contador(10) // thread 3
	contador(10)

	canal := make(chan string) //vazio -- canal preenchido

	go func() { // thread 4
		canal <- "Olá, canal!"
	}()

	fmt.Println(<-canal) // thread 1 -- canal esvaziado

	// load balancer
	// Gerenciado de Tarefas -- Balancer
	// Worker

	canal2 := make(chan int)
	qtdWorkers := 5
	// go worker(1, canal2)
	// go worker(2, canal2)
	for i := range qtdWorkers {
		go worker(i, canal2)
	}

	for i := range 10 {
		canal2 <- i // travado aqui até o canal ser esvaziado.
		//quando esvaziar, o loop continua.
	}

	carro1 := Carro{Modelo: "Fusca", Ano: 1969}
	fmt.Printf(carro1.Modelo, carro1.Ano)

	carro2 := Carro{Modelo: "BMW", Ano: 2000}
	fmt.Printf(carro2.Modelo, carro2.Ano)

	carro1.Andar()

	VemAndarComigo(carro1)

	db, err := sql.Open("sqllite3", "teste.db")

	if err != nil {
		panic(err)
	}
	defer db.Close()
	insertCarro(db, carro1)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func contador(qtd int) {
	for i := range qtd {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerID int, data chan int) {
	for x := range data { //esvazia o canal
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func insertCarro(db *sql.DB, carro Carro) {
	_, err := db.Exec("INSERT INTO carro (modelo, ano) VALUES (?, ?)", carro.Modelo, carro.Ano)
	if err != nil {
		panic(err)
	}
}
