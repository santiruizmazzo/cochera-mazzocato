#!/bin/bash

# Optimiza una carpeta de imágenes

INPUT=$1

if [ -z "$1" ]; then
    INPUT="public/assets/icons"
fi

OUTPUT=$INPUT

bun svgo -f $INPUT -o $OUTPUT --config=svgo.config.js --multipass --pretty -p 2

