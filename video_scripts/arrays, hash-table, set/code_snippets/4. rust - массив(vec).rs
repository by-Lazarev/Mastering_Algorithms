// 4. rust - массив(vec)
fn main() {
    let mut nums = vec![10, 20, 30, 40];

    println!("{:?}", nums); // [10, 20, 30, 40]
    println!("{}", nums[1]); // 20

    nums.push(50);
    nums.insert(1, 15);

    println!("{:?}", nums);
}