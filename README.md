# Goclient

# Support Methods
## Get
`GoClient(TestStruct{}).Url("https://jsonplaceholder.typicode.com/todos/1").Get()`

## Post
`GoClient(TestStruct{}).Url("https://jsonplaceholder.typicode.com/posts").Post(TestStruct{
Title:  "Testing",
Body:   "Testing json",
UserId: 1,
})`