@echo off
SET TEST_DIR=./src

echo Running tests...
gotestsum --format testdox "%TEST_DIR%/..."