package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
)

func vphash_file(videopath string) []string {
	cmd := exec.Command("./hash_vid.py", videopath, "/tmp/vidhasher.out")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		Log.Fatal(err)
	}

	hashes := make([]string, 0)
	outfile, err := os.Open("/tmp/vidhasher.out")
	if err != nil {
		Log.Fatal(err)
	}
	reader := bufio.NewReader(outfile)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			Log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		hashes = append(hashes, line)
	}

	return hashes
}
