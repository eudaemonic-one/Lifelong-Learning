# Append Text to a File

## Append Text to A File

```go
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()
if _, err := f.WriteString("text to append\n"); err != nil {
	log.Println(err)
}
```
