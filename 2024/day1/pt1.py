import timeit
 
def run():
    f = open("input.txt", "r")
    data = f.read()
    lines = data.split("\n")
    delim = "   "
    left = []
    right = []
    for x in lines:
        br = x.split(delim)
        left.append(br[0])
        right.append(br[1])
 
    #fixing order
    left.sort()
    right.sort()
 
    loop = len(left)
    distance = 0
    for x in range(0, loop):
        l = int(left[x])
        r = int(right[x])
        distance += abs(l-r)
 
    print("distance: " + str(distance))
 
if __name__ == "__main__":
    timing = timeit.timeit(run, number=1)
    print("Execution time: " + str(timing) + " seconds")

#Execution time: 0.0012259000213816762 seconds