package main

import (
	ftp "./goftp"
	"github.com/pkg/sftp"
)

func main() {
	ftpClient, err := ftp.NewFtp(getConnectionString())
	if err != nil {
		panic(err)
	}
	defer ftpClient.Close()

	if err = ftpClient.Login("ftper", "1q2w3e4r"); err != nil {
		panic(err)
	}

	if err = ftpClient.OpenDirectory("/mt4/"); err != nil {
		panic(err)
	}

	if err = ftpClient.Upload("abc.txt", "abc.txt"); err != nil {
		panic(err)
	}
}

func getConnectionString() string {
	return "localhost:21"
}
