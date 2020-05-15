# Haskell Project

## IntelliJ (repository creation)

 - lol
 - mdr

## Repository setup.
 - Copy Makefile
 - Modify variable for Project name and Binary Name in Makefile
 - Copy **Profiling** folder into **test** folder


## Package

 - change name for "executables" and tests in package.yaml

## Stack
change resolver (WARNING: break IntelliJ):
```
resolver: lts-14.11
```

## Cabal

 - add Library packages to exposed Modules:
```
  exposed-modules:
      Lib
      Args
      Help
```

 - change name of executable and test in $(project).cabal:
```
executable poll
  main-is: Main.hs
  other-modules:
      Paths_209poll
...

test-suite poll-test
  type: exitcode-stdio-1.0
  main-is: Spec.hs
...
```

 - change cabal version to 2.2 
```
cabal-version: 2.2
```
 - change license type to BSD-3-Clause
```
license:        BSD-3-Clause
```

 - add the following default options:
```
common default
    default-language:    Haskell2010
    build-depends:       base >= 4.7 && < 5
    ghc-options:         -Wall
```
 - remove defaulted options