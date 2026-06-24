# 14. python - пример задачи(решение используя множество)
def has_duplicates_set(arr):
    seen = set()
    for x in arr:
        if x in seen:
            return True
        seen.add(x)
    return False