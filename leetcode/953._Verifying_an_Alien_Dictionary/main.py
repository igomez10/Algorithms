class Solution:

    # isBigger returns true if w1 is bigger than w2
    def isSmaller(self, w1, w2, order):

        for i in range(min(len(w1), len(w2))):
            if order.find(w1[i]) > order.find(w2[i]):
                return False
            elif order.find(w1[i]) < order.find(w2[i]):
                return True

        if len(w1) > len(w2):
            return False
        else:
            return True

    def isAlienSorted(self, words, order) -> bool:
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
