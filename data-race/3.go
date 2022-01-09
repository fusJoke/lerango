package main

import "sync"

var service = map[string]string{}

func RegisterService(name, addr string){
	service[name] = addr
}

func LookupService(name string) string {
	return service[name]
}

//少用或尽量避免用全局变量，但是使用map作为全局变量一种比较常见的情况就是配置信息。全局变量的话一般的做法就是加锁。也可以使用sync.Map

var serviceMu sync.Mutex

func RegisterServiceSafe(name, addr string) {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	service[name] = addr
}

func LookupServiceSafe(name string) string {
	serviceMu.Lock()
	defer serviceMu.Unlock()
	return service[name]
}

