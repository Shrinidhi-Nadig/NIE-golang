//FUNCTIONS IN JS ARE ALSO CALLED METHODS

// function sum(a,b){
//     return a+b;
// }
// console.log("The sum of the values are: ");
// console.log(sum(12,13));
// console.log(sum(14,15));
// console.log(sum(23,24));

// const add = (a,b) =>a+b;
// console.log("Sum using the fat arrow function: ")
// console.log(add(12,24));
// console.log(add(25,36));
// console.log(add(47,55));


// const cube = (a) =>a*a*a;
// console.log(cube(2))
// console.log(cube(5))
// console.log(cube(8))

// const greet = (name) =>("Hello, "+name);
// console.log(greet("Shrinidhi"));
// console.log(greet("Sandesh"));
// console.log(greet("Sakaleshwar"));

//FILTER METHODS

const numbers=[1,2,3,4,5,6];
const multipleof2=numbers.filter(num=>num%2==0);
console.log("Multiple of 2 are:"+ multipleof2);

//REDUCE METHODS

const somenumbers=[1,3,4,7,9,13,15];
const sumofnumbers=somenumbers.reduce((temp,present)=>temp+present,0);
console.log("Sum of given numbers are: "+ sumofnumbers)