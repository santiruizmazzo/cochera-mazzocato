#!/bin/bash

# Genera un sprite SVG a partir de una carpeta de SVGs

INPUT=$1

if [ -z "$1" ]; then
    INPUT="public/assets/icons"
fi

OUTPUT=public/assets/icons.svg

echo '<svg xmlns="http://www.w3.org/2000/svg">' >>$OUTPUT
sed s/svg/symbol/g $INPUT/*.svg >>$OUTPUT
echo '</svg>' >>$OUTPUT