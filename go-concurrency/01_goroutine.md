


# Go 동시성 모델

## 01 동시성과 병렬성

### Concurrency

> 여러 작업을 동시에 다루는것. 기본적으로 구조(Structure 의 문제)


```go

func main() {
  go task1()
  go task2()
  go task3()
}

```


---

**_Concept_**

- **Concurrency** :  여러 작업을 논리적으로 동시에 다루는 구조
- **Parallelism** :  여러 작업을 물리적으로 동시에 실행.  멀티코어 필요
- **Green Tread** : OS가 아닌 런타임이 관리하는 경량 스레드 개념

---
