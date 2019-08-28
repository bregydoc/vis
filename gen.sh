#!/bin/sh

solc --abi --overwrite rgy.sol -o .
solc --bin --overwrite rgy.sol -o .

if [ -d "rgy" ]; then
  echo "rgy dir already exist"
else
  echo "crating rgy dir"
  mkdir rgy
fi

abigen --bin=RGYx.bin --abi=RGYx.abi --pkg=rgy --out=rgy/rgy.go


