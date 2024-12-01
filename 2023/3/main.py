import re

test = [
'.3......34',
'.....!....',
'..35..633.',
'......#...',
'617*......',
'.....+.58.',
'..592.....',
'......755.',
'...$.*....',
'.664.598..'
]

def main():
    # lines = getFile()
    lines = test
    nrs = []
    getlinelocs(lines)
    for i, line in enumerate(lines):
        # print(line)
        linenrs = getlinenumbers(line)

def issymboladjacent(lines, linenr: int, start: int, end: int):
    startline = max(linenr - 1, 0)
    endline = min(linenr + 1, len(lines) - 1)
    starti = max(start - 1, 0)
    endi = min(end, len(lines[linenr]) - 1)
    # print(startline, endline)
    # print(starti, endi)
    for lineindex in range(startline, endline + 1):
        for charindex in range(starti, endi + 1):
            char = lines[lineindex][charindex]
            if not char.isalnum() and char != '.':
                # print(lineindex, charindex, char)
                # print('dingdingding')
                return True


def getlinelocs(lines):
        locs = {}
    for linenr, line in enumerate(lines):
        print(linenr, line)
        matches = re.finditer('\\d+', line)
        matchspans = (match.span() for match in matches)
        for begin, end in matchspans:
            print(line[begin], line[end])
        return


def getlinenumbers(line: str):
    # print(line)
    matches = re.finditer('\\d+', line)
    return [(match.group(), match.span()) for match in matches]


def getFile():
    with open('data', encoding='utf-8') as f:
        lines = f.readlines()
        return list(line.strip() for line in lines)

if __name__ == '__main__':
    main()
