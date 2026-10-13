# Compilador de Linguagem e Paradigmas, realizado em Go

[![Compilation Status](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)

![Diagrama Sintático](https://compiler-tester.insper-comp.com.br/ds?version=v2.2)

```ebnf
PROGRAM = { STATEMENT } ;
BLOCK = "{", "\n", { STATEMENT }, "}" ;
STATEMENT = BLOCK | (("Println", "(", BOOLEXPRESSION, ")") | ("for", BOOLEXPRESSION, BLOCK) | ("if", BOOLEXPRESSION, BLOCK, ("else", BLOCK) | ε) | (IDENTIFIER, "=", BOOLEXPRESSION) | VARDEC | ε), "\n" ;
VARDEC = "var", IDENTIFIER, TYPE, [ "=", BOOLEXPRESSION ] ;
BOOLEXPRESSION = BOOLTERM, { "||", BOOLTERM } ;
BOOLTERM = RELEXPRESSION, { "&&", RELEXPRESSION } ;
RELEXPRESSION = EXPRESSION, [ ("<" | "==" | ">"), EXPRESSION ] ;
EXPRESSION = TERM, { ("+" | "-"), TERM } ;
TERM = FACTOR, { ("*" | "/"), FACTOR } ;
FACTOR = ("!" | "+" | "-"), FACTOR | POWER ;
POWER = PRIMARY, [ "**", FACTOR ] ;
PRIMARY = NUMBER | STRING | BOOLEAN | IDENTIFIER | "(", BOOLEXPRESSION, ")" | "Scanln", "(", ")" ;
TYPE = "int" | "string" | "bool" ;
BOOLEAN = "true" | "false" ;
STRING = "\"", { "..." }, "\"" ;
IDENTIFIER = LETTER, { LETTER | DIGIT | "_" } ;
NUMBER = DIGIT, { DIGIT } ;
LETTER = "a" | "b" | ... | "z" | "A" | "B" | ... | "Z" ;
DIGIT = "0" | "1" | ... | "9" ;
```