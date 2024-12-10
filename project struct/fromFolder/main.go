package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func getIndentation(depth int) string {
    return strings.Repeat("  ", depth)
}

func shouldSkipDir(name string) bool {
    return strings.HasPrefix(name, ".")
}

func walkDirectory(path string, depth int, output *os.File) error {
    entries, err := os.ReadDir(path)
    if err != nil {
        return err
    }

    for _, entry := range entries {
        if entry.IsDir() && shouldSkipDir(entry.Name()) {
            continue
        }

        fullPath := filepath.Join(path, entry.Name())
        indentation := getIndentation(depth)
        
        if entry.IsDir() {
            fmt.Fprintf(output, "%s%s/\n", indentation, entry.Name())
            if err := walkDirectory(fullPath, depth+1, output); err != nil {
                return err
            }
        } else {
            fmt.Fprintf(output, "%s%s\n", indentation, entry.Name())
        }
    }
    return nil
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: program <directory_path> <output_file>")
        os.Exit(1)
    }

    root := os.Args[1]
    outputPath := os.Args[2]

    output, err := os.Create(outputPath)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
        os.Exit(1)
    }
    defer output.Close()

    if err := walkDirectory(root, 0, output); err != nil {
        fmt.Fprintf(os.Stderr, "Error walking directory: %v\n", err)
        os.Exit(1)
    }
}