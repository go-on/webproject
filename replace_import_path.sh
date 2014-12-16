# run inside bash
sed --in-place='.orig' -e 's|"github\.com/metakeule/watcher/|"gopkg.in/metakeule/watcher.v1/|' $(find . -name '*.go')
find . -name '*.go.orig' -type f -delete
