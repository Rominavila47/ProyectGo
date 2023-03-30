package log

import (
	"log"
	"os"
)

type Logger struct {
}

func Log(*Logger) {
	logger, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Close()

	log.SetOutput(logger)

	for i := 0; i <= 10; i++ {
		log.Println("error linea %v", i)
	}
	return (nil)
}

func Log2(){
	var 
	buf := new(bytes.Buffer)
	logger := log.New(buf, "logger: ", log.Lshortfile)

	info := func(info string){
		logger.Output(2, info)
	}

	fmt.Println(logger)
}