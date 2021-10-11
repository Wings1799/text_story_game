package miniproject

import (
	"fmt"
	"strings"
)

var items []string

func main() {
	items = append(items, "phone", "ID")
	fmt.Println("Welcome to the Interactive Story!" +
		"In this game you will be confronted with choices, and the choices you make" +
		"will decide the outcome of your story." +
		"To continue, please enter your name.")
	fmt.Print("Enter name here: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("\nHello ", name, ". It seems you have wandered into a busy city,"+
		"similar to New York. How do you figure out where you are? Do you:"+
		"A. Ask the nearest person?"+
		"B. Take out your phone?"+
		"C. Try to use your surroundings?")
	choice1()
}
func choice1() {
	fmt.Printf("Enter choice here: ")
	var mychoice1 string
	fmt.Scanln(&mychoice1)
	if strings.ToUpper(mychoice1) == "A" {
		stranger_choice()
	} else if strings.ToUpper(mychoice1) == "B" {
		phone_choice()
	} else if strings.ToUpper(mychoice1) == "C" {
		srdgs_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice1()
	}
}

func stranger_choice() {
	items = append(items, "knowledge")
	fmt.Println("\nITEM ADDED: knowledge" +
		"You decide to ask the nearest stranger, which is a busy-looking" +
		"man around his mid 30's. 'Where is this place?' you ask, hoping to get" +
		"some information. The man hurriedly replies, 'The Ground', before he" +
		"resumes his hurried pace down the busy street. Knowing this information, " +
		"do you:" +
		"A. Check your phone?" +
		"B. Try to rush after the man?" +
		"C. Begin to explore the city?")
	choice2()
}

func choice2() {
	fmt.Printf("Enter choice here: ")
	var mychoice2 string
	fmt.Scanln(&mychoice2)
	if strings.ToUpper(mychoice2) == "A" {
		phone2_choice()
	} else if strings.ToUpper(mychoice2) == "B" {
		man_choice()
	} else if strings.ToUpper(mychoice2) == "C" {
		explore_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice2()
	}
}

func phone_choice() {
	fmt.Println("\nYou take out your phone and check the screen." +
		"'No Service' it reads. What do you do?" +
		"A. Ask the nearest person?" +
		"C. Try to use your surroundings?")
	choice1()
}

func srdgs_choice() {
	fmt.Println("\nYou look around your immediate area and find many small shops" +
		"inside large skyscrapers. There are food vendors on the sidewalk, and " +
		"many people bustling around. Your immediate thought is that you are " +
		"in fact in New york, but the absence of famous landmarks of that city " +
		"tell you otherwise. While you're glancing around like a true tourist," +
		"a young college-aged woman seems to be laughing at you from a bench " +
		"across the street. What do you decide to do?" +
		"A. Approach the woman?" +
		"B. Ignore the woman and begin to explore the city?")
	choice3()
}

func choice3() {
	fmt.Printf("Enter choice here: ")
	var mychoice3 string
	fmt.Scanln(&mychoice3)
	if strings.ToUpper(mychoice3) == "A" {
		approach_choice()
	} else if strings.ToUpper(mychoice3) == "B" {
		explore_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice3()
	}
}

func phone2_choice() {
	fmt.Println("\nYou take out your phone and check the screen." +
		"'No Service' it reads. What do you do?" +
		"B. Try to rush after the man?" +
		"C. Begin to explore the city?")
	choice2()
}

func man_choice() {
	fmt.Println("\nYou run after the man, but a large crowd of people appears in" +
		"your way as they try to enter the bus that just stopped. Giving up on " +
		"catching the man, what is your next move?" +
		"A. Explore the city?" +
		"B. Try to get on the bus?")
	choice4()
}

func explore_choice() {
	for _, v := range items {
		if v == "knowledge" {
			fmt.Println("\nYou begin to explore 'The Ground', taking in as much" +
				"information as you can.")
			end3b()
		} else if v == "phone" {
			continue
		} else if v == "ID" {
			continue
		} else {
			fmt.Println("\nYou begin to explore the unknown city, taking in as much " +
				"information as you can.")
			end3a()
		}
	}
}

func choice4() {
	fmt.Printf("Enter choice here: ")
	var mychoice4 string
	fmt.Scanln(&mychoice4)
	if strings.ToUpper(mychoice4) == "A" {
		explore_choice()
	} else if strings.ToUpper(mychoice4) == "B" {
		bus_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice4()
	}
}

func bus_choice() {
	fmt.Println("\nYou attempt to get on the bus, but the driver stops you before" +
		"you can take more than 1 step up the stairs. 'Where's your ticket?' he " +
		"asks. You don't have a ticket, but you ask if you can ride anyways and pay " +
		"a different way. Do you:" +
		"A. Try to bribe the driver with an item?" +
		"B. Pay the driver for a ticket?")
	choice6()
}

func choice6() {
	fmt.Printf("Enter choice here: ")
	var mychoice6 string
	fmt.Scanln(&mychoice6)
	if strings.ToUpper(mychoice6) == "A" {
		bribe_choice()
	} else if strings.ToUpper(mychoice6) == "B" {
		pay_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice6()
	}
}

