package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const timeDelay = 5

func main(){

	intro()
  
	for {
		menu()

		comando := readCommand()	
	
		switch comando {
		case 1:
			start()
		case 2:
			fmt.Println("Exibindo Logs...")
			osLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Desculpe! Não consigo reconhecer este comando. Saindo...")
			os.Exit(-1)
		}
	}
}

func intro(){
	nome := "Luís"
	version := 1.19
	idade := 21

	reflect.TypeOf(nome)

	fmt.Println("Olá Sr.", nome)
	fmt.Println("O presente programa encontra-se na versão", version)
	fmt.Println("Com os meus conhecimentos, descobri sua idade, que é:", idade)
}

func menu(){
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func readCommand() int {
	var comando int
	fmt.Scan( &comando )
	fmt.Println("O comando selecionado foi:", comando)
	fmt.Println("")

	return comando
}

func start(){
	fmt.Println("Iniciando Monitoramento...")

	sites:= leSitesDoArquivo()

	for i:= 0; i < monitoramentos ; i++ {
		for i, site := range sites {
			fmt.Println("Estou passando no indice:", i, "no site:", site)
			testSite(site)
		}
		time.Sleep(timeDelay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string){
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	
	if resp.StatusCode == 200 {
		fmt.Println("O perfil de Luís foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("O perfil:", site, "está com problemas! Error/Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("Sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool){
	
	arquivo, err := os.OpenFile("Log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func osLogs(){
	arquivo, err := ioutil.ReadFile("Log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}