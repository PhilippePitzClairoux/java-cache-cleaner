package main

import (
	"flag"
	"log"
	"os"
)

var (
	userHomeDir        = flag.String("user-home-dir", "", "User where the cache files are location")
	DirectoriesToClear = []string{
		".m2/repository",
		".gradle/daemon",
		".gradle/jdks",
		".gradle/build-scan-data",
		".gradle/caches",
		".gradle/wrapper",
		".gradle/native",
	}
)

func main() {
	flag.Parse()
	var currentUser string
	var err error

	if *userHomeDir == "" {
		currentUser, err = os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		currentUser = *userHomeDir
	}

	for _, dir := range DirectoriesToClear {
		currentDir := currentUser + "/" + dir
		err = os.RemoveAll(currentDir)
		if err != nil {
			log.Printf("Could not delete %s due to error : %s\n", dir, err)
		}
		log.Printf("%s has been deleted!\n", currentDir)
	}

	log.Println("Done clearing directories!")
}
