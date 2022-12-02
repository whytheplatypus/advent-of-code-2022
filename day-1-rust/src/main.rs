use std::str::FromStr;
use std::fs;

fn main() {
    let file_path = "input.txt";
	println!("In file {}", file_path);

    let input = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");
    //println!("{}", input);
    println!("{}", biggest_horder(&input));
}

fn biggest_horder(input: &str) -> i32 {
    let raw_elves = split_elves(input);
    let mut elves = Vec::<Vec::<i32>>::new();
    for raw_elf in raw_elves {
        elves.push(read_elf(raw_elf).unwrap());
    }
    let max_calories = elves.iter().map(|elf| elf_total_cals(elf)).max();
    max_calories.unwrap()
}

fn split_elves(input: &str) -> impl Iterator::<Item = &str> {
    input.split("\n\n")
}

fn read_elf(input: &str) -> Result<Vec::<i32>, Box<dyn std::error::Error>> {
    let items = input.split('\n');
    let mut calories = Vec::<i32>::new();
    for item in items {
        match FromStr::from_str(item) {
            Ok(amount) => calories.push(amount),
            Err(error) => println!("{}", error),
        }
    }
    Ok(calories)
}

fn elf_total_cals(elf: &Vec::<i32>) -> i32 {
    elf.iter().sum()
}

const TEST_INPUT: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000";


#[test]
fn test_biggest_horder() {
    assert_eq!(biggest_horder(TEST_INPUT), 24000);
}

#[test]
fn test_split_elves() {
    let mut raw_elves = split_elves(TEST_INPUT);
    assert_eq!(raw_elves.next().unwrap(), "1000
2000
3000");
    assert_eq!(raw_elves.last().unwrap(), "10000");

}

#[test]
fn test_read_elf() {
    let test_elf = [1000, 2000, 3000];
    assert_eq!(read_elf("1000
2000
3000").unwrap(), test_elf);
}
