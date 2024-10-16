package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func startApi() {

	cmd := exec.Command("go", "run", "api/main.go", "-f", "etc/api.yaml")

	// 获取标准输出和标准错误
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating StdoutPipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating StderrPipe: %v", err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting command: %v", err)
	}

	// 创建一个 goroutine 来将标准输出写入标准输出
	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	// 创建一个 goroutine 来将标准错误写入标准错误
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Error waiting for command: %v", err)
	}

	fmt.Println("Command executed successfully.")
}

func startLogin() {

	cmd := exec.Command("go", "run", "loginsvc/main.go", "-f", "etc/loginsvc.yaml")

	// 获取标准输出和标准错误
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating StdoutPipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating StderrPipe: %v", err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting command: %v", err)
	}

	// 创建一个 goroutine 来将标准输出写入标准输出
	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	// 创建一个 goroutine 来将标准错误写入标准错误
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Error waiting for command: %v", err)
	}

	fmt.Println("Command executed successfully.")
}

func startLogin1() {

	cmd := exec.Command("go", "run", "loginsvc/main.go", "-f", "etc/loginsvc1.yaml")

	// 获取标准输出和标准错误
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating StdoutPipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating StderrPipe: %v", err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting command: %v", err)
	}

	// 创建一个 goroutine 来将标准输出写入标准输出
	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	// 创建一个 goroutine 来将标准错误写入标准错误
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Error waiting for command: %v", err)
	}

	fmt.Println("Command executed successfully.")
}

func startUser() {

	cmd := exec.Command("go", "run", "usersvc/main.go", "-f", "etc/usersvc.yaml")

	// 获取标准输出和标准错误
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Error creating StdoutPipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Error creating StderrPipe: %v", err)
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error starting command: %v", err)
	}

	// 创建一个 goroutine 来将标准输出写入标准输出
	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	// 创建一个 goroutine 来将标准错误写入标准错误
	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		log.Fatalf("Error waiting for command: %v", err)
	}

	fmt.Println("Command executed successfully.")
}

func main() {

	s := flag.String("s", "http", "启动服务rpc,http,default(:http)")

	flag.Parse()

	go startLogin()
	// go startLogin1()
	go startUser()

	if *s != "rpc" {
		time.Sleep(15 * time.Second)
		startApi()
	} else {
		select {}
	}

}
