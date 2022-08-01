package main

import (
	"fmt"
	"net/http"
	"strings"
)

func SimpleSelectTag(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>

<body>
	<script>
	function speak(words){
		// console.log("working");
		let utterance = new SpeechSynthesisUtterance(words);
		let voicesArray = speechSynthesis.getVoices();
		utterance.voice = voicesArray[2];
		speechSynthesis.speak(utterance);
	}
	function speakChoices(option) {
		if (option == 1) {
			var selectBox = document.getElementById("q1");
			var selectedValue = selectBox.options[selectBox.selectedIndex].value;
			speak(selectedValue);
		}
		if (option == 2) {
			var selectBox = document.getElementById("q2");
			var selectedValue = selectBox.options[selectBox.selectedIndex].value;
			speak(selectedValue);
		}
		if (option == 3) {
			var selectBox = document.getElementById("q3");
			var selectedValue = selectBox.options[selectBox.selectedIndex].value;
			speak(selectedValue);
		}
		if (option == 4) {
			var selectBox = document.getElementById("q4");
			var selectedValue = selectBox.options[selectBox.selectedIndex].value;
			speak(selectedValue);
		}
	}
	</script>


	<form action="/quiz" method="post">
		<p onclick="speak('What sounds the most fun to you?');">What sounds the most fun to you?</p>
		
		<select id="q1" name="q1" onclick = "speakChoices(1)" )>
		<option id = "petting_zoo" value="Going to the petting zoo">Going to a petting zoo
		</option>
		<option id = "soccer" value="Playing soccer with your friends">Playing soccer with your friends</option>
		<option id="hike" value = "Taking a hike">Taking a hike</option>
		<option id="eating_out" value="Eating out with friends">Eating out with friends</option>
		</select>
		<p onclick="speak('What would people describe you as?')">What would people describe you as?</p>
		<select id = "q2" name = "q2" onclick = "speakChoices(2)" >
		<option id="animal_lover" value = "Animal Lover" >Animal Lover
		</option>
		<option id="athlete" value = "athlete">Athlete</option>
		<option id="nature_lover" value="Nature Lover">Nature Lover</option>
		<option id="foodie" value="foodie">Foodie</option>
		</select>
		<p onclick="speak('What do you associate the color orange with?')">What do you associate the color orange with?
		</p>
		<select id = "q3" name = "q3" onclick = "speakChoices(3)" >
		<option id="fox" value="fox">Fox
		</option>
		<option id = "basketball" value="basketball">Basketball</option>
		<option id="leaves" value="fall leaves">Fall Leaves</option>
		<option id="fruit" value="Orange Fruit">Orange Fruit</option>
		</select>
		<p onclick="speak('Would you rather spend your day inside or outside?')">Would you rather spend your day inside or outside?
		</p>
		<select id = "q4" name="q4" onclick = "speakChoices(4)" >

		<option id="inside" value="inside">Inside</option>
		<option id = "outside" value="outside">Outside</option>
		</select>
		<p></p>
		<input onclick="speak('submit')" type = "submit" value = "submit">
	</form>
	<div id="audio"></div>
</body>
</html>`

	w.Write([]byte(html))

}

func quiz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")

	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {

		fmt.Println("get method")
	} else {
		r.ParseForm()
		fmt.Println("Q1:", r.Form["q1"])
		fmt.Println("Q2:", r.Form["q2"])
		fmt.Println("Q3:", r.Form["q3"])
		fmt.Println("Q4:", r.Form["q4"])
		var animal int = 0
		var sport int = 0
		var nature int = 0
		var food int = 0

		if strings.Join(r.Form["q1"], "") == "petting_zoo" {
			animal += 1
		} else if strings.Join(r.Form["q1"], "") == "soccer" {
			sport += 1
		} else if strings.Join(r.Form["q1"], "") == "hike" {
			nature += 1
		} else if strings.Join(r.Form["q1"], "") == "eating_out" {
			food += 1
		}

		if strings.Join(r.Form["q2"], "") == "animal_lover" {
			animal += 1
		} else if strings.Join(r.Form["q2"], "") == "athlete" {
			sport += 1
		} else if strings.Join(r.Form["q2"], "") == "nature_lover" {
			nature += 1
		} else if strings.Join(r.Form["q2"], "") == "foodie" {
			food += 1
		}

		if strings.Join(r.Form["q3"], "") == "fox" {
			animal += 1
		} else if strings.Join(r.Form["q3"], "") == "basketball" {
			sport += 1
		} else if strings.Join(r.Form["q3"], "") == "leaves" {
			nature += 1
		} else if strings.Join(r.Form["q3"], "") == "fruit" {
			food += 1
		}

		if strings.Join(r.Form["q4"], "") == "inside" {
			food += 1
		} else if strings.Join(r.Form["q4"], "") == "outside" {
			sport += 1
			nature += 1
			animal += 1
		}

		fmt.Println("sport: ", sport, " nature: ", nature, " animal: ", animal, " food: ", food)
		if sport > nature && sport > animal && sport > food {
			fmt.Println("Max sport")
			http.Redirect(w, r, "/sport", http.StatusFound)
			// sport(w,r)
		} else if nature > sport && nature > food && nature > animal {
			fmt.Println("Max nature")
		} else if food > nature && food > animal && food > sport {
			fmt.Println("Max food")
		} else if animal > sport && animal > food && animal > nature {
			fmt.Println("Max animal")
		} else {
			fmt.Println("error")
		}
	}
	html := `<!DOCTYPE html>
	<html>
	<body>
		<p>Quiz Results Proccessing...</p>
	</body>
	</html>`

	w.Write([]byte(html))
}

func sport(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<body>

<p> sports </p>

</body>
</html>`
	w.Write([]byte(html))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", SimpleSelectTag)
	mux.HandleFunc("/quiz", quiz)
	mux.HandleFunc("/sport", sport)

	http.ListenAndServe(":8080", mux)
}