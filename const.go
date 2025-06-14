package thaibath

var thaiNumeral = map[string]string{
	"0": "ศูนย์",
	"1": "หนึ่ง",
	"2": "สอง",	
	"3": "สาม",
	"4": "สี่",
	"5": "ห้า",
	"6": "หก",
	"7": "เจ็ด",
	"8": "แปด",
	"9": "เก้า",
}

var thaiPowerOfTen = map[int]string{
	0: "ล้าน",
	1: "สิบ",	
	2: "ร้อย",	
	3: "พัน",	
	4: "หมื่น",
	5: "แสน",
}

const noFractionSuffix = "ถ้วน"
const thaiBath = "บาท"
const fractionSuffix = "สตางค์"
const onePos0 = "เอ็ด"
const twoPos1 = "ยี่"