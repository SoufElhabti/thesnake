package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func takescreenshot(url string, index string) {

	app := "google-chrome-stable"

	agruments := "--headless  --timeout 30000 --disable-gpu --enable-logging --screenshot=\""

	cmd := app + " " + agruments + "./out/" + index + ".jpg" + "\" " + url
	fmt.Println("[+] executing : ", cmd)

	out := exec.Command("bash", "-c", cmd)
	/*      if err != nil {
	        log.Fatal(err)
	}*/
	err1 := out.Run()
	if err1 != nil {
		log.Fatal(err1)
	}

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
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		S1 := scanner.Text()
		S2 := scanner.Text()
		S2 = strings.Replace(S2, "http:", "", -1)
		S2 = strings.Replace(S2, "https:", "", -1)
		S2 = strings.Replace(S2, "/", "", -1)

		takescreenshot(S1, S2)

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
