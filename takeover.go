package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func banner() {
	fmt.Println(`  ________                                                                   
 /__  ___/                                                           
   / /   ___     / ___      ___      ___              ___      __    
  / /  //   ) ) //\ \     //___) ) //   ) ) ||  / / //___) ) //  ) ) 
 / /  //   / / //  \ \   //       //   / /  || / / //       //       
/ /  ((___( ( //    \ \ ((____   ((___/ /   ||/ / ((____   //        `)
}

func getCNAME(domain string, results chan<- string) {
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		log.Println(err)
	}

	results <- cname
}

func main() {
	fileName := flag.String("file", "", "The input file containing a list of domain names")
	perf := flag.Bool("perf", false, "Enable performance measurement")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Error: Please provide an input file using the -file flag")
		flag.Usage()
		os.Exit(1)
	}

	banner()
	fmt.Println("\033[0;32m  @sulemanmalik_3                         		 v1\033[0m")
	fmt.Println("-------------------------------------------------------------")

	var lines []string

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalLines := len(lines)

	resultsChan := make(chan string, totalLines)
	var wg sync.WaitGroup

	start := time.Now()
	for _, domain := range lines {
		wg.Add(1)
		go func(domain string) {
			defer wg.Done()
			getCNAME(domain, resultsChan)
		}(domain)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var cnames []string
	for cname := range resultsChan {
		cnames = append(cnames, cname)
	}
	elapsed := time.Since(start)

	fmt.Println("\n\n")
	for i := 0; i < totalLines; i++ {
		fmt.Printf("%d - %s --> %s\n", i+1, lines[i], cnames[i])
	}

	fmt.Println("-------------------------------------------------------------")

	if *perf {
		fmt.Printf("Time taken: %s\n", elapsed)
	}
}
