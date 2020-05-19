# Write Log to File (or `/dev/null`)

* This code appends a log message to the file `text.log`
  * It creates the file if it doesn’t already exist

```go
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()

logger := log.New(f, "prefix", log.LstdFlags)
logger.Println("text to append")
logger.Println("more text to append")
```

* Contents of `text.log`

```text
prefix: 2017/10/20 07:52:58 text to append
prefix: 2017/10/20 07:52:58 more text to append
```

* `log.New` creates a new `log.Logger` that writes to `f`
* The `prefix` appears at the beginning of each generated log line
* The `flag` argument defines which text to prefix to each log entry

## Disable Logging

* To turn off all output from a `log.Logger`, set the output destination to `ioutil.Discard`, a writer on which all calls succeed without doing anything

```go
log.SetOutput(ioutil.Discard)
```
