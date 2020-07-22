# filt

Filter lines while watching current line every second.

This is a command you can use in your pipeline to search for a given string in the output while watching the output every second.

As the command processes lines from stdin it will output the line number and current line every second to stderr.

When the command finds a line containing the supplied substring, it will output that line to stdout.

## Installing

```
go get github.com/norganna/filt
go install github.com/norganna/filt
```

## Using

```
command | filt 'findText'
```

