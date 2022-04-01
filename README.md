# Search Whole Directory
Search for a string hiding somewhere maybe in multiple files or deeply nested in the current directory.

## Usage
        go run . SEARCH_TERM
where SEARCH_TERM is the text to match in the current directory, and the source code is in the root of the current directory. Search term is <strong>case-sensitive</strong>

### Example
        go run . search
produced this output when run from the root of the developer's local repo:

![example screenshot](https://live.staticflickr.com/65535/51976097728_f0521f996b_z.jpg)

## Caveats
* This package relies on [golang.org/x/tools/godoc/util.IsText()](https://pkg.go.dev/golang.org/x/tools/godoc/util#IsText), to only search text files. One unusual aspect of IsText() is that it returns false for files that contain newlines, even .txt files which contain vertical spaces that are newline characters. Unfortunately for fans, forking util.IsText is way beyond the original scope of this dinky little project.
* If two matches are close enough together that they overlap in the first one's preview window, only that first one will be matched. The improvement to match both would be a great first issue to tackle in the bizarre event this package becomes useful to anyone.