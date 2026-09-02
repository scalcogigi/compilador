# Compilador de Linguagem e Paradigmas, realizado em Go

[![Compilation Status](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)](https://compiler-tester.insper-comp.com.br/svg/scalcogigi/compilador)

![Diagrama Sintático](https://compiler-tester.insper-comp.com.br/ds?version=v1.1)

```ebnf

EXPRESSION = TERM, { ("+" | "-"), TERM } ;
TERM = FACTOR, { ("*" | "/"), FACTOR } ;
FACTOR = NUMBER | ("+" | "-"), FACTOR | "(", EXPRESSION, ")" ;
NUMBER = DIGIT, {DIGIT} ;
DIGIT = "0" | "1" | ... | "9" ;

```