#!/bin/sh

echo "Installing a clean Sass project, won't be a sec..."

rm -rf .git/            && \
npm install             && \
rm go

echo "...And we're done. Enjoy :)"

gulp
