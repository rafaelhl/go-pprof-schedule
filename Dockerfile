# Use uma imagem base adequada para o seu aplicativo Go.
# Neste exemplo, usamos a imagem oficial do Go.
FROM golang:1.20

# Configuração do diretório de trabalho dentro do contêiner.
WORKDIR /app

# Copia o código fonte para o diretório de trabalho no contêiner.
COPY . .

# Compila o código Go e gera o binário do aplicativo.
RUN go build -o profiler ./cmd/profiler

# Define a variável de ambiente para o arquivo de CPU profiling (se necessário).
# Se você não quiser usar o CPU profiling, remova esta linha ou deixe-a vazia.
ENV CPU_PROFILE_FILE /app/cpu.prof

# Exponha a porta 8080 para acessar a interface web.
EXPOSE 8080

# Comando para iniciar o servidor web.
CMD ["./profiler"]
