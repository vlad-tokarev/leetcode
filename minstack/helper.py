import sys
import json

if __name__ == '__main__':
    instructions = json.loads(next(sys.stdin))
    values = json.loads(next(sys.stdin))

    for i, v in zip(instructions, values):
        if i == "MinStack":
            continue

        if i == "push":
            print(f"Push({v[0]}),")
        if i == "pop":
            print("Pop(),")
        if i == "pop":
            print("Pop(),")
        if i == "getMin":
            print("GetMin(),")
        if i == "top":
            print("Top(),")

