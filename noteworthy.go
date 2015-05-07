package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const PORT = 8090

const faqURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/FAQ.wiki"
const changelogURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/Changelog.md"
const sourceURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/SourceCode.md"

// Doesn't properly fork, but starts same program again then exits
func daemonize() {
	err := os.MkdirAll("run", 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat("run/pid")
	if os.IsNotExist(err) {
		pidFile, err := os.Create("run/pid")
		if err != nil {
			log.Fatal(err)
		}

		err = pidFile.Sync()
		if err != nil {
			log.Fatal(err)
		}

		cmd := exec.Command(os.Args[0])
		err = cmd.Start()
		if err != nil {
			log.Fatal(err)
		}

		_, err = fmt.Fprintln(pidFile, cmd.Process.Pid)
		if err != nil {
			log.Fatal(err)
		}
		err = pidFile.Sync()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Daemon started with pid", cmd.Process.Pid)
		os.Exit(0)
	} else if err != nil {
		log.Fatal(err)
	}
}

func main() {
	daemonize()

	http.HandleFunc("/faq", makeRedirect(faqURL))
	http.HandleFunc("/changelog", makeRedirect(changelogURL))
	http.HandleFunc("/sourcecode", makeRedirect(sourceURL))

	http.Handle("/", http.FileServer(http.Dir("static/")))

	logF, _ := os.Create("run/log") // If error, can't write it to log, might as well proceed
	
	fmt.Fprintln(logF, http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}

func makeRedirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
