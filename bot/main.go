package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
var (
	cp              []cord
	cpi             []int
	cpidist         []int
	checkpointCount int
	laps            int
	globalchn       int
)

const (
	MAXT = 200
	MEDT = 100
	MINT = 0
)

type cord struct {
	x, y int
}

type controller struct {
	p      *pod
	xd, yd int
	boost  bool
	shield bool
	thrust int
}

type pod struct {
	x, y                    int
	vx, vy                  int
	angle, nextCheckPointId int

	pv    *pod
	boost bool
	laps  int
	rank  int
}

type idon struct {
	chn, xx, yy, tthrust int
	fail, lock           bool
}

func (p *pod) debuginfo(t int) {
	cx, cy := p.nextchp()

	fmt.Fprintln(os.Stderr, "x y", p.x, p.y)
	fmt.Fprintln(os.Stderr, "vx vy", p.x, p.y)
	fmt.Fprintln(os.Stderr, "nextCheckPointId", p.nextCheckPointId)
	fmt.Fprintln(os.Stderr, "boost", p.boost)
	fmt.Fprintln(os.Stderr, "laps", p.laps)
	fmt.Fprintln(os.Stderr, "rank", p.rank)
	fmt.Fprintln(os.Stderr, "speed", p.speed())
	fmt.Fprintln(os.Stderr, "actualacc", p.actualacc())
	fmt.Fprintln(os.Stderr, "maxacce", p.maxacce(t))
	fmt.Fprintln(os.Stderr, "diffspeedmax", p.difspeed(t))
	fmt.Fprintln(os.Stderr, "distchp", p.distchp())
	fmt.Fprintln(os.Stderr, "angle", p.angle)

	fmt.Fprintln(os.Stderr, "checkpointangle", p.checkpointang(cx, cy))
	fmt.Fprintln(os.Stderr, "vectorangle", p.vectorang())
	fmt.Fprintln(os.Stderr, "frontchpdiff", p.checkpointang(cx, cy))

	fmt.Fprintln(os.Stderr, "diffvectorangle", p.difangleppv())

}

func (p *pod) init(x, y, vx, vy, angle, nextCheckPointId int) {
	prev := *p
	prev.pv = nil

	p.x = x
	p.y = y
	p.vx = vx
	p.vy = vy
	p.angle = angle
	p.nextCheckPointId = nextCheckPointId
	p.pv = &prev

	p.lapsup()

}
func (p *pod) lapsup() {
	pv := p.pv
	if p.nextCheckPointId == 0 {
		if p.nextCheckPointId != pv.nextCheckPointId {
			p.laps += 1
		}
	}

}
func (p *pod) isturn(cx, cy int) bool {
	if Abs(p.frontchpdiff(cx, cy)) > 15 {
		return true
	}
	return false

}

func (p *pod) nextchp() (int, int) {
	x := cp[p.nextCheckPointId].x
	y := cp[p.nextCheckPointId].y
	return x, y
}

func (p *pod) maxacce(t int) int {
	pv := p.pv
	m := pv.speed()
	mx := Nspeed(m, t)
	return mx
}

func (p *pod) difspeed(t int) int {
	s := p.speed()
	m := p.maxacce(t)
	df := m - s
	return df
}

func (p *pod) actualacc() int {
	d := p.speed() - p.pv.speed()
	return d
}

func (p *pod) ranksum() int {
	cx, cy := p.nextchp()

	sum := (p.laps * 1000000) + ((p.nextCheckPointId + 1) * 100000) + (20000 - p.dist(cx, cy))
	return sum

}

func (p *pod) speed() int {
	s := dist(p.vx, p.vy)
	return s
}

func (p *pod) vectorang() int {
	ca := Calcangle(p.vx, -1*p.vy)
	return ca

}

func (p *pod) vectorangf() float64 {
	ca := Calcanglerad(p.vx, -1*p.vy)
	return ca

}

