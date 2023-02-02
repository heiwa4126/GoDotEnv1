package main

// an example code from https://github.com/joho/godotenv
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	// 0個以上のファイル名が引数にとれる。デフォルトは".env"
	// 既に存在している環境変数をオーバーライドしない
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")

	fmt.Printf("S3_BUCKET=%s\nSECRET_KEY=%s\n", s3Bucket, secretKey)
}
