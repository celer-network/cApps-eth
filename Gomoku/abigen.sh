#!/bin/sh
# Generate go binding
abigen --sol contracts/Gomoku.sol --pkg gomoku --out gomoku.go