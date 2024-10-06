package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"linux-3.9/utils"
)

func main() {
    flags := map[string]string{
        "/home/brandon/flag.txt": "OS{time_to_test_our_linux_skills}",
        "/home/katy/flag.txt":    "OS{soon_this_whole_box_will_be_yours}",
        "/mnt/usb/flag.txt":     "OS{Regulators_mount_up!}",
        "/root/flag.txt":    "OS{keys_to_the_kingdom}",
    }

    createFlagFiles(flags)

    if err := addUserWithSamePermissions("go_user", "jesse", "password"); err != nil {
        log.Printf("Failed to add user: %v", err)
    }

    runNewUserCheck()

    if err := simulateUSBMount("/tmp/usb"); err != nil {
        log.Printf("Failed to simulate USB mount: %v", err)
    }

    countFiles()

    setupBackupScript()
}

func createFlagFiles(flags map[string]string) {
    for path, flag := range flags {
        dir := filepath.Dir(path)
        if err := os.MkdirAll(dir, 0755); err != nil {
            log.Printf("Failed to create directory %s: %v", dir, err)
            continue
        }
        if err := utils.WriteFile(path, flag); err != nil {
            log.Printf("Failed to write flag to file: %v", err)
        } else {
            fmt.Printf("Flag written to %s\n", path)
        }
    }
}

func addUserWithSamePermissions(existingUser, newUser, _ string) error {
    fmt.Printf("Simulating adding user %s with same permissions as %s\n", newUser, existingUser)
    return nil
}

func runNewUserCheck() {
    fmt.Println("Simulating newuser_check")
}

func simulateUSBMount(mountPoint string) error {
    if err := os.MkdirAll(mountPoint, 0755); err != nil {
        return fmt.Errorf("failed to create mount point: %v", err)
    }
    fmt.Printf("Simulated mounting USB to %s\n", mountPoint)
    return nil
}

func countFiles() {
    count := 0
    filepath.Walk("/tmp", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && (filepath.Ext(path) == ".doc" || filepath.Ext(path) == ".csv") {
            count++
        }
        return nil
    })
    fmt.Printf("Total .doc and .csv files: %d\n", count)
}

func setupBackupScript() {
    fmt.Println("Simulating backup script setup")
}