fn bubble_short_asc(mut in_arr: [i32; 10]) -> [i32; 10] {
    for i in 0..9 {
        for j in i+1..10 {
            if in_arr[j] <= in_arr[i] {
                let temp: i32 = in_arr[i];
                in_arr[i] = in_arr[j];
                in_arr[j] = temp;
            }
        }
    }
    return in_arr;
}

fn bubble_short_desc(mut in_arr: [i32; 10]) -> [i32; 10] {
    for i in 0..9 {
        for j in i+1..10 {
            if in_arr[j] >= in_arr[i] {
                let temp: i32 = in_arr[i];
                in_arr[i] = in_arr[j];
                in_arr[j] = temp;
            }
        }
    }
    return in_arr;
}

fn main() {
    let rand_arr: [i32; 10] = [233302, 32, 23923, 3421, 2, 889, 3, 4839, 32123, 19998];
    print!("initial array : ");
    for el in rand_arr {
        print!("{} ", el);
    }
    println!();

    let asc_arr: [i32; 10] = bubble_short_asc(rand_arr);
    for el in asc_arr {
        print!("{} ", el);
    }
    println!();

    let desc_arr: [i32; 10] = bubble_short_desc(rand_arr);
    for el in desc_arr {
        print!("{} ", el);
    }
    println!();
}
