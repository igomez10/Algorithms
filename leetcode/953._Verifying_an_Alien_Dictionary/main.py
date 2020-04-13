class Solution:

    # isBigger returns true if w1 is bigger than w2
    def isSmaller(self, w1, w2, order):
        if w1 == w2:
            return True

        longer = w1
        shorter = w2
        if len(w2) > len(w1):
            longer = w2
            shorter = w1

        for i in range(len(shorter)):
            charInLonger = longer[i]
            charInShorter = shorter[i]

            if charInLonger != charInShorter:
                # words are not the same and the current character will tell us which one is bigger
                rankCharLonger = order.find(charInLonger)
                rankCharShorter = order.find(charInShorter)

                if rankCharLonger < rankCharShorter:
                    # longer one should go first
                    if longer == w1:
                        return True
                    return False
                else:
                    # longer one should go second
                    if longer == w1:
                        return False
                    return True

        if longer == w1:
            return False
        else:
            return True

        # if w1 == w2:
        #     return True

        # for i in range(min(len(w1), len(w2))):
        #     if order.find(w1[i]) > order.find(w2[i]):
        #         print("found misplaced character", w1[i], w2[i])
        #         return False

        # print("using legth to determine position")
        # if len(w1) > len(w2):
        #     return False
        # else:
        #     return True

    def isAlienSorted(self, words, order) -> bool:
        self.order = order
        for i in range(1, len(words)):
            if not self.isSmaller(words[i-1], words[i], order):
                return False
        return True


s = Solution()


cases = {
    "hlabcdefgijkmnopqrstuvwxyz": (["hello", "leetcode"], True),
    "worldabcefghijkmnpqstuvxyz": (["word", "world"], False),
    "abcdefghijklmnopqrstuvwxyz": (["apple", "app"], False),
}


for k in cases.keys():
    w1 = cases[k][0][0]
    w2 = cases[k][0][1]
    expected = cases[k][1]

    if s.isAlienSorted(cases[k][0], k) != expected:
        print("failed with", w1, w2, expected)
    else:
        print("worked with", w1, w2, expected)
    print()
    print()

print("------END---------")
