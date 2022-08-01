package main

import (
	"fmt"
	"net/http"
	"strings"
)

func SimpleSelectTag(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head> 
<title>Code for Change</title>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css" integrity="sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2" crossorigin="anonymous">
<link rel="stylesheet" type="text/css" href="main.css">
	<style> 
	#page
	{
		border: 1px solid white;
		width: 1175px;
		height: 600px;
		padding: 20px;
		background-color: #fffff0;
	}

	#form1
		  {
			padding: 5px;
			font-size:120%;
			font-family:'Trebuchet MS', sans-serif;
			background-color: #f5f5dc;
		  }
		  #headr
		  {
			border: 4px solid navy;
			border-radius: 10px;
			box-shadow: 10px #C0C0C0;
			text-align: center;
			color:white;
			font-size:300%;
			font-family: Verdana, sans-serif;
			background-color: #000080;
		  }
		</style>
	</head>
	<script src="https://code.jquery.com/jquery-3.6.0.js" integrity="sha256-H+K7U5CnXl1h5ywQfKtSj8PCmoN9aaq30gDh27Xc0jk=" crossorigin="anonymous"></script>
	<script src="https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.min.js" integrity="sha384-w1Q4orYjBQndcko6MimVbzY0tgp4pWB4lZ7lr30WKz0vr/aWKhXdBNmNb5D92v7s" crossorigin="anonymous"></script>	
	   <body>
		  <div style="background-color=#00008b;">
			<nav class="navbar navbar-expand-md">
			<a class="navbar-brand" href="#">VM-Innovators</a>
			<button class="navbar-toggler navbar-dark" type="button" data-toggle="collapse" data-target="#main-navigation">
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="main-navigation">
				<ul class="navbar-nav">
					<li class="nav-item">
						<a class="nav-link" href="#">Quiz</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#">Charities</a>
					</li>
					<li class="nav-item">
						<a class="nav-link" href="#">About</a>
					</li>
				</ul>
			</div>
			</nav>
		  </div>
		  <header>
		  	<div id="headr">
		 		<p>Personality Quiz!</p> 
			</div> 
		  </header>
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

	<div id="page">
	  <div id="form1">
	    <form action="/quiz" method="post">
	    	<div id="question1" style="background-color: #4b8e23;">
		<p onclick="speak('What sounds the most fun to you?');">What sounds the most fun to you?</p>
		
		<select id="q1" name="q1" onclick = "speakChoices(1)" )>
		<option id = "petting_zoo" value="Going to the petting zoo">Going to a petting zoo
		</option>
		<option id = "soccer" value="Playing soccer with your friends">Playing soccer with your friends</option>
		<option id="hike" value = "Taking a hike">Taking a hike</option>
		<option id="eating_out" value="Eating out with friends">Eating out with friends</option>
		</select> <p><br></p> </div>
		<div id="question1" style="background-color: #5cab53;">
		<p onclick="speak('What would people describe you as?')">What would people describe you as?</p>
		<select id = "q2" name = "q2" onclick = "speakChoices(2)" >
		<option id="animal_lover" value = "Animal Lover" >Animal Lover
		</option>
		<option id="athlete" value = "athlete">Athlete</option>
		<option id="nature_lover" value="Nature Lover">Nature Lover</option>
		<option id="foodie" value="foodie">Foodie</option>
		</select> <p><br></p> </div>
		<div id="question1" style="background-color: #7fbc6f;">
		<p onclick="speak('What do you associate the color orange with?')">What do you associate the color orange with?
		</p>
		<select id = "q3" name = "q3" onclick = "speakChoices(3)" >
		<option id="fox" value="fox">Fox
		</option>
		<option id = "basketball" value="basketball">Basketball</option>
		<option id="leaves" value="fall leaves">Fall Leaves</option>
		<option id="fruit" value="Orange Fruit">Orange Fruit</option>
		</select> <p><br></p> </div>
		<div id="question1" style="background-color: #a4d4a3;">
		<p onclick="speak('Would you rather spend your day inside or outside?')">Would you rather spend your day inside or outside?
		</p>
		<select id = "q4" name="q4" onclick = "speakChoices(4)" >

		<option id="inside" value="inside">Inside</option>
		<option id = "outside" value="outside">Outside</option>
		</select> <p><br></p> </div>
		<p></p>
		<input onclick="speak('submit')" type = "submit" value = "submit">
	    </form>
	  </div>
	</div>
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
			http.Redirect(w, r, "/index", http.StatusFound)
			// sport(w,r)
		} else if nature > sport && nature > food && nature > animal {
			fmt.Println("Max nature")
			http.Redirect(w, r, "/index", http.StatusFound)
		} else if food > nature && food > animal && food > sport {
			fmt.Println("Max food")
			http.Redirect(w, r, "/index", http.StatusFound)
		} else if animal > sport && animal > food && animal > nature {
			fmt.Println("Max animal")
			http.Redirect(w, r, "/index", http.StatusFound)
		} else {
			fmt.Println("error")
			http.Redirect(w, r, "/index", http.StatusFound)
		}
	}
	
}



// func index(w http.ResponseWriter, r *http.Request) {
// 	html := index.html(r)

// }

// func tmp(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("here")
// }
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./start_page"))))
	mux.HandleFunc("/q", SimpleSelectTag)
	mux.HandleFunc("/quiz", quiz)

	mux.Handle("/index/", http.StripPrefix("/index/", http.FileServer(http.Dir("./static"))))

	mux.Handle("/output/", http.StripPrefix("/output/", http.FileServer(http.Dir("./charity"))))
	// charity_finder("")
	// mux.HandleFunc("/output", output)
	// mux.HandleFunc("/index", tmp)
	http.ListenAndServe(":8080", mux)

}
