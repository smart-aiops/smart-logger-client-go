package client

type Level int

// 一般 低危 中危 高危
const (
	General Level = iota // Index = 1
	Low                  // Index = 2
	Medium               // Index = 3
	High                 // Index = 4
)
