# Error Handling in GO
These are just my notes to Error Handling in Go. For the full context I would suggest to checkout Go Guild Meetup Video which will tell you how to process code and following notes.

## 01 The Official Way
```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
```
Or throw new err directly `errors.New("msg")` / `fmt.Errorf("something")`


## 02 Predefined Errors
```go
type ResultError struct {}

func (e *ResultError) Error() string {
	return "Wrong Result"
}
```
Then Invoke it like so `&ResultError{}`

## 03 Custom Errors
```go
type CustomError struct {
  data string
}

func (e *CustomError) Error() string {
  return fmt.Sprintf("Error occured due to... %s", e.data)
}
```

Then log the custom err `log.Fatal(&CustomError{"nothing"})`


## 04 - Simplifying Error Handling
Error handling can be quite repetitive

**Problem:** http package doesn't understand function that returns err

So we will implement new type which returns errror like this: `type appHandler func(http.ResponseWriter, *http.Request) error`


we will extend ServerHTTP interface with reciever of type appHandler
```go
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), 500)
	}
}
```

* everytime ServeHTTP fails inside appHandler it will do the following error handling
* then rewrite `http.HandleFunc()` to `http.Handle()` and wrap errorGiver to appHandler

---
Presentation can be found here: [Google Slides](https://docs.google.com/presentation/d/1PJH9xIYHuRr-cF9kixnmsebNMZMa82vcjVFNMqR5_-A/edit?usp=sharing) 
 
Youtube Video: TBA
