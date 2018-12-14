package controllers

import (
	"fmt"
	"github.com/GoProjectGroupForEducation/Go-Blog/utils"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func DownloadFile(w http.ResponseWriter, req *http.Request, next utils.NextFunc) error  {
	//file, _, err := req.FormFile("uploadFile")
	//if err != nil {
	//	return utils.SendData(w, "","INVALID_FILE", http.StatusBadRequest)
	//}
	//defer file.Close()
	//fileBytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	return utils.SendData(w, "","INVALID_FILE", http.StatusBadRequest)
	//}
	////判定文件类型
	//filetype := http.DetectContentType(fileBytes)
	//
	////产生随机的token用作文件名
	////var checksum = []byte("heng-is-a-very-handsome-boy")
	////h := md5.New()
	////token := fmt.Sprintf("%x", h.Sum(checksum))
	//
	//fileName := "1"
	//fileEndings, err := mime.ExtensionsByType(filetype)
	//if err != nil {
	//	return utils.SendData(w, "","CANT_READ_FILE_TYPE", http.StatusInternalServerError)
	//}
	//root, _ := os.Getwd()
	//downloadFilePath := root + "/static/"
	//
	//fmt.Printf("fking\n filetype: " )
	//
	//newPath := downloadFilePath + fileName + ".ico"
	//
	//fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)
	//
	//newFile, err := os.Create(newPath)
	//if err != nil {
	//	return utils.SendData(w, "","CANT_WRITE_FILE", http.StatusInternalServerError)
	//}
	//defer newFile.Close()
	//if _, err := newFile.Write(fileBytes); err != nil {
	//	return utils.SendData(w, "","CANT_WRITE_FILE", http.StatusInternalServerError)
	//}


	getMultiPart(req)
	return utils.SendData(w, "New file name: " , "Upload file successfully", http.StatusOK)
}


//字节解析multi-part

//通过MultipartReader

func getMultiPart(r *http.Request)()  {



	mr,err := r.MultipartReader()

	if err != nil{

		fmt.Println("r.MultipartReader() err,",err)

		return

	}

	form ,_ := mr.ReadForm(128)

	getFormData(form)

}


func getFormData(form *multipart.Form) {
	//获取 multi-part/form body中的form value
	for k, v := range form.Value {
		fmt.Println("value,k,v = ", k, ",", v)
	}

	fmt.Println()
	//获取 multi-part/form中的文件数据
	for _, v := range form.File {
		for i := 0; i < len(v); i++ {
			fmt.Println("file part ", i, "-->")
			fmt.Println("fileName   :", v[i].Filename)
			fmt.Println("part-header:", v[i].Header)
			f, _ := v[i].Open()
			buf, _ := ioutil.ReadAll(f)
			fmt.Println("file-content", string(buf))
			fmt.Println()
		}
	}
}