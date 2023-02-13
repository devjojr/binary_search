var binarySearch = function (arr, data) {
  var first = 0;
  var last = arr.length - 1;
  var found = false;
  while (first <= last && !found) {
    var mid = Math.floor((first + last) / 2);
    if (arr[mid] === data) {
      found = true;
    } else {
      if (data < arr[mid]) {
        last = mid - 1;
      } else {
        first = mid + 1;
      }
    }
  }
  return found;
};
var arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13];
console.log(binarySearch(arr, 11));
console.log(binarySearch(arr, 25));
console.log(binarySearch(arr, 2));
