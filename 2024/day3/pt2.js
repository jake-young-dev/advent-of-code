const fs = require("fs");

//find and return all data until next instance of str
function findNext(data, str) {
    let idx = data.indexOf(str)
    if(idx === -1) {
        return idx;
    }
    idx += str.length;

    const firstChunk = data.substring(0, idx);
    return firstChunk;
}

//instruction constants
const stop = "don't()";
const start = "do()";

console.time("Execution Time");
fs.readFile("./input.txt", "utf-8", (err, data) => {
    if(err) {
        console.error(err);
        return;
    }

    const instructions = [];

    //check for instructions before first don't()
    const l = findNext(data, stop);
    instructions.push(l);
    data = data.substring(l.length);

    //grab enabled commands until there is no data left
    while(true) {
        const x = findNext(data, start);
        if(x === -1) {
            break;
        }
        data = data.substring(x.length);
        const y = findNext(data, stop);
        if(y === -1) {
            instructions.push(data);
            break;
        }
        instructions.push(y)
        data = data.substring(y.length);
    }

    //parse mul commands with regex and calculate
    let finalCount = 0;
    const mulrgx = /mul\(\d{1,3}\,\d{1,3}\)/g;
    instructions.forEach((e) => {
        const inst = e.match(mulrgx);
        inst.forEach((ee) => {
            const f = ee.indexOf("(")
            const l = ee.indexOf(")")

            const digits = ee.substring(f+1, l);
            const [fn, ln] = digits.split(",")
            finalCount += (parseInt(fn)*parseInt(ln));
        });
    });

    console.log(finalCount);
    console.timeEnd("Execution Time");
});

//Execution Time: 8.808ms