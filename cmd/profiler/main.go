// cmd/profiler/main.go

package main

import (
	"log"
	"time"

	"github.com/spf13/cobra"

	"github.com/rafaelhl/go-pprof-schedule/internal/profiling"
	"github.com/rafaelhl/go-pprof-schedule/web"
)

func main() {
	var (
		period     uint
		date       string
		timeStr    string
		cpuProfile string
		appURL     string
	)

	var rootCmd = &cobra.Command{
		Use:   "profiler",
		Short: "Agenda a execução do profiling",
		Run: func(cmd *cobra.Command, args []string) {

			// Agende o profiling com o período, data e horário especificados pelo usuário.
			profilesDir := profiling.ScheduleProfiler(time.Duration(period)*time.Minute, date, timeStr, cpuProfile, appURL)

			// Inicie o servidor HTTP para a interface HTML do profiling.
			web.StartServer(profilesDir)
		},
	}

	// Defina as flags para o período, data e horário do profiling.
	rootCmd.Flags().UintVarP(&period, "period", "p", 5, "Período para executar o profiling (em minutos)")
	rootCmd.Flags().StringVarP(&date, "date", "d", time.Now().Format(time.DateOnly), "Data para executar o profiling (formato YYYY-MM-DD)")
	rootCmd.Flags().StringVarP(&timeStr, "time", "t", "00:00", "Horário para executar o profiling (formato HH:mm)")
	rootCmd.Flags().StringVar(&cpuProfile, "cpuprofile", "", "Arquivo para salvar o CPU profiling")
	rootCmd.Flags().StringVar(&appURL, "appurl", "", "URL da aplicação externa com profiling habilitado")

	// Execute o comando root.
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
