package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	var hash string
	fmt.Printf("		MD5 Hasher \n")
	fmt.Printf("[+] Veuillez entrer vos données à crypter : ")
	fmt.Scanf("%v", hash)

	hashed := []byte(hash)
	fmt.Printf("Voici votre texte chiffré [ %x ]", md5.Sum(hashed))
}
