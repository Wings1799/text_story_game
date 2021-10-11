package main

import (
	"fmt"
	"strings"
)

var items []string

func main() {
	items = append(items, "phone", "ID")
	fmt.Println("Welcome to the Interactive Story!\n" +
		"In this game you will be confronted with choices, and the choices you make\n" +
		"will decide the outcome of your story.\n" +
		"To continue, please enter your name.")
	fmt.Print("Enter name here: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("\nHello ", name, ". It seems you have wandered into a busy city,\n"+
		"similar to New York. How do you figure out where you are? Do you:\n"+
		"A. Ask the nearest person?\n"+
		"B. Take out your phone?\n"+
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
	fmt.Println("\nITEM ADDED: knowledge\n" +
		"You decide to ask the nearest stranger, which is a busy-looking\n" +
		"man around his mid 30's. 'Where is this place?' you ask, hoping to get\n" +
		"some information. The man hurriedly replies, 'The Ground', before he\n" +
		"resumes his hurried pace down the busy street. Knowing this information, \n" +
		"do you:\n" +
		"A. Check your phone?\n" +
		"B. Try to rush after the man?\n" +
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
	fmt.Println("\nYou take out your phone and check the screen.\n" +
		"'No Service' it reads. What do you do?\n" +
		"A. Ask the nearest person?\n" +
		"C. Try to use your surroundings?")
	choice1()
}

func srdgs_choice() {
	fmt.Println("\nYou look around your immediate area and find many small shops\n" +
		"inside large skyscrapers. There are food vendors on the sidewalk, and \n" +
		"many people bustling around. Your immediate thought is that you are \n" +
		"in fact in New york, but the absence of famous landmarks of that city \n" +
		"tell you otherwise. While you're glancing around like a true tourist,\n" +
		"a young college-aged woman seems to be laughing at you from a bench \n" +
		"across the street. What do you decide to do?\n" +
		"A. Approach the woman?\n" +
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
	fmt.Println("\nYou take out your phone and check the screen.\n" +
		"'No Service' it reads. What do you do?\n" +
		"B. Try to rush after the man?\n" +
		"C. Begin to explore the city?")
	choice2()
}

func man_choice() {
	fmt.Println("\nYou run after the man, but a large crowd of people appears in\n" +
		"your way as they try to enter the bus that just stopped. Giving up on \n" +
		"catching the man, what is your next move?\n" +
		"A. Explore the city?\n" +
		"B. Try to get on the bus?")
	choice4()
}

func explore_choice() {
	for _, v := range items {
		if v == "knowledge" {
			fmt.Println("\nYou begin to explore 'The Ground', taking in as much\n" +
				"information as you can.")
			end3b()
		} else if v == "phone" {
			continue
		} else if v == "ID" {
			continue
		} else {
			fmt.Println("\nYou begin to explore the unknown city, taking in as much \n" +
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
	fmt.Println("\nYou attempt to get on the bus, but the driver stops you before\n" +
		"you can take more than 1 step up the stairs. 'Where's your ticket?' he \n" +
		"asks. You don't have a ticket, but you ask if you can ride anyways and pay \n" +
		"a different way. Do you:\n" +
		"A. Try to bribe the driver with an item?\n" +
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
	fmt.Println("\nYou see a building with a book symbol on the front. Thinking \n" +
		"it must be a library, you enter it. Immediately upon entering you are \n" +
		"grabbed and thrown to the ground by two men in suits wearing earpieces \n" +
		"and shades. 'This is restricted property. You're coming with us' said \n" +
		"one of the men as they dragged you down a long hallway. THE END.\n")
}

func end3a() {
	fmt.Println("\nYou see a sign that says at the top 'THE GROUND'. Underneath \n" +
		"the sign is a map, but not a regular map. This map is 3-dimensional, \n" +
		"and as you look at it, new places appear from the image you looked at.\n" +
		"You now know all the basics about this place, and can begin to delve \n" +
		"deeper into what's going on. THE END\n")
}

func bribe_choice() {
	for _, v := range items {
		if v == "image" {
			fmt.Println("\nYou give the image you found earlier to the bus driver.\n" +
				"He begins to cry and explains the story of how he's related to the \n" +
				"woman, but she passed away many years ago. It was his wife, and \n" +
				"she haunts the streets of 'The Ground'. You were very lucky to \n" +
				"have seen the woman as she only appears to those who have the \n" +
				"potential to bring happiness to the world. You hug the driver, \n" +
				"and he allows you onto the bus to ride until you want to  get off.")
			end1()
		} else if v == "phone" {
			continue
		} else if v == "ID" {
			continue
		} else if v == "image" {
			continue
		} else {
			fmt.Println("\nYou try to use your phone to bribe the bus driver, and \n" +
				"he promptly kicks you off the bus and tells you to come back \n" +
				"with a ticket. Do you:\n" +
				"A. Pay the driver?\n" +
				"B. Give up and explore the city?")
			choice9()
		}
	}
}

func pay_choice() {
	fmt.Println("\nYou try to pay the bus driver for a ticket, but he says that \n" +
		"tickets can not be purchased on the bus. Dejectedly, you get off and \n" +
		"try to find another way to get around town.")
	end2()
}

func end1() {
	fmt.Println("\nYou ride the bus for hours, waiting to hear a familiar stop. \n" +
		"After many more hours you realize that you haven't moved at all, \n" +
		"but rather you've been travelling through time while being in the same \n" +
		"position. You look outside and there is nothing there except wilderness.\n" +
		"The bus you're on starts to disappear, and not long after that you \n" +
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
	fmt.Println("\nITEM ADDED: image\n" +
		"You decide to approach the woman and begin walking across the\n" +
		"street as soon as you're able. She suddenly stops laughing and stares you\n" +
		"in the eyes with the most powerful glare you've ever experienced. You \n" +
		"can't hold eye contact that long so you look to your left to avoid her \n" +
		"gaze. When you look back up, she's gone. Where she was sitting is now a \n" +
		"photograph depicting two birds sitting on a branch. What do you do now?\n" +
		"A. Examine the photo again?\n" +
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
	fmt.Println("\nYou examine the photo of the birds on the branch. On the back \n" +
		"of the photo is a sketch of a bus, which draws your curiosity. Do you:\n" +
		"A. Search for a bus stop?\n" +
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
	fmt.Println("\nYou look around to try and find a tree where the birds from \n" +
		"the picture could have been sitting.\n")
	end4()
}

func end4() {
	fmt.Println("\nWhile you aren't paying attention, a Lamborghini Gallardo \n" +
		"comes screaming down the street. The car collides with you, and you \n" +
		"can finally see the tree where the birds were sitting. It's right \n" +
		"in front of you! You reach out and grab it. THE END.\n")
}

func end2() {
	fmt.Println("\nYou see a taxi fast approaching, you wave it down in hopes \n" +
		"it can give you a ride somewhere. The man driving looks at you and \n" +
		"tells you to get in. He then begins driving rapidly before you can \n" +
		"even tell him where you want to go. The man approaches a small shack,\n" +
		"where he lets you out. 'This is your new home' the man says. Handcuffs\n" +
		"appear on your wrists, and you're ushered into the basement. THE END\n")
}
