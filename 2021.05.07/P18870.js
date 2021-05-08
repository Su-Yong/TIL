const fs = require('fs');
const input = fs.readFileSync('/dev/stdin').toString().split(/\s/g);

input.shift();

let result = [];
loop:
for (let value of input) {
  value = Number(value);

  let v = 0;
  for (let i = 0; i < result.length; i++) {
    if (input[i] > value) {
      result[i] += 1;
    } else if (input[i] == value) {
      result.push(result[i]);
      continue loop;
    } else {
      v += 1;
    }
  }

  result.push(v);
}

console.log(result.join(' '));