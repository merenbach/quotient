import qualified System.Environment as Environment

divide :: Integral a => a -> a -> [a]
divide a b = quot a b : divide (10 * rem a b) b

-- divide a b = fst result : divide (10 * snd result) b
--     where result = quotRem a b
--     where result = divMod a b
divideToString :: Integer -> Integer -> String
divideToString a b =
  show (head result) ++ "." ++ concat [show n | n <- tail result]
  where
    result = divide a b

main = do
  args <- Environment.getArgs
  let intArgs = [read x :: Integer | x <- args]
  putStrLn $ divideToString (intArgs !! 0) (intArgs !! 1)
