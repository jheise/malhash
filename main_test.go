package main

import "testing"

var (
	sampleFile    = "sample.txt"
	headerCorrect = "text/plain"
	ssdeepCorrect = "192:9ZoTal7mup1MBBEgZK5diWOxhC6ES5Ob3WIZ7JaPw+yxy/fMOA/B5fKRDBMrkCka:LKWAgCC7srttuPn"
	sha256Correct = "78e60831aa21c95cea85bf2d13cf729b48ff218295951729ff7bcdc11a1ca227"
	sha1Correct   = "86c16498987e8bcc248dc0ca709ec29350ea1d49"
	md5Correct    = "a965303227cdaa2798a12bf5895ae673"
)

func TestHeaders(t *testing.T) {
	headers := _grabFileHeaders(sampleFile)
	if headers != headerCorrect {
		t.Errorf("Headers received are incorrect, expected '%s', received '%s'", headerCorrect, headers)
	}
}

func TestSsdeep(t *testing.T) {
	ssdeep, err := _grabSsdeep(sampleFile)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if ssdeep != ssdeepCorrect {
		t.Errorf("ssdeep value incorrect\n expected: %s\n received: %s", ssdeepCorrect, ssdeep)
	}
}

func TestSHA256(t *testing.T) {
	sha256, err := _grabSHA256(sampleFile)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if sha256 != sha256Correct {
		t.Errorf("sha256 value incorrect\n expected: %s\n received: %s", sha256Correct, sha256)
	}
}

func TestSHA1(t *testing.T) {
	sha1, err := _grabSHA1(sampleFile)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if sha1 != sha1Correct {
		t.Errorf("sha1 value incorrect\n expected: %s\n received: %s", sha1Correct, sha1)
	}
}

func TestMD1(t *testing.T) {
	md5, err := _grabMD5(sampleFile)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if md5 != md5Correct {
		t.Errorf("md5 value incorrect\n expected: %s\n received: %s", md5Correct, md5)
	}
}
