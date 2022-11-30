/*
Package font implements a helper for writing digital clock symbols.

It is inspired by [painter].

[painter]: https://github.com/guillaumebreton/gone/painter
*/
package font

import "strings"

var colon = `
..
#.
..
#.
..
`

var zero = `
######.
#....#.
#....#.
#....#.
######.
`
var one = `
.....#.
.....#.
.....#.
.....#.
.....#.
`
var two = `
######.
.....#.
######.
#......
######.
`

var three = `
######.
.....#.
...###.
.....#.
######.
`

var four = `
#......
#......
#...#..
######.
....#..
`

var five = `
######.
#......
######.
.....#.
######.
`

var six = `
######.
#......
######.
#....#.
######.
`

var seven = `
######.
.....#.
.....#.
.....#.
.....#.
`

var height = `
######.
#....#.
######.
#....#.
######.
`

var nine = `
######.
#....#.
######.
.....#.
######.
`

// smallFont defines the font use to display the timer on termbox
var smallFont = map[rune][][]rune{
	':': asArray(colon),
	'1': asArray(one),
	'2': asArray(two),
	'3': asArray(three),
	'4': asArray(four),
	'5': asArray(five),
	'6': asArray(six),
	'7': asArray(seven),
	'8': asArray(height),
	'9': asArray(nine),
	'0': asArray(zero),
}

// Convert a character as an array of rune
func asArray(chars string) [][]rune {
	result := [][]rune{}
	line := []rune{}
	str := strings.TrimPrefix(chars, "\n")
	for _, c := range str {
		if c == '\n' {
			result = append(result, line)
			line = []rune{}
		} else {
			line = append(line, c)
		}
	}
	return result
}
