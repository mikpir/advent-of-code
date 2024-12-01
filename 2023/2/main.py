config = {
    'red':12,
    'green':13,
    'blue':14
}

def main():
    lines = getFile()
    possible = []
    powers = []
    for line in lines:
        linedata = getlinedata(line)
        # print(linedata)
        power = getpower(linedata)
        # print(power)
        powers.append(power)
        if ispossible(linedata):
            # print(linedata)
            possible.append(linedata['game'])
    print(sum(powers))

def getlinedata(line: str):
    [gamenumber, showsstr] = line.split(':')
    gamenr = int(gamenumber.split(' ').pop())
    shows = showsstr.strip().split(';')
    parsedshows = []
    for show in shows:
        parsedshow = {}
        for color in show.split(','):
            color = color.strip()
            [nr, colorname] = color.split(' ')
            parsedshow[colorname] = int(nr)
            # print(nr, colorname)
        parsedshows.append(parsedshow)
    # print(parsedshows)
    return {
        'game':gamenr,
        'shows': parsedshows
    }

def getpower(linedata):
    mins = {
        'blue': 0,
        'red': 0,
        'green': 0
    }
    for show in linedata['shows']:
        for color in ['red', 'green', 'blue']:
            if color not in show:
                continue
            if show[color] > mins[color]:
                mins[color] = show[color]
    return mins['red'] * mins['blue'] * mins['green']

def ispossible(linedata):
    for show in linedata['shows']:
        for color in ['red', 'green', 'blue']:
            if color not in show:
                continue
            if show[color] > config[color]:
                return False
    return True

def getFile():
    with open('data', encoding='utf-8') as f:
        lines = f.readlines()
        return list(line.strip() for line in lines)

if  __name__ == '__main__':
    main()
