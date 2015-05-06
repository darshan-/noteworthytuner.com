package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8090

const faqURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/FAQ.wiki"
const changelogURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/Changelog.md"
const sourceURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/SourceCode.md"

func main() {
	http.HandleFunc("/faq", makeRedirect(faqURL))
	http.HandleFunc("/changelog", makeRedirect(changelogURL))
	http.HandleFunc("/sourcecode", makeRedirect(sourceURL))

	http.Handle("/", http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil))
}

func makeRedirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
