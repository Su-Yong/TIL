const fs = require('fs');
const input = /*fs.readFileSync('/dev/stdin')*/`5
10111
1 2
2 3
2 4
4 5`.toString().split('\n');

const nodeTypes = input[1].split('').map(Number);
const nodes = [];

for (let i = 2; i < input.length; i++) {
  const [index1, index2] = input[i].split(' ').map(Number);

  if (!nodes[index1 - 1]) nodes[index1 - 1] = [];
  if (!nodes[index2 - 1]) nodes[index2 - 1] = [];
  nodes[index1 - 1].push(index2 - 1);
  nodes[index2 - 1].push(index1 - 1);
}

let result = 0;
let visited = 0b0;
for (let i = 0; i < nodeTypes.length; i++) {
  if (nodeTypes[i]) {
    visited = 0b0;

    mDFS(nodes, visited, i);
  }
}

function mDFS(array, visited, index) {
  visited = visited | (1 << index);
  console.log(visited.toString(2));

  for (let i = 0; i < array[index].length; i++) {
    const value = array[index][i];

    if (!(visited & (1 << value))) {
      if (!nodeTypes[value]) {
        mDFS(array, visited, value);
      } else {
        result += 1;
      }
    }
  }
}

console.log(result);
