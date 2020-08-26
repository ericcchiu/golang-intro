## Lab 1 Arrays and Slices

Work in `arrayslice.go`.

Make all the tests pass by following the instructions in the comments of the file.

To test, in the directory with the file, run: `go test .`

---
Continuous testing options:

1. `watch go test .`
2. `ginkgo watch test .`
___

## Lab 2 Maps

Work in `lab2/maps.go`

Make all the tests pass.

Lab 1.1 State Capitals

```sh
Given the state and capitals
    Columbus, Ohio
    Austin, Texas
    Atlanta, Georgia
    Juneau, Alaska

Then a map of the state as a key, and the capital as a value, should be returned.
```

Lab 1.2 Find Missing Keys

```sh
Given a map of integers called keycheck
And a string called key

When the key is found in keycheck
Then Return true

When the key is not found in keycheck
Then return false
```

The `return` followed by a literal or value will return a value to the test

To test, in the directory with the file, run: `go test .`

Continuous testing options:

1. `watch go test .`
2. `ginkgo watch test .`
___

## Lab 3: Map String Buzz

Work in `lab3/mapstringbuzz.go`

Make the tests pass

```sh
Scenario: iterate over a string variable input and map fizz, buzz or fizzbuzz
based on the rune values

Given a string named input and a map named stringbuzz
When the string is iterated over
And its character's rune value is divisible by 3
Then add the character as a string to the map with fizz as the value

When the character's rune value is divisible by 5
Then add the character as a string to the map with buzz as the value

When the character's rune value is divisible by both 3 & 5
Then add the character as a string to the map with fizzbuzz as the value

When the character's rune value is not divisible by either 3 or 5
Then the character should be omitted from the map
```

You can use `string(runevar)` to convert a rune to a string.

To test, in the directory with the file, run: `go test .`

___
Continuous testing options:

1. `watch go test .`
2. `ginkgo watch test .`
