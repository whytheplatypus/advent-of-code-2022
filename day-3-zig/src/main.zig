const std = @import("std");
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();
    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [2046]u8 = undefined;
    var total: usize = 0;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var err = try find_error(line);
        var score = error_score(err);
        std.debug.print("line score {c} {d} {d}.\n", .{ err, score, total });
        total += score;
        // do something with line...
    }
    std.debug.print("total {d}.\n", .{ total });
}



fn find_error(rucksack: []const u8) !u8 {
    var i: usize = 0;
    while (i < rucksack.len / 2) : (i += 1) {
        var first = rucksack[i]; 
        var j: usize = rucksack.len / 2;
        while (j < rucksack.len) : (j += 1) {
            var second = rucksack[j]; 
            if (first == second) {
                return first;
            }
        }
    }
    std.debug.print("trouble with rucksack {s}.\n", .{ rucksack });
    // update to return an error
    return error.InvalidChar;
}

fn error_score(char: u8) usize {
    for (alpha) |letter,index| {
        if (letter == char) {
            return index+1;
        }
    }
    return 0;
}

test "test error" {
    var rucksack = "vJrwpWtwJgWrhcsFMMfFFhFp";
    const match = try find_error(rucksack);

    std.debug.print("Error found {c}.\n", .{ match });
    try std.testing.expectEqual(@as(u8, 'p'), match);

    var edge_rucksack = "rwgRGdpGprNNLQLsbZJPsn";
    const edge_match = try find_error(edge_rucksack);
    try std.testing.expectEqual(@as(u8, 'N'), edge_match);

    var far_edge_rucksack = "RwgRGdpGprabLQLsbZJPsR";
    const far_edge_match = try find_error(far_edge_rucksack);
    try std.testing.expectEqual(@as(u8, 'R'), far_edge_match);
}

test "test error score" {
    const score = error_score('p');
    try std.testing.expectEqual(@as(usize, 16), score);
}
