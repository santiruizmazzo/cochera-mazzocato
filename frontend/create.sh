#!/bin/bash

# Crea un mapa de SVG con todos los iconos de la carpeta

INPUT=$1
OUTPUT=$1.svg

echo '<?xml version="1.0" encoding="utf-8"?>' >$OUTPUT
echo '<svg xmlns="http://www.w3.org/2000/svg">' >>$OUTPUT
echo '    <defs>' >>$OUTPUT
sed s/svg/symbol/g $INPUT/*.svg >>$OUTPUT
echo '    </defs>' >>$OUTPUT
echo '</svg>' >>$OUTPUT