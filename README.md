# mindfck

A simple language that transpiles to brainfuck:

```
byte a
byte b
a = 3 + a
a = 33 + a
a = a + 2
byte c
a = a + 21
a = a + 2
b = 10 + a
c = a + b
print c
```

Transpiles to:

```brainfuck
>>[-]+++>>[-]<[-]<<<[>>>+>+<<<<-][-]>>>>[<<<<+>>>>-]>>[-]<<[-]<<[>>+>>+<<<<-][-]>>>>[<<<<+>>>>-]>[-]<<<<[>>>>+<<<+<-][-]>>>>[<<<<+>>>>-]<<[-]<[>+<-]<<<<[-]>>>>>[<<<<<+>>>>>-]>[-]+++++++++++++++++++++++++++++++++<<[-]>>>[-]<<<<<<<[>>>>>>>+<<<+<<<<-][-]>>>>[<<<<+>>>>-]<[-]>>[-]>[<+<<+>>>-][-]<<<[>>>+<<<-]>[-]>>>[<<<+>+>>-][-]<<<[>>>+<<<-]<<[-]>>>[<<<+>>>-]<<<<<[-]>>[<<+>>-]>>[-]<[-]<<<[>>>+>+<<<<-][-]>>>>[<<<<+>>>>-]>>>[-]++<<[-]<<<[-]>[<+>>>+<<-][-]>>[<<+>>-]<[-]>>>[<<<+<<+>>>>>-][-]<<<[>>>+<<<-]>>[-]<<<<[>>>>+<<<<-]<<[-]>>>>>>[<<<<<<+>>>>>>-]<<<<[-]>>[-]<<<<[>>>>+<<+<<-][-]>>[<<+>>-]>[-]+++++++++++++++++++++<[-]>>>>>[-]<<<[>>>+<<<<<+>>-][-]<<[>>+<<-]>>>>>>[-]<<<<<[>>>>>+<+<<<<-][-]>>>>>[<<<<<+>>>>>-]<<[-]>[<+>-]<<<<<<<[-]>>>>>>[<<<<<<+>>>>>>-]>>[-]<<<<<<[-]<<[>>+>>>>>>+<<<<<<<<-][-]>>>>>>>>[<<<<<<<<+>>>>>>>>-]<<<<<[-]++>>>>[-]<[-]<<<<[>>>>+>+<<<<<-][-]>>>>>[<<<<<+>>>>>-]>[-]<<<<<[>>>>>+<<+<<<-][-]>>>>>[<<<<<+>>>>>-]<<<<[-]>>[<<+>>-]<<<<<<[-]>>>>[<<<<+>>>>-]>>>[-]++++++++++<[-]>>[-]<<<<<<<<[>>>>>>>>+<<+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]<<<[-]>[-]>>>[<<<+<+>>>>-][-]<<<<[>>>>+<<<<-]>>>[-]>>[<<+<<+>>>>-][-]<<[>>+<<-]<<<<[-]>>[<<+>>-]<<<[-]>[<+>-]>>>>[-]<<<[-]<<<[>>>+>>>+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]>[-]>[-]<<<<<<<[>>>>>>>+<+<<<<<<-][-]>>>>>>[<<<<<<+>>>>>>-]<<<<<[-]>>>>[-]<<<[>>>+<<<<+>-][-]<[>+<-]>>>>>[-]>[<+<+>>-][-]<[>+<-]<<<[-]>>[<<+>>-]<[-]<[>+<-]>>>[-]<<<<<[-]>>>[<<<+>>>>>+<<-][-]>>[<<+>>-]<<<<<.
```

## Usage

To compile:

```
minfck examples/abc.mf
```

To compile and execute the brainfuck code:

```
minfck example.mf --run
```

## FAQ

Q: Is this useful?  
A: No  
Q: Is it fun?  
A: Also no

## Development

Mindfck is written in go.

**Build**

```
go build .
```

**Run**

```
go run .
```

**Tests**

```
go test ./...
```

### Grammar

Grammar is written in antlr4 (`parser/mindfck.g4`). If it is updated, files need to be regenerated.

Install antlr4-tools to work with the parser

```
pip3 install antlr4-tools
```

You'll also need java-11-jre or openjdk-11-jre installed

To regenerate the parser, do:

```
go generate
```

This will execute antlr4, equivalent to: `antlr4 parser/antlr/mindfck.g4 -Dlanguage=Go`

To parse something manually in the command line:

```
antlr4-parse parser/antlr/mindfck.g4 statements -gui
```

## Resources

These resources where used as inspiration / tools for this project

-   https://github.com/LucasMW/Headache
-   https://esolangs.org/wiki/Brainfuck_algorithms
-   https://www.nayuki.io/page/brainfuck-interpreter-javascript
-   https://gist.github.com/roachhd/dce54bec8ba55fb17d3a
