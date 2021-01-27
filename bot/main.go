
package main

import (
	"fmt"
	"os"
	"sync"
	 "math"
)

var onlyOnce sync.Once

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
type field struct {
	x, y int
	lap  int
	cp   int
}

type info struct {
    ang int
    dist int
}

type input struct {
	x, y                             int
	nextCheckpointX, nextCheckpointY int
	nextCheckpointDist               int
	nextCheckpointAngle              int
	opponentX, opponentY             int
}

type drone struct {
	x, y         int
	distx, disty int
	// speedx,speedy int
	angle int
	dist  int
	boost bool
	first bool
	path  int
	//checkpoint int
}

type checkpoint struct {
	x, y         int
	pathx, pathy int
	num          int
	dist         int
	ang          int
}

func (s *checkpoint) exist(x, y int) bool {
	if s.x == x && s.y == y {
		return true
	}
	return false
}

func existch(s *[]checkpoint, i input) bool {
	for _, v := range *s {
		if v.exist(i.nextCheckpointX, i.nextCheckpointY) {
			return true
		}
	}

	return false

}

func chp(ax, ay *int, i input) bool {
	if *ax != i.nextCheckpointX && *ay != i.nextCheckpointY {
		*ax = i.nextCheckpointX
		*ay = i.nextCheckpointY
		return true
	}
	return false
}
func initcheckpoint(i input, cou int) checkpoint {
	t := checkpoint{
		x:     i.nextCheckpointX,
		y:     i.nextCheckpointY,
		pathx: Abs(i.nextCheckpointX - i.x),
		pathy: Abs(i.nextCheckpointY - i.y),
		num:   cou,
		dist:  i.nextCheckpointDist,
		ang:   Abs(i.nextCheckpointAngle),
	}
	return t
}
func equalxy(c checkpoint, x, y int) bool {
	if c.x == x && c.y == y {
		return true

	}
	return false
}

func trimc(c *[]checkpoint, ang []int) {
	t := (*c)[len(*c)-1]

	*c = append([]checkpoint{t}, (*c)[:len(*c)-2]...)
	for i, v := range *c {
		v.ang = ang[i]
		v.num = i
	}
}

func search(c []checkpoint, cou int) int {
	for i, v := range c {
		if v.num == cou {
			return i
		}
	}
	return -1

}
func collide(x, opx, difx, odifx, y, opy, dify, odify int) bool {

	//fmt.Fprintln(os.Stderr, "collide ...", pvx, pvy, dx, dy)

	ax := x + difx
	ay := y + dify
	ox := opx + odifx
	oy := opy + odify

	fx := ax - ox
	fy := ay - oy
	//	fmt.Fprintln(os.Stderr, "collideab ...", a, b)
	
	if fx*fx+fy*fy <= 700*700 {
fmt.Fprintln(os.Stderr, "coll ...", fx*fx+fy*fy)
			fmt.Fprintln(os.Stderr, "collide ...")

			return true
		
	}

	return false
}

func detectcollide(x, opx, difx, odifx, y, opy, dify, odify int) bool {

	//fmt.Fprintln(os.Stderr, "collide ...", pvx, pvy, dx, dy)

	ax := x +difx*2-(difx*8/10)*2
	ay := y +dify*2-(dify*8/10)*2
	ox := opx + odifx*2
	oy := opy + odify*2
	if difx == 0 || dify==0 || odifx==0 || odify==0 {
		return false
	}
    
	fx := ax - ox
	fy := ay - oy
	//	fmt.Fprintln(os.Stderr, "collideab ...", a, b)
		fmt.Fprintln(os.Stderr, "detect c ...", fx,fy)

	if fx*fx+fy*fy <= 800*800 {

		return true

	}

	return false
}

func boost(i int, f bool) int {
	c := 100
	l := 750

	for s := 1; s < i; s++ {
		l = l - c
		c -= 14

	}
	if f {
		return l * (-1)
	}
	return l
}

