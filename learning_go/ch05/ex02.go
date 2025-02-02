package main

import "os"

func fileLen(filename string) (int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return -1, err
	}
	return stat.Size(), err
}
