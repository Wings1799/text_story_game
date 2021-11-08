x :: Int
x = 4
-- this is a comment!
 {-this is
also a comment-}
biggestInt, smallestInt :: Int
biggestInt  = maxBound
smallestInt = minBound
--main = print biggestInt

y :: Double
y = 3.14
--main = print y

a :: Char
a = 'z'
--b = "x"
--main = print (a+b) --doesnt work

b :: Bool
b = False
--main = print b

mybool :: Bool
mybool = True || False
--main = print mybool

--main = print (5 < 3)

i, j :: Int
i = 4
j = 6
--if (i > j) then (i+j) else (i-j)

add3 :: Int -> Int
add3 n = n+3
--main = print (add3 7)

add3list :: [Int] -> [Int]
add3list [] = []
add3list (x:xs) = (x+3) : add3list xs
--main = print (add3list [1, 2, 3])

--___________________________________________________________
--Exercise 1
toDigits :: Integer -> [Integer]
toDigits 0 = []
toDigits (n) = (toDigits (n `div` 10) ++ [n `mod` 10])

toDigitsRev :: Integer -> [Integer]
toDigitsRev 0 = []
toDigitsRev (n) = n `mod` 10 : (toDigitsRev (n `div` 10))

--main = print (toDigits 1234)
--main = print (toDigitsRev 1234)

--Exercise 2
doubleEveryOther :: [Integer] -> [Integer]
doubleEveryOther [] = []
doubleEveryOther [n] = n : []
doubleEveryOther (n : m : ns) = n : m*2 : doubleEveryOther ns

--main = print (doubleEveryOther [8, 8, 3, 2, 1, 1])

--Exercise 3
sumDigits :: [Integer] -> Integer
sumDigits [n] = n
sumDigits (n : ns) = n + (sumDigits ns)

--main = print (sumDigits [1, 2, 3, 4])

--Exercise 4
isTens :: Integer -> Bool
isTens n = if (n `mod` 10 == 0) then (True) else (False)


validate :: Integer -> Bool
validate n = (isTens (sumDigits (doubleEveryOther (toDigitsRev n))))
