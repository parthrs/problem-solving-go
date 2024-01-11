Solutions to problems and puzzles on the web

#### Testing
To test solutions against test cases
`make test`

The test cases for most problems are in the same package
since we want to test unexported methods as well.
https://stackoverflow.com/questions/19998250/proper-package-naming-for-testing-with-the-go-language

##### Why is there main?
There are exercises that don't necessarily require tests but rather a working "demo".
For e.g. the package installer with dependencies (check `exercise misc/package-installer-with-deps-*.go`)
