#lang racket/base

(define (divide a b scale)
  (let-values ([(q r) (quotient/remainder a b)])
	  (if (> scale 0)
		  (cons q (divide (* 10 r) b (- scale 1)))
		  (list q))
	  ))

(displayln (divide 17 7 150))
