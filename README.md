# Huffman Al{Go}

This is Go implementation of the Huffman coding algorithm. It reads a text file, compresses the text using Huffman encoding, and then decompresses it back to the original text.

## Algorithm

![Screenshot 2024-12-10 182243](https://github.com/user-attachments/assets/3a6ac800-7f66-4c89-8d76-7f3ab101664a)

## Features
- Build a Huffman tree based on character frequencies.
- Encode text into a binary string using the Huffman codes.
- Decode the binary string back to the original text.
- Save the encoded data to a file.

## How It Works
1. Reads the input text file and calculates character frequencies.
2. Constructs a Huffman tree using a priority queue (min-heap).
3. Generates unique binary codes for each character based on the tree structure.
4. Encodes the text using the generated codes.
5. Decodes the encoded text back to its original form.
6. Verifies that the decoding matches the original text.

## File Structure
- `input.txt`: Input text file to be encoded.
- `encoded.txt`: Output file containing the encoded text.

## How to Run
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/huffman-go.git
   cd huffman-go
   ```

2. Place your input file as `input.txt` in the project directory.

3. Build and run the program:
   ```bash
   go run huffman.go
   ```

4. View the encoded text in the console and in the saved `encoded.txt`

## Example 
Given an `input.txt` with the context
```text
hello world
```

The program outputs:
```text
Encoded Text: 1010101110001110... (binary string)
Decoded Text: hello world
Encoded data saved to encoded.txt
Decoding successful!
```

## Requirements
<ul> Go 1.20 or higher </ul>

## To-Do
- [ ] To implement full text file compression 
