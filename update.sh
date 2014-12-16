#!/bin/bash
rm -rf .vendor
rm Goopfile.lock 
goop install
find .vendor -name .git -exec rm -fr '{}' \; 2>/dev/null
find .vendor -name .hg -exec rm -fr '{}' \; 2>/dev/null
find .vendor -name .bzr -exec rm -fr '{}' \; 2>/dev/null
find .vendor -name '.gitignore' -type f -delete 2>/dev/null