func linemove(x, opx, difx, odifx, y, opy, dify, odify, cx, cy int) bool {
	if difx == 0 && dify == 0 {
		return false
	}
	xx := x
	yy := y

	aaa := opx
	bbb := opy

  //  inc:=5
	//cont:=4
//    h:=hypdist(cx-opx,cy-opy)
   // l:=hypoth(cx-x,cy-y)
  

	d := (Abs(difx) * 100) / (Abs(difx) + Abs(dify))


	for i := 1; i < 4; i++ {
		if difx < 0 {
			xx += difx + ((boost(i, true) * d) / 100)
			aaa += odifx

		} else {
			xx += difx + ((boost(i, false) * d) / 100)

			aaa += (odifx)

		}
		if dify < 0 {
			yy += dify + ((boost(i, true) * (100 - d)) / 100)
			bbb += odify

		} else {
			yy += dify + ((boost(i, false) * (100 - d)) / 100)
			bbb += odify

		}
		aa := xx
		a := Abs(aa - aaa)

		bb := yy

		b := Abs(bb - bbb)

		// fmt.Fprintln(os.Stderr, "linemove ...", aa, aaa,bb,bbb)

	//if collide(aa, aaa, 0, 0, bb, bbb,0, 0) {
		// fmt.Fprintln(os.Stderr, "linemove ...", aa, aaa,bb,bbb)
         		 fmt.Fprintln(os.Stderr, "hyp ...", hypdist(a,b))

          if hypdist(a,b)<650 {

			  if Abs(cx-aa)<Abs(cx-aaa) && Abs(cy-bb)<Abs(cy-bbb) || hypdist(cx-aaa,cy-bbb)<700 {

			 
			return true
		}
        }
	}

	return false
}

/*
func linepredict(x, opx, difx,odifx ,cx,cy int) (bool) {
	if difx == 0 && dify == 0 {
		return false
	}
    xx:=x
    yy:=y

    xxx:=opx
    yyy:=opy




    d:=(Abs(difx)*100)/(Abs(difx)+Abs(dify))



 difx=-100*d/100
 dify=-100*(100-d)/100

	for i := 1; i < 4; i++ {
        xx+= difx
              xxx+= odifx


                yy+= dify
                        yyy+= odify

difx-=18*d/100
dify-=18*(100-d)/100

a:=xxx-xx
b:=yyy-yy


       // fmt.Fprintln(os.Stderr, "linemove ...", aa, aaa,bb,bbb)

		if (a*a+b*b<= 800*800)  {
           if (cx-x)*(cx-x)+(cy-y)*(cy-y)<1500*1500 {
			//
            //fmt.Fprintln(os.Stderr, "step ...", a,b)

			return true
                     }
    }
    }

	return false
}
*/
func linesatgn(x, difx, y, dify, cx, cy int) bool {
	if difx == 0 && dify == 0 {
		return false
	}
	xx := x
	yy := y

	for i := 1; i < 4; i++ {
		xx += difx * i

		yy += dify * i

		a := cx - xx
		b := cy - yy

		// fmt.Fprintln(os.Stderr, "linemove ...", aa, aaa,bb,bbb)

		if a*a+b*b <= 600*600 {

			//	fmt.Fprintln(os.Stderr, "step ...", i)

			return true
		}
	}

	return false
}
func chox(xo, yo, x, y int) bool {
	if Abs(xo-x) < 1500 && Abs(xo-x) > 600 {

		return true

	}
	return false
}
func choy(xo, yo, x, y int) bool {
	if Abs(yo-y) < 1500 && Abs(yo-y) > 600 {

		return true
	}
	return false
}

