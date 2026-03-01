//
// This is only a SKELETON file for the 'Line Up' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const ending = (number) => {
  if (number % 10 === 1 && number % 100 !== 11) {
    return "st";
  } else if (number % 10 === 2 && number % 100 !== 12) {
    return "nd";
  } else if (number % 10 === 3 && number %100 !== 13) {
    return "rd";
  } else {
    return "th";
  }
}

export const format = (name, number) => {
  let end = ending(number);
  return `${name}, you are the ${number}${end} customer we serve today. Thank you!`
};
