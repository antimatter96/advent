import re
from copy import deepcopy
class PartRange:
        def __init__(self, ranges=None):
                self.ranges = deepcopy(ranges) or { l:{"min": 0,"max": 4001} for l in "xmas"}

        # rule
        #       what        what
        # against     int
        # greaterThan bool
        #

        def splitOn(self, lhs):
                inRange = PartRange(self.ranges)
                outRange = PartRange(self.ranges)
                against = int(lhs[2:])

                existingRange = inRange.ranges[what]
                if existingRange["min"] <= against <= existingRange["max"]:
                        if greaterThan:
                                inRange.ranges[what]["min"] = against
                                outRange.ranges[what]["max"] = against+1
                        else: # <
                                inRange.ranges[what]["max"] = against
                                outRange.ranges[what]["min"] = against-1
                elif against < existingRange["min"] and greaterThan:
                        outRange = None
                elif against > existingRange["max"] and not greaterThan:
                        outRange = None
                else:
                        inRange = outRange = None
                return inRange, outRange

        def score(self):
                ans = 1
                for r in self.ranges.values():
                        ans*=r["max"]-r["min"]-1
                return ans

def main(input):
        workflows_raw, parts = input.split("\n\n")

        workflows = {}
        for l in workflows_raw.splitlines():
                name, rules = l[:-1].split("{")
                workflows[name] = rules.split(",")

        validRanges = []
        def testRange(partRange, name):
                if partRange is None: return
                if name in "AR":
                        if name == "A":
                                validRanges.append(partRange)
                else:
                        rules = workflows[name]
                        for r in rules:
                                if ":" not in r:
                                        testRange(partRange, r)
                                else:
                                        lhs, next = r.split(":")
                                        inR, outR = partRange.splitOn(lhs)
                                        testRange(inR, next)
                                        partRange = outR
                                        if partRange == None: return

        testRange(PartRange(), "in")

        return 0, sum(pr.score() for pr in validRanges)

print()
