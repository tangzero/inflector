# Inflector 

[![Build Status](https://travis-ci.org/tangzero/inflector.svg?branch=master)](https://travis-ci.org/tangzero/inflector) [![GoDoc](https://godoc.org/tangzero/inflector?status.svg)](https://godoc.org/github.com/tangzero/inflector) [![Coverage Status](https://coveralls.io/repos/github/tangzero/inflector/badge.svg?branch=master)](https://coveralls.io/github/tangzero/inflector?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/tangzero/inflector)](https://goreportcard.com/report/github.com/tangzero/inflector)

Inflector is a set of functions aimed to transform strings.
You can pluralize, singularize and do other transformations.

### Examples

```go
inflector.Pluralize("person") // "people"
inflector.Singularize("posts") // "post"
```

For more examples, take a look on the [examples_test.go](http://github.com/tangzero/inflector/blob/master/examples_test.go).

### Installation

`$ go get github.com/tangzero/inflector`

### Update

`$ go get -u github.com/tangzero/inflector`


### License

	MIT License
	
	Copyright (c) 2017 Jairo Luiz aka TangZero
	
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	
	The above copyright notice and this permission notice shall be included in all
	copies or substantial portions of the Software.
	
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
	SOFTWARE.
