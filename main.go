package main

import (
	"flag"
)

func main() {

	// c := client.NewClient("http://localhost:3000")
	// resp, err := c.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(resp)

	// return

	listenAddr := flag.String("listenAddr", ":3000", "port the api listens on.")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	jsonAPIServer := NewJSONAPIServer(*listenAddr, svc)
	jsonAPIServer.Run()

}
