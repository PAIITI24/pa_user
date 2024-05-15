package helper

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var (
	ProdukStoragePublicURL = ""
	ObatStoragePublicURL   = ""
	ObatBucketName         = "gambarobat"
	ProdukBucketName       = "gambarproduk"
	endpoint               = ""
	accessKeyID            = ""
	secretKey              = ""
	useSSL                 = true
)

func S3Connect() *minio.Client {
	Client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return Client
}
