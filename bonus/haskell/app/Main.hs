module Main where

------------------------
-- Imported Lib
import           System.Environment
import           System.Exit        (ExitCode (ExitFailure, ExitSuccess), exitWith)

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
    handleArgs (_, Options pSize sSize p) = putStrLn "good"
    exitHelp = getProgName >>= help >> exit
    exit = exitWith $ ExitSuccess
    exitErr = exitWith $ ExitFailure 84
