module Main where

------------------------
-- Imported Lib
import           System.Environment
import           System.Exit        (ExitCode (ExitFailure), exitWith)

------------------------
-- Project
import           Lib

import           Args               (ArgsState (HELP, NOK), Options(..), checkArgs)
import           Help

main :: IO ()
main = getArgs >>= handleArgs . checkArgs
  where
    handleArgs (HELP, _)   = exitHelp
    handleArgs (NOK, _)    = exitErr
--    handleArgs (_, Options pSize sSize p) =
    exitHelp = help "" >> exit
    exit = exitWith $ ExitFailure 0
    exitErr = exitWith $ ExitFailure 84
