package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var crypter string
	fmt.Printf("		MD5 CRYPTER\n")
	fmt.Printf("[+] Veuillez entrer vos données à crypter : ")
	fmt.Scanf("%v", crypter)

	crypted := []byte(crypter)
	fmt.Printf("Voici votre texte chiffré [ %x ]", md5.Sum(crypted))
}
