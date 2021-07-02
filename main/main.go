package main

func main() {
	var paths []string

	file, dir, err := parseFlags()
	if err != nil {
		panic(err)
	}

	if file != "" {
		appendFilePath(file, &paths)
	} else if dir != "" {
		appendDirPaths(dir, &paths)
	}

	for _, path := range paths {
		generateHTML(path)
	}
}
