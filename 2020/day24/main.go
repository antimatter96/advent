package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "github.com/emirpasic/gods/stacks/arraystack"
)

type directions struct {
	i, j, k int
}

/*
  n   i
m   X  j
  l   k
*/

func (d *directions) reduce() {
	if d.k >= d.i {
		d.k -= d.i
		d.j += d.i
		d.i = 0
	} else {
		d.i -= d.k
		d.j += d.k
		d.k = 0
	}
}

func getIJK(s string) (i, j, k int) {
	fmt.Sscanf(s, "%d,%d,%d", &i, &j, &k)
	return
}
func (d *directions) String() string {
	return fmt.Sprintf("%d,%d,%d", d.i, d.j, d.k)
}

func (d *directions) update(dir string) {
	switch dir {
	case "nw":
		d.i++
	case "w":
		d.j++
	case "sw":
		d.k++
	case "se":
		d.i--
	case "e":
		d.j--
	case "ne":
		d.k--
	}

	d.reduce()

	d.expandMega(markAsWhite)
}

func main() {
	var inp []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp = append(inp, scanner.Text())

	}

	if scanner.Err() != nil {
		fmt.Errorf(scanner.Err().Error())
	}

	day22(inp)
}

var mp = make(map[string]bool)

func day22(inps []string) {

	for _, inp := range inps {
		inp = strings.TrimSpace(inp)

		ss := strings.Split(inp, "")

		dd := &directions{}

		for i := 0; i < len(ss); i++ {
			if ss[i] == "n" || ss[i] == "s" {
				dd.update(ss[i] + ss[i+1])
				i++
			} else {
				dd.update(ss[i])
			}

		}

		mp[dd.String()] = !mp[dd.String()]
	}

	// fmt.Println(len(mp), countBlack())

	part2()
}

func countBlack() int {
	tots := 0

	for _, v := range mp {
		if v {
			tots++
		}
	}
	return tots
}

func part2() {
	for i := 0; i < 100; i++ {
		fill := make(map[string]bool)
		empty := make(map[string]bool)

		copy := make(map[string]bool)
		for key, value := range mp {
			copy[key] = value
		}

		// fmt.Println(copy)
		for key, v := range copy {
			i, j, k := getIJK(key)

			isReduced(i, j, k)
			occupied := count(i, j, k)

			if v {
				if occupied == 0 || occupied > 2 {
					empty[key] = true
				}
			} else {
				if occupied == 2 {
					fill[key] = true
				}
			}
		}

		for s := range fill {
			copy[s] = true
		}
		for s := range empty {
			copy[s] = false
		}

		// mp = map[string]bool{}
		for key, value := range copy {
			mp[key] = value
		}

	}

	fmt.Println(countBlack())

	// fmt.Println(len(mp), countBlack())

}

var temp = 0

func updateTemp(d *directions) {
	d.reduce()
	s := d.String()
	// fmt.Println("checking", s, mp[s])

	mp[s] = true && mp[s]

	if mp[s] {
		temp++
	}
}

func isReduced(i, j, k int) {
	l, m, n := i, j, k
	if k >= i {
		k -= i
		j += i
		i = 0
	} else {
		i -= k
		j += k
		k = 0
	}

	if l != i || m != j || n != k {
		panic(fmt.Sprintf("%d,%d,%d != %d,%d,%d", l, m, n, i, j, k))
	}
}

func markAsWhite(d *directions) {
	d.reduce()
	s := d.String()
	mp[s] = true && mp[s]
}

func count(i, j, k int) int {
	// fmt.Println("counting around", i, j, k)
	dir := &directions{i, j, k}

	temp = 0
	// fmt.Println(temp)

	dir.expandMega(updateTemp)

	return temp
}

func (d *directions) expandMega(ff func(*directions)) {
	d.i++
	ff(d)
	d.i--

	d.j++
	ff(d)
	d.j--

	d.k++
	ff(d)
	d.k--

	d.i--
	ff(d)
	d.i++

	d.j--
	ff(d)
	d.j++

	d.k--
	ff(d)
	d.k++

	d.reduce()
}
