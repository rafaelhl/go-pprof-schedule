// internal/profiling/profiler.go

package profiling

import (
	"log"
	"time"
)

func ScheduleProfiler(period time.Duration, execDateStr, execTimeStr string) {
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
		runProfiler()

		// Agende o próximo profiling com o período especificado.
		ticker := time.NewTicker(period)
		for range ticker.C {
			runProfiler()
		}
	})
}

func runProfiler() {
	// Coloque aqui a lógica específica do profiling que você deseja executar.
	// Por exemplo, pode ser algo como coletar os dados do programa e salvá-los em um arquivo.
	// Você pode usar as funções do pacote "pprof" para coletar dados como CPU profiling, memory profiling, etc.
	log.Println("Profiling executado.")
}
