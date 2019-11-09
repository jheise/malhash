package main

import "testing"

var (
	sample_file    = "sample.txt"
	header_correct = "text/plain"
	ssdeep_correct = "192:9ZoTal7mup1MBBEgZK5diWOxhC6ES5Ob3WIZ7JaPw+yxy/fMOA/B5fKRDBMrkCka:LKWAgCC7srttuPn"
	sha256_correct = "78e60831aa21c95cea85bf2d13cf729b48ff218295951729ff7bcdc11a1ca227"
	sha1_correct   = "86c16498987e8bcc248dc0ca709ec29350ea1d49"
	md5_correct    = "a965303227cdaa2798a12bf5895ae673"
)

func TestHeaders(t *testing.T) {
	headers := _grabFileHeaders(sample_file)
	if headers != header_correct {
		t.Errorf("Headers received are incorrect, expected '%s', received '%s'", header_correct, headers)
	}
}

func TestSsdeep(t *testing.T) {
	ssdeep, err := _grabSsdeep(sample_file)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if ssdeep != ssdeep_correct {
		t.Errorf("ssdeep value incorrect\n expected: %s\n received: %s", ssdeep_correct, ssdeep)
	}
}

func TestSHA256(t *testing.T) {
	sha256, err := _grabSHA256(sample_file)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if sha256 != sha256_correct {
		t.Errorf("sha256 value incorrect\n expected: %s\n received: %s", sha256_correct, sha256)
	}
}

func TestSHA1(t *testing.T) {
	sha1, err := _grabSHA1(sample_file)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if sha1 != sha1_correct {
		t.Errorf("sha1 value incorrect\n expected: %s\n received: %s", sha1_correct, sha1)
	}
}

func TestMD1(t *testing.T) {
	md5, err := _grabMD5(sample_file)
	if err != nil {
		t.Errorf("Excepted nil, received: %s", err)
	}

	if md5 != md5_correct {
		t.Errorf("md5 value incorrect\n expected: %s\n received: %s", md5_correct, md5)
	}
}