func (p *pod) vectfrontang() int {
	a := p.vectorang()
	b := p.angle
	ca := difangli(a, b)
	return ca
}
func (p *pod) checkpointang(cx, cy int) int {
	ca := Calcangle(cx-p.x, -1*(cy-p.y))
	return ca

}
func (p *pod) checkpointangf(cx, cy int) float64 {
	ca := Calcanglerad(cx-p.x, -1*(cy-p.y))
	return ca

}
func (p *pod) checkpointdist(cx, cy int) (int, int) {
	x, y := cx-p.x, cy-p.y
	return x, y

}
func (p *pod) dist(cx, cy int) int {
	d := dist(cx-p.x, cy-p.y)
	return d

}
func (p *pod) distchp() int {
	x, y := p.nextchp()
	d := dist(x-p.x, y-p.y)
	return d

}

func (p *pod) vectorchpdiff(cx, cy int) int {
	a := p.vectorang()
	b := p.checkpointang(cx, cy)
	ca := difangli(a, b)
	return ca
}

func (p *pod) difangleppv() int {
	pv := p.pv
	a := pv.vectorangf()
	b := p.vectorangf()
	c := difanglef(a, b)

	return c
}
func (p *pod) difangleppvf() float64 {
	pv := p.pv
	a := pv.vectorangf()
	b := p.vectorangf()
	c := difanglerad(a, b)

	return c
}
func (p *pod) frontchpdifff(cdx, cdy int) float64 {

	a := angletorad(p.angle)
	b := p.checkpointangf(cdx, cdy)
	ca := difanglerad(a, b)
	return ca
}
func (p *pod) frontchpdiff(cdx, cdy int) int {
	a := radtoangle(p.frontchpdifff(cdx, cdy))
	return a

}

func (p *pod) inifinte(cdx, cdy int, length float64) (int, int) {
	a := p.checkpointangf(cdx, cdy)

	ax, ay := radtocord(a, length)
	return p.x + ax, p.y + ay

}

func (p *pod) modinifinte(i int) bool {
	a := cpi[pastring(i)]
	b := cpi[i]
	c := cpi[nextring(i)]

	res := difangli(a, b)
	res1 := difangli(b, c)

	prod := res * res1
	if prod < 0 {

		return true

	}
	return false

}

func (p pod) collide(op pod, finalx, finaly, thrust, force int) bool {

	if Abs(p.frontchpdiff(finalx, finaly)) > 16 {
		p = predictcurve(p, finalx, finaly, thrust)

	} else {

		p = predictline(p, finalx, finaly, thrust)

	}
	oponentx, oponenty := 0, 0
	if op.preturn(100) {
		oponentx, oponenty = p.preturnchp()
	}

	if Abs(op.frontchpdiff(oponentx, oponenty)) > 16 {
		op = predictcurve(op, oponentx, oponenty, 100)

	} else {

		op = predictline(op, oponentx, oponenty, 100)

	}

	if podcollide(p, op, force) {
		return true
	}
	return false
}

func (p pod) avoidcollide(op pod, finalx, finaly, final2x, final2y, thrust1, thrust2, force int) bool {
	for i := 0; i < 4; i++ {
		if Abs(p.frontchpdiff(finalx, finaly)) > 16 {
			p = predictcurve(p, finalx, finaly, thrust1)

		} else {

			p = predictline(p, finalx, finaly, thrust1)

		}
		if Abs(op.frontchpdiff(final2x, final2y)) > 16 {
			op = predictcurve(op, final2x, final2y, thrust2)

		} else {

			op = predictline(op, final2x, final2y, thrust2)

		}
		if podcollide(p, op, force) {
			return true
		}
	}

	return false
}

