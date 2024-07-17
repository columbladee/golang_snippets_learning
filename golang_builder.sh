#!/bin/bash

#DOCS: https://golang.org/doc/install/source#environment

#Shell script to build golang files on different arch/os

show_usage() {
	echo "Usage: $0 <source_file> <binary_name> (e.g. main.go)"  
	echo "<binary_name>: The name of your binary (e.g. main)" 
	exit 1
}

if [ $# -ne 2 ]; then
	show_usage
fi

options=("macOS (ARM64)" "macOS (AMD64)" "Windows (AMD64)" "Linux (AMD64)" "Linux (ARM64)")
platforms=("darwin" "darwin" "windows" "linux" "linux")
arches=("arm64" "amd64" "amd64" "amd64" "arm64")

echo "Select the target platform and architecture:"
for i in "${!options[@]}"; do
    echo "$((i+1)). ${options[$i]}"
done

read -p "Enter the option number: " option

if [[ $option -lt 1 || $option -gt ${#options[@]} ]]; then
    echo "Invalid option"
    exit 1
fi

# Set GOOS and GOARCH based on the user selection
goos=${platforms[$((option-1))]}
goarch=${arches[$((option-1))]}

target="$1"

# Validate the provided source file
if ! [ -f "$target" ]; then
    echo "Source file '$target' does not exist."
    exit 1
fi

binaryName="$2"

output="$binaryName"
output="$(basename $target | sed 's/\.go//')"

[[ "windows" == "$goos" ]] && output="$output.exe"

destination="$(dirname $target)/builds/$goos/$goarch/$output"

echo -e "\e[00;33mGOOS=$goos GOARCH=$goarch go build -ldflags \"-w\" -o $destination $target\e[00m"
GOOS=$goos GOARCH=$goarch go build -ldflags "-w" -o "$destination" "$target"

if [ $? -eq 0 ]; then
    echo "Build successful! Binary created at: $destination"
else
    echo "Build failed!"
    exit 1
fi
