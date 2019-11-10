package main

import (
	// standard
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"

	// external
	"github.com/alecthomas/kingpin"
	"github.com/glaslos/ssdeep"
	"github.com/vimeo/go-magic/magic"
)

type MalHeader struct {
	Header string `json:"header"`
	Ssdeep string `json:"ssdeep"`
	SHA256 string `json:"sha256"`
	SHA1   string `json:"sha1"`
	MD5    string `json:"md5"`
}

func (self MalHeader) printHuman() {
	fmt.Printf("header: %s\n", self.Header)
	fmt.Printf("ssdeep: %s\n", self.Ssdeep)
	fmt.Printf("sha1:   %s\n", self.SHA1)
	fmt.Printf("sha256: %s\n", self.SHA256)
	fmt.Printf("md5:    %s\n", self.MD5)
}

func (self MalHeader) printJSON() error {
	jsonStr, err := json.MarshalIndent(self, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(jsonStr))
	return nil
}

var (
	//target
	target     = kingpin.Arg("target", "What file to parse").Required().String()
	jsonOutput = kingpin.Flag("json", "Enable json output").Bool()
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

	malheader := MalHeader{}
	var err error

	malheader.Header = _grabFileHeaders(*target)

	malheader.Ssdeep, err = _grabSsdeep(*target)
	if err != nil {
		panic(err)
	}

	malheader.SHA1, err = _grabSHA1(*target)
	if err != nil {
		panic(err)
	}

	malheader.SHA256, err = _grabSHA256(*target)
	if err != nil {
		panic(err)
	}

	malheader.MD5, err = _grabMD5(*target)
	if err != nil {
		panic(err)
	}

	if *jsonOutput {
		malheader.printJSON()
	} else {
		malheader.printHuman()
	}

}
