from concurrent.futures import ProcessPoolExecutor
import time

def fn(x, y):
    print("calculating...")
    time.sleep(30)
    arg_sum = x + y
    return arg_sum


if __name__ == '__main__':

    with ProcessPoolExecutor(max_workers=5) as executor:
        future = executor.submit(fn, 2, 11)
        future = executor.submit(fn, 3, 12)
        print(future.running())
        print(future.result())
