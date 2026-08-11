# HTTP

> Part 1. Return codes

Create a script `part-1.sh` which runs an http request and parses the return code.

The script must print the answer in the format `<DESC>(<CODE>)`

Where `DESC` is a textual description of the return code:
- SUCCESS
- REDIRECT
- CLIENT_ERROR
- SERVER_ERROR

And `CODE` is the return code

To prevent waiting indefinitely for the response, set the maximum execution time to 5 sec.

**Synopsis:**
```cmd
./part-1.sh <URL>
```

**Output:**
```cmd
SUCCESS(200)
```
## Error handling

### Unable to connect

Exit with code 77

### Timeout

Exit with code 88


---

* What are the 4 types of http codes?
* Why is a timeout necessary?
