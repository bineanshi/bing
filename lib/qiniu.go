package lib

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

const (
	AccessKey = "d4qFWb7EDiDzC5wmKwV5wvQKpuHxzPpDK3tyMCPp"
	SecretKey = "gIsoiA91y4IuXkUpo8aiR2aWQikOaVOWd6F-ynKz"
	Bucket    = "cupw"
)

type PutRet struct {
	Key    string
	Hash   string
	FSize  int
	Bucket string
	Name   string
}

func QiNiuUpload(filePath string, fileName string)(bool,error) {

	// 使用 returnBody 自定义回复格式
	putPolicy := storage.PutPolicy{
		Scope:      Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","f_size":$(f_size),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": fileName,
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, fileName + ".jpg", filePath, &putExtra)
	fmt.Println(ret.Bucket, ret.Key, ret.FSize, ret.Hash, ret.Name)
	if err != nil {
		//fmt.Println(err)
		return true, err
	} else {
		//fmt.Println(ret.Bucket, ret.Key, ret.FSize, ret.Hash, ret.Name)
		return false, nil
	}

}

