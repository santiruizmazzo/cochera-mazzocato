#!/bin/bash

# Optimiza una imagen

INPUT=$1

# Si no se proporciona OUTPUT, generarlo automáticamente
if [ -z "$2" ]; then
    # Extraer el directorio, nombre base y extensión
    DIR=$(dirname "$INPUT")
    BASENAME=$(basename "$INPUT" | sed 's/\.[^.]*$//')
    EXTENSION="${INPUT##*.}"
    
    # Generar nombre automático con sufijo _optimized
    OUTPUT="$DIR/${BASENAME}_optimized.$EXTENSION"
else
    OUTPUT=$2
fi

echo "Optimizando: $INPUT -> $OUTPUT"
bun svgo $INPUT --pretty -p 0 -o $OUTPUT


