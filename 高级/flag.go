package main

import (
	"flag"
	"fmt"
)

func main() {
	backup_dir := flag.String("b", "/home/default_dir", "backup path")
	debug_mode := flag.Bool("d", false, "debug mode")
	flag.Parse()

	fmt.Println("backup_dir:", *backup_dir)
	fmt.Println("debug_mode:", *debug_mode)
}

/**
E:\go_workspace\src\learn-Go\高级>go run flag.go -b /home/backup
backup_dir: /home/backup
debug_mode: false
 */