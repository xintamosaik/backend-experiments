package main

import (
	"fmt"
	"os"
)

func get_file_content(filename string) (string, error) {

	file, err:= os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	file_content := make([]byte, 1024)

	_, err = file.Read(file_content)

	if err != nil {
		return "", err
	}

	return string(file_content), nil

	// print file body.html
}

func main() {
	fmt.Println("<!DOCTYPE html>")
	fmt.Println("<html lang=\"en\">")
	fmt.Println("<head>")
	fmt.Println("<meta charset=\"UTF-8\">")
	fmt.Println("<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">")
	fmt.Println("<title>Document</title>")
	fmt.Println("</head>")

	fmt.Println("<body>")
	// get body.html
	filename := "body.html";
	file_content, err := get_file_content(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(file_content)





	fmt.Println("</body>")
	fmt.Println("</html>")
}
