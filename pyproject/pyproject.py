def main():
    items = ['phone', 'ID']
    print("Welcome to the Interactive Story!\nIn this game you will be confronted with choices, and the choices you make will decide the outcome of your story.\nTo continue, please enter your name.")
    name = input('Name here: ')
    print("\nHello ", name, ". It seems you have wandered into a busy city, similar to New York. How do you figure out where you are?\nDo you: \nA. Ask the nearest person?\nB. Take out your phone?\nC. Try to use your surroundings?")
    choice1()

def choice1():
    mychoice1 = input('Enter choice here: ')
    if mychoice1.capitalize()=='A':
        stranger_choice()
    elif mychoice1.capitalize()=='B':
        phone_choice()
    elif mychoice1.capitalize()=='C':
        srdgs_choice()
    else:
        print("INVALID ANSWER. PLEASE TRY AGAIN")
        choice1()

def stranger_choice():
    print("\nYou decide to ask the nearest stranger, which is a busy-looking man around his mid 30's. 'Where is this place?' you ask, hoping to get some information. \nThe man hurriedly replies, 'The Ground', before he resumes his hurried pace down the busy street. Knowing this information, do you: \nA. Check your phone?\nB. Try to rush after the man?\nC. Begin to explore the city?")
    choice2()

def choice2():
    mychoice2 = input('Enter choice here: ')
    if mychoice2.capitalize()=='A':
        phone2_choice()
    elif mychoice2.capitalize()=='B':
        man_choice()
    elif mychoice2.capitalize()=='C':
        explore_choice()
    else:
        print("INVALID ANSWER. PLEASE TRY AGAIN")
        choice2()

def phone2_choice():
    pass

def man_choice():
    pass

def explore_choice():
    pass

def phone_choice():
    pass
def srdgs_choice():
    pass

main()