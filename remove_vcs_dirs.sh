#!/bin/bash
find .vendor -name .git -exec rm -fr '{}' \;
find .vendor -name .hg -exec rm -fr '{}' \;
find .vendor -name .bzr -exec rm -fr '{}' \;
find .vendor -name '.gitignore' -type f -delete