func (p pod) dribble(op pod, finalx, finaly int) bool {

	final2x, final2y := (finalx-p.x)/2, (finaly-p.y)/2
	thrust1 := 100
	thrust2 := 100

	for i := 0; i < 3; i++ {
		if Abs(p.frontchpdiff(finalx, finaly)) > 16 {
			p = predictcurve(p, finalx, finaly, thrust1)

		} else {

			p = predictline(p, finalx, finaly, thrust1)

		}
		if Abs(op.frontchpdiff(final2x, final2y)) > 16 {
			op = predictcurve(op, final2x, final2y, thrust2)

		} else {

			op = predictline(op, final2x, final2y, thrust2)

		}
		if podcollide(p, op, 760) {
			return true
		}
	}
	return false

}

func (p pod) calchunt(op pod) (int, int) {
	distp := 1
	distop := 0
	cumul := 0
	ax, ay := op.nextchp()
	dx, dy := angletocord(op.checkpointang(ax, ay), 50)
	px, py := op.x, op.y
	fmt.Fprintln(os.Stderr, "dsto", distp < distop)

	for distp > distop {
		px += dx
		py += dy
		distp = p.dist(px, py)
		distop = op.dist(px, py) + cumul
		if op.dist(px, py) > op.distchp() {
			cumul += op.dist(px, py)

			op.nextCheckPointId = nextring(op.nextCheckPointId)

			op.x = px
			op.y = py
			ax, ay := op.nextchp()

			dx, dy = angletocord(op.checkpointang(ax, ay), 50)
		}
		fmt.Fprintln(os.Stderr, distp, distop, px, py)

	}

	rx, ry := p.inifinte(px, py, 20000)

	return rx, ry
}

func (p *pod) preturnchp() (int, int) {
	i := nextring(p.nextCheckPointId)
	cx, cy := cp[i].x, cp[i].y
	return cx, cy
}

func (p pod) preturn(thrust1 int) bool {
	ax, ay := p.preturnchp()

	for i := 0; i < 10; i++ {

		p = predictcurve(p, ax, ay, thrust1)

		if p.distchp() < 600 {
			return true
		}

	}
	return false

}

func (p *pod) altredchp() (int, int) {
	a := nextring(p.nextCheckPointId)
	ax, bx := angletocord(a, 400)
	return ax, bx
}

