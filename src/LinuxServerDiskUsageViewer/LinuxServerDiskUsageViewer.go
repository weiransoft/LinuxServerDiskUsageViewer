package main
import (
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"bytes"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/viewDiskUsage/", viewHandler)
	http.ListenAndServe(":8080", nil)
}


func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Linux Server Disk Usage Viewer:\n")
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "URL: %s\n", r.URL.Path)
}

func exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true}
    if os.IsNotExist(err) { return false}
    return false
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "viewDiskUsage: \n")
	var validPath = regexp.MustCompile("^/(viewDiskUsage)/(view)$")
	m := validPath.FindString(request.URL.Path)

	if m!="" {
			viewPath := request.URL.RawQuery
			if viewPath:=strings.Replace(viewPath,"path=","",-2);exists(viewPath) {
					fmt.Fprintf(writer, "%s", executeShellCmd(exec.Command("du","-h","--max-depth=1",viewPath)))
			}
	} else {
			fmt.Fprintf(writer, "%s", executeShellCmd(exec.Command("df","-h")))
	}
}

func executeShellCmd(cmd *exec.Cmd) string{
    var out bytes.Buffer
    cmd.Stdout=&out
    err:=cmd.Run()
    if(err!=nil) {
        err.Error()
    }
    return out.String()

}
