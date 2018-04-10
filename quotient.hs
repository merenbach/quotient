divide :: Integral a => a -> a -> [a]
divide a b = quot a b : divide (10 * rem a b) b
-- divide a b = fst result : divide (10 * snd result) b
--     where result = quotRem a b
