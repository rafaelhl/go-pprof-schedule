// internal/profiling/profiler.go

package profiling

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/rafaelhl/go-pprof-schedule/internal/appclient"
)

func ScheduleProfiler(period time.Duration, execDateStr, execTimeStr, cpuprofile, appURL string) {
	// Analise a data e horário de execução fornecidos pelo usuário.
	execDateTimeStr := execDateStr + " " + execTimeStr
	execDateTime, err := time.Parse("2006-01-02 15:04", execDateTimeStr)
	if err != nil {
		log.Fatalf("Erro ao analisar a data e horário de execução: %v", err)
	}

	// Calcule a duração até a próxima data e horário de execução.
	now := time.Now()
	nextExecDateTime := execDateTime
	if now.After(nextExecDateTime) {
		nextExecDateTime = nextExecDateTime.Add(24 * time.Hour)
	}
	durationUntilExec := nextExecDateTime.Sub(now)

	// Agende o primeiro profiling na próxima data e horário de execução especificados.
	time.AfterFunc(durationUntilExec, func() {
		// Execute o profiling.
		runProfiler(cpuprofile, appURL)

		// Agende o próximo profiling com o período especificado.
		ticker := time.NewTicker(period)
		for range ticker.C {
			runProfiler(cpuprofile, appURL)
		}
	})
}

func runProfiler(appURL, cpuProfile string) {
	appProfile := appclient.CollectAppProfile(appURL)
	resultProfile := fmt.Sprintf("/profile-%v.prof", time.Now().UnixNano())
	if cpuProfile != "" {
		resultProfile = cpuProfile
	}

	err := os.WriteFile(fmt.Sprintf("%s/go-pprof-schedule/%s", os.TempDir(), resultProfile), []byte(appProfile), fs.ModeAppend)
	if err != nil {
		log.Fatalf("Erro ao salvar o profile: %v", err)
	}
}
