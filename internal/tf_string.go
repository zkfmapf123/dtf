package internal

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	FIND_MATCH = "will be created"
)

func FindImportReturnArrs(output string) []string {

	findOutput := []string{}

	sc := bufio.NewScanner(strings.NewReader(string(output)))

	for sc.Scan() {
		line := strings.Trim(sc.Text(), "")

		if strings.Contains(line, FIND_MATCH) {

			line := removeStr(line, FIND_MATCH, "#")
			findOutput = append(findOutput, fmt.Sprintf("'%s'", line))
		}
	}

	return findOutput
}

func removeStr(line string, strs ...string) string {

	for _, str := range strs {

		line = strings.Replace(line, str, "", 1)
	}

	return line

}
