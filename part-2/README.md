# HTTP

> Part 2. POST requests

You are to create a script `part-2.sh` which sends a POST request to the server and parses its response.

Request body is defined by the following class:
```kotlin
data class pair(
    val x : Int,
    val y : Int,
)
```

You must support three content types:
- JSON
- YAML
- XML

> **Important!** XML body must be wrapped in `<req version="1.1">...</req>`.

Response body is defined by this class:
```kotlin
data class response(
    val result : Int,
)
```

The script must print only one number: the value of `result`.

**Synopsis:**
```cmd
./part-2.sh (-x|-y|-j) <URL> <x> <y>

Options:
   -x        XML format
   -y        YAML format
   -j        JSON format
```

**Output:**
```cmd
56
```
## Error handling

Is not needed. We guarantee the correctness of the input data.


---

* How does the server understand which format is used?
* Which method is used by default in `curl`?
* What is the difference between `GET` & `POST`?
