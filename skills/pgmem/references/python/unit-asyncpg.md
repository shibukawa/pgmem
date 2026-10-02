# asyncpg unit tests

Mark an async pytest test with `@shadow_pg` and the project's asyncio test marker. `await asyncpg.connect(...)` inside the test reaches a fresh fork; close the connection in `finally`. `asyncpg.create_pool(...)` also follows routing when created inside the marked scope. Finish all tasks using its connections before marker teardown. Use `fork=False` only for read-only tests. See [unit lifecycle](unit.md).
