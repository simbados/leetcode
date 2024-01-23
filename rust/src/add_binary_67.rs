struct Solution { }

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        let mut answer = String::from("");
        let longest_arr;
        if a.len() > b.len() {
            longest_arr = a.len();
        } else {
            longest_arr = b.len();
        }
        let a_rev = a.chars().rev().collect::<String>();
        let b_rev = b.chars().rev().collect::<String>();
        let mut over = 0;
        for i in 0..longest_arr {
            let first;
            let second;
            if a_rev.chars().nth(i).is_some() {
                first= a_rev.chars().nth(i).unwrap().to_digit(10).unwrap();
            } else {
                first = 0
            }
            if b_rev.chars().nth(i).is_some() {
                second = b_rev.chars().nth(i).unwrap().to_digit(10).unwrap();
            } else {
                second = 0;
            }
            let curr = over + first + second;
            if curr >= 2 {
                over = 1;
            } else {
                over = 0;
            }
            answer = (curr % 2).to_string() + &answer;
        }
        if over == 1 {
            answer = 1.to_string() + &answer;
        }
        return answer;
    }
}

mod tests {
    use super::*;

    #[test]
    fn test_word_pattern() {
        assert_eq!(Solution::add_binary("1010".to_string(), "1011".to_string()), "10101");
        assert_eq!(Solution::add_binary("11".to_string(), "1".to_string()), "100");
    }
}
