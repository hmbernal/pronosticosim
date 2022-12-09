package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func main() {

	var tipo string
	var fArr [25]float64
	var ss [1000]string
	var str1 string
	var fArr2 [25]float64 = [25]float64{2.4, 2.3, 2.2, 2.1, 2.1, 1.9, 8.8, 17.7, 16.6, 15.5, 14.4, 13.3, 12, 2, 11.1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1} //<<< FUNCIONA
	//var fArr2 [25]float64 = [25]float64{24.4, 23.3, 22.2, 21.1, 20.1, 19.9, 18.8, 17.7,16.6,15.5,14.4,13.3,12,2,11.1,10,9,8,7,6,5,4,3,2,1}  //<<< NO FUNCIONA

	for i := 0; i < 24; i++ {
		ss[i] = strconv.FormatFloat(fArr2[i], 'g', 5, 64)
	}

	tipo = "H"

	//Dias
	if tipo == "D" {
		//./pronosticoviento.bin" D 7.7 6.6 5.5 4.4 3.3 2.2 1.1  <<< Linea de comandos
		//prc := exec.Command("/ucm/middleware/go/src/ucmgo/applications/pronosticoviento/pronosticoviento",tipo, "7.7", "6.7", "5.7", "4.7", "3.7", "2.7", "1.7")  //<<< FUNCIONA
		//prc := exec.Command("/ucm/middleware/go/src/ucmgo/applications/pronosticoviento/pronosticoviento",tipo, "77.7", "66.7", "5.7", "4.7", "3.7", "2.7", "1.7")  //*** NO FUNCIONA

		prc := exec.Command("./pronosticoviento.bin", tipo, ss[0], ss[1], ss[2], ss[3], ss[4], ss[5], ss[6])

		out := bytes.NewBuffer([]byte{})
		prc.Stdout = out
		err := prc.Run()
		if err != nil {
			fmt.Println(err)
		}
		prc.Wait()

		/*if prc.ProcessState.Success() {
				fmt.Println("P, Proceso pronosticoviento ejecutado con exito con salida:")
				fmt.Println("P, out.String:",out.String())
		      }*/

		p := make([]byte, 1000)
		n, err := out.Read(p)
		fmt.Println("P, n:", n)
		if err != nil {
			fmt.Println("*** P, ERROR out.Read")
		}

		str1 = string(p)

		//stringLength := len(str1)
		//fmt.Println("str1:\n",str1," leng:",stringLength)

		fmt.Sscanf(str1, "%g\n%g\n%g\n%g\n%g\n%g\n%g", &fArr[0], &fArr[1], &fArr[2], &fArr[3], &fArr[4], &fArr[5], &fArr[6])

		for i := 0; i < 7; i++ {
			fmt.Println("P, salida viento:", fArr[i])
			ss[i] = strconv.FormatFloat(fArr[i], 'g', 5, 64)
		}

		//./pronosticogeneracion.bin" D 7.7 6.6 5.5 4.4 3.3 2.2 1.1
		prc2 := exec.Command("./pronosticogeneracion.bin", tipo, ss[0], ss[1], ss[2], ss[3], ss[4], ss[5], ss[6])

		out2 := bytes.NewBuffer([]byte{})
		prc2.Stdout = out2
		err = prc2.Run()
		if err != nil {
			fmt.Println(err)
		}
		prc.Wait()
		fmt.Println("Sali de pronosticogeneracion")

		//fmt.Println(string(out))
		/*      if prc2.ProcessState.Success() {
				fmt.Println("P, Proceso generacion ejecutado con exito con salida:")
				fmt.Println("P, generacion out2.String():", out2.String())
		      }*/

		p = make([]byte, 1000)
		n, err = out2.Read(p)
		fmt.Println("P, n:", n)
		if err != nil {
			fmt.Println("*** P generacion, ERROR out.Read")
		}

		str1 = string(p)

		//stringLength := len(str1)
		//fmt.Println("str1:\n",str1," leng:",stringLength)

		fmt.Sscanf(str1, "%g\n%g\n%g\n%g\n%g\n%g\n%g", &fArr[0], &fArr[1], &fArr[2], &fArr[3], &fArr[4], &fArr[5], &fArr[6])

		for i := 0; i < 7; i++ {
			fmt.Println("P, salida generacion:", fArr[i])
		}

		//Horas
	} else {
		//No acepta 8.0, 22.3
		prc := exec.Command("./pronosticoviento.bin", tipo, "4.7", "3.6", "2.5", "2.4", "2.3", "9.2", "8.1", "8.1", "8.1", "8.1", "2.3", "2.3", "1.2", "1.1", "1.5", "9.1", "8.3", "7.1", "6.2", "5.3", "4.4", "3.5", "2.2", "1.1")

		out := bytes.NewBuffer([]byte{})
		prc.Stdout = out
		err := prc.Run()
		if err != nil {
			fmt.Println(err)
		}
		prc.Wait()

		/*if prc.ProcessState.Success() {
				fmt.Println("P hora, Proceso pronosticoviento ejecutado con exito con salida:")
				fmt.Println("P hora, out.String:",out.String())
		      }*/

		p := make([]byte, 1000)
		n, err := out.Read(p)
		fmt.Println("P, n:", n)
		if err != nil {
			fmt.Println("*** P, ERROR out.Read")
		}

		str1 = string(p)

		fmt.Sscanf(str1, "%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g\n%g",
			&fArr[0], &fArr[1], &fArr[2], &fArr[3], &fArr[4], &fArr[5], &fArr[6], &fArr[7], &fArr[8], &fArr[9], &fArr[10], &fArr[11], &fArr[12], &fArr[13], &fArr[14], &fArr[15], &fArr[16], &fArr[17], &fArr[18], &fArr[19], &fArr[20], &fArr[21], &fArr[22], &fArr[23])

		for i := 0; i < 24; i++ {
			fmt.Println("P, salida viento:", fArr[i])
			ss[i] = strconv.FormatFloat(fArr[i], 'g', 5, 64)
		}

		//./pronosticogeneracion.bin" D 7.7 6.6 5.5 4.4 3.3 2.2 1.1
		prc2 := exec.Command("./pronosticogeneracion.bin", tipo, ss[0], ss[1], ss[2], ss[3], ss[4], ss[5], ss[6], ss[7], ss[8], ss[9], ss[10], ss[11], ss[12], ss[13], ss[14], ss[15], ss[16], ss[17], ss[18], ss[19], ss[20], ss[21], ss[22], ss[23])

		out2 := bytes.NewBuffer([]byte{})
		prc2.Stdout = out2
		err = prc2.Run()
		if err != nil {
			fmt.Println(err)
		}
		prc.Wait()

		//fmt.Println(string(out))
		if prc2.ProcessState.Success() {
			fmt.Println("P hora, Proceso pronosticogeneracion ejecutado con exito con salida:")
			fmt.Println("P hora, out2.String():", out2.String())
		}
		p = make([]byte, 1000)
		n, err = out2.Read(p)
		fmt.Println("P, n:", n)
		if err != nil {
			fmt.Println("*** P generacion, ERROR out.Read")
		}

		str1 = string(p)

		fmt.Sscanf(str1, "%g\n%g\n%g\n%g\n%g\n%g\n%g", &fArr[0], &fArr[1], &fArr[2], &fArr[3], &fArr[4], &fArr[5], &fArr[6], &fArr[7], &fArr[8], &fArr[9], &fArr[10], &fArr[11], &fArr[12], &fArr[13], &fArr[14], &fArr[15], &fArr[16], &fArr[17], &fArr[18], &fArr[19], &fArr[20], &fArr[21], &fArr[22], &fArr[23])

		for i := 0; i < 24; i++ {
			fmt.Println("P, salida generacion:", fArr[i])
			//ss[i] = strconv.FormatFloat(fArr[i], 'g', 5, 64)
		}

	}

}
