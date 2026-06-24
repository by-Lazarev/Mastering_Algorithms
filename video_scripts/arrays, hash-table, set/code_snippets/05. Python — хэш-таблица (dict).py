# 5. Python — хэш-таблица (dict)
ages = {
    "Alice": 28,
    "Bob": 32,
    "Carol": 25,
}

print(ages)          # {'Alice': 28, 'Bob': 32, 'Carol': 25}
print(ages["Bob"])   # 32

ages["Dave"] = 30
print("Alice" in ages)  # True
print(ages)          # {'Alice': 28, 'Bob': 32, 'Carol': 25, 'Dave': 30}