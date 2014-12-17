#!/bin/bash
rm -rf .vendor/src
rm Goopfile.lock 
goop install
find .vendor/src -name .git -exec rm -fr '{}' \; 2>/dev/null
find .vendor/src -name .hg -exec rm -fr '{}' \; 2>/dev/null
find .vendor/src -name .bzr -exec rm -fr '{}' \; 2>/dev/null
find .vendor/src -name '.gitignore' -type f -delete 2>/dev/null
