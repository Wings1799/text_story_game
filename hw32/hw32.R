#!/usr/bin/Rscript --vanilla

is_4cyl <- function(x) {
    if (x==4) {
        return(TRUE)
    } else {
       return(FALSE)
    }
}

#lapply(mtcars, is_4cyl)
#lapply(mtcars, mean)

#lapply(mtcars, function(x) if(mean(x)<10){print(mean(x))} else{print("nah")})

add <- function(x) {
    function(y) {
        y+x
    }
}

plus4 <- add(4)
#plus4(6)

#Exercise 2
#lapply(mtcars, function(x) sd(x)/mean(x))

#Exercise 6
pick <- function(i) {
    function(x) {
        x[[i]]
    }
}

#lapply(mtcars, pick(5))