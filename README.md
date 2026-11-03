# Compilador de Linguagem e Paradigmas, realizado em Go

[![Compilation Status](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)

![Diagrama Sintático](https://compiler-tester.insper-comp.com.br/ds?version=v2.3)

```ebnf
PROGRAM = { VARDECL | FUNCDECL | "\n" } ;
BLOCK = "{", "\n", { STATEMENT }, "}" ;
VARDECL = "var", IDENTIFIER, TYPE, (ε | "=", BOOLEXPRESSION), "\n" ;
FUNCDECL = "func", IDENTIFIER, "(", (ε | PARAMLIST), ")", (ε | TYPE), BLOCK, "\n" ;
PARAMLIST = PARAM, { ",", PARAM } ;
PARAM = IDENTIFIER, TYPE ;
STATEMENT = VARDECL | ("Println", "(", BOOLEXPRESSION, ")", "\n") | ("for", BOOLEXPRESSION, BLOCK, "\n") | ("if", BOOLEXPRESSION, BLOCK, (ε | "else", BLOCK), "\n") | BLOCK | (IDENTIFIER, "=", BOOLEXPRESSION, "\n") | (FUNCCALL, "\n") | ("return", BOOLEXPRESSION, "\n") | "\n" ;
FUNCCALL = IDENTIFIER, "(", (ε | ARGLIST), ")" ;
ARGLIST = BOOLEXPRESSION, { ",", BOOLEXPRESSION } ;
BOOLEXPRESSION = BOOLTERM, { "||", BOOLTERM } ;
BOOLTERM = RELEXPRESSION, { "&&", RELEXPRESSION } ;
RELEXPRESSION = EXPRESSION, (ε | ("<" | "==" | ">"), EXPRESSION) ;
EXPRESSION = TERM, { ("+" | "-"), TERM } ;
TERM = FACTOR, { ("*" | "/"), FACTOR } ;
FACTOR = ("!" | "+" | "-"), FACTOR | PRIMARY, (ε | "\*\*", FACTOR) ;
PRIMARY = NUMBER | BOOLEAN | STRING | IDENTIFIER | FUNCCALL | "(", BOOLEXPRESSION, ")" | "Scanln", "(", ")" ;
TYPE = "string" | "int" | "bool" ;
BOOLEAN = "true" | "false" ;
STRING = '"', { "..." }, '"' ;
IDENTIFIER = LETTER, { LETTER | DIGIT | "\_" } ;
NUMBER = DIGIT, { DIGIT } ;
LETTER = "a" | "b" | ... | "z" | "A" | "B" | ... | "Z" ;
DIGIT = "0" | "1" | ... | "9" ;
```