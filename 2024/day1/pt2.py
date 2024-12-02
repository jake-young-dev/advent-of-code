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
 
    similarity = 0
    for x in left:
        count = 0
        for y in right:
            if x == y:
                count += 1
 
        similarity += int(x) * count
 
    print("Similarity score: " + str(similarity))
 
 
if __name__ == "__main__":
    timing = timeit.timeit(run, number=1)
    print("Execution time: " + str(timing) + " seconds")

#Execution time: 0.02903360000345856 seconds