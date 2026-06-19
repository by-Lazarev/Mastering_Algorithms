# 13. python - пример задачи(решение используя список)
def has_duplicates_naive(arr):
    n = len(arr)
    for i in range(n):
        for j in range(i + 1, n):
            if arr[i] == arr[j]:
                return True
    return False