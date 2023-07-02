package main

func main() {
	bc := NewBlockChain()

	bc.AddBlock("send 1 BTC to Ivan")

	bc.AddBlock("send 2 BTC to Ivan")

	bc.Print()

}
