# Compilador de Linguagem e Paradigmas, realizado em Go

[![Compilation Status](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)

![Diagrama Sintático](https://compiler-tester.insper-comp.com.br/ds?version=v2.1)

```ebnf
PROGRAM = { STATEMENT } ;
BLOCK = "{", { STATEMENT }, "}" ;
STATEMENT = (("Println", "(", BOOLEXPRESSION, ")") | ("for", BOOLEXPRESSION, BLOCK) | ("if", BOOLEXPRESSION, BLOCK, ("else", BLOCK) | ε) | (IDENTIFIER, "=", BOOLEXPRESSION) | BLOCK | ε), "\n" ;
BOOLEXPRESSION = BOOLTERM, { "||", BOOLTERM } ;
BOOLTERM = RELEXPRESSION, { "&&", RELEXPRESSION } ;
RELEXPRESSION = EXPRESSION, [ ("==" | "<" | ">"), EXPRESSION ] ;
EXPRESSION = TERM, { ("+" | "-"), TERM } ;
TERM = FACTOR, { ("*" | "/"), FACTOR } ;
FACTOR = NUMBER | IDENTIFIER | ("!" | "+" | "-"), FACTOR | "(", BOOLEXPRESSION, ")" | "Scanln", "(", ")" ;
IDENTIFIER = LETTER, { LETTER | DIGIT | "_" } ;
NUMBER = DIGIT, { DIGIT } ;
LETTER = "a" | "b" | ... | "z" | "A" | "B" | ... | "Z" ;
DIGIT = "0" | "1" | ... | "9" ;
```