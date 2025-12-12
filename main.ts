/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-12-10
 * @fileoverview: This file plays a guessing game with the user 
 */

let MIN = 1;
let MAX = 10;
let randomNumber = Math.floor(Math.random() * (MAX - MIN + 1)) + MIN;

console.log("Welcome to the Guessing Game!");
console.log("Guess a number between 1 and 10. Enter 0 to quit.");

let guess = -1;

while (guess !== 0 && guess !== randomNumber) {
  guess = parseInt(prompt("Enter your guess: ") || "0");
  if (guess === 0) {
    console.log("Game terminated.");
  } else if (guess === randomNumber) {
    console.log(`Congratulations! You guessed the correct number: ${randomNumber}`);
  } else {
    console.log("Try again.");
  }
  console.log("\nDone.")
}