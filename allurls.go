package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func takescreenshot(url string, index string) string {

	app := "google-chrome-stable"

	agruments := "--headless  --timeout=30000 --disable-gpu --enable-logging --screenshot=\""

	cmd := app + " " + agruments + "./out/" + index + ".jpg" + "\" " + url
	fmt.Println("[+] executing : ", cmd)

	out := exec.Command("bash", "-c", cmd)
	/*      if err != nil {
	        log.Fatal(err)
	}*/
	err1 := out.Run()
	if err1 != nil {
		log.Fatal(err1)
		return "[!] error"
	}
	return "[+] Done"

}

func main() {
	crtout := "mkdir ./out "
	o := exec.Command("bash", "-c", crtout)
	err2 := o.Run()
	if err2 != nil {
		log.Fatal(err2)
	}
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)

	}
	c1 := make(chan string, 1)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		S1 := scanner.Text()
		S2 := scanner.Text()
		S2 = strings.Replace(S2, "http:", "", -1)
		S2 = strings.Replace(S2, "https:", "", -1)
		S2 = strings.Replace(S2, "/", "", -1)
		go func() {
			text := takescreenshot(S1, S2)
			c1 <- text
		}()
		select {
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(3 * time.Second):
			fmt.Println("[!] Timeout !! too many redirections or dead asset -_-'")
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
