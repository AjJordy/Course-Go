
### Commands

- Run: `go run main.go` 
- Build: `go build -o "hello.exe"`
- PProf: `go tool pprof -seconds 5 http://localhost:6060/debug/pprof/profile`
	- `top`
	- `list CPUIntensiveEndpoint`
- Test: `go test -bench=. -benchmen -memprofile mem.prof -cpuprofile cpu.prof -count 10 > 1.bench`
	- `benchstat 1.bench 2.bench`