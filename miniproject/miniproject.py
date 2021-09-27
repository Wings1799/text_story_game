class Story:
    def __init__(self):
        self.main()

    def main(self):
        self.items = ['phone', 'ID']
        print("""Welcome to the Interactive Story!
        In this game you will be confronted with choices, and the choices you make
        will decide the outcome of your story.
        To continue, please enter your name.""")
        name = input('Name here: ')
        print("\nHello ", name, """. It seems you have wandered into a busy city,
        similar to New York. How do you figure out where you are? Do you:
        A. Ask the nearest person?
        B. Take out your phone?
        C. Try to use your surroundings?""")
        self.choice1()

    def choice1(self):
        mychoice1 = input('Enter choice here: ')
        if mychoice1.capitalize()=='A':
            self.stranger_choice()
        elif mychoice1.capitalize()=='B':
            self.phone_choice()
        elif mychoice1.capitalize()=='C':
            self.srdgs_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice1()

    def stranger_choice(self):
        self.items.append('knowledge')
        print("""\nITEM ADDED: knowledge
        You decide to ask the nearest stranger, which is a busy-looking
        man around his mid 30's. 'Where is this place?' you ask, hoping to get 
        some information. The man hurriedly replies, 'The Ground', before he 
        resumes his hurried pace down the busy street. Knowing this information, 
        do you:
        A. Check your phone?
        B. Try to rush after the man?
        C. Begin to explore the city?""")
        self.choice2()

    def choice2(self):
        mychoice2 = input('Enter choice here: ')
        if mychoice2.capitalize()=='A':
            self.phone2_choice()
        elif mychoice2.capitalize()=='B':
            self.man_choice()
        elif mychoice2.capitalize()=='C':
            self.explore_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice2()

    def phone_choice(self):
        print("""\nYou take out your phone and check the screen.
        'No Service' it reads. What do you do?
        A. Ask the nearest person?
        C. Try to use your surroundings?""")
        self.choice1()

    def phone2_choice(self):
        print("""\nYou take out your phone and check the screen.
        'No Service' it reads. What do you do?
        B. Try to rush after the man?
        C. Begin to explore the city?""")
        self.choice2()

    def man_choice(self):
        print("""\nYou run after the man, but a large crowd of people appears in
        your way as they try to enter the bus that just stopped. Giving up on 
        catching the man, what is your next move?
        A. Explore the city?
        B. Try to get on the bus?""")
        self.choice4()

    def choice4(self):
        mychoice4 = input("Enter choice here: ")
        if mychoice4.capitalize() == 'A':
            self.explore_choice()
        elif mychoice4.capitalize() == 'B':
            self.bus_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice4()
        
    def explore_choice(self):
        if 'knowledge' in self.items:
            print("""\nYou begin to explore 'The Ground', taking in as much 
            information as you can.""")
            self.end3b()
        else:
            print("""\nYou begin to explore the unknown city, taking in as much 
            information as you can.""")
            self.end3a()

    def bus_choice(self):
        print("""\nYou attempt to get on the bus, but the driver stops you before 
        you can take more than 1 step up the stairs. 'Where's your ticket?' he 
        asks. You don't have a ticket, but you ask if you can ride anyways and pay 
        a different way. Do you:
        A. Try to bribe the driver with an item?
        B. Pay the driver for a ticket?""")
        self.choice6()

    def choice6(self):
        mychoice6 = input("Enter choice here: ")
        if mychoice6.capitalize() == 'A':
            self.bribe_choice()
        elif mychoice6.capitalize() == 'B':
            self.pay_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice6()

    def bribe_choice(self):
        if 'image' in self.items:
            print("""\nYou give the image you found earlier to the bus driver. 
            He begins to cry and explains the story of how he's related to the 
            woman, but she passed away many years ago. It was his wife, and 
            she haunts the streets of 'The Ground'. You were very lucky to 
            have seen the woman as she only appears to those who have the 
            potential to bring happiness to the world. You hug the driver, 
            and he allows you onto the bus to ride until you want to  get off.""")
            self.end1()
        else:
            print("""\nYou try to use your phone to bribe the bus driver, and 
            he promptly kicks you off the bus and tells you to come back 
            with a ticket. Do you:
            A. Pay the driver?
            B. Give up and explore the city?""")
            self.choice9()

    def choice9(self):
        mychoice9 = input("Enter choice here: ")
        if mychoice9.capitalize() == 'A':
            self.pay_choice()
        elif mychoice9.capitalize() == 'B':
            self.explore_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice9()

    def end1(self):
        print("""\nYou ride the bus for hours, waiting to hear a familiar stop. 
        After many more hours you realize that you haven't moved at all, 
        but rather you've been travelling through time while being in the same 
        position. You look outside and there is nothing there except wilderness.
        The bus you're on starts to disappear, and not long after that you 
        disappear as well. Into the void. THE END.""")
        return 0

    def srdgs_choice(self):
        print("""\nYou look around your immediate area and find many small shops
            inside large skyscrapers. There are food vendors on the sidewalk, and 
            many people bustling around. Your immediate thought is that you are 
            in fact in New york, but the absence of famous landmarks of that city 
            tell you otherwise. While you're glancing around like a true tourist,
            a young college-aged woman seems to be laughing at you from a bench 
            across the street. What do you decide to do?
            A. Approach the woman?
            B. Ignore the woman and begin to explore the city?""")
        self.choice3()

    def choice3(self):
        mychoice3 = input("Enter choice here: ")
        if mychoice3.capitalize() == 'A':
            self.approach_choice()
        elif mychoice3.capitalize() == 'B':
            self.explore_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice3()

    def approach_choice(self):
        self.items.append('image')
        print("""\nITEM ADDED: image
        You decide to approach the woman and begin walking across the
        street as soon as you're able. She suddenly stops laughing and stares you
        in the eyes with the most powerful glare you've ever experienced. You 
        can't hold eye contact that long so you look to your left to avoid her 
        gaze. When you look back up, she's gone. Where she was sitting is now a 
        photograph depicting two birds sitting on a branch. What do you do now?
        A. Examine the photo again?
        B. Ask someone around you?""")
        self.choice5()

    def choice5(self):
        mychoice5 = input("Enter choice here: ")
        if mychoice5.capitalize() == 'A':
            self.photo_choice()
        elif mychoice5.capitalize() == 'B':
            self.around_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice5()

    def photo_choice(self):
        print("""\nYou examine the photo of the birds on the branch. On the back 
        of the photo is a sketch of a bus, which draws your curiosity. Do you:
        A. Search for a bus stop?
        B. Look around you for where the picture could've been taken?""")
        self.choice7()

    def choice7(self):
        mychoice7 = input("Enter choice here: ")
        if mychoice7.capitalize() == 'A':
            self.bus_choice()
        elif mychoice7.capitalize() == 'B':
            self.around_choice()
        else:
            print("INVALID ANSWER. PLEASE TRY AGAIN")
            self.choice7()

    def around_choice(self):
        print("""\nYou look around to try and find a tree where the birds from 
        the picture could have been sitting.""")
        self.end4()

    def pay_choice(self):
        print("""\nYou try to pay the bus driver for a ticket, but he says that 
        tickets can not be purchased on the bus. Dejectedly, you get off and 
        try to find another way to get around town.""")
        self.end2()

    def end2(self):
        print("""\nYou see a taxi fast approaching, you wave it down in hopes 
        it can give you a ride somewhere. The man driving looks at you and 
        tells you to get in. He then begins driving rapidly before you can 
        even tell him where you want to go. The man approaches a small shack,
        where he lets you out. 'This is your new home' the man says. Handcuffs
        appear on your wrists, and you're ushered into the basement. THE END""")
        return 0

    def end3a(self):
        print("""\nYou see a sign that says at the top 'THE GROUND'. Underneath 
        the sign is a map, but not a regular map. This map is 3-dimensional, 
        and as you look at it, new places appear from the image you looked at.
        You now know all the basics about this place, and can begin to delve 
        deeper into what's going on. THE END.""")
        return 0

    def end3b(self):
        print("""\nYou see a building with a book symbol on the front. Thinking 
        it must be a library, you enter it. Immediately upon entering you are 
        grabbed and thrown to the ground by two men in suits wearing earpieces 
        and shades. 'This is restricted property. You're coming with us' said 
        one of the men as they dragged you down a long hallway. THE END.""")
        return 0

    def end4(self):
        print("""\nWhile you aren't paying attention, a Lamborghini Gallardo 
        comes screaming down the street. The car collides with you, and you 
        can finally see the tree where the birds were sitting. It's right 
        in front of you! You reach out and grab it. THE END.""")
        return 0

s = Story()