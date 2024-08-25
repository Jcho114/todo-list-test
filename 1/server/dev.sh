#!/bin/bash

echo "Starting tailwind cli"
npx tailwindcss -i ./modules/templates/input.css -o ./modules/templates/public/output.css --watch --minify