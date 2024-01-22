use std::collections::HashMap;

pub fn word_pattern(pattern: String, s: String) -> bool {
    let mut p_map = HashMap::new();
    let mut s_map = HashMap::new();
    let s_arr: Vec<&str> = s.split(" ").collect();
    if s_arr.len() != pattern.len() {
        return false;
    }
    for i in 0..pattern.len() {
        let p_val = pattern.chars().nth(i).unwrap();
        let s_val = s_arr.get(i);
        let p_map_val = p_map.get(&p_val);
        if p_map_val.is_some() {
            if *p_map_val.unwrap() != s_val {
                return false;
            }
        } else {
            p_map.insert(p_val, s_val);
        }
        let s_map_val = s_map.get(&s_val);
        if s_map_val.is_some() {
            if *s_map_val.unwrap() != p_val {
                return false;
            }
        } else {
            s_map.insert(s_val, p_val);
        }
    }
    return true;
}

mod tests {
    use super::*;

    #[test]
    fn test_word_pattern() {
        assert_eq!(word_pattern(String::from("abba"), String::from("dog cat cat dog")), true);
        assert_eq!(word_pattern(String::from("abba"), String::from("dog cat cat fish")), false);
        assert_eq!(word_pattern(String::from("aaa"), String::from("aa aa aa aa")), false);
    }
}
