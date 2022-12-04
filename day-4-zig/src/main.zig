const std = @import("std");

pub fn main() !void {
    // Prints to stderr (it's a shortcut based on `std.io.getStdErr()`)
    std.debug.print("All your {s} are belong to us.\n", .{"codebase"});

    // stdout is for the actual output of your application, for example if you
    // are implementing gzip, then only the compressed bytes should be sent to
    // stdout, not any debugging messages.
    const stdout_file = std.io.getStdOut().writer();
    var bw = std.io.bufferedWriter(stdout_file);
    const stdout = bw.writer();

    try stdout.print("Run `zig build test` to run the tests.\n", .{});

    try bw.flush(); // don't forget to flush!

    const input_file = fetch_file_arg();
    const lines = try read_file_to_lines(input_file);
    var covered_count: usize = 0;
    var intersected_count: usize = 0;
    for (lines) |line| {
        const areas = split_line(line);
        const elf_one_area = try parse_assignment(areas[0]);
        const elf_two_area = try parse_assignment(areas[1]);
        if (check_coverage(elf_one_area, elf_two_area)) {
            covered_count += 1;
        }
        if (check_intersection(elf_one_area, elf_two_area)) {
            intersected_count += 1;
        }
    }
    std.debug.print("fully redunant elves {d}.\n", .{ covered_count });
    std.debug.print("mostly redunant elves {d}.\n", .{ intersected_count });
}

fn check_coverage(first: [2]i32, second: [2]i32) bool {
    const first_covers_second = first[0] <= second[0] and first[1] >= second[1]; 
    const second_covers_first = first[0] >= second[0] and first[1] <= second[1];
    return first_covers_second or second_covers_first;
}

fn check_intersection(first: [2]i32, second: [2]i32) bool {
    const first_separate = first[1] < second[0]; 
    const second_separate = second[1] < first[0];
    return !(first_separate or second_separate);
}

fn split_line(line: []const u8) [2][]const u8 {
    var iter = std.mem.split(u8, line, ",");
    return [2][]const u8{iter.next().?, iter.next().?};
}

test "split line test" {
    const areas = split_line("1-5,3-5");

    std.debug.print("{s} {s}.\n", .{ areas[0], areas[1] });
}

fn parse_assignment(assignment: []const u8) ![2]i32 {
    var iter = std.mem.split(u8, assignment, "-");
    var first = try std.fmt.parseInt(i32, iter.next().?, 0);
    var second = try std.fmt.parseInt(i32, iter.next().?, 0);
    return [2]i32{first, second};
}

test "parse assignment"{
    const area = try parse_assignment("1-5");

    std.debug.print("{d} {d}.\n", .{ area[0], area[1] });
}

fn fetch_file_arg() []const u8 {
    var args = std.process.args();
    _ = args.skip();
    return args.next().?;
}

fn read_file_to_lines(file_name: []const u8) ![][]const u8 {
    var file = try std.fs.cwd().openFile(file_name, .{});
    defer file.close();
    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [2046]u8 = undefined;
    var group = std.ArrayList([]u8).init(std.heap.page_allocator);
    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var copy = std.ArrayList(u8).init(std.heap.page_allocator);
        try copy.appendSlice(line);
        try group.append(copy.items);
    }
    return group.items;
}

test "simple test" {
    var list = std.ArrayList(i32).init(std.testing.allocator);
    defer list.deinit(); // try commenting this out and see if zig detects the memory leak!
    try list.append(42);
    try std.testing.expectEqual(@as(i32, 42), list.pop());
}
