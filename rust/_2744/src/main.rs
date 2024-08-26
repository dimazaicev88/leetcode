use std::collections::HashMap;


fn main() {
    let arr = vec![String::from("cd"), String::from("ac"), String::from("dc"),
                   String::from("ca"), String::from("zz")];

    println!("{}", Solution::maximum_number_of_string_pairs(arr));
}

impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
        let mut count = 0;
        let mut map_words: HashMap<String, bool> = HashMap::new();

        for i in 0..words.len() {
            if map_words.contains_key(&words[i]) {
                count += 1;
            }
            map_words.insert(words[i].chars().rev().collect::<String>(), true);
        }

        count
    }
}

struct Solution;