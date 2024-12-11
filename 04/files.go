package main

import _ "embed"

//go:embed example.txt
var example []byte

//go:embed input.txt
var input []byte