func main() {
	f := field{x: 15999, y: 8999, lap: 1}
	hs := make(map[int]info)
	//    cux:=0
	//  cuy:=0
	lock := false
	newc := false
	initx := 0
	inity := 0
	// prox:=0
    finishx:=0
    finishy:=0

	cc := 0
	//proy:=0
pr:=0
pro:=0
//    sec:=false
bdist:=0
	prev := input{}
                     detectco:=false

	thr := make([]int, 4)
	ss := false

	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		inp := input{x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle, opponentX, opponentY}
		fmt.Fprintln(os.Stderr, "step ...", inp)

		onlyOnce.Do(func() {
			initx = nextCheckpointX
			inity = nextCheckpointY
			prev = inp
            finishx=(x+opponentX)/2
            finishy=(y+opponentY)/2

		})

		sst := ""

		thrust := 100

		if (prev.nextCheckpointX + prev.nextCheckpointY) != (nextCheckpointX + nextCheckpointY) {
                 if (nextCheckpointX + nextCheckpointY)==initx+inity {
                   				f.lap++
                 }
			if _, ok := hs[prev.nextCheckpointX+prev.nextCheckpointY]; !ok {

				hs[prev.nextCheckpointX+prev.nextCheckpointY] = info{nextCheckpointAngle,nextCheckpointDist}
                if nextCheckpointDist>bdist {
                    bdist=nextCheckpointDist
                }
			}  else {
                if hs[prev.nextCheckpointX+prev.nextCheckpointY].ang==0 {
                    hs[prev.nextCheckpointX+prev.nextCheckpointY]=info{nextCheckpointAngle,nextCheckpointDist}
                }
            }
			
		}

		fmt.Fprintln(os.Stderr, "hs ...", hs, f.lap)

		if collide(x, opponentX, x-prev.x, opponentX-prev.opponentX, y, opponentY, y-prev.y, opponentY-prev.opponentY) {
			dx := opponentX - prev.opponentX
			dy := opponentY - prev.opponentY
            xd:=x-prev.x
            yd:=y-prev.y
		fmt.Fprintln(os.Stderr, "abs ...", Abs(xd-dx)+Abs(yd-dy))

            if Abs(xd-dx)+Abs(yd-dy)>600 {

				if sst == "" {
					sst = "SHIELD"
                    lock = false
				}
				
            }
        }
		

		if f.lap == 1 {
			
            if nextCheckpointDist > 5000 && Abs(nextCheckpointAngle) < 3 && !ss   {
				
            if v, ok := hs[nextCheckpointX+nextCheckpointY]; ok {
               if v.dist==bdist {
            
				if sst == "" {
					sst = "BOOST"
                    ss = true
				lock = false
				}
               }
            }
			}
		}


		if f.lap == 2 {
			if nextCheckpointDist > 5000 && Abs(nextCheckpointAngle) < 3 && !ss   {
				
            if v, ok := hs[nextCheckpointX+nextCheckpointY]; ok {
               if v.dist==bdist {
            
				if sst == "" {
					sst = "BOOST"
                    ss = true
				lock = false
				}
               }
            }
			}
		}

		if f.lap == 3 {

			if nextCheckpointX == initx && nextCheckpointY == inity && Abs(nextCheckpointAngle) < 3 && !ss {
				
				  
            
				if sst == "" {
					sst = "BOOST"
                    lock = false
				ss = true
				
               }
            
			}

		}
		 fmt.Fprintln(os.Stderr, "ss ...",ss)

		if linemove(x, opponentX, x-prev.x, opponentX-prev.opponentX, y, opponentY, y-prev.y, opponentY-prev.opponentY, nextCheckpointX, nextCheckpointY) && !ss && Abs(nextCheckpointAngle)<10 {

			if sst == "" {
				sst = "BOOST"
                lock = false

			ss = true
			}

			

		}



		if prev.nextCheckpointX != nextCheckpointX && prev.nextCheckpointY != nextCheckpointY {

			newc = true
		}

		//	thrust = 100 - (Abs(nextCheckpointAngle) * 30 / 180)
        distop:=hypdist(nextCheckpointX-opponentX,nextCheckpointY-opponentY)
        ds:=hypdist(x-opponentX,y-opponentY)
    if distop<1200 && ds < 1600 && ds> 800 {
            // 		fmt.Fprintln(os.Stderr, "distop  ...", distop)

     	//	fmt.Fprintln(os.Stderr, "src  ...", true)

   //    sec=true
    }
     if nextCheckpointAngle<1200 && ds < 1600 && ds> 800 {
             		fmt.Fprintln(os.Stderr, "distop  ...", distop)

     		fmt.Fprintln(os.Stderr, "src  ...", true)

//       sec=true
    }

		if newc {
			lock = true

			thr = []int{100, 100, 100, 100}


			if Abs(nextCheckpointAngle) > 30 {
                
				thr = []int{90, 95, 100, 100}
                if nextCheckpointDist <6500 {
            	thr = []int{40, 50, 60, 70}

                }

			}

			if Abs(nextCheckpointAngle) > 60 {
				thr = []int{80, 90, 100, 100}
                if nextCheckpointDist <6500 {
            	thr = []int{40, 50, 60, 70}

                }

			}
             


			

			if Abs(nextCheckpointAngle) > 90 {
				thr = []int{60, 70, 80, 80}
                
                    if nextCheckpointDist <6500 {
            	thr = []int{20, 30, 40, 50}

                }



			}

            
           
            

			if Abs(nextCheckpointAngle) > 130 {
				thr = []int{20, 40, 60, 80}
                     if nextCheckpointDist <6500 {
            	thr = []int{0, 10, 20, 40}

                }
                
                                 if detectco {
detectco=false
                			thr = []int{60, 50, 40 , 30}

                                 }
			}

          
        

				
			
		}

		//	sss = false

		newc = false

		//  thr = []int{100,100,100,100}

		if lock && cc < 4 {

			thrust = thr[cc]
			cc++

		} else {
			lock = false
			cc = 0

		}


		if nextCheckpointDist < 1500 && hypdist( x-prev.x,  y-prev.y)>500  {
		
                    			thrust = 60

            if   Abs(nextCheckpointAngle) <10 &&  Abs(nextCheckpointAngle) >3 {
              // 			thrust = 40
            }
            if v, ok := hs[nextCheckpointX+nextCheckpointY]; ok {
	fmt.Fprintln(os.Stderr, "thrud ...",v)
if Abs(v.ang)>0{
              thrust=100    
			} 
              if Abs(v.ang)>10{
              thrust=90    
			} 
              if Abs(v.ang)>60 {
              thrust=80    
			} 
              if Abs(v.ang)>90 {
              thrust=40    
			} 
              if Abs(v.ang)>130 {
              thrust=0    
			} 
            

            if v.dist<6500  {
                          thrust=40
                          }  

            }
            if f.lap==1 {
            			thrust = 80

            }
            if f.lap==3 && nextCheckpointX==finishx && nextCheckpointY==finishy {
                           			thrust = 100
            }

            if detectcollide(x, opponentX, x-prev.x, opponentX-prev.opponentX, y, opponentY, y-prev.y, opponentY-prev.opponentY) {
                           			thrust = 100
                     detectco=true
        
        }

		
		}
           
              if nextCheckpointDist <3000 && nextCheckpointDist >1500 && Abs(nextCheckpointAngle) >3 {
     			thrust = 60
                 


                }          // thrust=100    


if prev.nextCheckpointX != nextCheckpointX && prev.nextCheckpointY != nextCheckpointY {

//			sec = false
		}
		
        

		/*
		   hx:=x-opponentX
		       hy:=y-opponentY
		       subnx:=(nextCheckpointX-opponentX)
		       subny:=(nextCheckpointY-opponentY)
		       nextCheckpointDistop:=subnx*subnx+subny*subny
		       if nextCheckpointDistop<0 {
		       }

		*/
		/*
			                   if linepredict(x, opponentX, x-prevx,opponentX-prevopx, y, opponentY, y-prevy,opponentY-prevopy,nextCheckpointX,nextCheckpointY){
			                    if nextCheckpointAngle<4 {


			                   collid=true
			                                			thrust = 0
			                    }

					}


		*/

		if sst == "" {
			fmt.Printf("%d %d %d\n", nextCheckpointX, nextCheckpointY, thrust)
		} else {
			fmt.Printf("%d %d %s\n", nextCheckpointX, nextCheckpointY, sst)
		}

		fmt.Fprintln(os.Stderr, "dif xy ...", x-prev.x, y-prev.y, " = ",hypdist( x-prev.x,  y-prev.y))
		fmt.Fprintln(os.Stderr, "dif oppon ...", opponentX-prev.opponentX, opponentY-prev.opponentY, " = ", hypdist( opponentX-prev.opponentX,  opponentY-prev.opponentY))




		fmt.Fprintln(os.Stderr, "dif speed ...",hypdist( x-prev.x,  y-prev.y)-pr)
		fmt.Fprintln(os.Stderr, "dif speed op ...", hypdist( opponentX-prev.opponentX,  opponentY-prev.opponentY)-pro)
      
     pr=hypdist( x-prev.x,  y-prev.y)

      pro=hypdist( opponentX-prev.opponentX,  opponentY-prev.opponentY)

		//fmt.Fprintln(os.Stderr, "speed ...", nextCheckpointDist-prevdist)
		/*
			                        		fmt.Fprintln(os.Stderr, "hashmap...", hs)

			                		fmt.Fprintln(os.Stderr, "me  angle...", nextCheckpointAngle)

					fmt.Fprintln(os.Stderr, "dist ...",nextCheckpointDist )

					fmt.Fprintln(os.Stderr, "cordxy ...", x, y)

					fmt.Fprintln(os.Stderr, "dif xy ...", x-prevx, y-prevy," = ",Abs(x-prevx)+ Abs(y-prevy))
			         fmt.Fprintln(os.Stderr, "dif oppon ...", opponentX-prevopx,opponentY-prevopy," = ",Abs(opponentX-prevopx)+ Abs(opponentY-prevopy))

			        gx:=Abs(x-prevx)
			        gy:=Abs(y-prevy)

			        gox= Abs(opponentX-prevopx)
			        goy=Abs(opponentY-prevopy)


			            	fmt.Fprintln(os.Stderr, "xggy ...",gx-prx,gy-pry," = ",gx-prx+gy-pry )


					fmt.Fprintln(os.Stderr, "goxgoy ...", gox-prox,goy-proy," = ",gox-prox+goy-proy)







		*/

		newc = false
		prev = inp

	}

}


func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}




/*
booststart
boostfinish
boostavoid
boostbeat

checkpoint
detectstraigt
detectturn
detectturndirect
detectlead
detectopponentboost
detetctcollision
turn
*/

func power(a int) int {
	return a * a
}

func Circle(x, y, r int) {
	xi := r
	yi := 0
	for yi <= xi {
		fmt.Println(xi+x, yi+y)
		/*
			fmt.Println(yi + x, xi + y);
			fmt.Println(-xi + x, yi + y);
			fmt.Println(-yi + x, xi + y);
			fmt.Println(-xi + x, -yi + y);
			fmt.Println(-yi + x, -xi + y);
			fmt.Println(xi + x, -yi + y);
			fmt.Println(yi + x, -xi + y);
		*/
		yi++

		re := xi*xi + yi*yi - r*r
		yc := 2*yi + 1
		xc := 1 - 2*xi
		if 2*(re+yc)+xc > 0 {
			xi--

		}

	}
}


func hypdist(x, y int) int {
	a := float64(power(x) + power(y))
	b := math.Sqrt(a)
	return int(b)
}
