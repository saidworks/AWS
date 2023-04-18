package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
)

func main() {
	region := "us-east-1"
	config := aws.Config{Region: aws.String(region), CredentialsChainVerboseErrors: aws.Bool(true)}

	movies, err := readMovies("movies.json")
	if err != nil {
		log.Fatal(err)
	}
	for _, movie := range movies {
		fmt.Println("Inserting : " + movie.Name)
		err = insertMovie(config, movie)
		if err != nil {
			log.Fatal(err)
			fmt.Println(err)
		}
	}
}