func main() {

	var p1, p2, op1, op2 pod

	p1.laps = 1
	p2.laps = 1
	op1.laps = 1
	op2.laps = 1

	var bsh, bsh2 string

	thrust1 := 100
	thrust2 := 100
	coll := false

	fmt.Scan(&laps)

	fmt.Scan(&checkpointCount)

	for i := 0; i < checkpointCount; i++ {
		var checkpointX, checkpointY int
		fmt.Scan(&checkpointX, &checkpointY)
		cp = append(cp, cord{checkpointX, checkpointY})

	}
	cpinit()
	cpdistinit()
	fmt.Fprintln(os.Stderr, " cpidist : ", cpidist)

	finalx2, finaly2 := 0, 0

	for {

		var x, y, vx, vy, angle, nextCheckPointId int
		fmt.Scan(&x, &y, &vx, &vy, &angle, &nextCheckPointId)

		p1.init(x, y, vx, vy, angletransf(angle), nextCheckPointId)

		fmt.Scan(&x, &y, &vx, &vy, &angle, &nextCheckPointId)

		p2.init(x, y, vx, vy, angletransf(angle), nextCheckPointId)

		var xo, yo, vxo, vyo, angleo, nextCheckPointIdo int
		fmt.Scan(&xo, &yo, &vxo, &vyo, &angleo, &nextCheckPointIdo)

		op1.init(xo, yo, vxo, vyo, angletransf(angleo), nextCheckPointIdo)

		fmt.Scan(&xo, &yo, &vxo, &vyo, &angleo, &nextCheckPointIdo)

		op2.init(xo, yo, vxo, vyo, angletransf(angleo), nextCheckPointIdo)

		s := []int{p1.ranksum(), p2.ranksum(), op1.ranksum(), op2.ranksum()}
		sort.Ints(s)

		p1.rank = rankcalc(s, p1.ranksum())
		p2.rank = rankcalc(s, p2.ranksum())
		op1.rank = rankcalc(s, op1.ranksum())
		op2.rank = rankcalc(s, op2.ranksum())

		fmt.Fprintln(os.Stderr, " p1 : ", p1)
		fmt.Fprintln(os.Stderr, " p2 : ", p2.speed(), p2.maxacce(thrust2), p2.difspeed(100))
		op1x, op1y := op1.nextchp()

		fmt.Fprintln(os.Stderr, " op1 : ", op1.speed(), op1.distchp(), op2.checkpointang(op1x, op1y))
		op2x, op2y := op2.nextchp()

		fmt.Fprintln(os.Stderr, " op2 : ", op2.speed(), op2.distchp(), op2.checkpointang(op2x, op2y))

		bsh = ""
		bsh2 = ""

		thrust1 = 100
		thrust2 = 100

		think := p1.modinifinte(nextring(p1.nextCheckPointId))

		finalx, finaly := p1.nextchp()

		finalx, finaly = p1.inifinte(finalx, finaly, 10)

		anglelim := 130

		if cpidist[p1.nextCheckPointId] < 5500 {
			anglelim = 90

			if think {
				anglelim = 30
			}

		}

		fmt.Fprintln(os.Stderr, " angleli : ", anglelim)

		if Abs(p1.frontchpdiff(finalx, finaly)) > anglelim && p1.speed() > 200 {

			thrust1 = 0
		}

		if p1.preturn(thrust1) {
			finalx, finaly = p1.preturnchp()
			finalx, finaly = p1.inifinte(finalx, finaly, 10)
		}

		tar := pod{}
		enemy := pod{}
		if op1.rank < op2.rank {
			tar = op1

			enemy = op2
		} else {
			enemy = op1

			tar = op2
		}

		finalx2, finaly2 = p2.calchunt(tar)
		if finalx2+finaly2 == 0 {
			finalx2, finaly2 = predict(tar, p2, 20)
		}

		if coll {

		}

		thrust2 = 100

		if p2.avoidcollide(p1, finalx2, finaly2, finalx, finaly, thrust2, thrust1, 800) {
			if p2.speed() > 250 {
				thrust2 = 0

				fmt.Fprintln(os.Stderr, " avoiid : ")
				angr := radtoangle(p1.vectorangf() + angletorad(180))
				finalx2, finaly2 = angletocord(angr, 20000)
			} else {

				thrust2 = 100

				fmt.Fprintln(os.Stderr, " avoiid2 : ")
				angr := radtoangle(p1.vectorangf())
				finalx2, finaly2 = angletocord(angr, 20000)
			}
		}

		if p1.dribble(enemy, finalx, finaly) && p1.rank == 1 {

		}

		if p2.collide(op1, finalx2, finaly2, thrust2, 700) {
			coll = true
			bsh2 = "SHIELD"

		}
		if p2.collide(op2, finalx2, finaly2, thrust2, 700) {

			coll = true
			bsh2 = "SHIELD"

		}

		if p1.collide(op1, finalx, finaly, thrust1, 680) {
			fmt.Fprintln(os.Stderr, " shie : ", Abs(difangli(p1.angle, op1.angle)))

			ax, ay := p1.nextchp()
			di := p1.dist(ax, ay)
			odi := op1.dist(ax, ay)
			if Abs(difangli(p1.angle, op1.angle)) > 60 || di > odi {

				coll = true
				bsh = "SHIELD"
			}
		}
		if p1.collide(op2, finalx, finaly, thrust1, 680) {

			fmt.Fprintln(os.Stderr, " shie : ", Abs(difangli(p1.angle, op2.angle)))

			ax, ay := p1.nextchp()
			di := p1.dist(ax, ay)
			odi := op2.dist(ax, ay)
			if Abs(difangli(p1.angle, op2.angle)) > 60 && di > odi {

				coll = true
				bsh = "SHIELD"
			}
		}

		if p1.distchp() > 6000 && !p1.boost && Abs(p1.frontchpdiff(finalx, finaly)) < 5 {
			bsh = "BOOST"
			p1.boost = true
		}

		if bsh == "" {
			fmt.Printf("%d %d %d\n", finalx, finaly, thrust1)
		} else {
			fmt.Printf("%d %d %s\n", finalx, finaly, bsh)
		}
		if bsh2 == "" {
			fmt.Printf("%d %d %d\n", finalx2, finaly2, thrust2)
		} else {

			fmt.Printf("%d %d %s\n", finalx2, finaly2, bsh2)

		}

	}

}

