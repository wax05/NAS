package FileLog

import (
	"os"

	"errors"

	"time"
)

var (
	errFileNotFound = errors.New("file not found")
	errCantWrite    = errors.New("can't cile write")
)

func Log(Type, log, Path string) error {
	now := time.Now()
	t1 := now.String() + "\nType:" + Type + "|" + log + "\n----------------------------------------------------------------"
	w1 := []byte(t1)
	f, err := os.OpenFile(Path, os.O_APPEND|os.O_RDWR, 0755)
	if err != nil {
		return errFileNotFound
	}
	defer f.Close()
	_, err = f.Write(w1)
	if err != nil {
		return errCantWrite
	}
	return nil
}

func ErrLog(Path string, err error) error {
	err = Log("Error", err.Error(), Path)
	if err != nil {
		return err
	}
	return nil
}
