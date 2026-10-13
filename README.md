# Compilador de Linguagem e Paradigmas, realizado em Go

[![Compilation Status](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)

![Diagrama Sintático](https://compiler-tester.insper-comp.com.br/ds?version=v2.2)

```ebnf
PROGRAM = { STATEMENT } ;
BLOCK = "{", "\n", { STATEMENT }, "}" ;
STATEMENT = ("var", IDENTIFIER, TYPE, (ε | "=", BOOLEXPRESSION) | ("Println", "(", BOOLEXPRESSION, ")") | ("for", BOOLEXPRESSION, BLOCK) | ("if", BOOLEXPRESSION, BLOCK, (ε | "else", BLOCK) | ) | BLOCK | IDENTIFIER, "=", BOOLEXPRESSION | ε), "\n" ;
BOOLEXPRESSION = BOOLTERM, { "||", BOOLTERM } ;
BOOLTERM = RELEXPRESSION, { "&&", RELEXPRESSION } ;
RELEXPRESSION = EXPRESSION, { ("<" | "==" | ">"), EXPRESSION } ;
EXPRESSION = TERM, { ("+" | "-"), TERM } ;
TERM = FACTOR, { ("*" | "/"), FACTOR } ;
FACTOR = NUMBER | BOOLEAN | STRING | IDENTIFIER | ("!" | "+" | "-"), FACTOR | "(", BOOLEXPRESSION, ")" | "Scanln", "(", ")" ;
TYPE = "string" | "int" | "bool" ;
BOOLEAN = "true" | "false" ;
STRING = '"', { "..." }, '"' ;
IDENTIFIER = LETTER, { LETTER | DIGIT | "_" } ;
NUMBER = DIGIT, { DIGIT } ;
LETTER = "a" | "b" | ... | "z" | "A" | "B" | ... | "Z" ;
DIGIT = "0" | "1" | ... | "9" ;
```