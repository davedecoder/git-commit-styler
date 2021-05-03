package main

import (
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_getFileData(t *testing.T)  {
	tests := []struct {
		name string
		want inputFile
		wantErr bool
		osArgs []string
	}{
		{"Default parameters", inputFile{"test.csv", "comma", false}, false, []string {"cmd", "test.csv"}},
		{"No Parameters", inputFile{}, true, []string {"cmd"}},
		{"Semicolon enabled", inputFile{"test.csv", "semicolon", false}, false, []string {"cmd" , "--separator=semicolon", "test.csv"}},
		{"Pretty enabled", inputFile{"test.csv", "comma", true}, false, []string{"cmd", "--pretty", "test.csv"}},
		{"Pretty and semicolon enabled", inputFile{"test.csv", "semicolon", true}, false, []string{"cmd", "--pretty", "--separator=semicolon", "test.csv"}},
		{"Separator not identified", inputFile{}, true, []string{"cmd", "--separator=pipe", "test.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args

			defer func (){
				os.Args = actualOsArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs
			got, err := getFileData()
			if (err != nil) != tt.wantErr {
				t.Errorf("getFileData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFileData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIfValidFile(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "test*.csv")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())

	tests := []struct {
		name    string
		filename    string
		want    bool
		wantErr bool
	}{
		{ "File does exists", tmpFile.Name(), true, false},
		{ "File does not exists", "nowhere/file.csv", false, true},
		{ "File is not csv", "text.txt", false, true},

	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := checkIfValidFile(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkIfValidFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("checkIfValidFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}