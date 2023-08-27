// cmd/profiler/main.go

package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	"github.com/rafaelhl/go-pprof-schedule/internal/profiling"
	"github.com/rafaelhl/go-pprof-schedule/web"
)

var cpuprofile = flag.String("cpuprofile", "", "Arquivo para salvar o CPU profiling")
var appURL = flag.String("appurl", "", "URL da aplicação externa com profiling habilitado")

func main() {
	var period uint
	var date string
	var timeStr string

	var rootCmd = &cobra.Command{
		Use:   "profiler",
		Short: "Agenda a execução do profiling",
		Run: func(cmd *cobra.Command, args []string) {
			// Inicie o servidor HTTP para profiling em segundo plano.
			go func() {
				log.Println(http.ListenAndServe("localhost:6060", nil))
			}()

			// Agende o profiling com o período, data e horário especificados pelo usuário.
			profiling.ScheduleProfiler(time.Duration(period)*time.Minute, date, timeStr)

			// Inicie o servidor HTTP para a interface HTML do profiling.
			web.StartServer()
		},
	}

	// Defina as flags para o período, data e horário do profiling.
	rootCmd.Flags().UintVarP(&period, "period", "p", 5, "Período para executar o profiling (em minutos)")
	rootCmd.Flags().StringVarP(&date, "date", "d", "", "Data para executar o profiling (formato YYYY-MM-DD)")
	rootCmd.Flags().StringVarP(&timeStr, "time", "t", "00:00", "Horário para executar o profiling (formato HH:mm)")
	rootCmd.Flags().StringVar(&*cpuprofile, "cpuprofile", "", "Arquivo para salvar o CPU profiling")
	rootCmd.Flags().StringVar(&*appURL, "appurl", "", "URL da aplicação externa com profiling habilitado")

	// Execute o comando root.
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
