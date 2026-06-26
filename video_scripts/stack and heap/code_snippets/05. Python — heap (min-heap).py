#5. Python — heap (min-heap)
import heapq

heap = []

heapq.heappush(heap, 30)
heapq.heappush(heap, 10)
heapq.heappush(heap, 20)

print(heap) #[10,30,20]

top = heapq.heappop(heap)
print(top) #10
print(heap) #[20, 30]