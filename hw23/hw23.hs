data Fruit = Apple
           | Banana
           | Grape
           | Watermelon
  deriving Show

w1 :: Fruit
w1 = Watermelon
--main = print w1

color :: Fruit -> String
color Apple = "red"
color Banana = "yellow"
color Grape = "purple"
color Watermelon = "green"

--main = print (color w1)

data Point = Point Int Int
  deriving Show

p1 :: Point
p1 = Point 1 2
--main = print p1

slope :: Point -> Int
slope (Point x y) = y `div` x

--main = print (slope p1)

data Box = Parent Char Box
         | Child Char
  deriving Show

box :: Box
box = Parent 'A' (Parent 'B' (Child 'C' ))
--main = print box