// Command mimefold 是邮件头解析的场景入口。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/mime-fold/mimefold"
)

func emit(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	sample := flag.String("sample", "", "fold|case|multi|default|work")
	flag.Parse()

	for _, s := range mimefold.Samples() {
		if s.Name == *sample {
			emit(s.Run())
			return
		}
	}
	if *sample == "work" {
		emit(mimefold.Work(2000))
		return
	}
	fmt.Fprintln(os.Stderr, "未知场景："+*sample)
	os.Exit(2)
}
