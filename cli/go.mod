module github.com/johannesheesterman/lambda/cli

go 1.18

require github.com/antlr/antlr4/runtime/Go/antlr v0.0.0-20220527190237-ee62e23da966 // indirect

require github.com/johannesheesterman/lambda/parser v0.0.0
replace github.com/johannesheesterman/lambda/parser => ../parser/src
