package echoer

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// GetValue is a testing function
func GetValue() string {
	return "Hello from this another package"
}

// EchoParseMultipartForm is an multipart echoer using parse multipart form
func EchoParseMultipartForm(r *http.Request) {

	if err := r.ParseMultipartForm(32 << 20); err != nil { // maxMemory 32MB)
		fmt.Println("r.ParseMultipartForm() err,", err)
		return
	}

	fmt.Println("r.Form: ", r.Form)
	fmt.Println("r.PostForm: ", r.PostForm)
	fmt.Println("r.MultipartFrom: ", r.MultipartForm)
}

// EchoMultipartReader is an multipart echoer using multipart reader
func EchoMultipartReader(r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Println("r.MultipartReader() err,", err)
		return
	}

	form, _ := mr.ReadForm(32 << 20) // maxMemory 32MB
	getFormData(form)
}

// EchoMultipart3 is an multipart echoer using multipart reader next part
func EchoMultipart3(r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Println("r.MultipartReader() err,", err)
		return
	}

	fmt.Println()
	fmt.Println("----------------------------------------")
	for {
		p, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("mr.NextPart() err", err)
			break
		}

		fmt.Println("Part header: ", p.Header)
		formName := p.FormName()
		fileName := p.FileName()

		if formName != "" && fileName == "" {
			formValue, _ := ioutil.ReadAll(p)
			fmt.Printf("Form Name: %s, Form Value: %s\n", formName, formValue)
		}

		if fileName != "" {
			fileData, _ := ioutil.ReadAll(p)
			fmt.Printf("File Name: %s, File Data: %s \n", fileName, fileData)
		}

		fmt.Println()
		fmt.Println("----------------------------------------")
	}
}

func getFormData(form *multipart.Form) {
	//Iterate Value Map
	for k, v := range form.Value {
		fmt.Println("value Form k,v = ", k, " ", v)
	}

	fmt.Println()

	//Iterate File Map
	for k, v := range form.File {
		for i := 0; i < len(v); i++ {
			fmt.Println("File Form Key :", k)
			fmt.Println("File part ", i, "-->")
			fmt.Println("File Name: ", v[i].Filename)
			fmt.Println("part-header: ", v[i].Header)

			//Read File contain
			f, _ := v[i].Open()
			buf, _ := ioutil.ReadAll(f)

			fmt.Println("file-content", string(buf))
			fmt.Println()
		}
	}
}
