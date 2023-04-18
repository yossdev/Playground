use std::{env, fs, path::Path, time::SystemTime};

fn main() {
    let start = SystemTime::now();

    let args: Vec<String> = env::args().collect();
    let mut cmd_args = "";
    if args.len() > 1 {
        cmd_args = &args[1];
    }

    const ROOT: &str = "D:/Code/Playground";
    const CWD: &str = "rust/copy-file";
    const DIR: &str = "dst";

    if cmd_args == "clean" {
        let path_dir = format!("{}/{}/{}", &ROOT, &CWD, &DIR);
        clean_up(&path_dir);
    } else {
        const COUNT: i16 = 10_000;
        let src = format!("{}/common/src/test.md", ROOT);
        copy_file(COUNT, &src, &DIR);
    }

    let end = SystemTime::now();
    let elapsed = end.duration_since(start);

    println!(
        "Time took in rust: {}ms\n",
        elapsed.unwrap_or_default().as_millis()
    )
}

fn clean_up(path_dir: &str) {
    let path = Path::new(path_dir);
    let total_size = dir_size(path);

    fs::remove_dir_all(path_dir).unwrap_or_default();

    fs::create_dir(path_dir).unwrap_or_default();

    print!("Total deleted: {} bytes\n", total_size)
}

fn copy_file(count: i16, src: &str, dir: &str) {
    let start = SystemTime::now();

    let mut sum_n_bytes: u64 = 0;

    for i in 0..=count - 1 {
        let dst = format!("./{}/test{}.md", dir, i);
        let size = fs::copy(src, &dst).unwrap_or_default();
        // Ok(());

        sum_n_bytes += &size;
    }

    let end = SystemTime::now();
    let elapsed = end.duration_since(start);

    print!(
        "Time took in rust create_file: {}ms\n",
        elapsed.unwrap_or_default().as_millis()
    );
    print!("Total written: {} bytes\n", sum_n_bytes);
}

fn dir_size(path_dir: &Path) -> u64 {
    let mut total: u64 = 0;

    if path_dir.is_dir() {
        let entries = fs::read_dir(path_dir).unwrap();
        for entry in entries {
            let meta = fs::metadata(entry.unwrap().path());
            total += meta.unwrap().len();
        }
    }

    return total;
}
