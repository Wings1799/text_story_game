#!/usr/bin/Rscript --vanilla

#?print

x <- c(2.1, 4.2, 3.3, 5.4)
y <- c(TRUE, FALSE, TRUE, TRUE)
#print(x[1:2])
#print(x[c(4,2)])
#print(x[order(x)])
#print(x[-c(-1,-2)])
#print(x[y==TRUE]) #this is crazy right here
#print(x[])
z <- setNames(x, letters[1:4])
#print(z[c("d")])

a <- matrix(9:1, nrow = 3)
colnames(a) <- c("A", "B", "C")
#print(a[])
#print(a[c(2:3),])

df <- data.frame(test = letters[1:8], result = c(TRUE, FALSE), date = c("Sunday", "Monday"))
print(df$result)
#print(df[df$result == TRUE, ])
#print(?mtcars)

#Exercise 1
#mtcars[mtcars$cyl == 4, ]
#mtcars[1:4, ]
#mtcars[mtcars$cyl <= 5, ]
#mtcars[mtcars$cyl == c(4,6), ]

#Exercise 4
#Because 1:20 is not defined to be columns in the data frame, so you must transform it to this form
#It's different because it specifies the rows as well

#Exercise 1
#You must use the sample function to randomly change aspects of a data frame. You can not change both rows and columns at once.

#Exercise 2
#You also use sample function here and then specify how many rows you want for the sample. If it's contiguous then you can do that too.

#Exercise 3
#You use the order function which would order them alphabetically