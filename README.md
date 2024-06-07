[![Go Report Card](https://goreportcard.com/badge/github.com/nikunjy/rules?style=flat-square)](https://goreportcard.com/report/github.com/nikunjy/rules)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/github.com/nikunjy/rules)](https://pkg.go.dev/github.com/nikunjy/rules)
[![Release](https://img.shields.io/github/v/release/nikunjy/rules?sort=semver&style=flat-square)](https://github.com/nikunjy/rules/releases/latest)
[![Release](https://img.shields.io/github/go-mod/go-version/nikunjy/rules?style=flat-square)](https://github.com/nikunjy/rules/releases/latest)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg?style=flat-square)](https://github.com/nikunjy/rules/commits)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub stars](https://img.shields.io/github/stars/nikunjy/rules?style=flat-square&label=Star&maxAge=2592000)](https://github.com/nikunjy/rules/stargazers/)

# Golang Rules Engine

Rules engine written in golang with the help of antlr.

This package will be very helpful in situations where you have a generic rule and want to verify if your values (specified using `map[string]interface{}`) satisfy the rule.

Here are some examples:

```go
  parser.Evaluate("x eq 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x == 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x lt 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x < 1", map[string]interface{}{"x": 1})
  parser.Evaluate("x gt 1", map[string]interface{}{"x": 1})

  parser.Evaluate("x.a == 1 and x.b.c <= 2", map[string]interface{}{
    "x": map[string]interface{}{
       "a": 1,
       "b": map[string]interface{}{
          "c": 2,
       },
    },
  })


  parser.Evaluate("y == 4 and (x > 1)", map[string]interface{}{"x": 1})

  parser.Evaluate("y == 4 and (x IN [1,2,3])", map[string]interface{}{"x": 1})

  parser.Evaluate("y == 4 and (x eq 1.2.3)", map[string]interface{}{"x": "1.2.3"})
```

### Extra function of this fork
This `rules` fork has an extra function which is more suitable for Firewall environment.  
The following operations were added.

* IP Address types
   - Supports both IPv4 and IPv6
* CIDR Notation support
* Regular Expression support

```go
// regular expression support
parser.Evaluate("x matches /^wel(.+)world$/g", map[string]interface{}{"x": "welcome to the mungtaeng-i world"}) // true
parser.Evaluate("x matches /^wel(.+)world$/g", map[string]interface{}{"x": "good-bye from the mungtaeng-i world!"}) // false

// IP Address types
parser.Evaluate("x eq 1.1.1.1", map[string]interface{}{"x": net.ParseIP("1.1.1.1")}) // true
parser.Evaluate("x == 2001:0db8:85a3:0000:0000:8a2e:0370:7334", map[string]interface{}{"x": net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")}) // true
parser.Evaluate("x ne 1.1.1.1", map[string]interface{}{"x": net.ParseIP("1.1.1.1")}) // false
parser.Evaluate("x != 2001:0db8:85a3:0000:0000:8a2e:0370:7334", map[string]interface{}{"x": net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")}) // false
parser.Evaluate("x in 10.0.0.0/8", map[string]interface{}{"x": net.ParseIP("10.0.0.1")}) // true
parser.Evaluate("x in 2001::8a2e:1/8", map[string]interface{}{"x": net.ParseIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334")}) // true

// Right side operands
parser.Evaluate("x eq y", map[string]interface{}{"x": 1, "y": 1}) // true
parser.Evaluate("x ne y.a", map[string]interface{}{"x": 1, "y": map[string]interface{}{ "a": 2 }}) // true
```

## Operations

All the operations can be written capitalized or lowercase (ex: `eq` or `EQ` can be used)

Logical Operations supported are `and` `or`

Compare Expression and their definitions

| expression | meaning                                         | 
------------|-------------------------------------------------
| eq         | equals to                                       |
| ==         | equals to                                       |
| ne         | not equal to                                    |
| !=         | not equal to                                    |
| lt         | less than                                       |
| <          | less than                                       |
| gt         | greater than                                    |
| >          | greater than                                    |
| le         | less than or equal to                           |
| <=         | less than or equal to                           |
| ge         | greater than or equal to                        |
| >=         | greater than or equal to                        |
| co         | contains                                        |
| sw         | starts with                                     |
| ew         | ends with                                       |
| in         | in a list or CIDR range (IP)                    |
| pr         | present, will be true if you have a key as true |
| not        | not of a logical expression                     |
| mt         | regular expression match (string)               |

## How to use it

Use your dependency manager to import `github.com/nikunjy/rules/parser`. This will let you parse a rule and keep the parsed representation around.
Alternatively, you can also use `github.com/nikunjy/rules` directly to call the root `Evaluate(string, map[string]interface{})` method.

I would recommend importing `github.com/nikunjy/rules/parser`

## How to extend the grammar

1. Please look at this [antlr tutorial](https://tomassetti.me/antlr-mega-tutorial/#setup-antlr), the link will show you how to setup antlr.
   The article has a whole lot of detail about antlr I encourage you to read it, you might also like [my blog post](https://medium.com/@nikunjyadav/generic-rules-engine-in-golang-using-antlr-d30a0d0bb565) about this repo.
2. After taking a look at the antlr tutorial, you can extend the [JsonQuery.g4 file](https://github.com/nikunjy/rules/blob/master/parser/JsonQuery.g4).
3. Compile the parser `antlr4 -Dlanguage=Go -visitor -no-listener JsonQuery.g4 -o ./` (Note: `-o` is the output directory, make sure all the stuff it generates is in the `parser` directory of the root repo folder)

[ci-img]: https://api.travis-ci.org/nikunjy/rules.svg?branch=master
[ci]: https://travis-ci.org/nikunjy/rules

## Note for myself for extending rules
[Note](NOTE.md)
