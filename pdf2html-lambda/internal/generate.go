package internal

import (
	"log"
	"os/exec"
)

// Executes the pd2html CLI with the given input and output filepaths
func pdf2html(inputPath string, outputPath string) error {
	cmd := exec.Command("pdf2html", "--zoom=2", inputPath, outputPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("command error: ", string(out), err)
		return err
	}

	return nil
}
