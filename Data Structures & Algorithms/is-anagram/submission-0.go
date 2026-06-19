func isAnagram(s string, t string) bool {
        if len(s) != len(t) {
                return false
        }

        count := make(map[rune]int)
        for _, ch := range s {
                count[ch]++ // буквы из s прибавляем
        }
        for _, ch := range t {
                count[ch]-- // буквы из t вычитаем
        }

        for _, c := range count {
                if c != 0 {
                        return false
                }
        }
        return true
  }