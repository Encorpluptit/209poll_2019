import Test.HUnit

import Args(checkArgs)



--module Tests.Args(argsSpec) where
--
--import Test.HUnit
--import Args
--
isOk :: Test
isOk = TestCase (assertEqual name expected out)
    where
        name = "CheckArgs: OK"
        expected = fromEnum OK
        out  = fromEnum status
        (status, _) = checkArgs ["2", "0.1", "tests/exemple"]

argsSpec :: Test
argsSpec = TestList []

main :: IO ()
main = putStrLn "Test suite not yet implemented"
