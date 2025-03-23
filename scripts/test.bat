@echo off
SET TEST_DIR=./src

echo Running Go tests...
go test -v "%TEST_DIR%/..."

echo Running tests with coverage report...
go test -cover "%TEST_DIR%/..."