func cpdistinit() {

	for i := 0; i < checkpointCount; i++ {
		ax := cp[pastring(i)].x - cp[i].x
		ay := cp[pastring(i)].y - cp[i].y

		cpidist = append(cpidist, dist(ax, ay))

	}
}

func podcollide(p1, p2 pod, w int) bool {
	ax := p1.x - p2.x
	ay := p1.y - p2.y

	r := dist(ax, ay)
	elas := w
	fmt.Fprintln(os.Stderr, " elsas : ", r)

	if r < elas {
		return true
	}
	return false
}
func cpinit() {

	for i := 0; i < checkpointCount; i++ {
		cpi = append(cpi, vectlineang(cp[pastring(i)], cp[i]))

	}

}
func nextmove(p pod) pod {
	return pod{}
}
func rankcalc(s []int, vp int) int {
	for i, v := range s {
		if v == vp {
			return len(s) - i
		}
	}
	return 0
}

func nextring(i int) int {
	i++
	if i == checkpointCount {
		i = 0
	}
	return i
}

func pastring(i int) int {
	i--
	if i == -1 {
		i = checkpointCount - 1
	}
	return i
}

func angletransform(a int) int {
	if a < 180 {
		a = a * (-1)
	} else {
		a = 360 - a
	}
	return a
}

func iscurve(a int) bool {
	if a > 17 {
		return true
	}
	return false
}

func Heuristicl(diffvel int) int {
	offset := 0
	switch {
	case diffvel > 60:
		offset = 20
		break
	case diffvel > 40 && diffvel <= 60:
		offset = 15
		break

	case diffvel > 20 && diffvel <= 40:
		offset = 10
		break
	case diffvel <= 20 && diffvel > 10:
		offset = 5
		break

	case diffvel <= 10 && diffvel > 5:
		offset = 2

	case diffvel <= 5:
		offset = 1

		break
	}
	return offset

}

func Heuristicc(diffvel int) int {
	offset := 0
	switch {
	case diffvel <= 8:
		offset = 8
		break

	case diffvel > 8 && diffvel <= 58:
		offset = 10
		break

	case diffvel > 58 && diffvel <= 66:
		offset = 4

	case diffvel > 66:
		offset = 2

		break
	}
	return offset
}

func predictcurve(p pod, cx, cy, t int) pod {

	r2 := angletorad(2)
	r18 := angletorad(18)

	vectang := p.vectorangf()
	difc := p.frontchpdiff(cx, cy)

	if cx+cy == 0 {
		difc = difangli(p.pv.angle, p.angle)
	}
	pang := angletorad(p.angle)

	dfs := p.difspeed(t)

	difa := p.difangleppvf()

	Heuristicv := 0
	if t != 0 {
		Heuristicv = Heuristicc(dfs)
	}

	var tan float64
	var pa float64

	if difc > 0 {
		difa += r2
		pa = pang + r18
	} else {
		difa -= r2
		pa = pang - r18

	}

	tan = vectang + difa

	xcc := math.Cos(tan)
	ycc := math.Sin(tan)
	ycc = -1 * ycc
	disf := float64(Nspeed(p.speed()-(dfs+Heuristicv), t))
	sdx, sdy := square(xcc, ycc, disf)

	p.pv = nil

	np := pod{p.x + sdx, p.y + sdy, sdx, sdy, radtoangle(pa), p.nextCheckPointId, &p, p.boost, 0, 0}

	return np

}

