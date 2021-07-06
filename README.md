# Compiling and testing
just usual `go build` and `go test`

# Run
```./hashing google.com http://whatever.com```

The hashing program supports the `parallel` option which limits the concurrency to the value passed in this argument. By default the limit is 10
```./hashing -parellel 2 google.com http://whatever.com```

Also, the URLs can be `google.com` or `http://google.com`, there's a small preprocessing to check for URL format. Note that it only prepends http if it's not there (i.e., doesn't support https)
