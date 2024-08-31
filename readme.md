Нужно написать простую библиотеку in-memory cache.

Для простоты считаем, что у нас бесконечная память и нам не нужно задумываться об удалении ключей из него.

Реализация должна удовлетворять интерфейсу:


type Cache interface {
    Set(k, v string)
    Get(k string) (v string, ok bool)
}