func predictline(p pod, cx, cy, t int) pod {

	r2 := angletorad(2)
	r1 := angletorad(1)

	vectang := p.vectorangf()

	cpangle := p.checkpointang(cx, cy)
	difc := p.frontchpdiff(cx, cy)
	offang := p.frontchpdifff(cx, cy)

	if cx+cy == 0 {
		cpangle = p.angle + (difangli(p.pv.angle, p.angle))
		difc = difangli(p.pv.angle, p.angle)

		offang = angletorad(difangli(p.pv.angle, p.angle))

	}

	dfs := p.difspeed(t)
	difa := p.difangleppvf()
	difi := p.difangleppv()

	Heuristicv := 0
	if t != 0 {
		Heuristicv = Heuristicl(dfs)
	}

	var tan float64

	if Abs(difi) > 7 {
		if difc < 0 {
			difa += r2
		} else {
			difa -= r2

		}
	} else {
		if difc > 0 {
			difa = offang + r1
		} else {
			difa = offang - r1

		}

	}
	tan = vectang + difa
	xcc := math.Cos(tan)
	ycc := math.Sin(tan)
	ycc = -1 * ycc
	disf := float64(Nspeed(p.speed()-(dfs-Heuristicv), t))
	sdx, sdy := square(xcc, ycc, disf)
	p.pv = nil
	np := pod{p.x + sdx, p.y + sdy, sdx, sdy, cpangle, p.nextCheckPointId, &p, p.boost, 0, 0}

	return np

}

