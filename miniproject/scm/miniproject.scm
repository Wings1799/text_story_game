#!/usr/local/dept/bin/mzscheme
#lang scheme/base

(define items (list 'phone 'ID))

(define main
    (lambda()
        (display "\tWelcome to the Interactive Story!
    In this game you will be confronted with choices, and the choices you make
    will decide the outcome of your story.
    To continue, please enter your name.\n")
        (display "Enter name here: ")
        (define name (read-line))
        (display "\n\tHello ") 
        (display name)
        (display ". It seems you have wandered into a busy city,
    similar to New York. How do you figure out where you are? Do you:
    A. Ask the nearest person?
    B. Take out your phone?
    C. Try to use your surroundings?\n")
        (choice1)))

(define choice1
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (stranger_choice)]
            [(equal? (string-upcase choice) "B")(phone_choice)]
            [(equal? (string-upcase choice) "C")(srdgs_choice)]
            [else ((display "INVALID ANSWER. PLEASE TRY AGAIN")(choice1))])))

(define stranger_choice
    (lambda()
        (cons "knowledge" items)
        (display "\nITEM ADDED: knowledge
    You decide to ask the nearest stranger, which is a busy-looking
    man around his mid 30's. 'Where is this place?' you ask, hoping to get 
    some information. The man hurriedly replies, 'The Ground', before he 
    resumes his hurried pace down the busy street. Knowing this information, 
    do you:
    A. Check your phone?
    B. Try to rush after the man?
    C. Begin to explore the city?")
    (choice2)))

(define choice2
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (phone2_choice)]
            [(equal? (string-upcase choice) "B")(man_choice)]
            [(equal? (string-upcase choice) "C")(explore_choice)]
            [else ((display "INVALID ANSWER. PLEASE TRY AGAIN")(choice2))])))

(define phone_choice
    (lambda()
        (display "\nYou take out your phone and check the screen.
        'No Service' it reads. What do you do?
        A. Ask the nearest person?
        C. Try to use your surroundings?")
        (choice1)))

(define phone2_choice
    (lambda()
        (display "\nYou take out your phone and check the screen.
        'No Service' it reads. What do you do?
        B. Try to rush after the man?
        C. Begin to explore the city?")
        (choice2)))

(define man_choice
    (lambda()
        (display "\nYou run after the man, but a large crowd of people appears in
        your way as they try to enter the bus that just stopped. Giving up on 
        catching the man, what is your next move?
        A. Explore the city?
        B. Try to get on the bus?")
        (choice4)))

(define choice4
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (explore_choice)]
            [(equal? (string-upcase choice) "B")(bus_choice)]
            [else ((display "INVALID ANSWER. PLEASE TRY AGAIN")(choice4))])))

(define explore_choice
    (lambda()
        (cond
            [(memq "knowledge" items) ((display "\nYou begin to explore 'The Ground', taking in as much 
            information as you can.")(end3b))]
            [else ((display "\nYou begin to explore the unknown city, taking in as much 
            information as you can.")(end3a))])))

(define bus_choice
    (lambda()
        (display "\nYou attempt to get on the bus, but the driver stops you before 
        you can take more than 1 step up the stairs. 'Where's your ticket?' he 
        asks. You don't have a ticket, but you ask if you can ride anyways and pay 
        a different way. Do you:
        A. Try to bribe the driver with an item?
        B. Pay the driver for a ticket?")
        (choice6)))

(define choice6
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))
        (cond
            [(equal? (string-upcase choice) "A") (bribe_choice)]
            [(equal? (string-upcase choice) "B")(pay_choice)]
            [else ((display "INVALID ANSWER. PLEASE TRY AGAIN")(choice6))])))

(define bribe_choice
    (lambda()
        (null)))

(define pay_choice
    (lambda()
        (null)))

(define srdgs_choice
    (lambda()
        (null)))

(define end3a
    (lambda()
        (null)))

(define end3b 
    (lambda()
        (null)))

(main)