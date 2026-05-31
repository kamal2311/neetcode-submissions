func isPalindrome(s string) bool {

    for i, j := 0, len(s) -1; i < j; i, j = i+1, j-1{

        for i < j && !isAlphanumeric(s[i]){
            i++
        }

        for i < j && !isAlphanumeric(s[j]){
            j--
        }  

        if toLower(s[i]) != toLower(s[j]) {
            return false
        }
    }

    return true

}

func isAlphanumeric(b byte) bool {
    return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLower(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        b += 'a'- 'A' 
    }
    return b
}