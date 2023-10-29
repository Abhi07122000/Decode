# Decode
Go Packet Decoder
This Go program demonstrates how to decode a binary packet according to a specific format and create a structured representation of the decoded data.

# Overview
The program performs the following tasks:

1.Imports necessary packages from the Go standard library: encoding/binary for binary data manipulation and fmt for printing.

2.Defines a DecodedStruct struct to represent the decoded data with fields for various data types.

3.Implements a decodePacket function that takes a byte slice (packet) and decodes it into a DecodedStruct instance. It uses constants to define the packet format and field sizes.

4.The decodePacket function checks the length of the provided packet to ensure it matches the expected size and handles potential errors.

5.A decodeString helper function is used to decode strings from the packet at specific offsets.

6.The main function sets the packet data and calls decodePacket to decode it. It then prints the decoded struct in a specified format.
