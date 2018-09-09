package popcount


var pc [256]byte

func init() {
  for i := range pc {
    pc[i] = pc[i/2] + byte(i&1)
  }
}

func PopCount(x uint64) int {
  return int(pc[byte(x>>(0*8))] +
  pc[byte(x>>(1*8))] +
  pc[byte(x>>(2*8))] +
  pc[byte(x>>(3*8))] +
  pc[byte(x>>(4*8))] +
  pc[byte(x>>(5*8))] +
  pc[byte(x>>(6*8))] +
  pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
  res := 0
  for i := 0; i < 8; i++ {
    res += int(pc[byte(x>>(uint64(i)))])
  }
  return res;
}

func PopCountBitshift(x uint64) int {
  res := 0
  for i := 0; i < 64; i++ {
    if x & uint64(1) == uint64(1) {
      res++
    }
    x = x >> 1
  }
  return res
}

func PopCountBitclear(x uint64) int {
  res := 0
  for x > 0 {
    x = x & (x - 1)
    res++
  }
  return res
}
