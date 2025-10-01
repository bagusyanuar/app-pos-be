package config

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func NewMinioClient(viper *viper.Viper) *minio.Client {
	host := viper.GetString("MINIO_HOST")
	port := viper.GetString("MINIO_PORT")
	endpoint := fmt.Sprintf("%s:%s", host, port)
	accessKeyID := viper.GetString("MINIO_USERNAME")
	secretAccessKey := viper.GetString("MINIO_PASSWORD")
	useSSL := viper.GetBool("MINIO_SSL")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatal("[Minio] ❌ Gagal koneksi ke Minio")
	}

	log.Printf("[Minio] ✅ Koneksi berhasil ke Minio")
	return minioClient
}
