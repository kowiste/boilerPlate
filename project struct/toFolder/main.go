package main

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

type ProjectItem struct {
    path     string
    isFolder bool
    comment  string
    indent   int
    parent   string
}

func getIndentLevel(line string) int {
    return len(line) - len(strings.TrimLeft(line, " "))
}

func parseProjectStructure(input string) []ProjectItem {
    var items []ProjectItem
    var pathStack []string
    var indentStack []int
    
    scanner := bufio.NewScanner(strings.NewReader(input))
    
    for scanner.Scan() {
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            continue
        }

        indent := getIndentLevel(line)

        // Extract path and comment
        parts := strings.Split(line, "#")
        pathPart := strings.TrimSpace(parts[0])
        comment := ""
        if len(parts) > 1 {
            comment = strings.TrimSpace(parts[1])
        }

        // Clean up path
        path := strings.TrimSpace(pathPart)
        if path == "" {
            continue
        }

        // Update path stack based on indentation
        for len(indentStack) > 0 && indent <= indentStack[len(indentStack)-1] {
            indentStack = indentStack[:len(indentStack)-1]
            pathStack = pathStack[:len(pathStack)-1]
        }

        // Determine if it's a folder or file
        isFolder := strings.HasSuffix(path, "/")
        if isFolder {
            path = strings.TrimSuffix(path, "/")
        } else {
            isFolder = !strings.Contains(path, ".")
        }

        // Build full path
        fullPath := path
        if len(pathStack) > 0 {
            fullPath = filepath.Join(pathStack...)
            fullPath = filepath.Join(fullPath, path)
        }

        item := ProjectItem{
            path:     fullPath,
            isFolder: isFolder,
            comment:  comment,
            indent:   indent,
        }

        items = append(items, item)

        if isFolder {
            pathStack = append(pathStack, path)
            indentStack = append(indentStack, indent)
        }
    }
    return items
}

func createProjectStructure(baseDir string, items []ProjectItem) error {
    for _, item := range items {
        fullPath := filepath.Join(baseDir, item.path)
        
        if item.isFolder {
            if err := os.MkdirAll(fullPath, 0755); err != nil {
                return fmt.Errorf("failed to create directory %s: %w", fullPath, err)
            }
            fmt.Printf("Created directory: %s\n", fullPath)
        } else {
            // Create parent directories if they don't exist
            parentDir := filepath.Dir(fullPath)
            if err := os.MkdirAll(parentDir, 0755); err != nil {
                return fmt.Errorf("failed to create parent directory %s: %w", parentDir, err)
            }

            // Create the file
            file, err := os.Create(fullPath)
            if err != nil {
                return fmt.Errorf("failed to create file %s: %w", fullPath, err)
            }
            
            // If there's a comment, write it as a package comment
            if item.comment != "" {
                _, err = file.WriteString(fmt.Sprintf("// Package %s %s\npackage %s\n",
                    filepath.Base(filepath.Dir(fullPath)),
                    item.comment,
                    filepath.Base(filepath.Dir(fullPath))))
                if err != nil {
                    file.Close()
                    return fmt.Errorf("failed to write to file %s: %w", fullPath, err)
                }
            } else {
                // Write default package declaration
                _, err = file.WriteString(fmt.Sprintf("package %s\n", 
                    filepath.Base(filepath.Dir(fullPath))))
                if err != nil {
                    file.Close()
                    return fmt.Errorf("failed to write to file %s: %w", fullPath, err)
                }
            }
            
            file.Close()
            fmt.Printf("Created file: %s\n", fullPath)
        }
    }
    return nil
}

func readStructureFile(filename string) (string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return "", fmt.Errorf("failed to read file %s: %w", filename, err)
    }
    return string(data), nil
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: program <structure_file>")
        fmt.Println("Example: program structure.txt")
        os.Exit(1)
    }

    filename := os.Args[1]
    input, err := readStructureFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
        os.Exit(1)
    }

    items := parseProjectStructure(input)

    // Create the project structure in the current directory
    if err := createProjectStructure(".", items); err != nil {
        fmt.Fprintf(os.Stderr, "Error creating project structure: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Project structure created successfully!")
}