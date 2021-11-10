data MessageType = Info
                 | Warning
                 | Error Int
                 deriving (Show, Eq)

type TimeStamp = Int

data LogMessage = LogMessage MessageType TimeStamp String
                | Unknown String
                deriving (Show, Eq)

data MessageTree = Leaf
                 | Node MessageTree LogMessage MessageTree
                 deriving (Show, Eq)

--Exercise 1
parseMessage :: String -> LogMessage
parseMessage str = let lis = words str in 
    case lis of
        ("I":time:msg) -> LogMessage Info (read time) (unwords msg)
        ("W":time:msg) -> LogMessage Warning (read time) (unwords msg)
        ("E":code:time:msg) -> LogMessage (Error (read code)) (read time) (unwords msg)
        _ -> Unknown (unwords lis)

parse :: String -> [LogMessage]
parse str = map parseMessage (words str)

--Exercise 2
insert :: LogMessage -> MessageTree -> MessageTree
{-creates a new MessageTree with time of LogMessage in the middle and MessageTrees
on each side with lower on the left and greater on the right-}
insert ((LogMessage Unknown _) tree) = tree
insert ((LogMessage _ time _) (Node less (LogMessage _ time1 _) greater)) = 
    case time of
        (time > time1) -> --if time is greater than time1, use greater tree's time1 as new time1 and check again
        (time < time1) -> --if time is less than time1, use less tree's time1 as new time1 and check again
            {-In both cases, creates a new MessageTree with the lesser tree on the left, the original LogMessage in the 
            middle, and the greater tree on the right-}

--Exercise 3
build :: [LogMessage] -> MessageTree
build (log) = map insert log

--Exercise 4
inOrder :: MessageTree -> [LogMessage]
inOrder 