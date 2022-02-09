# pizzabot
A technical challenge I did as part of a job interview. The idea is that you should be able to pass pizzabot dimensions
of a grid, and a series of coordinates, and pizzabot will print out directions to each coordinate pair and deliver
pizza. 
## Building pizzabot

### Install Go

Install Go for your platform by following the instructions
[here](https://golang.org/doc/install#install)

### Compiling pizzabot

With Go installed and the pizzabot source in your Go workspace, you
should be able to build pizzabot by running the following command
inside the `pizzabot` folder:

```bash
go build -a -o pizzabot cmd/main.go
```

This will create a `pizzabot` binary inside the current folder

## Running pizzabot

Run pizzabot with:

```bash
./pizzabot "5x5 (0, 0) (1, 3) (4, 4) (4, 2) (4, 2) (0, 1) (3, 2) (2, 3) (4, 1)"
```

## Testing pizzabot

To execute unit tests, run:

```bash
go test ./...
```
