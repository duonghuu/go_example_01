package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func createS3Client(region string) *s3.Client {
	// Load the default AWS configuration (reads credentials and region from standard sources)
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	// Create a new S3 service client
	client := s3.NewFromConfig(cfg)
	return client
}

func S3_client() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		s3Client := createS3Client("ap-southeast-1")
		log.Println(s3Client.ListBuckets(c, &s3.ListBucketsInput{}))
		// c.JSON(http.StatusOK, gin.H{
		// 	"message": "Hello from Gin!",
		// })
	})
	router.Run(":8080")
}
