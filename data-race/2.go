package main

import "os"

func main(){
	ParallelWrite([]byte("xxx"))
}

func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("/tmp/file1")
	if err != nil {
		res <- err
	} else {
		go func(){
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("/tmp/file2")
	if err != nil {
		res <- err
	} else {
		go func(){
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}
	return res

	//共享了err变量,导致数据竞争
}
