# WIP

Rewrite of Predikit in Golang, gonna take awhile

(It's so much easier with Antlr4 + Go)

A note on testing:
- The code has changed so much that I got tired of constantly rewriting my tests. Once things solidify a bit, I will go back and test as much as possible.

## TODO

- [ ] recursive defaults (broken!)
- [ ] ALL the tests
- [ ] NOT
- [ ] move all error handling "to the top" (or at least in one place)
- [ ] Path type + env var handling
- [ ] acceptsChildren for groups
- [ ] terminology: test / check / assert / control
- [ ] = foo etc
- [ ] main exit code
- [ ] can't have empty checks `test foo {}`
- [ ] hook results to ChkResult
- [ ] hook params
- [ ] global config?
- [ ] retries / retry_delay validation at compile

# License

This project is licensed under the Apache License 2.0

see [LICENSE](LICENSE)
