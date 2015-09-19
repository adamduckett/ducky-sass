#!/bin/sh

echo "Installing clean Sass project, won't be a sec!"

mv .git/modules .       && \
rm -rf .git/            && \
npm install             && \
rm go

echo "...Aaaand done! Have fun with your new Sass project :)"

gulp
