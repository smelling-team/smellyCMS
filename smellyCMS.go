package smellyCMS

import(
	"net/http"
	"os"
	"bufio"
	"strings"
	)

func main() {
	settings := make(map[string]string)
	settings = parseConfig()
	root := settings["root"]
	http.HandleFunc(root, handleCall)
}

func handleCall(writer http.ResponseWriter, request *http.Request) {

}

func parseConfig() map[string]string {
	settings := make(map[string]string)
	file, err := os.Open("conf/config.txt")
	var lineContent string
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			lineContent = scanner.Text()
			lineParts := strings.Split(lineContent, "=")
			settings[lineParts[0]] = lineParts[1]
		}
	}
	return settings
}