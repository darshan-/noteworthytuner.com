package goapp

import (
	"net/http"
)

const faqURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/FAQ.wiki"
const changelogURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/Changelog.md"
const sourceURL = "https://github.com/darshan-/Noteworthy-Tuner-Support/blob/master/SourceCode.md"

func init() {
	http.HandleFunc("/faq", makeRedirect(faqURL))
	http.HandleFunc("/changelog", makeRedirect(changelogURL))
	http.HandleFunc("/sourcecode", makeRedirect(sourceURL))
}

func makeRedirect(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
