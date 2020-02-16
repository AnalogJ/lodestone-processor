package api

import (
	"fmt"
	"github.com/analogj/lodestone-processor/pkg/model"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func GenerateStoragePath(event model.S3Event) (string, string, error) {
	/*
		{
			"Records": [{
				"eventVersion": "2.0",
				"eventSource": "lodestone:publisher:fs",
				"awsRegion": "",
				"eventTime": "2019-11-16T23:46:21.1467633Z",
				"eventName": "s3:ObjectRemoved:Delete",
				"userIdentity": {
					"principalId": "lodestone"
				},
				"requestParameters": {
					"sourceIPAddress": "172.19.0.5"
				},
				"responseElements": {},
				"s3": {
					"s3SchemaVersion": "1.0",
					"configurationId": "Config",
					"bucket": {
						"name": "documents",
						"ownerIdentity": {
							"principalId": "lodestone"
						},
						"arn": "arn:aws:s3:::documents"
					},
					"object": {
						"key": "filetypes/fIoiDm",
						"size": 0,
						"urlDecodedKey": "",
						"versionId": "1",
						"eTag": "",
						"sequencer": ""
					}
				}
			}]
		}
	*/
	bucketName := event.Records[0].S3.Bucket.Name
	documentPath := event.Records[0].S3.Object.Key

	return bucketName, documentPath, nil
}

func GenerateThumbnailStoragePath(storagePath string) string {
	storagePath = storagePath + ".jpg"
	return storagePath
}

//File CRUD Operations
func CreateFile(apiEndpoint *url.URL, storageBucket string, storagePath string, localFilepath string) error {

	localFile, err := os.Open(localFilepath)
	if err != nil {
		return err
	}
	defer localFile.Close()

	//manipulate the path
	apiEndpoint.Path = fmt.Sprintf("/api/v1/storage/%s/%s", storageBucket, storagePath)

	resp, err := http.Post(apiEndpoint.String(), "binary/octet-stream", localFile)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func ReadFile(apiEndpoint *url.URL, storageBucket string, storagePath string, outputDirectory string) (string, error) {

	//secureProtocol := apiEndpoint.Scheme == "https"
	//
	//s3Client, err := minio.New(apiEndpoint.Host, os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_SECRET_KEY"), secureProtocol)
	//if err != nil {
	//	return "", err
	//}

	fileName := filepath.Base(storagePath)
	localFilepath := filepath.Join(outputDirectory, fileName)

	localFile, err := os.Create(localFilepath)
	if err != nil {
		return "", err
	}
	defer localFile.Close()

	//manipulate the path
	apiEndpoint.Path = fmt.Sprintf("/api/v1/storage/%s/%s", storageBucket, storagePath)

	resp, err := http.Get(apiEndpoint.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	_, err = io.Copy(localFile, resp.Body)
	if err != nil {
		return "", err
	}

	return localFilepath, err
}

func DeleteFile(apiEndpoint *url.URL, storageBucket string, storagePath string) error {

	//manipulate the path
	apiEndpoint.Path = fmt.Sprintf("/api/v1/storage/%s/%s", storageBucket, storagePath)

	_, err := http.NewRequest(http.MethodDelete, apiEndpoint.String(), nil)
	return err
}
