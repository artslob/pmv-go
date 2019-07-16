package main

import (
	"fmt"
	"github.com/artslob/pmv-go/lab2"
)

func main() {
	var input = `
def func()
	do {
		a = a + b;
		if b > 10 then
			break;
		end
		do {
			b = b + a;
			if b == 10 then
				break;
			end
		}
		until b < 42;
		b = b * 2;
	} until a < 3;
	print(t);
end

def qwe()
	a = 1;
end
`
	head := lab2.ParseInputToCFG(input)
	printer := lab2.NewCfgPrinter()
	printer.Print(head)
	fmt.Print(printer.String())
}
