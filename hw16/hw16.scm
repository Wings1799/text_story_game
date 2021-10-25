(define list-copy
    (lambda (lis)
        (if (null? lis) '()
            (cons (car lis) (list-copy (cdr lis))))))

(define sub-sum
;returns the sum of the list, but subtracts the values instead of adding them
    (lambda (lis)
        (if (null? lis) 0
            (- (sub-sum (cdr lis)) (car lis)))))