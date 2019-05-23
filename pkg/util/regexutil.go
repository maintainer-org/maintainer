package util

import "regexp"

// GetParams get the params from regexp.
// See http://stackoverflow.com/questions/30483652/how-to-get-capturing-group-functionality-in-golang-regular-expressions.
func GetParams(regEx, url string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}
