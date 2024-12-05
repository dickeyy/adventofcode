#!/bin/bash

# Check if a day parameter is provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <year> <day>"
    exit 1
fi

# Store the parameters
year=$1
day=$2

# Create folder name with the day
folder_name="${year}/day-${day}"
input_folder_name="inputs/${year}/day-${day}"

echo -e "\033[90mSetting up ${year} day ${day}"

# Create the folder
mkdir -p "$folder_name"
mkdir -p "$input_folder_name"
echo -e "Created folder ${folder_name}"

# Create our files
echo "package main" > "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "import (" >> "$folder_name/main.go"
echo "    \"github.com/dickeyy/adventofcode/${year}/utils\"" >> "$folder_name/main.go"
echo ")" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "func main() {" >> "$folder_name/main.go"
echo "    utils.ParseFlags()" >> "$folder_name/main.go"
echo "    p := utils.GetPart()" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "    i := utils.ReadFile(\"../../${input_folder_name}/input.txt\")" >> "$folder_name/main.go"
echo "    utils.Output(day${day}(i, p))" >> "$folder_name/main.go"
echo "}" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo "func day${day}(input string, part int) int {" >> "$folder_name/main.go"
echo "    return 0" >> "$folder_name/main.go"
echo "}" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"

touch "$input_folder_name/input.txt"

echo -e "\033[90mCreated boilerplate for ${year} day ${day} in ${folder_name}"

# Done
cd "$folder_name"
echo -e "\n\033[35mGood luck \033[90mon day ${day}!"