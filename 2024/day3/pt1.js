const fs = require("fs");

console.time("Execution Time");
fs.readFile("./input.txt", "utf-8", (err, data) => {
    if(err) {
        console.error(err);
        return;
    }

    const rgx = /mul\(\d{1,3}\,\d{1,3}\)/g;
    const instructions = data.match(rgx);

    let finalCount = 0;
    instructions.forEach((e) => {
        const f = e.indexOf("(")
        const l = e.indexOf(")")

        const digits = e.substring(f+1, l);
        const [fn, ln] = digits.split(",")
        finalCount += (parseInt(fn)*parseInt(ln));
    });

    console.log(finalCount);
    console.timeEnd("Execution Time");
});

//Execution Time: 9.168ms