func stupidangle(a int) int {
	if a > 359 {
		a = 360 - a
	}
	if a < 0 {
		a = 360 + a
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func squaref(x, y, h float64) (float64, float64) {
	x = x * 10000
	y = y * 10000

	hy := math.Sqrt(x*x + y*y)

	l := hy / h

	return x / l, y / l
}

func square(x, y, h float64) (int, int) {
	sx, sy := squaref(x, y, h)

	return int(sx), int(sy)
}

func Circle(x, y, r int) {
	xi := r
	yi := 0
	for yi <= xi {
		fmt.Println(xi+x, yi+y)
		yi++

		re := xi*xi + yi*yi - r*r
		yc := 2*yi + 1
		xc := 1 - 2*xi
		if 2*(re+yc)+xc > 0 {
			xi--

		}

	}
}

func dist(x, y int) int {
	a := float64(x*x + y*y)
	b := math.Sqrt(a)
	return int(b)
}

func distf(x, y float64) float64 {
	a := x*x + y*y
	b := math.Sqrt(a)
	return b
}

func vectlineang(a, b cord) int {
	dx := b.x - a.x
	dy := b.y - a.y

	r := Calcangle(dx, -1*dy)
	return r
}

func angletransf(a int) int {
	a = 360 - a
	ang := angletorad(a)
	sx := math.Cos(ang)
	sy := math.Sin(ang)
	r := Calcanglef(sx, sy)
	return r
}

func Calcangle(x, y int) int {
	s := Calcanglef(float64(x), float64(y))
	return s
}

func Calcanglef(x, y float64) int {
	a := math.Atan2(y, x)
	s := radtoangle(a)
	return s
}

func Calcanglerad(x, y int) float64 {
	s := Calcangleradf(float64(x), float64(y))
	return s
}

func Calcangleradf(x, y float64) float64 {
	s := math.Atan2(y, x)
	return s
}

func cordang(ang, hd int) (int, int) {
	rad := float64(ang) * math.Pi / 180
	dx := math.Cos(rad)
	dy := dx * rad

	for i := 0; i < 100; i++ {
		dy = dx * rad
		if dist(int(dx), int(dy)) >= hd {
			return int(dx), int(dy)
		}
		dx += 5
	}
	return int(dx), int(dy)

}
func difangle(x, y, xi, yi int) int {
	a := Calcanglerad(x, y)
	b := Calcanglerad(xi, yi)

	res := difanglef(a, b)

	return res
}
func difangli(a, b int) int {
	ra := angletorad(a)
	rb := angletorad(b)
	r := difanglerad(ra, rb)
	res := radtoangle(r)
	return res
}
func difanglef(a, b float64) int {

	r := difanglerad(a, b)
	res := radtoangle(r)
	return res
}

func difanglerad(a, b float64) float64 {

	c := b - a
	dx := math.Cos(c)
	dy := math.Sin(c)
	r := math.Atan2(dy, dx)

	return r
}
func radtoangle(a float64) int {
	r := int(a * 180 / math.Pi)
	return r
}
func angletorad(a int) float64 {
	r := float64(a) * math.Pi / 180
	return r
}

func radtocord(a, h float64) (int, int) {
	sx := math.Cos(a)
	sy := -1 * (math.Sin(a))

	dx, dy := square(sx, sy, h)
	return dx, dy
}
func angletocord(a, h int) (int, int) {
	r := angletorad(a)
	hf := float64(h)
	dx, dy := radtocord(r, hf)
	return dx, dy
}

func mirror(x, y, xc, yc int) (int, int) {
	ax := x - xc
	ay := y - yc

	return x + ax, y + ay

}

func reverse90(x, y int) (int, int) {
	return y, x
}
func reverse90n(x, y int) (int, int) {
	return -y, -x
}

func Speed(s, t int) int {
	if t == 0 {
		t = 1
	}

	max := (650 * t / 100)
	sp := s * 100 / max
	per := 100 - sp
	thr := t * per / 100

	return thr

}
func Nspeed(s, t int) int {

	res := (s + t) * 85 / 100

	return res

}

func (p *pod) distrand(c int) int {
	res := p.distchp()
	i := p.nextCheckPointId
	for {
		if i == c {
			break
		}
		i = nextring(i)
		res += cpidist[i]

	}
	return res
}
func (p *pod) anglerand(c int) int {
	ax, ay := p.nextchp()
	i := p.nextCheckPointId

	res := p.frontchpdiff(ax, ay)
	for {
		if i == c {
			break
		}
		i = nextring(i)
		res += cpi[i]

	}
	return res
}

func (p *pod) blockn(op pod, nod idon) idon {
	x, y, thrust := 0, 0, 0
	lock := nod.lock

	chn := nod.chn

	cnx, cny := cp[chn].x, cp[chn].y

	opdistn := op.distrand(chn)

	ofx, ofy := radtocord(angletorad(cpi[chn]), 2000)

	cnx -= ofx
	cny -= ofy

	pdist := p.dist(cnx, cny)
	anglep := Abs(p.frontchpdiff(cnx, cny))

	fmt.Fprintln(os.Stderr, " isttt : ", pdist, opdistn)

	pos := pdist < opdistn //&& anglep<angleopn
	succ := pdist < 1200 || lock
	attack := opdistn < 5000
	rightang := Abs(p.frontchpdiff(op.x, op.y)) < 100 && chn == op.nextCheckPointId

	fmt.Fprintln(os.Stderr, " lock : ", lock)

	if succ {
		globalchn = chn
		fmt.Fprintln(os.Stderr, " suc : ", succ)

		x, y = angletocord(op.angle, 800)
		x, y = p.inifinte(op.x+x, op.y+y, 20000)

		lock = true

		thrust = 100
		if attack && rightang {
			thrust = 100
		}

	}
	if pos && !succ {
		x, y = cnx, cny
		fmt.Fprintln(os.Stderr, " pos : ", succ)

		thrust = calcstop(p.speed(), pdist)

		if anglep > 60 {

		}
	}
	if op.laps == 3 && op.nextCheckPointId == 0 {

		x, y = p.inifinte(op.x, op.y, 20000)
		thrust = 100
	}

	fail := (x+y == 0)

	return idon{nextring(chn), x, y, thrust, fail, lock}

}

func calcstop(s int, dis int) int {
	res := 0
	for i := 0; i < 20; i++ {
		s = s * 85 / 100
		res += s
		if res > dis {
			return 0

		}
	}
	return 100
}

func predict(p, op pod, step int) (int, int) {
	pp := p
	fxx, fyy := angletocord(p.angle, 800)
	if Abs(difangli(p.angle, op.angle)) > 150 {
		return p.x + fxx, p.y + fyy
	}
	ax, ay := 0, 0

	thrust := 100
	thrustp := 100
	cx, cy := p.nextchp()

	for i := 0; i < step; i++ {
		thrust = 100
		diffang := Abs(p.frontchpdiff(cx, cy))
		if diffang > 130 {
			thrust = 0
		}

		if p.preturn(thrust) {
			cx, cy = p.preturnchp()
		}
		if Abs(diffang) > 16 {
			p = predictcurve(p, cx, cy, thrust)
			fmt.Fprintln(os.Stderr, "curve : ")

		} else {
			p = predictline(p, cx, cy, thrust)
			fmt.Fprintln(os.Stderr, "line : ")

		}
		if p.dist(cx, cy) < 600 {
			cx, cy = cp[nextring(p.nextCheckPointId)].x, cp[nextring(p.nextCheckPointId)].y
			p.nextCheckPointId = nextring(p.nextCheckPointId)

		}
		fmt.Fprintln(os.Stderr, "tarrrget : ", p, cx, cy)

		pb := op
		if i/5 == 0 {
			for ip := 0; ip < step; ip++ {

				fx, fy := angletocord(p.angle, 800)
				ax, ay = p.x+fx, p.y+fy

				thrustp = 100
				diffangp := Abs(pb.frontchpdiff(ax, ay))
				if diffangp > 130 {
					thrustp = 0
				}

				if Abs(diffangp) > 16 {
					pb = predictcurve(pb, ax, ay, thrustp)

				} else {
					pb = predictline(pb, ax, ay, thrustp)

				}

				if pb.dist(ax, ay) < 800 && i == ip {
					fmt.Fprintln(os.Stderr, "hunnnter : ", pb, ax, ay)

					ax, ay = op.inifinte(ax, ay, 20000)
					return ax, ay
				}
			}
		}
	}

	return op.inifinte(pp.x+fxx, pp.y+fyy, 20000)

}

func predictcol(p, op pod, px, py, step int) (int, int, bool) {

	thrust := 100
	thrustp := 100
	cx, cy := p.nextchp()
	ax, ay := px, py

	for i := 0; i < step; i++ {
		thrust = 100
		diffang := Abs(p.frontchpdiff(cx, cy))
		if diffang > 130 {
			thrust = 0
		}
		if Abs(diffang) > 16 {
			p = predictcurve(p, cx, cy, thrust)

		} else {
			p = predictline(p, cx, cy, thrust)

		}
		if p.dist(cx, cy) < 600 {
			cx, cy = cp[nextring(p.nextCheckPointId)].x, cp[nextring(p.nextCheckPointId)].y
			p.nextCheckPointId = nextring(p.nextCheckPointId)
			//	finalx, finaly = p1.modinifinte(nextring(p1.nextCheckPointId))

		}

		thrustp = 100
		diffangp := Abs(op.frontchpdiff(ax, ay))
		if diffangp > 16 {
			thrustp = 100
		}
		if Abs(diffangp) > 16 {
			op = predictcurve(op, ax, ay, thrustp)

		} else {
			op = predictline(op, ax, ay, thrustp)

		}

		if dist(op.x-p.x, op.y-p.y) < 700 {
			return ax, ay, true
		}

	}

	return p.x, p.y, false

}

func boost(i int, f bool) int {
	c := 100
	l := 650

	for s := 1; s < i; s++ {
		l = l - c
		c -= 14

	}
	if f {
		return l * (-1)
	}
	return l
}
