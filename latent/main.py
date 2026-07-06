# design a key value store with transactions

from collections import deque


class KVStore:
    def __init__(self):
        self.transactions = deque()
        self.transactions.append({})

    def GET(self, key) -> str | None:
        if key in self.transactions[-1]:
            return self.transactions[-1][key]
        return None

    def SET(self, key, value) -> None:
        self.transactions[-1][key] = value

    def DELETE(self, key) -> None:
        if key in self.transactions[-1]:
            del self.transactions[-1][key]
        return None

    def BEGIN_TRANSACTION(self) -> None:
        copied = self.transactions[-1].copy()
        self.transactions.append(copied)
        return None

    def COMMIT(self) -> None:
        if len(self.transactions) == 1:
            return None

        popped = self.transactions[-1]
        self.transactions.pop()
        self.transactions.pop()
        self.transactions.append(popped)
        return None

    def ROLLBACK(self) -> None:
        if len(self.transactions) == 1:
            return None

        self.transactions.pop()
        return None


kv = KVStore()

assert kv.GET("somekey") is None
assert kv.SET("somekey", "somevalue") is None
assert kv.GET("somekey") == "somevalue"
assert kv.DELETE("somekey") is None
assert kv.GET("somekey") is None
# begin transaction
assert kv.BEGIN_TRANSACTION() is None
assert kv.SET("insidetransaction", "inside_transaction_value") is None
assert kv.GET("insidetransaction") == "inside_transaction_value"
assert kv.COMMIT() is None
assert kv.GET("insidetransaction") == "inside_transaction_value"
# commited

# begin transaction
assert kv.BEGIN_TRANSACTION() is None
assert kv.SET("will_be_rolled_back", "somevalue") is None
assert kv.GET("will_be_rolled_back") == "somevalue"
assert kv.ROLLBACK() is None
assert kv.GET("will_be_rolled_back") is None
# rolledback
