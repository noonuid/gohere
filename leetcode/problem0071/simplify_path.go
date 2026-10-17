package problem0071

import "strings"

func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	fileNames := []string{}
	for _, str := range strs {
		switch str {
		case "", ".":

		case "..":
			if len(fileNames) > 0 {
				fileNames = fileNames[:len(fileNames)-1]
			}
		default:
			fileNames = append(fileNames, str)
		}
	}
	result := "/" + strings.Join(fileNames, "/")
	return result
}
