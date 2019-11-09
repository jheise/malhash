package main

import (
	// standard
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	// external
	"github.com/alecthomas/kingpin"
	"github.com/glaslos/ssdeep"
	"github.com/vimeo/go-magic/magic"
)

var (
	//target
	target = kingpin.Arg("target", "What file to parse").Required().String()
)

func _grabFileHeaders(filepath string) string {
	return magic.MimeFromFile(filepath)
}

func _grabSsdeep(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h, err := ssdeep.FuzzyFile(f)
	if err != nil {
		return "", err
	}
	return h, nil
}

func _grabMD5(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func _grabSHA1(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func _grabSHA256(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes), nil
}

func main() {
	kingpin.Parse()

	header := _grabFileHeaders(*target)

	ssdeep, err := _grabSsdeep(*target)
	if err != nil {
		panic(err)
	}

	sha1sum, err := _grabSHA1(*target)
	if err != nil {
		panic(err)
	}

	sha256sum, err := _grabSHA256(*target)
	if err != nil {
		panic(err)
	}

	md5sum, err := _grabMD5(*target)
	if err != nil {
		panic(err)
	}

	fmt.Printf("header: %s\n", header)
	fmt.Printf("ssdeep: %s\n", ssdeep)
	fmt.Printf("sha1:   %s\n", sha1sum)
	fmt.Printf("sha256: %s\n", sha256sum)
	fmt.Printf("md5:    %s\n", md5sum)
}
