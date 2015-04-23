package main

import (
	"fmt"
	"go-resource"
	"log"
	"os"
	"os/user"
	"path"
)

var resource_dir string = os.Getenv("RESOURCE_DIR")

func check_resource_dir() {
	if resource_dir == "" {
		log.Println("RESOURCE_DIR not set")
		os.Exit(100)
	}
	_, err := os.Stat(resource_dir)
	if nil != err {
		log.Println("RESOURCE_DIR Directory not found: " + resource_dir)
		log.Println(err)
		os.Exit(101)
	}
}

func syntax() {
	fmt.Fprintln(os.Stderr, "Syntax: create_resource PATH")
	os.Exit(99)
}

func main() {
	check_resource_dir()

	if 2 != len(os.Args) {
		syntax()
	} else {
		resource_path := os.Args[1]
		fmt.Println("OK: path=" + resource_path)
		resource_full_path := path.Join(resource_dir, resource_path)
		os.MkdirAll(path.Dir(resource_full_path), 0777)
		_, err := os.Stat(resource_full_path)
		if nil == err {
			fmt.Fprintln(os.Stderr, "Resouce "+resource_full_path+" already exists")
			os.Exit(120)
		}
		os.Create(resource_full_path)
		resource_host := os.Getenv("RESOURCE_HOST")
		if "" == resource_host {
			resource_host, _ = os.Hostname()
		}
		resource_user := os.Getenv("RESOURCE_USER")
		if "" == resource_user {
			usr, _ := user.Current()
			resource_user = usr.Username
		}
		resource_uri := resource_host + "@" + resource_user + ":" + resource_path
		log.Println("created resource " + resource_uri)
		fmt.Print(resource_uri)
	}
}
