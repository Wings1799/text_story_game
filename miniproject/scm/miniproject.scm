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
        #(choice1)))

(define choice1
    (lambda()
        (display "Enter choice here: ")
        (define choice (read-line))

        #USE cond()
        (if (equal? (string-upcase choice) "A")((stranger_choice))
        (if (equal? (string-upcase choice) "B")((phone_choice))
        (if (equal? (string-upcase choice) "C")((srdgs_choice))
        ((display "INVALID ANSWER. PLEASE TRY AGAIN")(choice1)))))))

(main)