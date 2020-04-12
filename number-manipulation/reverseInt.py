cases = {
    123: 321,
    1: 1,
    0: 0,
    987654321: 123456789,
    -1: -1,
}


def reverseInt(initial) -> int:
    final = 0
    negative = False
    if initial < 0:
        initial = -initial
        negative = True

    while initial != 0:
        r = initial % 10
        initial = initial // 10
        final *= 10
        final += r

    if negative:
        final = -final

    return final


for k in cases:
    if cases[k] != reverseInt(k):
        print("failed during", k, ", Expected", cases[k], "got", reverseInt(k))
        print()
    else:
        print("worked for", k, " Expected", cases[k], "got", reverseInt(k))
