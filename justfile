alias b := build
alias r := run

default:
    @just --list

run:
    go run main.go

build:
    go build -o build/main main.go