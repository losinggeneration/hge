#!/bin/bash 
set -o nounset                              # Treat unset variables as an error
mkdir -p docs/helpers

for i in . gfx ini input rand resource sound timer; do
	[ $i = '.' ] && i="" && pkg='hge'
	[ "$i" ] && pkg="hge_$i"
	echo "creating $pkg.html"
	cat << EOF > docs/$pkg.html
---
title: $pkg
layout: godocs
---
EOF
	godoc --html=true github.com/losinggeneration/hge-go/${pkg/_/\/} >> docs/$pkg.html
done

# helpers
for i in animation  color  distortionmesh  font  gui  guictrls  particle  rect  sprite  strings  vector; do
	echo "creating helpers/$i.html"
	cat << EOF > docs/helpers/$i.html
---
title: helpers/$i
layout: godocs
---
EOF
	godoc --html=true github.com/losinggeneration/hge-go/helpers/$i >> docs/helpers/$i.html
done

# Legacy
echo "creating legacy.html"
cat << EOF > docs/legacy.html
---
title: legacy
layout: godocs
---
EOF
godoc --html=true github.com/losinggeneration/hge-go/legacy >> docs/legacy.html
