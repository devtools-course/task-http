# HTTP

> Part 4. Download files

You are to create a script `part-4.sh` which downloads a file from a URL. The file must be saved to the provided directory.

**Synopsis:**
```cmd
./part-4.sh <URL> <DIRPATH>
```

The name is defined in the response:

```cmd
< HTTP/1.1 200 OK
< Accept-Ranges: bytes
< Content-Disposition: attachment; filename="file.txt"
```

## Error handling

Is not needed. We guarantee the correctness of the input data.

---

* How can the filename be obtained?
* What happens if you try accessing the URL without explicitly downloading the file?
