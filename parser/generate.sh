#!/bin/sh

if [ ! -f "./antlr-4.13.2-complete.jar" ]; then
   wget https://www.antlr.org/download/antlr-4.13.2-complete.jar
fi


alias antlr4='java -Xmx500M -cp "./antlr-4.13.2-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -no-visitor -listener -package parser *.g4
