package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "flag"
    "math"
)

type struct_code struct{
  code string
  num int
}

func fun_struct(in int, code_len int)(struct_code, int){
  var strc_co string
  mod := 0
  strc_num := 0
  fub0 := 0
  fub1 := 1
  for i:=0; i<code_len; i++{
    strc_co = strc_co + strconv.Itoa(in%2)
    strc_num += fub1*(in%2)
    fub0, fub1 = fub1, fub0
    fub1 = fub0 + fub1 + 1
    in /= 2
  }
  mod = fub1
  fmt.Println("infometion word:",strc_co,"  sum:",strc_num,"  set:",strc_num%mod)//check output
  out := struct_code{
    code: strc_co,
    num: strc_num%mod,
  }
  return out, mod
}


func main(){
  flag.Parse()//input n
  n,_ := strconv.Atoi(flag.Arg(0))
  N := int(math.Pow(2,float64(n)))
  code_slice :=make([]struct_code,N)
  var mod int
  for i:=0; i<N; i++{
    code_slice[i], mod = fun_struct(i,n)
  }

  check_slice :=make([]int, mod)
  maxCheck :=0
  for i:=0; i<N; i++{
    check_slice[code_slice[i].num]++
    if check_slice[code_slice[i].num]>maxCheck{
      maxCheck = check_slice[code_slice[i].num]
    }
  }
  fmt.Printf("--------------------------\nThe set(s) have %d codes:  ",maxCheck)
  for i:=0; i<mod; i++{
    if check_slice[i] == maxCheck{
      fmt.Printf("%d  ",i)
    }
  }
  fmt.Printf("\nWhat do you want as the output?\n")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  chosen_num,_ := strconv.Atoi(scanner.Text())
  if check_slice[chosen_num] != maxCheck{
    fmt.Println("This number is denied.")
    os.Exit(1)
  }
  fmt.Println("--------------------------")
  out_fl, e2 := os.Create("output.txt")
  if e2 != nil{
    fmt.Println(os.Stderr,e2)
    os.Exit(1)
  }
  defer out_fl.Close()
  writer := bufio.NewWriter(out_fl)
  for i:=0; i<N; i++{
    if code_slice[i].num == chosen_num{
      if _, e3 := fmt.Fprintln(writer,code_slice[i].code); e3 != nil{
        fmt.Println(os.Stderr,e3)
        os.Exit(1)
      }
      fmt.Println(code_slice[i].code)
    }
  }
  writer.Flush()
}
