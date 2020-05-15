module Help
    ( help
    ) where

help :: String -> IO ()
help pgrm =
    putStrLn "Usage\n" >> putStrLn ("\t" ++ pgrm ++ " pSize sSize p\n\n") >> putStrLn "DESCRIPTION:\n" >>
    putStrLn "\tpSize\tsize of the population\n" >>
    putStrLn "\tsSize\tsize of the sample (supposed to be representative)\n" >>
    putStrLn "\tp\tpercentage of voting intentions for a specific candidate"
