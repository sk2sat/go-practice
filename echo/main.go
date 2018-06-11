package main

// importは複数やるときはこうする
import (
	"os"
	"flag" // コマンドラインオプションパーサ
)


var omitNewline = flag.Bool("n", false, "don't print final newline")

// 定数
const(
	Space	= " "
	Newline	= "\n"
)

func main(){
	flag.Parse()
	var s string = "" // string型の変数s

	// "i := 0"は"var i := 0"の省略形
	for i := 0; i< flag.NArg(); i++ {
		if i > 0 {
			s += Space
		}
		s += flag.Arg(i)
	}
	if ! *omitNewline {
		s += Newline
	}
	os.Stdout.WriteString(s)
}
