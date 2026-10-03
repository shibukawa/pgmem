import pytest
import pgmem


@pytest.mark.parametrize("base", [pgmem.Pgmem, pgmem.Server, pgmem.Snapshot])
def test_context_preserves_body_and_cleanup_failures(base):
    original = AssertionError("test failed")
    cleanup = RuntimeError("close failed")
    class Broken(base):
        def __init__(self):
            pass
        def close(self):
            raise cleanup
    with pytest.raises(pgmem.CleanupError) as caught:
        with Broken():
            raise original
    assert caught.value.primary_error is original
    assert caught.value.cleanup_errors == (cleanup,)
    assert caught.value.__cause__ is original
