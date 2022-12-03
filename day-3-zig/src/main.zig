const std = @import("std");
const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();
    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [2046]u8 = undefined;
    var group = std.ArrayList([]u8).init(std.heap.page_allocator);
    var err_total: usize = 0;
    var badge_total: usize = 0;
    var i: usize = 0;
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var err = try find_error(line);
        var score = error_score(err);
        err_total += score;
        
        // did this to keep the memory space separate, otherwise the pointer to line would just get shifted.
        var experiment = std.ArrayList(u8).init(std.heap.page_allocator);
        try experiment.appendSlice(line);
        try group.append(experiment.items);
        std.debug.print("{s}.\n", .{ line });

        if (i%3 == 2) {
            var group_slice = group.items[i-2..i+1];
            std.debug.print("finding badge total {d}.\n", .{ i });
            std.debug.print("{s}.\n", .{ group_slice[0] });
            std.debug.print("{s}.\n", .{ group_slice[1] });
            std.debug.print("{s}.\n", .{ group_slice[2] });
            var badge = try find_badge(group_slice);
            badge_total += error_score(badge);
            std.debug.print("-------{d} {c}\n", .{ badge_total, badge });
        }
        i += 1;
    }
    std.debug.print("error total {d}.\n", .{ err_total });
    std.debug.print("badge total {d}.\n", .{ badge_total });
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

fn find_badge(group: [][]const u8) !u8 {
    for (group[0]) |one| {
        for (group[1]) |two| {
            for (group[2]) |three| {
                if (one == two and two == three) {
                    return one;
                }
            }
        }
    }
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

test "find badge" {
    var badge = try find_badge([3][]const u8{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"});
    try std.testing.expectEqual(@as(u8, 'r'), badge);
}

test "test error score" {
    const score = error_score('p');
    try std.testing.expectEqual(@as(usize, 16), score);
}
