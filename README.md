# Go Pprof Schedule

This is a profiling application that allows you to schedule and collect profiling data both locally and from an external application via HTTP. It also includes a web interface to visualize the results.

## Prerequisites

Make sure you have Go (1.20 or above) installed on your system. You will also need Docker if you want to run the application as a Docker image.

## Steps to Run the Application

### 1. Clone the Repository

```bash
git clone https://github.com/rafaehl/go-pprof-schedule.git
cd go-pprof-schedule
```

### 2. Run the Web Application

```bash
go run cmd/profiler/main.go --cpuprofile cpu.prof --appurl http://external-app-url:port
```

Access `http://localhost:8080` in your browser to view the profiling web interface.

### 3. Run the Command-line Application (CMD)

```bash
go run cmd/profiler/main.go --cpuprofile cpu.prof --appurl http://external-app-url:port
```

This will run the application in command-line mode.

### 4. Run the Application using Docker

```bash
docker build -t go-pprof-schedule .
docker run -p 8080:8080 go-pprof-schedule
```

This will create a Docker image and run the application inside a container.

## Configuration

- `--cpuprofile`: Specify the filename to save the CPU profile (optional).
- `--appurl`: Specify the URL of the external application with profiling enabled (optional).

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
