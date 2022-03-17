sub bubble_sort_asc {
    $arr = $_[0];
    @in_arr = @arr;
    $max_arr = $_[1];

    for ($i = 0; $i < $max_arr-1; $i++) {
        for($j = $i+1; $j < $max_arr; $j++) {
            if (@in_arr[$j] <= @in_arr[$i]) {
                $temp = @in_arr[$i];
                @in_arr[$i] = @in_arr[$j];
                @in_arr[$j] = $temp;
            }
        }
    }

    return (@in_arr);
}

sub bubble_sort_desc {
    $arr = $_[0];
    @in_arr = @arr;
    $max_arr = $_[1];

    for ($i = 0; $i < $max_arr-1; $i++) {
        for($j = $i+1; $j < $max_arr; $j++) {
            if (@in_arr[$j] > @in_arr[$i]) {
                $temp = @in_arr[$i];
                @in_arr[$i] = @in_arr[$j];
                @in_arr[$j] = $temp;
            }
        }
    }

    return (@in_arr);
}

@arr = (123, 3129, 32331, 319320, 2931, 3, 454, 223, 123, 129339);
$max_arr = 10;
print join(" ", @arr);
print "\n";

@desc_arr = bubble_sort_desc(@arr, $max_arr);
@asc_arr = bubble_sort_asc(@arr, $max_arr);

print join(" ", @desc_arr);
print "\n";
print join(" ", @asc_arr);
print "\n";
