((lambda(n) (* n 3)) 4)

(define x3
    (lambda(n) 
        (* n 3)))
(x3 2)

(define mystring "This is my string\n")
(display mystring)

(define myfunction
    (lambda(x y) 
        (* x y)))

(define caddr
    (lambda(lis)
        (car (cdr (cdr lis)))))

(define mylist '(a b c d))

(define greaterthan
    (lambda(x y)
        (cond
            [(> x y) true]
            [(< x y) false]
            [(= x y) false])))