func end3b() {
	fmt.Println("\nYou see a building with a book symbol on the front. Thinking " +
		"it must be a library, you enter it. Immediately upon entering you are " +
		"grabbed and thrown to the ground by two men in suits wearing earpieces " +
		"and shades. 'This is restricted property. You're coming with us' said " +
		"one of the men as they dragged you down a long hallway. THE END.")
}

func end3a() {
	fmt.Println("\nYou see a sign that says at the top 'THE GROUND'. Underneath " +
		"the sign is a map, but not a regular map. This map is 3-dimensional, " +
		"and as you look at it, new places appear from the image you looked at." +
		"You now know all the basics about this place, and can begin to delve " +
		"deeper into what's going on. THE END")
}

func bribe_choice() {
	for _, v := range items {
		if v == "image" {
			fmt.Println("\nYou give the image you found earlier to the bus driver." +
				"He begins to cry and explains the story of how he's related to the " +
				"woman, but she passed away many years ago. It was his wife, and " +
				"she haunts the streets of 'The Ground'. You were very lucky to " +
				"have seen the woman as she only appears to those who have the " +
				"potential to bring happiness to the world. You hug the driver, " +
				"and he allows you onto the bus to ride until you want to  get off.")
			end1()
		} else if v == "phone" {
			continue
		} else if v == "ID" {
			continue
		} else if v == "image" {
			continue
		} else {
			fmt.Println("\nYou try to use your phone to bribe the bus driver, and " +
				"he promptly kicks you off the bus and tells you to come back " +
				"with a ticket. Do you:" +
				"A. Pay the driver?" +
				"B. Give up and explore the city?")
			choice9()
		}
	}
}

func pay_choice() {
	fmt.Println("\nYou try to pay the bus driver for a ticket, but he says that " +
		"tickets can not be purchased on the bus. Dejectedly, you get off and " +
		"try to find another way to get around town.")
	end2()
}

func end1() {
	fmt.Println("\nYou ride the bus for hours, waiting to hear a familiar stop. " +
		"After many more hours you realize that you haven't moved at all, " +
		"but rather you've been travelling through time while being in the same " +
		"position. You look outside and there is nothing there except wilderness." +
		"The bus you're on starts to disappear, and not long after that you " +
		"disappear as well. Into the void. THE END.")
}

func choice9() {
	fmt.Printf("Enter choice here: ")
	var mychoice9 string
	fmt.Scanln(&mychoice9)
	if strings.ToUpper(mychoice9) == "A" {
		pay_choice()
	} else if strings.ToUpper(mychoice9) == "B" {
		explore_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice9()
	}
}

func approach_choice() {
	items = append(items, "image")
	fmt.Println("\nITEM ADDED: image" +
		"You decide to approach the woman and begin walking across the" +
		"street as soon as you're able. She suddenly stops laughing and stares you" +
		"in the eyes with the most powerful glare you've ever experienced. You " +
		"can't hold eye contact that long so you look to your left to avoid her " +
		"gaze. When you look back up, she's gone. Where she was sitting is now a " +
		"photograph depicting two birds sitting on a branch. What do you do now?" +
		"A. Examine the photo again?" +
		"B. Ask someone around you?")
	choice5()
}

func choice5() {
	fmt.Printf("Enter choice here: ")
	var mychoice5 string
	fmt.Scanln(&mychoice5)
	if strings.ToUpper(mychoice5) == "A" {
		photo_choice()
	} else if strings.ToUpper(mychoice5) == "B" {
		around_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice5()
	}
}

func photo_choice() {
	fmt.Println("\nYou examine the photo of the birds on the branch. On the back " +
		"of the photo is a sketch of a bus, which draws your curiosity. Do you:" +
		"A. Search for a bus stop?" +
		"B. Look around you for where the picture could've been taken?")
	choice7()
}

func choice7() {
	fmt.Printf("Enter choice here: ")
	var mychoice7 string
	fmt.Scanln(&mychoice7)
	if strings.ToUpper(mychoice7) == "A" {
		bus_choice()
	} else if strings.ToUpper(mychoice7) == "B" {
		around_choice()
	} else {
		fmt.Println("INVALID ANSWER. PLEASE TRY AGAIN")
		choice7()
	}
}

func around_choice() {
	fmt.Println("\nYou look around to try and find a tree where the birds from " +
		"the picture could have been sitting.")
	end4()
}

func end4() {
	fmt.Println("\nWhile you aren't paying attention, a Lamborghini Gallardo " +
		"comes screaming down the street. The car collides with you, and you " +
		"can finally see the tree where the birds were sitting. It's right " +
		"in front of you! You reach out and grab it. THE END.")
}

func end2() {
	fmt.Println("\nYou see a taxi fast approaching, you wave it down in hopes " +
		"it can give you a ride somewhere. The man driving looks at you and " +
		"tells you to get in. He then begins driving rapidly before you can " +
		"even tell him where you want to go. The man approaches a small shack," +
		"where he lets you out. 'This is your new home' the man says. Handcuffs" +
		"appear on your wrists, and you're ushered into the basement. THE END")
}
