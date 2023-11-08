#!/bin/bash

# Check if three arguments are provided
if [ $# -ne 3 ]; then
    echo "Usage: create_folder.sh <leet_id> <difficulty> <language>"
    exit 1
fi

# Get the command-line arguments
leet_id="$1"
difficulty="$2"
language="$3"

#Create the directory structure
mkdir -p "$difficulty/$leet_id"

# Change directory to the new folder
cd "$difficulty/$leet_id" || exit 1

# Create a Go file with a basic template
cat > "$leet_id""$language" << EOF
package $difficulty

EOF

# Create a notes.md file with additional content
cat > notes.md << EOF
# Notes   

Runtime   
_ms   
Beats __% of users with __   

Memory   
_MB   
Beats % of users with __   

Notes   
EOF

echo "Folder '$leet_id' created with '$leet_id''$language' and a notes.md"
echo "test"

exit 0