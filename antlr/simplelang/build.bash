#!/bin/bash

antlr4 -Dlanguage=Go -visitor *.g4 -o parser
