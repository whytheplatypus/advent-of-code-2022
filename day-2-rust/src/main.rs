use std::env;
use std::fs;
use std::collections::HashMap;


fn main() {
    let args: Vec<String> = env::args().collect();
    let file_path = dbg!(&args[1]);
	println!("In file {}", file_path);

    let input = fs::read_to_string(file_path)
        .expect("Should have been able to read the file");

    let matches = input.split('\n');
    let total: i32 = matches.map(|game| game_score(game)).sum();
    println!("total: {}", total);
}

fn game_score(game: &str) -> i32 {
    let mut scores = HashMap::new();
    scores.insert("A", 1);
    scores.insert("B", 2);
    scores.insert("C", 3);

    let mut results = HashMap::new();
    results.insert("X", 0);
    results.insert("Y", 3);
    results.insert("Z", 6);

    let mut games = HashMap::new();
    // draws
    games.insert("A X", "C");
    games.insert("B Y", "B");
    games.insert("C Z", "A");
    // wins
    games.insert("A Y", "A");
    games.insert("B Z", "C");
    games.insert("C X", "B");
    // losses
    games.insert("A Z", "B");
    games.insert("B X", "A");
    games.insert("C Y", "C");

	match games.get(game) {
		None => 0,
		Some(shape) => {
			let shape_score = dbg!(scores.get(shape).unwrap());
			let result = dbg!(game.split(" ").last().unwrap());
			let result_score = dbg!(results.get(result).unwrap());
			shape_score + result_score
		}
	}
}

#[test]
fn test_match_score() {

    let score = game_score("A X");
    assert_eq!(score, 3);
}
