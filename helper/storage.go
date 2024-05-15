package helper

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var (
	StoragePublicURL = "https://pub-a6d879b3e38f4fe0b2f3beb986236466.r2.dev"
	ObatBucketName   = "gambarobat"
	ProdukBucketName = "gambarproduk"
	endpoint         = "23e6cc54af99f15b0d42f07eb55e0db7.r2.cloudflarestorage.com"
	accessKeyID      = "bd59c36fa19d23d8520f17035326c060"
	secretKey        = "e9e958d5d105a62c9488c1d7afd25a5ccdae344689a071ecdab11748a8af4c33"
	useSSL           = true
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
