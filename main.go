package main

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

var (
	defaultServers = []string{"149.154.175.50:443", "149.154.167.51:443", "149.154.175.100:443", "149.154.167.91:443", "149.154.171.5:443"}
)

func main() {
	secret := getSecret()
	if "" == secret {
		log.Fatalln("Can not load secret file.")
	}

	secret_hex, err := hex.DecodeString(secret)
	if nil != err {
		log.Fatalln(secret_hex)
	}

	laddr, err := net.ResolveTCPAddr("tcp", ":8822")
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}

	network := NewNetwork(defaultServers)

	log.Println("[*]MTProxy is running...")
	log.Println("[*]")
	log.Println("[*] Secret %s", secret)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Failed to accept connection '%s'", err)
			continue
		}

		client := NewClient(conn, network, secret_hex)
		go client.Do()
	}
}

func getSecret() string {
	if _, err := os.Stat("/data/secret"); err != nil {
		writeSecret()
	}
	return readSecret()
}

func readSecret() string {
	f, err := os.Open("/data/secret")
	if nil != err {
		log.Fatalln(err)
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if nil != err {
		log.Fatalln(err)
	}
	return string(fd)
}

func writeSecret() string {
	secret := getRandom()
	var be = []byte(secret)
	err := ioutil.WriteFile("/data/secret", be, 0666)
	if nil != err {
		log.Fatalln(err)
	}
	return secret
}

func getRandom() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`~!@#$%^&*()-=_+[]{}|;:',<.>/?"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return hex.EncodeToString(result)
}
