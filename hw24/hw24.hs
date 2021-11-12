mylist :: [Int]
mylist = [1, 2, 3, 4]

mult2 :: Int -> Int
mult2 n = 2*n

--main = print (map mult2 mylist)

data List t = Emptyy | Conss t (List t)
    deriving Show
lst1 :: List Char
lst1 = Conss 'a' (Conss 'b' (Conss 'c' Emptyy))
--main = print lst1
lst2 :: List Int
lst2 = Conss 1 (Conss 2 (Conss 3 Emptyy))
--main = print lst2

safeDiv :: Int -> Int -> Int
safeDiv (_, 0) = Nothing
safeDiv (x, y) = Just (x / y)

--main = print (safeDiv 2 1)

--_____________________________________________________
--Exercise 1
skips :: [a] -> [[a]]
skips [] -> [[]]
skips (x:xs) = 

--Exercise 2
localMaxima :: [Integer] -> [Integer]
localMaxima [] = []
localMaxima 