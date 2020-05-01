# go-web-server-tips

This project demonstrates some code practices I use when writing web servers in Go.

To start, run:

```shell script
go run main.go
```

Blog post: [dev.to/chidiwilliams/writing-cleaner-go-web-servers-3oe4](https://dev.to/chidiwilliams/writing-cleaner-go-web-servers-3oe4)

## Tips

- [x] Use clean architecture

  - Good code organization/folder structure
  - Decouple dependencies

- [x] Extend HTTP handler

  - Handle errors in one location

- [x] Standardized response format

- [x] Create custom errors for client errors

  - Clean error handling

- [ ] ozzo-validator with custom validator?

  - Struct validation outside controller

- [ ] Integration testing with testify?
