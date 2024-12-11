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
    let count = 0;
    let newspots = {};

    const lines = data.split("\r\n");
    for(let x = 0; x < lines.length; x++) {
        const fields = lines[x].split("");
        board.push([]);
        for(let y = 0; y < fields.length; y++) {
            if(fields[y].match(rgx)) {
                points.push({x: x, y: y, value: fields[y]});
                board[x].push(fields[y]);
                const keystr = x + "-" + y;
                if(newspots[keystr] == null) {
                    newspots[keystr] = 1;
                    count++;
                }
            }
            else {
                board[x].push('.');
            }
        }
    }

    const boardw = board.length;
    const boardh = board[0].length;
    
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

            let iter = 1;
            while(true) {
                const rf = reflectOverPoint(first, second, iter);

                if(rf.x < boardw && rf.y < boardh && rf.x >= 0 && rf.y >= 0) {
                    const keystr = rf.x + "-" + rf.y
                    if(newspots[keystr] == null) {
                        board[rf.x][rf.y] = 'X';
                        newspots[keystr] = 1;
                        count++;
                    }
                }
                else {
                    break;
                }
                iter++;
            }
        }
    }

    // printBoard(board)
    console.log("unique location count: " + count);
    console.timeEnd("Execution Time");
});

//reflects p over r
function reflectOverPoint(p, r, iter) {
    const rx = Math.abs(p.x - r.x)*iter;
    const ry = Math.abs(p.y - r.y)*iter;

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

//Execution Time: 12.175ms
