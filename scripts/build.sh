#!/usr/bin/env bash

dep ensure -v

echo "Compiling functions to bin/ ..."

for service in services/*; do
  for file in ${service}/*; do
    dirname=$(basename -- "$service")
    filename=$(basename -- "$file")
    extension="${filename##*.}"
    filename="${filename%.go}"
    if [ $extension != "yml" ]; then
      if GOOS=linux go build -ldflags="-s -w" -o ${service}/bin/${filename} ${file}; then
        echo "✓ Compiled $service"
      else
        echo "✕ Failed to compile $service!"
        exit 1
      fi
    fi
  done
done

echo "Done."