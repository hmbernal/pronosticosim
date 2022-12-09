///ucm/middleware/go/src/ucmgo/applications/pronosticoviento
package main

import (
	"fmt"
	//"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func funcionRandomSem(fEnt [7]float64, fSal *[7]float64) {

	var min, max float64

	for ii := 0; ii < 7; ii++ {
		if fEnt[ii] == 0 {
			fEnt[ii] = 0.1
		}
		min = fEnt[ii] * 0.9
		max = fEnt[ii] * 1.1
		fSal[ii] = min + rand.Float64()*(max-min)
	}

}

func funcionRandomHr(fEnt [24]float64, fSal *[24]float64) {

	var min, max float64

	for ii := 0; ii < 24; ii++ {
		if fEnt[ii] == 0 {
			fEnt[ii] = 0.1
		}
		min = fEnt[ii] * 0.9
		max = fEnt[ii] * 1.1
		fSal[ii] = min + rand.Float64()*(max-min)
	}

}

//./pronosticoviento D 7.7 6.6 5.5 4.4 3.3 2.2 1.1
func main() {
	indexCmd := 0
	var valHr [24]float64
	var valHrSal [24]float64
	var valSemana [7]float64
	var valSemanaSal [7]float64

	cmdAux := "" //???
	//f, err := strconv.ParseFloat("3.1415926535", 8) //???
	//fmt.Println(f)                             // ???

	for _, cmd := range os.Args {
		cmdAux = cmd
		cmdAux = os.Args[1]
		indexCmd++
		if (cmdAux == "H") || (cmdAux == "D") { //???
			if (os.Args[1] == "H") && (indexCmd > 1) && (indexCmd < 24) {
				// Horas >>>
				//fmt.Println(">>> EXITO, datos horarios (24)")
				valHr[indexCmd-2], _ = strconv.ParseFloat(os.Args[indexCmd], 64)
				//if err != nil {
				//	fmt.Println(err)
				//}
				//fmt.Println("ValHr[",indexCmd-2, "]:",valHr[indexCmd-2])
				//fmt.Println(valHr[indexCmd-2])
				if indexCmd == 25 {
					break
					//fmt.Println("Aquí se rompe el bucle")
				}
			} else if (os.Args[1] == "D") && (indexCmd > 1) && (indexCmd < 9) {
				// Dias >>>
				//fmt.Println(">>> EXITO, datos semanales (24)")
				valSemana[indexCmd-2], _ = strconv.ParseFloat(os.Args[indexCmd], 64)
				//if err != nil {
				//	fmt.Println(err)
				//}
				//fmt.Println("valSemana[",indexCmd-2, "]:",valSemana[indexCmd-2])
				//valSemana[indexCmd-2]=valSemana[indexCmd-2]+10
				//Es lo que envia de regreso a apipronostico
				//fmt.Println(valSemana[indexCmd-2])
				if indexCmd == 8 {
					break
					//fmt.Println("Aquí se rompe el bucle")
				}
			}
		} else {
			//fmt.Println("*** ERROR, ni semanal, ni horario")
			break
			//fmt.Println("Aquí se rompe el bucle")
		}
	}

	//time.Sleep(60 * time.Second)

	rand.Seed(time.Now().UnixNano())
	if cmdAux == "D" {
		funcionRandomSem(valSemana, &valSemanaSal)
		for ii := 0; ii < 7; ii++ {
			fmt.Println(valSemanaSal[ii])
		}
	} else if cmdAux == "H" {
		funcionRandomHr(valHr, &valHrSal)
		for ii := 0; ii < 24; ii++ {
			fmt.Println(valHrSal[ii])
		}
	}

}
