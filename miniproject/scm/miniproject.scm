#!/usr/local/dept/bin/mzscheme
#lang scheme/base

(define items (list 'phone 'ID))

(define main
    (lambda()
        (display "\n  Welcome to the Interactive Story!
    In this game you will be confronted with choices, and the choices you make
    will decide the outcome of your story. 
    To continue, please enter your name.
    (You can type 'items' to view current items.)\n")
        (display "Enter name here: ")
        (define name (read-line))
        (display "\n  Hello ") 
        (display name)
        (display ". It seems you have wandered into a busy city,
similar to New York. How do you figure out where you are? Do you:
    A. Ask the nearest person?
    B. Take out your phone?
    C. Try to use your surroundings?
    (Type 'items' to view current items.)\n")
        (choice1)))

(define choice1
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (stranger_choice)]
            [(equal? (string-upcase choice) "B")(phone_choice)]
            [(equal? (string-upcase choice) "C")(srdgs_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice1))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice1))])))

(define stranger_choice
    (lambda()
        (set! items(cons "knowledge" items))
        (display "\nITEM ADDED: knowledge
    You decide to ask the nearest stranger, which is a busy-looking
man around his mid 30's. 'Where is this place?' you ask, hoping to get 
some information. The man hurriedly replies, 'The Ground', before he 
resumes his hurried pace down the busy street. Knowing this information, 
do you:
    A. Check your phone?
    B. Try to rush after the man?
    C. Begin to explore the city?
    (Type 'items' to view current items.)\n")
    (choice2)))

(define choice2
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (phone2_choice)]
            [(equal? (string-upcase choice) "B")(man_choice)]
            [(equal? (string-upcase choice) "C")(explore_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice2))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice2))])))

(define phone_choice
    (lambda()
        (display "\n  You take out your phone and check the screen.
'No Service' it reads. What do you do?
    A. Ask the nearest person?
    C. Try to use your surroundings?
    (Type 'items' to view current items.)\n")
        (choice1)))

(define phone2_choice
    (lambda()
        (display "\n  You take out your phone and check the screen.
'No Service' it reads. What do you do?
    B. Try to rush after the man?
    C. Begin to explore the city?
    (Type 'items' to view current items.)\n")
        (choice2)))

(define man_choice
    (lambda()
        (display "\n  You run after the man, but a large crowd of people appears in
your way as they try to enter the bus that just stopped. Giving up on 
catching the man, what is your next move?
    A. Explore the city?
    B. Try to get on the bus?
    (Type 'items' to view current items.)\n")
        (choice4)))

(define choice4
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (explore_choice)]
            [(equal? (string-upcase choice) "B")(bus_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice4))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice4))])))

