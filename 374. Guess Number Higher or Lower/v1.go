/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

 func guessNumber(n int) int {
    l := 1
    r := n

    for  l <= r {
        c := ((r-l)/2) + l

        v := guess(c)
        
        switch v {
        case 0: // num == pick
            return c
        case -1: // num > pick
            r = c - 1
        case 1: // num < pick
            l = c + 1
        }
    }

    return -1
}