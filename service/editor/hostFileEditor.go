package editor

import (
	"bufio"
	"hostsEditor/config"
	"hostsEditor/enums"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// MakeFileByDeletionOrAddition
func MakeFileByDeletionOrAddition(file string, actionTargetString string, action enums.ActionsEnum) {
	var lines []string

	var backupFile string = config.VarBackupFiles + file + "_Backup_" + getTime().Local().GoString()
	var tempFile string = config.VarTempFile

	copyFile(file, backupFile)

	// open file with access level READ & WRITE

	var newFile, err = os.Create(tempFile)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()

	switch action {
	case enums.ActionRemove:
		lines = removeLinesFromFileForNewContent(file, actionTargetString)
	case enums.ActionAdd:
		lines = addLinesToFileForNewContent(file, actionTargetString)
	}

	for element := range lines {
		_, err = newFile.WriteString(lines[element] + "\n")
		if err != nil {
			panic(err)
		}
	}
	copyFile(tempFile, file)
}

func removeLinesFromFileForNewContent(file string, searchStr string) []string {
	lines := getLinesFromFile(file)
	s := removeLine(lines, searchStr)
	return s
}

func addLinesToFileForNewContent(file string, searchStr string) []string {
	lines := getLinesFromFile(file)
	s := addLine(lines, searchStr)
	return s
}

func addLine(s []string, r string) []string {
	var d []string
	var len int = 1
	for element := range s {
		if strings.EqualFold(s[element], r) {
			len--
		}
		d = append(d, s[element])
	}
	if !(len < 1) {
		d = append(d, r)
	}
	return d
}

func removeLine(s []string, r string) []string {
	var d []string
	for element := range s {
		if element <= config.VarOffsetFileHost {
			d = append(d, s[element])
		}
		if !strings.EqualFold(s[element], r) && element > config.VarOffsetFileHost {
			d = append(d, s[element])
		}
	}
	return d
}

func getTime() time.Time {
	return time.Now()
}

func copyFile(src string, dst string) {
	fin, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}

func getLinesFromFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}

func GetLines(fileName string) []string {
	return getLinesFromFile(fileName)
}

func GetLinesForHtmlView(fileName string) string {
	var d string
	var s = getLinesFromFile(fileName)
	for element := range s {
		d = d + s[element] + "<br />"
	}
	return d
}
