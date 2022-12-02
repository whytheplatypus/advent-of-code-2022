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
    scores.insert("X", 1);
    scores.insert("Y", 2);
    scores.insert("Z", 3);

    let mut games = HashMap::new();
    // draws
    games.insert("A X", 3);
    games.insert("B Y", 3);
    games.insert("C Z", 3);
    // wins
    games.insert("A Y", 6);
    games.insert("B Z", 6);
    games.insert("C X", 6);
    // losses
    games.insert("A Z", 0);
    games.insert("B X", 0);
    games.insert("C Y", 0);

    let game_result = games.get(game);
	match game_result {
		None => 0,
		Some(score) => score + scores.get(dbg!(game.split(" ").last().unwrap())).unwrap()
	}
}

#[test]
fn test_match_score() {

    let score = game_score("A X");
    assert_eq!(score, 4);
}
