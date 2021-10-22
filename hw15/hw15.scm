;Exercise 2.1
(define power
    (lambda(base exponent)
        (if (= exponent 0) 1 (* base (power base (- exponent 1))))))

;Exercise 2.3
(define square ; another version that doesn’t work
    (lambda (n)
        (if (= n 0) 0 (+ (square (- n 2)) (- (* 4 n) 4)))))
;The problem with this definition is that n will never reach 0 unless it is an even integer. That's 
;why all the tests that use even integers return the correct result. In the nested call of (square),
;n is being subtracted by 2 everytime, and unless n is even it will never reach 0

;Exercise 2.4
(define square
    (lambda (n)
        (if (= n 0) 0 
            (if (even? n)(* (square (/ n 2)) 4)
                (+ (square (- n 1))(- (+ n n) 1))))))

;Exercise 2.5
(define multiply
    (lambda (x y)
        (if (= y 0) 0
            (+ x (multiply x (- y 1))))))

;Exercise 2.17
(define presents-on-day
    (lambda (n)
        (if (= n 1) 1
            (+ n (presents-on-day (- n 1))))))

(define presents-through-day
    (lambda (n)
        (if (= n 1) 1
            (+ (presents-on-day n) (presents-on-day (- n 1))))))