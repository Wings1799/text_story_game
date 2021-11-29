data Team = Team string int

instance Monad Team where
    return = Team

    Team string _ >>= k = k string