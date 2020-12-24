var inp []string
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
  inp = append(inp, scanner.Text())
}

if scanner.Err() != nil {
	panic(scanner.Err().Error())
}
