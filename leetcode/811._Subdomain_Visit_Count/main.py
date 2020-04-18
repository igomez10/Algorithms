import collections


class Solution:
    def subdomainVisits(self, cpdomains):
        ans = {}

        for entry in cpdomains:
            count, domain = entry.split(" ")
            count = int(count)

            fragments = domain.split(".")

            for i in range(len(fragments)):
                current = ".".join(fragments[i:])
                if current in ans:
                    ans[current] += count

                else:
                    ans[current] = count

        return ans

    # iterate xi over array

    # split xi into pieces, "."

    # parse number

    # add current parsed values to all the components in a map


cases = [
    [["9001 discuss.leetcode.com"], [
        "9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]],
    [["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"], ["901 mail.com",
                                                                                 "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"]]
]


s = Solution()

for i in range(len(cases)):
    ans = s.subdomainVisits(cases[i][0])
    if sorted(ans) == sorted(cases[i]):
        print("worked for", cases[i])
    else:
        print("failed with", cases[i])
