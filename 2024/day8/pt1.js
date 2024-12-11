const fs = require("fs");

//start timer
console.time("Execution Time");

//read file and parse data
fs.readFile("./input.txt", "utf-8", (err, data) => {
    if(err) {
        console.error(err);
        return;
    }

    //existing points
    const points = [];
    //2d array
    let board = [];
    const rgx = /[a-zA-Z0-9]/g;

    const lines = data.split("\r\n");
    for(let x = 0; x < lines.length; x++) {
        const fields = lines[x].split("");
        board.push([]);
        for(let y = 0; y < fields.length; y++) {
            if(fields[y].match(rgx)) {
                points.push({x: x, y: y, value: fields[y]});
                board[x].push(fields[y]);
            }
            else {
                board[x].push('.');
            }
        }
    }

    const boardw = board.length;
    const boardh = board[0].length;
    
    //reflect all points over each other
    let count = 0;
    let newspots = {};
    for(let a = 0; a < points.length; a++) {
        for(let b = 0; b < points.length; b++) {
            const first = points[a];
            const second = points[b];
            if(first == second) {
                continue;
            }
            if(first.value != second.value) {
                continue;
            }

            const rf = reflectOverPoint(first, second);

            if(rf.x < boardw && rf.y < boardh && rf.x >= 0 && rf.y >= 0) {
                const keystr = rf.x + "-" + rf.y
                if(newspots[keystr] == null) {
                    board[rf.x][rf.y] = 'X';
                    newspots[keystr] = 1;
                    count++;
                }
            }
        }
    }

    // printBoard(board)
    console.log("unique location count: " + count);
    console.timeEnd("Execution Time");
});

//reflects p over r
function reflectOverPoint(p, r) {
    const rx = Math.abs(p.x - r.x);
    const ry = Math.abs(p.y - r.y);

    if(p.x > r.x) {
        if(p.y > r.y) {
            return {x: r.x-rx, y:r.y-ry};
        }
        else {
            return {x: r.x-rx, y: r.y+ry};
        }
    }
    else {
        if(p.y > r.y) {
            return {x: r.x+rx, y:r.y-ry};
        }
        else {
            return {x: r.x+rx, y: r.y+ry};
        }
    }
}

function printBoard(board) {
    for(let x = 0; x < board.length; x++) {
        const bl = board[x];
        for(let y = 0; y < board[x].length; y++) {
            process.stdout.write(bl[y].toString());
        }
        console.log();
    }
}

//Execution Time: 10.903ms
