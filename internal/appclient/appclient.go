// internal/appclient/appclient.go

package appclient

import (
	"bytes"
	"net/http"
)

func CollectAppProfile(appURL string) string {
	// Coleta os dados de profiling da aplicação externa via HTTP.
	resp, err := http.Get(appURL + "/debug/pprof/profile")
	if err != nil {
		return "Erro ao coletar o perfil da aplicação externa: " + err.Error()
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "Erro ao coletar o perfil da aplicação externa. Status: " + resp.Status
	}

	// Leitura dos dados do corpo da resposta HTTP e retorna como string.
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "Erro ao ler os dados do perfil da aplicação externa: " + err.Error()
	}

	return buf.String()
}
