package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
    "regexp"
    "strconv"
)

type coord struct {
    x, y *big.Int
}

type config struct {
    a, b    coord
    prize   coord
}

func newCoord(x, y int64) coord {
    return coord{
        x: big.NewInt(x),
        y: big.NewInt(y),
    }
}

func (c config) solve() *big.Int {
    zero := big.NewInt(0)
    
    if c.a.x.Cmp(zero) == 0 {
        return zero
    }

    det := new(big.Int)
    det.Mul(c.a.x, c.b.y)
    temp := new(big.Int).Mul(c.b.x, c.a.y)
    det.Sub(det, temp)
    
    if det.Cmp(zero) == 0 {
        return zero
    }

    bNum := new(big.Int)
    bNum.Mul(c.prize.y, c.a.x)
    temp.Mul(c.prize.x, c.a.y)
    bNum.Sub(bNum, temp)

    if new(big.Int).Mod(bNum, det).Cmp(zero) != 0 {
        return zero
    }

    b := new(big.Int).Div(bNum, det)
    
    aNum := new(big.Int)
    aNum.Mul(b, c.b.x)
    aNum.Sub(c.prize.x, aNum)
    
    if new(big.Int).Mod(aNum, c.a.x).Cmp(zero) != 0 {
        return zero
    }
    
    a := new(big.Int).Div(aNum, c.a.x)

    if a.Cmp(zero) < 0 || b.Cmp(zero) < 0 {
        return zero
    }

    result := new(big.Int).Mul(a, big.NewInt(3))
    result.Add(result, b)
    
    return result
}

func main() {
    f, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    re := regexp.MustCompile(`[-+]?\d+`)
    
    var cfgs []config
    var cfg config
    lineCount := 0
    
    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            continue
        }
        
        numStrs := re.FindAllString(text, -1)
        if len(numStrs) < 2 {
            continue
        }
        
        switch lineCount % 3 {
        case 0:
            x, _ := strconv.ParseInt(numStrs[0], 10, 64)
            y, _ := strconv.ParseInt(numStrs[1], 10, 64)
            cfg.a = newCoord(x, y)
        case 1:
            x, _ := strconv.ParseInt(numStrs[0], 10, 64)
            y, _ := strconv.ParseInt(numStrs[1], 10, 64)
            cfg.b = newCoord(x, y)
        case 2:
            x, _ := strconv.ParseInt(numStrs[0], 10, 64)
            y, _ := strconv.ParseInt(numStrs[1], 10, 64)
            cfg.prize = newCoord(x, y)
            
            offset := big.NewInt(10000000000000)
            cfg.prize.x.Add(cfg.prize.x, offset)
            cfg.prize.y.Add(cfg.prize.y, offset)
            
            cfgs = append(cfgs, cfg)
            cfg = config{}
        }
        lineCount++
    }

    result := big.NewInt(0)
    
    for _, cfg := range cfgs {
        solution := cfg.solve()
        result.Add(result, solution)
    }
    
    fmt.Println(result.String())
}