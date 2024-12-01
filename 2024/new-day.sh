#!/bin/bash

# Check if a number parameter is provided
if [ $# -ne 1 ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

# Store the number parameter
num=$1

# Create folder name with the number
folder_name="day-${num}"

echo -e "\033[90mSetting up day ${num}"

# Create the folder
mkdir -p "$folder_name"
echo -e "Created folder ./${folder_name}"

# Create our files
echo "package main" > "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "import (" >> "$folder_name/main.go"
echo "    \"github.com/dickeyy/adventofcode/2024/utils\"" >> "$folder_name/main.go"
echo ")" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "func main() {" >> "$folder_name/main.go"
echo "    utils.ParseFlags()" >> "$folder_name/main.go"
echo "    p := utils.GetPart()" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "    i := utils.ReadFile(\"./input.txt\")" >> "$folder_name/main.go"
echo "    utils.Output(day${num}(i, p))" >> "$folder_name/main.go"
echo "}" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "func day${num}(input string, part int) int {" >> "$folder_name/main.go"
echo "    return 0" >> "$folder_name/main.go"
echo "}" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"

touch "$folder_name/input.txt"

echo -e "\033[90mCreated boilerplate for day ${num} in ./${folder_name}"

# done :)
cd "$folder_name"
echo -e "\n\033[35mGood luck \033[90mon day ${num}!"