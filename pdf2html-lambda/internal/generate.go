package internal

import (
	"log"
	"os/exec"
)

// Executes the pd2html CLI with the given input and output filepaths
func pdf2html(inputPath string, outputPath string) (bool, error) {
	cmd := exec.Command("pdf2html", "--zoom=2", "--debug=1", inputPath, outputPath)
	out, err := cmd.CombinedOutput()
	log.Println("command output: ", string(out))
	if err != nil {
		log.Println("command error: ", err)
		return false, err
	}

	return true, nil
}

func fallbackpdf2html(inputPath string, outputPath string) (bool, error) {
	cmd := exec.Command("pdf2html", "--fallback=1", "--zoom=2", inputPath, outputPath)
	out, err := cmd.CombinedOutput()
	log.Println("command output: ", string(out))
	if err != nil {
		log.Println("command error: ", string(out), err)
		return false, err
	}

	return true, nil
}
