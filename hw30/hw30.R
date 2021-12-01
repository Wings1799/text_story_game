#!/usr/bin/Rscript --vanilla

x <- 4
y <- 3
#x*y
#print(str('hello world!'))
x <- c(1L, 2, 3L)
#print(x)
#x <- c(NA, x)
#print(x)
x <- c(TRUE, 8L, x)
#print(x)
#print(x[2:4])
#print(sum(x))
#print(mean(x))

y <- c("5", "6", "7")
#print(y)
#print(sum(y))
#print(sum(as.numeric(y)))

z <- list("y", 32, 1L, TRUE)
#print(z)
xz <- c(x, z)
#print(xz) #xz is a list, not a vector
x <- c(FALSE, NA) 
#print(x)

#Exercises
#1. logical, integer, double, character, complex, and raw. Lists don't care what type is inside it.
#2. is.vector() checks if the object uses c(), and is.list() checks if it uses list().
#is.numeric() checks if the elements are all numeric, while is.character() checks if they are characters.
#3. [2] 1.0 1.0
#   [2] "a" "1"
#   [[1]]
#       [1] 1.0
#       [1] "a"
#   [2] 1 1
#4. the elements of a list are just single vectors, so they must be turned into that form first
#5. Because the types are coerced by the operators, but "one" can't be coerced
#6. it's a logical vector because they are either 1 or 0, which can be converted to all types.

attr(x, "name") <- "Test List"
#print(attr(x, "name"))
#print(attributes(x))
names(x) <- c("Test List")
#print(names(x))

df <- data.frame(x = 10:18, y = c("a", "b", "c"), stringsAsFactors = FALSE)
#print(str(df))
#print(df)
#print(as.matrix(df))

df1 <- data.frame()
print(df1)

#Exercises
#1. data frames have names(), colnames(), rownames(), ncol(), and nrow()
#2. it tries to coerce the types first
#3. yes you can have a data frame with 0 rows and 0 columns, or both altogether.