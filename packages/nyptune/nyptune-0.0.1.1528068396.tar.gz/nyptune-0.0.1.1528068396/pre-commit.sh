#!/bin/bash
set -ex
chmod +x pre-commit.sh
ln -nfs ../../pre-commit.sh .git/hooks/pre-commit
chmod +x .git/hooks/pre-commit
black .
python -m pytest