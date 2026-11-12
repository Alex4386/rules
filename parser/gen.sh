# needs ANTLR 4.13.x to match the antlr4-go/antlr/v4 runtime (distro packages are often older)
antlr4 -Dlanguage=Go -visitor -no-listener JsonQuery.g4 -o ./