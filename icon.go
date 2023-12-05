package main

var ICON = []byte{
	0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x30, 0x30, 0x10, 0x00, 0x01, 0x00,
	0x04, 0x00, 0x68, 0x06, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
	0x00, 0x00, 0x30, 0x00, 0x00, 0x00, 0x60, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x28, 0x00, 0x00, 0x02,
	0x58, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x03, 0x97, 0x00, 0x00, 0x00,
	0xad, 0x00, 0x00, 0x00, 0xd0, 0x00, 0x00, 0x00, 0xfe, 0x00, 0x01, 0x05,
	0xf9, 0x00, 0x02, 0x2f, 0xfa, 0x00, 0x01, 0x41, 0xf6, 0x00, 0x01, 0x54,
	0xf2, 0x00, 0x00, 0x5e, 0xed, 0x00, 0x00, 0x66, 0xea, 0x00, 0x01, 0x6d,
	0xe8, 0x00, 0x00, 0x70, 0xe4, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x0f, 0x00, 0x00, 0x00, 0xff, 0xfe, 0xff, 0xf0, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xff, 0xef, 0xff, 0xff, 0xe0, 0x0f, 0xff, 0xff, 0xff, 0xff, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x0f, 0xfe, 0xef, 0xff, 0xfe, 0xef, 0xff, 0xfe, 0xef, 0xfe, 0xef, 0xe0,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xfe, 0xef, 0xff, 0xff, 0xef, 0xf0, 0x00, 0xff, 0xff, 0xef, 0xff, 0xef,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xff, 0xfe, 0xef, 0x00, 0x00, 0xee, 0xff, 0xfe, 0xff, 0xff,
	0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xff, 0xff, 0xef, 0xf0, 0x00, 0x00, 0x0f, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfe,
	0xee, 0xef, 0xfe, 0xef, 0x00, 0x00, 0x00, 0x00, 0xfe, 0xee, 0xfe, 0xba,
	0x87, 0x78, 0x78, 0x65, 0x41, 0x00, 0x00, 0x00, 0x02, 0x67, 0x77, 0x79,
	0xac, 0xee, 0xef, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xff, 0xc8, 0x77,
	0x77, 0x77, 0x77, 0x77, 0x74, 0x00, 0x00, 0x01, 0x68, 0x77, 0x77, 0x77,
	0x77, 0x8c, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xfa, 0x87, 0x7b,
	0xce, 0xfe, 0xe0, 0x77, 0x74, 0x00, 0x00, 0x16, 0x77, 0x7b, 0xdf, 0xed,
	0xb8, 0x77, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x79, 0xdf,
	0xff, 0xfe, 0xfe, 0x07, 0x74, 0x00, 0x00, 0x67, 0x77, 0x0f, 0xfe, 0xff,
	0xff, 0xa7, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x77, 0x8e, 0xff,
	0xef, 0xff, 0xff, 0xc7, 0x74, 0x00, 0x03, 0x87, 0x7f, 0xef, 0xff, 0xef,
	0xff, 0xe8, 0x77, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x0f, 0xef,
	0xef, 0xee, 0xff, 0xd7, 0x70, 0x00, 0x06, 0x77, 0x0e, 0xef, 0xff, 0xff,
	0xee, 0xe9, 0x77, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x00, 0xff,
	0xff, 0xff, 0xfe, 0xd7, 0x70, 0x00, 0x07, 0x78, 0xff, 0xff, 0xee, 0xff,
	0xff, 0xf0, 0x77, 0x70, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x00, 0x0f,
	0xef, 0xff, 0xfe, 0xd7, 0x7a, 0x00, 0x07, 0x8a, 0xef, 0xff, 0xff, 0xfe,
	0xf0, 0x00, 0x07, 0x70, 0x00, 0x00, 0x00, 0x00, 0x77, 0x70, 0x00, 0x00,
	0xff, 0xff, 0xff, 0xd7, 0x7c, 0xf0, 0x77, 0x7b, 0xff, 0xff, 0xef, 0xff,
	0xe0, 0x00, 0x77, 0x70, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x00, 0x00,
	0x0f, 0xa7, 0x77, 0x77, 0x7c, 0x00, 0x07, 0x7b, 0xef, 0xee, 0xef, 0xff,
	0x00, 0x00, 0x07, 0x70, 0x00, 0x00, 0x00, 0x00, 0x77, 0x70, 0x00, 0xef,
	0xef, 0xa7, 0x77, 0x77, 0x70, 0x00, 0x87, 0x7b, 0xff, 0xff, 0xff, 0xee,
	0xf0, 0x00, 0x77, 0x70, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x00, 0x0e,
	0xef, 0xff, 0xfe, 0xef, 0x00, 0x00, 0x38, 0x7a, 0xff, 0xff, 0xff, 0xff,
	0xfe, 0x00, 0x07, 0x70, 0x00, 0x00, 0x00, 0x00, 0x77, 0x77, 0xff, 0xff,
	0xff, 0xff, 0xef, 0xf0, 0x00, 0x00, 0x28, 0x80, 0xff, 0xef, 0xff, 0xef,
	0xff, 0xe0, 0x77, 0x70, 0x00, 0x00, 0x00, 0x00, 0x07, 0x77, 0x0e, 0xfe,
	0xff, 0xfe, 0xef, 0x00, 0x00, 0x00, 0x08, 0x77, 0x0e, 0xff, 0xee, 0xef,
	0xff, 0xff, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x77, 0xcf, 0xff,
	0xff, 0xef, 0xf0, 0x00, 0x00, 0x00, 0x06, 0x88, 0x80, 0xef, 0xff, 0xff,
	0xee, 0xfc, 0x77, 0x70, 0x00, 0x00, 0x00, 0x00, 0x00, 0x97, 0x8d, 0xfe,
	0xfe, 0xef, 0x00, 0x00, 0x00, 0x00, 0x03, 0x87, 0x80, 0x0f, 0xff, 0xff,
	0xff, 0xf9, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0xc7, 0x78, 0xdf,
	0xff, 0xff, 0xd0, 0x01, 0x10, 0x00, 0x00, 0x68, 0x86, 0x00, 0xff, 0xff,
	0xef, 0xa7, 0x7b, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x0e, 0xec, 0x87, 0x7a,
	0xcd, 0xe0, 0x05, 0x68, 0x40, 0x00, 0x00, 0x18, 0x87, 0x87, 0x0e, 0xed,
	0xb9, 0x78, 0xaf, 0xef, 0x00, 0x00, 0x00, 0x00, 0xef, 0xff, 0xda, 0x77,
	0x77, 0x78, 0x87, 0x78, 0x80, 0x00, 0x00, 0x01, 0x67, 0x87, 0x87, 0x77,
	0x77, 0x7b, 0xff, 0xff, 0xf0, 0x00, 0x00, 0x0f, 0xff, 0xff, 0xff, 0xdb,
	0x97, 0x78, 0x88, 0x64, 0x20, 0x00, 0x00, 0x00, 0x02, 0x58, 0x87, 0x78,
	0xab, 0xef, 0xef, 0xef, 0xff, 0x00, 0x00, 0xef, 0xff, 0xff, 0xff, 0xef,
	0xe0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xef, 0xff, 0xff, 0xff, 0xff, 0xf0, 0x0f, 0xff, 0xff, 0xfe, 0xee, 0xef,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x0f, 0xff, 0xff, 0xff, 0xee, 0xe0, 0xff, 0xef, 0xee, 0xef, 0xff, 0xf0,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0xee, 0xef, 0xef, 0xff, 0xf0, 0x0f, 0x0f, 0xff, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x0f, 0xff, 0xef, 0xef, 0x00, 0x00, 0x00, 0xe0, 0xf0, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xe0, 0xf0, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0xff, 0xf1, 0xff, 0xff, 0x57, 0xff,
	0x00, 0x00, 0xff, 0x80, 0xff, 0xfe, 0x03, 0xff, 0x00, 0x00, 0xff, 0x00,
	0x3f, 0xfc, 0x01, 0xff, 0x00, 0x00, 0xff, 0x80, 0x3f, 0xf8, 0x03, 0xff,
	0x00, 0x00, 0xff, 0xc0, 0x0f, 0xf0, 0x03, 0xff, 0x00, 0x00, 0xff, 0xe0,
	0x0f, 0xe0, 0x0f, 0xef, 0x00, 0x00, 0xc0, 0x70, 0x07, 0xc0, 0x0c, 0x01,
	0x00, 0x00, 0x80, 0x38, 0x03, 0x80, 0x18, 0x00, 0x00, 0x00, 0x00, 0x1c,
	0x01, 0x00, 0x10, 0x01, 0x00, 0x00, 0xc0, 0x0e, 0x00, 0x00, 0x70, 0x03,
	0x00, 0x00, 0xc0, 0x07, 0x00, 0x00, 0x40, 0x07, 0x00, 0x00, 0xe0, 0x03,
	0x80, 0x00, 0xc0, 0x0f, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x1f,
	0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0xfc, 0x00,
	0x40, 0x00, 0x00, 0x7f, 0x00, 0x00, 0xfe, 0x00, 0x20, 0x02, 0x00, 0x3f,
	0x00, 0x00, 0xfc, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0xf8, 0x80,
	0x04, 0x08, 0x00, 0x3f, 0x00, 0x00, 0xf8, 0xc0, 0x06, 0x00, 0x01, 0x1f,
	0x00, 0x00, 0xf8, 0xe0, 0x03, 0x80, 0x07, 0x9f, 0x00, 0x00, 0xf1, 0xf0,
	0x01, 0x00, 0x07, 0x1f, 0x00, 0x00, 0xf8, 0xf8, 0x03, 0x80, 0x0f, 0x9f,
	0x00, 0x00, 0xf1, 0xc0, 0x07, 0x00, 0x07, 0x1f, 0x00, 0x00, 0xf8, 0xe0,
	0x0e, 0x00, 0x03, 0x9f, 0x00, 0x00, 0xf0, 0x00, 0x1c, 0x10, 0x01, 0x1f,
	0x00, 0x00, 0xf8, 0x80, 0x38, 0x08, 0x00, 0xbf, 0x00, 0x00, 0xfc, 0x00,
	0x70, 0x04, 0x00, 0x1f, 0x00, 0x00, 0xfc, 0x00, 0xe0, 0x06, 0x00, 0x3f,
	0x00, 0x00, 0xf8, 0x00, 0x40, 0x03, 0x00, 0x1f, 0x00, 0x00, 0xf8, 0x01,
	0x80, 0x00, 0x80, 0x0f, 0x00, 0x00, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x07,
	0x00, 0x00, 0xe0, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0xc0, 0x04,
	0x00, 0x00, 0x70, 0x01, 0x00, 0x00, 0x80, 0x0c, 0x00, 0x80, 0x38, 0x01,
	0x00, 0x00, 0x00, 0x18, 0x01, 0xc0, 0x1c, 0x01, 0x00, 0x00, 0xa0, 0xf0,
	0x03, 0xe0, 0x0e, 0x03, 0x00, 0x00, 0xf5, 0xe0, 0x07, 0xf0, 0x07, 0x57,
	0x00, 0x00, 0xff, 0xe0, 0x0f, 0xf8, 0x03, 0xff, 0x00, 0x00, 0xff, 0xc0,
	0x1f, 0xf8, 0x01, 0xff, 0x00, 0x00, 0xff, 0x80, 0x3f, 0xfe, 0x00, 0xff,
	0x00, 0x00, 0xff, 0x00, 0x7f, 0xff, 0x01, 0xff, 0x00, 0x00, 0xff, 0xeb,
	0xff, 0xff, 0xaf, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x00, 0x00,
}
