arr = [1312, 1233, 3414214, 323, 3213, 2312313132, 3, 2312, 13221, 312331231321]

def bubble_short_asc(arr, max)
  for i in 0..max-2
    for j in i+1..max-1
      if arr[j]<=arr[i]
        temp = arr[i]
        arr[i] = arr[j]
        arr[j] = temp
      end
    end
  end
  return arr
end

def bubble_short_desc(arr, max)
  for i in 0..max-2
    for j in i+1..max-1
      if arr[j]>=arr[i]
        temp = arr[i]
        arr[i] = arr[j]
        arr[j] = temp
      end
    end
  end
  return arr
end

puts "#{arr.inspect}"
puts "#{bubble_short_asc(arr, 10).inspect}"
puts "#{bubble_short_desc(arr, 10).inspect}"
