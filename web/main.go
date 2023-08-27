// web/main.go

package web

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"
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

	profilesDir, err := os.ReadDir(fmt.Sprintf("%s/go-pprof-schedule", os.TempDir()))
	if err != nil {
		http.Error(w, "Erro ao ler os profiles", http.StatusInternalServerError)
		return
	}

	var allCollected bytes.Buffer
	for _, profile := range profilesDir {
		if profile.IsDir() {
			continue
		}

		result, err := os.ReadFile(profile.Name())
		if err != nil {
			http.Error(w, "Erro ao ler o profile", http.StatusInternalServerError)
			return
		}

		allCollected.Write(result)
	}

	// Renderiza o template HTML usando os dados do profiling
	data := struct {
		Timestamp  time.Time
		AppProfile string
	}{
		Timestamp:  time.Now(),
		AppProfile: allCollected.String(),
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
