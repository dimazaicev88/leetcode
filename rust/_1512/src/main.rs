use std::collections::HashMap;

fn main() {
    println!("{}", Solution::num_identical_pairs(vec![1,1,1,1]));
}

impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut map_nums: HashMap<i32, i32> = HashMap::new();
        let mut total_sum = 0;

        for num in nums {
            if map_nums.contains_key(&num) {
                let value = map_nums.get(&num);
                map_nums.insert(num, value.unwrap() + 1);
            } else {
                map_nums.insert(num, 1);
            }
        }

        for val in map_nums {
            total_sum += val.1 * (val.1 - 1) / 2
        }

        total_sum
    }
}

struct Solution;
