package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// UploadFileToS3 tải lên một file cục bộ lên S3
func UploadFileToS3(bucketName, filePath, objectKey string) error {
	// 1. Tải cấu hình mặc định (từ biến môi trường hoặc file ~/.aws/credentials)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return fmt.Errorf("không thể tải cấu hình AWS: %w", err)
	}

	// 2. Tạo Client S3
	client := s3.NewFromConfig(cfg)

	// 3. Mở file cục bộ để đọc
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("không thể mở file cục bộ: %w", err)
	}
	defer file.Close() // Đảm bảo file được đóng sau khi hàm kết thúc

	// 4. Gọi PutObject để tải file lên
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName), // Tên Bucket S3
		Key:    aws.String(objectKey),  // Tên Object (file) trong S3
		Body:   file,                   // Nội dung file (dạng io.Reader)
	})

	if err != nil {
		return fmt.Errorf("lỗi khi tải file lên S3: %w", err)
	}

	fmt.Printf("✅ Đã tải file '%s' lên s3://%s/%s thành công.\n", filePath, bucketName, objectKey)
	return nil
}

func Upload_file_s3() {
	// THAY THẾ CÁC GIÁ TRỊ NÀY
	const bucketName = "s3-eric-bucket"            // Thay bằng tên Bucket S3 của bạn
	const localFilePath = "./data.txt"             // Thay bằng đường dẫn file cục bộ của bạn
	const s3ObjectKey = "uploads/2025/data_v1.txt" // Tên file bạn muốn nó có trong S3 (bao gồm cả thư mục ảo)

	// Tạo một file dummy để test nếu cần
	// err := os.WriteFile(localFilePath, []byte("Nội dung thử nghiệm tải lên S3."), 0644)
	// if err != nil {
	// 	fmt.Println("Lỗi khi tạo file test:", err)
	// 	return
	// }

	err := UploadFileToS3(bucketName, localFilePath, s3ObjectKey)
	if err != nil {
		fmt.Println("LỖI:", err)
		os.Exit(1)
	}
}
