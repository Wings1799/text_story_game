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
    C. Try to use your surroundings?")
        #(choice1)))

(main)