//hello.go
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

func exibeIntroducao() {
	//Go can infer about the type.
	//var nome string = "Arthur"
	nome := "Arthur"
	versao := 1.1
	fmt.Println("Olá, sr. " + nome)
	fmt.Println("Este programa está na versão ", versao, reflect.TypeOf(versao))
}

//Retorna int
func leComando() int {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Log")
	fmt.Println("0 - Sair do programa")

	var comando int
	fmt.Scan(&comando)

	return comando
}

func exibeNomes() {
	nomes := []string{"Douglas", "Arthur", "Isadora", "Juliano", "Lilian", "Valentina"} //ArrayList, the SLICE!
	nomes = append(nomes, "Vó Maria")
	fmt.Println(nomes, len(nomes), cap(nomes)) //Slices usually double the 'cap'acity of the array.
}

func testaSite(sites string) {
	resp, _ := http.Get(sites)
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", sites, "foi carregado com sucesso")
		registraLog(sites, true)
	} else {
		fmt.Println("Site: ", sites, "está com problemas. Status code: ", resp.StatusCode)
		registraLog(sites, false)
	}
}

// func IniciarMonitoramento() {
// 	fmt.Println("Monitorando...")
// 	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}
// 	for i := 0; i < 5; i++ {
// 		for _, sites := range sites {
// 			testaSite(sites)
// 		}
// 		time.Sleep(5 * time.Second)
// 	}

// }

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

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

func IniciarMonitoramento() {
	fmt.Println("Monitorando...")
	for i := 0; i < 5; i++ {
		for _, sites := range leSitesDoArquivo() {
			testaSite(sites)
		}
		time.Sleep(5 * time.Second)
	}
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site +
		" - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}

func main() {
	//exibeIntroducao()
	// fmt.Println("O endereço de comando é", &comando), shows the memory address.
	// fmt.Print(comando)
	exibeNomes()

	for { //THERE'S NO WHILE!!!! WE CAN USE JUST for. This is equivalent to "while(true){}"
		switch leComando() {
		case 1:
			IniciarMonitoramento()
			//break //it's not necessary
		case 2:
			imprimeLogs()
			//break
		case 0:
			fmt.Print("Saindo...")
			os.Exit(0)
		default:
			fmt.Print("Não entendi...")
			os.Exit(-1)
			//break

		}
	}

}
