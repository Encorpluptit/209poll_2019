module Args
    ( ArgsState(..)
    , Options(..)
    , checkArgs
    ) where

import           Data.Maybe
import           Text.Read

data ArgsState
    = HELP
    | OK
    | NOK
    deriving (Enum)

data Options =
    Options
        { pSize :: Int
        , sSize :: Int
        , p     :: Float
        }
    deriving (Eq)

--        }deriving(Eq, Show)
instance Show Options where
    show (Options _pSize _sSize _p) = foldl1 (\x y -> x ++ " " ++ y) [show _pSize, show _sSize, show _p]

--instance Eq Options where
--    Options pSizeX sSizeX pX == Options pSizeY sSizeY pY = pSizeX == pSizeY && sSizeX == sSizeY && pX == pY
--    Options pSizeX sSizeX pX /= Options pSizeY sSizeY pY = pSizeX /= pSizeY || sSizeX /= sSizeY || pX /= pY
wrongOptions :: Options
--wrongOptions = Options 1 (-1) (-1)
wrongOptions = Options {pSize = -1, sSize = -1, p = -1}

readInt :: String -> Int
readInt s
    | isNothing nb = -1
    | otherwise = fromJust nb
  where
    nb = readMaybe s :: Maybe Int

readFloat :: String -> Float
readFloat s
    | isNothing nb = -1
    | otherwise = fromJust nb
  where
    nb = readMaybe s :: Maybe Float

checkArgs :: [String] -> (ArgsState, Options)
checkArgs av
    | "--help" `elem` av = (HELP, wrongOptions)
    | "-h" `elem` av = (HELP, wrongOptions)
checkArgs [pSzArg, sSzArg, pArg]
    | _pSize < 0 = (NOK, wrongOptions)
    | _sSize < 0 = (NOK, wrongOptions)
    | _p < 0 || _p > 100 = (NOK, wrongOptions)
    | otherwise = (OK, Options _pSize _sSize _p)
  where
    _pSize = readInt pSzArg
    _sSize = readInt sSzArg
    _p = readFloat pArg
checkArgs _ = (NOK, wrongOptions)
