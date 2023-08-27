// web/main.go

package web

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/rafaelhl/go-pprof-schedule/internal/appclient"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "Arquivo para salvar o CPU profiling")
	appURL     = flag.String("appurl", "", "URL da aplicação externa com profiling habilitado")
)

func StartServer() {
	// Inicie o servidor HTTP para profiling em segundo plano.
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Inicie o servidor HTTP para a interface HTML do profiling.
	http.HandleFunc("/", handleHTML)
	// Adicione a rota para o profiling do CPU.
	http.Handle("/debug/pprof/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// TODO:
	}))
	http.ListenAndServe(":8080", nil)
}

func handleHTML(w http.ResponseWriter, r *http.Request) {
	// Execute o CPU profiling
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			http.Error(w, "Erro ao criar o arquivo de CPU profiling", http.StatusInternalServerError)
			return
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Renderiza o template HTML usando os dados do profiling
	data := struct {
		Timestamp  time.Time
		AppProfile string
	}{
		Timestamp:  time.Now(),
		AppProfile: appclient.CollectAppProfile(*appURL),
	}

	t, err := template.New("profiling").ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao renderizar o template", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Erro ao executar o template", http.StatusInternalServerError)
		return
	}
}
