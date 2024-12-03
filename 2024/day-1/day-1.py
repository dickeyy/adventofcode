data = open("input.txt").read()

left = []
right = []

for line in data.split("\n"):
    left.append(int(line.split("   ")[0]))
    right.append(int(line.split("   ")[1]))

# sort the left and right
left.sort()
right.sort()

out = 0
part = int(input())

if part == 1:
    for i in range(len(left)):
        out += abs(left[i] - right[i])
else:
    freq = {}
    for num in right:
        freq[num] = freq.get(num, 0) + 1

    for num in left:
        if num in freq:
            out += num * freq[num]

print(out)