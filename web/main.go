// web/main.go

package web

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	baseDir, _    = os.Getwd()
	templateIndex = template.Must(template.New("index.html").ParseFiles(fmt.Sprintf("%s/web/templates/index.html", baseDir)))
)

func StartServer(profilesDir string) {
	// Inicie o servidor HTTP para profiling em segundo plano.
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Inicie o servidor HTTP para a interface HTML do profiling.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		handleHTML(profilesDir, writer, request)
	})
	// Adicione a rota para o profiling do CPU.
	http.Handle("/debug/pprof/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not supported yet", http.StatusNotImplemented)
	}))
	http.ListenAndServe(":8080", nil)
}

func handleHTML(dir string, w http.ResponseWriter, _ *http.Request) {
	profilesDir, err := os.ReadDir(dir)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var allCollected bytes.Buffer
	for _, profile := range profilesDir {
		if profile.IsDir() {
			continue
		}

		result, err := os.ReadFile(fmt.Sprintf("%s/%s", dir, profile.Name()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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

	if err := templateIndex.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Erro executando template: %s", err.Error()), http.StatusInternalServerError)
		return
	}
}
