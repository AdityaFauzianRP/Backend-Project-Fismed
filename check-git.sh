#!/bin/bash

# Fetch the latest updates from the remote repository
git fetch origin

# Get the latest commit hash from the remote repository
remote_commit=$(git rev-parse origin/main)

# Get the latest commit hash from the local repository
local_commit=$(git rev-parse HEAD)

# Compare the commit hashes
if [ "$remote_commit" != "$local_commit" ]; then
    echo "There are new commits in the remote repository."
else
    echo "Your local repository is up-to-date."
fi

