var inp string
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	inp += scanner.Text()
}

if scanner.Err() != nil {
	// handle error.
}
