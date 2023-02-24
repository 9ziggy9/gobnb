GC=go

all: main.go
	$(GC) build main.go

run: main.go
	$(GC) run main.go

clean:
	rm main
