const fs = require('fs');
const input = fs.readFileSync('/dev/stdin').toString().split('\n');

const nodeTypes = input[1].split('').map(Number);
const graphs = {};

let result = 0;

for (let i = 2; i < input.length; i++) {
  const [index1, index2] = input[i].split(' ').map(Number);

  if (nodeTypes[index1 - 1] && nodeTypes[index2 - 1]) {
    graphs[index1 - 1] = {
      out: [],
      in: [index2 - 1],
    };
  } else {
    const index = Object.values(graphs).findIndex((graph) => graph.out.includes(index1 - 1));

    if (index >= 0) {
      if (nodeTypes[index2 - 1]) {
        graphs[index].in.push(index2 - 1);
      } else {
        graphs[index].out.push(index2 - 1);
      }
    } else {
      
      graphs[index1 - 1] = {
        out: nodeTypes[index2 - 1] ? [] : [index2 - 1],
        in: nodeTypes[index2 - 1] ? [index2 - 1] : [],
      };
    }
  }
}

Object.values(graphs).forEach((graph) => {
  result += (graph.in.length + 1) * graph.in.length;
})

console.log(result);