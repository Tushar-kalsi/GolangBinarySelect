# GoSelect

GoSelect is a Go application that demonstrates the use of assembly language to perform bit manipulation operations in Go. It includes examples of selecting a specific bit from a binary sequence and counting trailing zeros in a 64-bit integer.

## Table of Contents

- [Usage](#usage)
- [How It Works](#how-it-works)
- [License](#license)
- [Author](#author)

## Usage

1. Ensure you have the Go programming language installed on your system.

2. Build and run the GoSelect application:

   ```bash
   go run main.go


This will execute the select_u64 function, which selects a specific bit from a binary sequence and prints the result.

1. The TrailingZeros64 function counts and prints the number of trailing zeros in a 64-bit integer. You can call this function from your own code as needed.


##  How It Works
 select_u64
The select_u64 function demonstrates the use of the "pdep" operation using assembly language to select a specific bit from a binary sequence. It takes a 64-bit integer and an index as input and returns the selected bit.


 `func select_u64(x uint64, i uint64) int`


 ## TrailingZeros64`
The TrailingZeros64 function calculates the number of trailing zeros in a 64-bit integer using an assembly implementation.

`func TrailingZeros64(x uint64) int`


## License
This project is open-source and available under the MIT License. 


## Author
Created for Mat Matheo
Feel free to modify this README file to provide more context or additional information about your project. Additionally, consider adding a license file (e.g., LICENSE) to specify the terms under which your code can be used.


If you have any questions or need further assistance, please don't hesitate to reach o
