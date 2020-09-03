# Example-MP-Solution
An example solution to MP assignments in CSCI3359

**Note: This solution is only tangentially related to the assignment and has almost no reflection on the steps _you_ need to take to solve MP0**

----
# To Run
Unlike MP0, this example uses goroutines instead of multiprocess parallelism, so it can be 
run relatively easily. 

```bash
go run main.go -f example.json
``` 

**EVEN BETTER** build your go files for darwin/amd64 and include the executable in your final solution.
This is the architecture of my laptop and will allow me to easily run your solution without installing the
dependencies you used. 

```bash
export GOOS=darwin 
export GOARCH=amd64 
go build main.go
```

Then to run it, use
```bash
./Example-MP-Solution -f example.json
```

should produce the following output:
```text
Title: Soll ich heute etwas Brot kaufen?
To: H. Muster
From: F. Muster
Date: 2018-01-22 14:28:00 +0000 GMT
Content: H- ich gehe gerade zum Laden. Brauchen wir Brot?
         -F Muster
```

----
# Structure and Design
### Message
There is a message struct in `/message/message.go` 
which contains the following fields:

```go
type Message struct{
	To string
	From string
	Title string
	Content string
	Date time.Time
}
```
It can also initialize a `Message` from a JSON file via `MessageFromJsonFile`. 

_Note:_ the `Date` field in the JSON file must conform to `Jan 2, 2006 at 3:04pm (MST)` format.

### Processes

Code can be found in `/processes/`

To simulate MP0, I created two mock processes with their own wrappers of Message. 
`ProcessA` is responsible for sending an initialized message over a channel
to `ProcessB`, which is responsible for receiving the message and printing it. 

Since we want to simulate TCP communication, I'm using channels with goroutines.

`ProcessA` is created from a JSON file via `ProcessAFromJsonFile` and `ProcessB` 
can remain unitialized as its message field gets populated from the channel. 