(define explore_choice
    (lambda()
        (cond
            [(memq "knowledge" items) (begin (display "\n  You begin to explore 'The Ground', taking in as much 
information as you can.")(end3b))]
            [else (begin (display "\n  You begin to explore the unknown city, taking in as much 
information as you can.")(end3a))])))

(define bus_choice
    (lambda()
        (display "\n  You attempt to get on the bus, but the driver stops you before 
you can take more than 1 step up the stairs. 'Where's your ticket?' he 
asks. You don't have a ticket, but you ask if you can ride anyways and pay 
a different way. Do you:
    A. Try to bribe the driver with an item?
    B. Pay the driver for a ticket?
    (Type 'items' to view current items.)\n")
        (choice6)))

(define choice6
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (bribe_choice)]
            [(equal? (string-upcase choice) "B")(pay_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice6))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice6))])))

(define bribe_choice
    (lambda()
        (cond 
            [(memq "image" items) (begin (display "\n  You give the image you found earlier to the bus driver. 
He begins to cry and explains the story of how he's related to the 
woman, but she passed away many years ago. It was his wife, and  
she haunts the streets of 'The Ground'. You were very lucky to 
have seen the woman as she only appears to those who have the 
potential to bring happiness to the world. You hug the driver, 
and he allows you onto the bus to ride until you want to get off.\n")(end1))]
            [else (begin (display "\n  You try to use your phone to bribe the bus driver, and 
he promptly kicks you off the bus and tells you to come back 
with a ticket. Do you:
    A. Pay the driver?
    B. Give up and explore the city?
    (Type 'items' to view current items.)\n")(choice9))])))

(define choice9
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (pay_choice)]
            [(equal? (string-upcase choice) "B")(explore_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice9))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice9))])))

(define pay_choice
    (lambda()
        (display "\n  You try to pay the bus driver for a ticket, but he says that 
tickets can not be purchased on the bus. Dejectedly, you get off and 
try to find another way to get around town.\n")
        (end2)))

(define srdgs_choice
    (lambda()
        (begin (display "\n  You look around your immediate area and find many small shops
inside large skyscrapers. There are food vendors on the sidewalk, and 
many people bustling around. Your immediate thought is that you are 
in fact in New york, but the absence of famous landmarks of that city 
tell you otherwise. While you're glancing around like a true tourist,
a young college-aged woman seems to be laughing at you from a bench 
across the street. What do you decide to do?
    A. Approach the woman?
    B. Ignore the woman and begin to explore the city?
    (Type 'items' to view current items.)\n")
            (choice3))))

(define choice3
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
            (cond
                [(equal? (string-upcase choice) "A") (approach_choice)]
                [(equal? (string-upcase choice) "B")(explore_choice)]
                [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice3))]
                [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice3))])))

(define approach_choice
    (lambda()
        (set! items (cons "image" items))
        (display "\nITEM ADDED: image
    You decide to approach the woman and begin walking across the
street as soon as you're able. She suddenly stops laughing and stares you
in the eyes with the most powerful glare you've ever experienced. You 
can't hold eye contact that long so you look to your left to avoid her 
gaze. When you look back up, she's gone. Where she was sitting is now a 
photograph depicting two birds sitting on a branch. What do you do now?
    A. Examine the photo again?
    B. Ask someone around you?
    (Type 'items' to view current items.)\n")
        (choice5)))

(define choice5
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (photo_choice)]
            [(equal? (string-upcase choice) "B")(around_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice5))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice5))])))

(define photo_choice
    (lambda()
        (display "\n  You examine the photo of the birds on the branch. On the back 
of the photo is a sketch of a bus, which draws your curiosity. Do you:
    A. Search for a bus stop?
    B. Look around you for where the picture could've been taken?
    (Type 'items' to view current items.)\n")
        (choice7)))

(define choice7
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (bus_choice)]
            [(equal? (string-upcase choice) "B")(around_choice)]
            [(equal? (string-upcase choice) "ITEMS")(begin (display items)(display "\n")(choice7))]
            [else (begin (display "INVALID ANSWER. PLEASE TRY AGAIN\n")(choice7))])))

(define around_choice
    (lambda()
        (display "\n  You look around to try and find a tree where the birds from 
the picture could have been sitting.\n")
        (end4)))

(define end1
    (lambda() 
        (display "\n  You ride the bus for hours, waiting to hear a familiar stop. 
After many more hours you realize that you haven't moved at all, 
but rather you've been travelling through time while being in the same 
position. You look outside and there is nothing there except wilderness.
The bus you're on starts to disappear, and not long after that you 
disappear as well. Into the void. \nTHE END.\n")))

(define end2
    (lambda()
        (display "\n  You see a taxi fast approaching, you wave it down in hopes 
it can give you a ride somewhere. The man driving looks at you and 
tells you to get in. He then begins driving rapidly before you can 
even tell him where you want to go. The man approaches a small shack,
where he lets you out. 'This is your new home' the man says. Handcuffs
appear on your wrists, and you're ushered into the basement. \nTO BE CONTINUED\n")))

(define end3a
    (lambda()
        (display "\n  You see a sign that says at the top 'THE GROUND'. Underneath 
the sign is a map, but not a regular map. This map is 3-dimensional, 
and as you look at it, new places appear from the image you looked at.
You now know all the basics about this place, and can begin to delve 
deeper into what's going on. \nTO BE CONTINUED.\n")))

(define end3b 
    (lambda()
        (display "\n  You see a building with a book symbol on the front. Thinking 
it must be a library, you enter it. Immediately upon entering you are 
grabbed and thrown to the ground by two men in suits wearing earpieces 
and shades. 'This is restricted property. You're coming with us' said 
one of the men as they dragged you down a long hallway. \nTO BE CONTINUED.\n")))

(define end4
    (lambda()
        (display "\n  While you aren't paying attention, a Lamborghini Gallardo 
comes screaming down the street. The car collides with you, and you 
can finally see the tree where the birds were sitting. It's right 
in front of you! You reach out and grab it. You cease to exist. \nTHE END.\n")))

(main)