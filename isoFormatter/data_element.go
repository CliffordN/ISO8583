package isoFormatter

type elem struct {
	typ   string
	len   int
	fixed bool
}

var dataElem = map[int]elem{
	1:   elem{"b", 64, true},
	2:   elem{"n", 19, false},
	3:   elem{"an", 6, true},
	4:   elem{"an", 12, true},
	5:   elem{"n", 12, true},
	6:   elem{"n", 12, true},
	7:   elem{"n", 10, true},
	8:   elem{"n", 8, true},
	9:   elem{"n", 8, true},
	10:  elem{"n", 8, true},
	11:  elem{"n", 6, true},
	12:  elem{"n", 6, true},
	13:  elem{"n", 4, true},
	14:  elem{"n", 4, true},
	15:  elem{"n", 4, true},
	16:  elem{"n", 4, true},
	17:  elem{"n", 4, true},
	18:  elem{"n", 4, true},
	19:  elem{"n", 3, true},
	20:  elem{"n", 3, true},
	21:  elem{"n", 3, true},
	22:  elem{"n", 3, true},
	23:  elem{"n", 3, true},
	24:  elem{"n", 3, true},
	25:  elem{"n", 2, true},
	26:  elem{"n", 2, true},
	27:  elem{"n", 1, true},
	28:  elem{"n", 8, true},
	29:  elem{"n", 8, true},
	30:  elem{"n", 8, true},
	31:  elem{"n", 8, true},
	32:  elem{"n", 11, false},
	33:  elem{"n", 11, false},
	34:  elem{"n", 28, false},
	35:  elem{"z", 37, false},
	36:  elem{"n", 104, false},
	37:  elem{"an", 12, true},
	38:  elem{"an", 6, true},
	39:  elem{"an", 2, true},
	40:  elem{"an", 3, true},
	41:  elem{"ans", 8, true},
	42:  elem{"ans", 15, true},
	43:  elem{"ans", 40, true},
	44:  elem{"an", 25, false},
	45:  elem{"an", 76, false},
	46:  elem{"an", 999, false},
	47:  elem{"an", 999, false},
	48:  elem{"an", 999, false},
	49:  elem{"a", 3, true},
	50:  elem{"an", 3, true},
	51:  elem{"a", 3, true},
	52:  elem{"b", 16, true},
	53:  elem{"n", 18, true},
	54:  elem{"an", 120, false},
	55:  elem{"ans", 999, false},
	56:  elem{"ans", 999, false},
	57:  elem{"ans", 999, false},
	58:  elem{"ans", 999, false},
	59:  elem{"ans", 999, false},
	60:  elem{"an", 7, false},
	61:  elem{"ans", 999, false},
	62:  elem{"ans", 999, false},
	63:  elem{"ans", 999, false},
	64:  elem{"b", 16, true},
	65:  elem{"b", 16, true},
	66:  elem{"n", 1, true},
	67:  elem{"n", 2, true},
	68:  elem{"n", 3, true},
	69:  elem{"n", 3, true},
	70:  elem{"n", 3, true},
	71:  elem{"n", 4, true},
	72:  elem{"ans", 999, false},
	73:  elem{"n", 6, true},
	74:  elem{"n", 10, true},
	75:  elem{"n", 10, true},
	76:  elem{"n", 10, true},
	77:  elem{"n", 10, true},
	78:  elem{"n", 10, true},
	79:  elem{"n", 10, true},
	80:  elem{"n", 10, true},
	81:  elem{"n", 10, true},
	82:  elem{"n", 12, true},
	83:  elem{"n", 12, true},
	84:  elem{"n", 12, true},
	85:  elem{"n", 12, true},
	86:  elem{"n", 15, true},
	87:  elem{"n", 15, true},
	88:  elem{"n", 15, true},
	89:  elem{"n", 15, true},
	90:  elem{"n", 42, true},
	91:  elem{"an", 1, true},
	92:  elem{"n", 2, true},
	93:  elem{"n", 5, true},
	94:  elem{"an", 7, true},
	95:  elem{"an", 42, true},
	96:  elem{"an", 8, true},
	97:  elem{"n", 16, true},
	98:  elem{"ans", 25, true},
	99:  elem{"n", 11, false},
	100: elem{"n", 11, false},
	101: elem{"ans", 17, true},
	102: elem{"ans", 28, false},
	103: elem{"ans", 28, false},
	104: elem{"ans", 100, false},
	105: elem{"ans", 999, false},
	106: elem{"ans", 999, false},
	107: elem{"ans", 999, false},
	108: elem{"ans", 999, false},
	109: elem{"ans", 999, false},
	110: elem{"ans", 999, false},
	111: elem{"ans", 999, false},
	112: elem{"ans", 999, false},
	113: elem{"n", 11, false},
	114: elem{"ans", 999, false},
	115: elem{"ans", 999, false},
	116: elem{"ans", 999, false},
	117: elem{"ans", 999, false},
	118: elem{"ans", 999, false},
	119: elem{"ans", 999, false},
	120: elem{"ans", 999, false},
	121: elem{"ans", 999, false},
	122: elem{"ans", 999, false},
	123: elem{"ans", 999, false},
	124: elem{"ans", 255, false},
	125: elem{"ans", 50, false},
	126: elem{"ans", 6, false},
	127: elem{"ans", 999, false},
	128: elem{"b", 16, true},
}