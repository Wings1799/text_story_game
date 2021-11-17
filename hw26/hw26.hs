import Parser

data Frog = Frog String Frog Frog
          | Kid String

instance Eq Frog where
    (Frog name kid1 kid2) == (Frog name_ kid1_ kid2_) = name == name_ 
    (Kid name1) == (Kid name1_) = name1 == name1_
    _ == _ = False

frog1 = Frog "James" (Kid "John") (Kid "Jane")
frog2 = Frog "Jamie" (Kid "John") (Kid "Jane")

--main = print (frog1 == frog2)

--____________________________________________________________

--I think the italicized text highlights the importance of commenting in large projects.
--Also since you have to be so specific with Haskell, it's very hard for it to do things that you didn't define it to do in the first place.

--____________________________________________________________
--Exercise 1
data ExprT = Lit Integer
           | Add ExprT ExprT
           | Mul ExprT ExprT
    deriving (Show, Eq)

eval :: ExprT -> Integer
eval (Lit x) = x
eval (Add x y) = (eval x) + (eval y)
eval (Mul x y) = (eval x) * (eval y)

--main = print (eval (Mul (Add (Lit 2) (Lit 3)) (Lit 4)))

--Exercise 2
evalStr :: String -> Maybe Integer
evalStr str = 
    case (parseExp str) of
        --Pretty sure I'm using pattern matching wrong here, not sure what to do
        (Nothing) -> Nothing
        (Just n) -> Just (eval n)

--main = print ( parseExp Lit Add Mul "(2+3)*4")

--Exercise 3
class Expr a where
    lit :: Integer -> a
    add :: a -> a -> a
    mul :: a -> a -> a

instance Expr ExprT where
    --lit :: Integer -> a
    lit x = Lit x
    --add :: a -> a -> a
    add (Lit a) (Lit b) = Add (Lit a) (Lit b)
    --mul :: a -> a -> a
    mul (Lit a) (Lit b) = Mul (Lit a) (Lit b)