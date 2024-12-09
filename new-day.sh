#!/bin/bash

# Load environment variables from .env file if it exists
if [ -f ".env" ]; then
    set -a
    source "./.env"
    set +a
else
    echo "Warning: .env file not found"
fi

# Check if a day parameter is provided
if [ $# -ne 2 ]; then
    echo "Usage: $0 <year> <day>"
    exit 1
fi

# Check if AOC_SESSION_COOKIE environment variable is set
if [ -z "$AOC_SESSION_COOKIE" ]; then
    echo "Error: AOC_SESSION_COOKIE environment variable is not set"
    echo "Please ensure it's set in your .env file or environment"
    exit 1
fi

# Rest of your script remains the same
year=$1
day=$2

# Create folder name with the day
folder_name="${year}/day-${day}"
input_folder_name="inputs/${year}/day-${day}"

# Create the folders
mkdir -p "$folder_name"
mkdir -p "$input_folder_name"

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
echo "    return -1" >> "$folder_name/main.go"
echo "}" >> "$folder_name/main.go"
echo "" >> "$folder_name/main.go"
echo -e "\033[90mGenerated boiler plate code"

# Fetch input data from Advent of Code and trim the trailing newline
input_file="${input_folder_name}/input.txt"
echo -e "\033[90mFetching input data for day ${day}..."
curl -s -H "Accept: application/json" --cookie "session=${AOC_SESSION_COOKIE}" "https://adventofcode.com/${year}/day/${day}/input" | perl -pe 'chomp if eof' > "$input_file"

if [ $? -eq 0 ] && [ -s "$input_file" ]; then
    echo -e "\033[90mSuccessfully downloaded input data to ${input_file}"
else
    echo -e "\033[91mFailed to download input data. Please check your session cookie and try again."
    echo -e "\033[90mYou may need to manually download the input data."
fi

# Done
cd "$folder_name"
echo -e "\n\033[35mGood luck \033[90mon day ${day}!"